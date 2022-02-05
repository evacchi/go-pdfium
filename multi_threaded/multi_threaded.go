package multi_threaded

import (
	goctx "context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/klippa-app/go-pdfium"
	"github.com/klippa-app/go-pdfium/internal/commons"
	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	pool "github.com/jolestar/go-commons-pool/v2"
)

type worker struct {
	plugin       commons.Pdfium
	pluginClient *plugin.Client
	rpcClient    plugin.ClientProtocol
}

type Config struct {
	MinIdle     int
	MaxIdle     int
	MaxTotal    int
	LogCallback func(string)
	Command     Command
}

type Command struct {
	BinPath string
	Args    []string
}

type pdfiumPool struct {
	workerPool   *pool.ObjectPool
	instanceRefs map[string]*pdfiumInstance
	poolRef      string
	closed       bool
	lock         *sync.Mutex
}

var poolRefs = map[string]*pdfiumPool{}
var multiThreadedMutex = &sync.Mutex{}

// Init will return a multi-threaded pool.
// It will launch a new worker for every requested instance as long as the limits
// allow it. If the pool has been exhausted. It will wait until a worker becomes
// available. So it's important that you close instances when you're done with them.
func Init(config Config) pdfium.Pool {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "BASIC_PLUGIN",
		MagicCookieValue: "hello",
	}

	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"pdfium": &commons.PdfiumPlugin{},
	}

	factory := pool.NewPooledObjectFactory(
		func(goctx.Context) (interface{}, error) {
			newWorker := &worker{}

			client := plugin.NewClient(&plugin.ClientConfig{
				HandshakeConfig: handshakeConfig,
				Plugins:         pluginMap,
				Cmd:             exec.Command(config.Command.BinPath, config.Command.Args...),
				Logger:          logger,
			})

			rpcClient, err := client.Client()
			if err != nil {
				return nil, err
			}

			raw, err := rpcClient.Dispense("pdfium")
			if err != nil {
				return nil, err
			}

			pdfium := raw.(commons.Pdfium)

			pong, err := pdfium.Ping()
			if err != nil {
				return nil, err
			}

			if pong != "Pong" {
				return nil, errors.New("Wrong ping/pong result")
			}

			newWorker.pluginClient = client
			newWorker.rpcClient = rpcClient
			newWorker.plugin = pdfium

			return newWorker, nil
		}, nil, func(ctx goctx.Context, object *pool.PooledObject) bool {
			worker := object.Object.(*worker)
			if worker.pluginClient.Exited() {
				config.LogCallback("Worker exited")
				return false
			}

			err := worker.rpcClient.Ping()
			if err != nil {
				config.LogCallback(fmt.Sprintf("Error on RPC ping: %s", err.Error()))
				return false
			}

			pong, err := worker.plugin.Ping()
			if err != nil {
				config.LogCallback(fmt.Sprintf("Error on plugin ping:: %s", err.Error()))
				return false
			}

			if pong != "Pong" {
				err = errors.New("Wrong ping/pong result")
				config.LogCallback(fmt.Sprintf("Error on plugin ping:: %s", err.Error()))
				return false
			}

			return true
		}, nil, nil)
	p := pool.NewObjectPoolWithDefaultConfig(goctx.Background(), factory)
	p.Config = &pool.ObjectPoolConfig{
		BlockWhenExhausted: true,
		MinIdle:            config.MinIdle,
		MaxIdle:            config.MaxIdle,
		MaxTotal:           config.MaxTotal,
		TestOnBorrow:       true,
		TestOnReturn:       true,
		TestOnCreate:       true,
	}

	multiThreadedMutex.Lock()
	defer multiThreadedMutex.Unlock()

	poolRef := uuid.New()

	// Create a new PDFium pool.
	newPool := &pdfiumPool{
		poolRef:      poolRef.String(),
		instanceRefs: map[string]*pdfiumInstance{},
		lock:         &sync.Mutex{},
		workerPool:   p,
	}

	poolRefs[newPool.poolRef] = newPool

	return newPool
}

func (p *pdfiumPool) GetInstance(timeout time.Duration) (pdfium.Pdfium, error) {
	if p.closed {
		return nil, errors.New("pool is closed")
	}

	timeoutCtx, cancel := goctx.WithTimeout(goctx.Background(), timeout)
	defer cancel()
	workerObject, err := p.workerPool.BorrowObject(timeoutCtx)
	if err != nil {
		return nil, err
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	newInstance := &pdfiumInstance{
		worker: workerObject.(*worker),
		lock:   &sync.Mutex{},
	}

	instanceRef := uuid.New()
	newInstance.instanceRef = instanceRef.String()
	newInstance.pool = p
	p.instanceRefs[newInstance.instanceRef] = newInstance

	return newInstance, nil
}

func (p *pdfiumPool) Close() (err error) {
	if p.closed {
		return errors.New("pool is already closed")
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	defer func() {
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("panic occurred in %s: %v", "Close", panicError)
		}
	}()

	// Close all instances
	for i := range p.instanceRefs {
		p.instanceRefs[i].Close()
		delete(p.instanceRefs, i)
	}

	multiThreadedMutex.Lock()
	delete(poolRefs, p.poolRef)
	multiThreadedMutex.Unlock()

	// Close the underlying pool.
	p.workerPool.Close(goctx.Background())

	return nil
}

type pdfiumInstance struct {
	worker      *worker
	pool        *pdfiumPool
	instanceRef string
	closed      bool
	lock        *sync.Mutex
}

// NewDocumentFromBytes creates a new PDFium document from a byte array.
// This will automatically select a worker and keep it for you until you execute
// the close method on the references.
func (i *pdfiumInstance) NewDocumentFromBytes(file *[]byte, opts ...pdfium.NewDocumentOption) (doc *references.FPDF_DOCUMENT, err error) {
	i.lock.Lock()
	if i.closed {
		i.lock.Unlock()
		return nil, errors.New("instance is closed")
	}
	i.lock.Unlock()

	defer func() {
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("panic occurred in %s: %v", "NewDocumentFromBytes", panicError)
		}
	}()

	openDocRequest := &requests.OpenDocument{File: file}
	for _, opt := range opts {
		opt.AlterOpenDocumentRequest(openDocRequest)
	}

	newDoc, err := i.worker.plugin.OpenDocument(openDocRequest)
	if err != nil {
		return nil, err
	}

	return &newDoc.Document, nil
}

// NewDocumentFromFilePath creates a new PDFium document from a file path.
func (i *pdfiumInstance) NewDocumentFromFilePath(filePath string, opts ...pdfium.NewDocumentOption) (doc *references.FPDF_DOCUMENT, err error) {
	i.lock.Lock()
	if i.closed {
		i.lock.Unlock()
		return nil, errors.New("instance is closed")
	}
	i.lock.Unlock()

	defer func() {
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("panic occurred in %s: %v", "NewDocumentFromFilePath", panicError)
		}
	}()

	openDocRequest := &requests.OpenDocument{FilePath: &filePath}
	for _, opt := range opts {
		opt.AlterOpenDocumentRequest(openDocRequest)
	}

	newDoc, err := i.worker.plugin.OpenDocument(openDocRequest)
	if err != nil {
		return nil, err
	}

	return &newDoc.Document, nil
}

// NewDocumentFromReader creates a new PDFium document from a reader.
func (i *pdfiumInstance) NewDocumentFromReader(reader io.ReadSeeker, size int, opts ...pdfium.NewDocumentOption) (*references.FPDF_DOCUMENT, error) {
	i.lock.Lock()
	if i.closed {
		i.lock.Unlock()
		return nil, errors.New("instance is closed")
	}
	i.lock.Unlock()

	// Since multi-threaded usage implements gRPC, it can't serialize the reader onto that.
	// To make it support the full interface, we just complete read the file into memory.
	fileData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return i.NewDocumentFromBytes(&fileData, opts...)
}

func (i *pdfiumInstance) Close() (err error) {
	i.lock.Lock()

	if i.closed {
		i.lock.Unlock()
		return errors.New("instance is already closed")
	}

	defer func() {
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("panic occurred in %s: %v", "Close", panicError)
		}
	}()

	defer func() {
		i.pool.workerPool.ReturnObject(goctx.Background(), i.worker)
		i.worker = nil
		delete(i.pool.instanceRefs, i.instanceRef)
		i.pool = nil
		i.closed = true
		i.lock.Unlock()
	}()

	return i.worker.plugin.Close()
}

func (i *pdfiumInstance) FPDF_CloseDocument(document references.FPDF_DOCUMENT) (err error) {
	if i.closed {
		return errors.New("instance is closed")
	}

	defer func() {
		if panicError := recover(); panicError != nil {
			err = fmt.Errorf("panic occurred in %s: %v", "FPDF_CloseDocument", panicError)
		}
	}()

	return i.worker.plugin.FPDF_CloseDocument(document)
}
