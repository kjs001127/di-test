package main

import (
	"di-test/app"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func TestWithFx(t *testing.T) {

	var h app.Route
	a := testApp(t, &h)
	a.RequireStart()

	assert.NotNil(t, h)
	assert.Equal(t, "/test", h.Pattern())

	a.RequireStop()
}

func testApp(t *testing.T, pop ...interface{}) *fxtest.App {
	return fxtest.New(
		t,
		testModule(),
		commonModule(),
		fx.Populate(pop...), // Populate 를 통해 컨테이너로부터 인젝션된 인스턴스를 꺼내올 수 있음
	)
}
