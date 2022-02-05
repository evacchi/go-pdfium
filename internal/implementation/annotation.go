package implementation

// #cgo pkg-config: pdfium
// #include "fpdfview.h"
import "C"
import (
	"github.com/google/uuid"
	"github.com/klippa-app/go-pdfium/references"
)

func (p *PdfiumImplementation) registerAnnotation(annotation C.FPDF_ANNOTATION, documentHandle *DocumentHandle) *AnnotationHandle {
	ref := uuid.New()
	handle := &AnnotationHandle{
		handle:      annotation,
		nativeRef:   references.FPDF_ANNOTATION(ref.String()),
		documentRef: documentHandle.nativeRef,
	}

	documentHandle.annotationRefs[handle.nativeRef] = handle
	p.annotationRefs[handle.nativeRef] = handle

	return handle
}