package app

import "net/http"

func NewServeMux(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, r := range routes {
		mux.Handle(r.Pattern(), r)
	}
	return mux
}

type Route interface {
	http.Handler
	Pattern() string
}
