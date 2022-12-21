//go:build wireinject
// +build wireinject

package main

import (
	"di-test/app"
	"github.com/google/wire"
	"go.uber.org/zap"
	"net/http"
)

var server = wire.NewSet(
	app.NewHTTPServer,
	app.NewServeMux,
)
var prodRoute = wire.NewSet(
	app.NewEchoHandler,
	app.NewHelloHandler,
	wire.Value("echoPath"),
	zap.NewProduction,
	NewProdRouters,
)
var devRoute = wire.NewSet(
	app.NewTestHandler,
	zap.NewDevelopment,
	NewDevRouters,
)

func NewTestServer() *http.Server {
	panic(
		wire.Build(
			server,
			devRoute,
		),
	)
}

func NewServer() *http.Server {
	panic(
		wire.Build(
			server,
			prodRoute,
		),
	)
}

func NewLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}

// group 대신 아래와 같은 방법으로 슬라이스 주입
func NewProdRouters(e *app.EchoHandler, h *app.HelloHandler) []app.Route {
	return []app.Route{e, h}
}

func NewDevRouters(e *app.TestHandler) []app.Route {
	return []app.Route{e}
}

func WithMockRouters(mocks []app.Route) *http.Server {
	panic(
		wire.Build(server),
	)
}
