package monitor_test

import (
	"github.com/miniyus/gofiber"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/monitor"
	"io"
	"net/http/httptest"
	"testing"
)

func TestNewAnalysis(t *testing.T) {
	a := gofiber.New()
	a.Bootstrap()
	ai := monitor.NewAnalysis(a)

	marshal, err := ai.Marshal(true)

	if err != nil {
		t.Log(ai)
		t.Error(err)
	}

	t.Log(marshal)
}

func TestNew(t *testing.T) {
	a := gofiber.New()
	a.Route("/monitor", func(router app.Router, app app.Application) {
		router.Route("/", monitor.New(app))
	})
	a.Bootstrap()
	req := httptest.NewRequest("GET", "/monitor/", nil)
	test, err := a.Test(req)
	if err != nil {
		t.Error(err)
	}
	body, err := io.ReadAll(test.Body)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(body))
}
