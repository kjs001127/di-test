package app

import (
	"net/http"
)

func NewHTTPServer(mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	return srv
}
