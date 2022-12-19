package app

import (
	"go.uber.org/fx"
	"net/http"
)

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	return srv
}
