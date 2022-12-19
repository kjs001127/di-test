package main

import (
	"context"
	"di-test/app"
	"fmt"
	"go.uber.org/dig"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

func main() {

	var module fx.Option
	switch os.Getenv("GO_ENV") {
	case "prod":
		module = prodModule()
	case "test":
		module = testModule()
	default:
		panic("unknown env")
	}

	fx.New(
		module,
		commonModule(),
		fx.Invoke(func(s *http.Server, lc fx.Lifecycle) { // invoke 는 fx.New (Initialize) 시점에 호출됌
			lc.Append(fx.Hook{ // fx..Hook 을 통해 추가된 lifecycle hook invoke 이후 execution 에서 실행됌
				OnStart: func(ctx context.Context) error {
					return startHttpSvr(s)
				},
				OnStop: func(ctx context.Context) error {
					return stopHttpSvr(ctx, s)
				},
			})
		}),
	).Run()

	fx.Populate()
}

func startHttpSvr(s *http.Server) error {
	fmt.Println("Starting server")
	go s.ListenAndServe()
	return nil
}

func stopHttpSvr(ctx context.Context, s *http.Server) error {
	fmt.Println("ending server", s)
	return s.Shutdown(ctx)
}

func testModule() fx.Option {
	return fx.Module("test",
		fx.Provide(
			fx.Annotate(
				app.NewTestHandler,
				fx.As(new(app.Route)),
				fx.ResultTags(`name:"test"`),
			),
			zap.NewDevelopment,
		),
	)
}

func prodModule() fx.Option {
	return fx.Module("prod",
		fx.Supply( // fx.Supply -> concrete value 를 제공, fx.Annotated 를 통해 name, group 지정 가능
			fx.Annotated{
				Name:   "echoPath",
				Target: "/echo", // 제공할 값
			},
		),
		fx.Provide(
			fx.Annotate(
				app.NewHelloHandler,
				fx.As(new(app.Route)),
				fx.ResultTags(`group:"route"`),
			),
			fx.Annotate(
				app.NewEchoHandler,
				fx.As(new(app.Route)), // fx.As 로 Route 인터페이스로 바인딩하겠다고 선언, 여러개 선언 가능
				fx.ResultTags(`group:"route"`),
				fx.ParamTags(`name:"echoPath"`), // fx.ParamTags 로 name 지정
			),
			zap.NewProduction,
		),
	)
}

func commonModule() fx.Option {
	return fx.Module("common", fx.Provide(
		app.NewHTTPServer, // 기본적으로는 이렇게 생성자 함수를 제공
		fx.Annotate( // fx.Annotate 로 부가적인 정보 제공
			app.NewServeMux,
			fx.ParamTags(`group:"route"`), // fx.ParamTags 로 group 지정
			// 파라미터 []Route 에 group:route 로 태깅된 인스턴스들을 주입
		),
	))
}

func main2() {
	c := dig.New()

	err := c.Provide("db", dig.Name("a"))
	if err != nil {
		panic(err)
	}

	err = c.Provide(app.NewHelloHandler, dig.As(new(app.Route)), dig.Group("route"))
	if err != nil {
		panic(err)
	}

	err = c.Provide(app.NewEchoHandler, dig.As(new(app.Route)), dig.Group("route"))
	if err != nil {
		panic(err)
	}

	err = c.Provide(app.NewServeMux, dig.Group("route"))
	err = c.Invoke(func(r app.Route) {
		fmt.Print(r.Pattern())
	})
	if err != nil {
		return
	}
	time.Sleep(time.Hour)
}
