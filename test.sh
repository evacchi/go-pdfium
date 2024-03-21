#!/usr/bin/env bash

# Exit the script if any subcommand fails.
set -e

# The Wasm file is given as the first and only argument to the script.
WASM=$1

# Run the Wasm in Wasmtime and `grep` for our target bug's panic
# message.

cp $WASM webassembly/pdfium.wasm
# go test -timeout 5s ./internal/implementation_webassembly -v --ginkgo.v --ginkgo.focus alpha channel -count 1 -args 2>&1 
go test -timeout 5s ./internal/implementation_webassembly -v --ginkgo.v --ginkgo.focus alpha channel -count 1 -args 2>&1 | egrep --quiet '(timed out)|(wasm error)'
