package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type EchoHandler struct {
	path string
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(path string) *EchoHandler {
	return &EchoHandler{path: path}
}

func (h *EchoHandler) Pattern() string {
	return "/echo"
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (*EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to handle request:", err)
	}
}
