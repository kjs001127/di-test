package app

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type TestHandler struct {
	log *zap.Logger
}

// NewTestHandler builds a new TestHandler.
func NewTestHandler(log *zap.Logger) *TestHandler {
	return &TestHandler{log: log}
}

func (*TestHandler) Pattern() string {
	return "/test"
}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
