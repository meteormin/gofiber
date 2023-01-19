package register_test

import (
	"github.com/miniyus/gofiber/internal/core"
	"github.com/miniyus/gofiber/internal/core/register"
	"github.com/miniyus/gofiber/pkg/worker"
	"log"
	"testing"
)

func TestRegister(t *testing.T) {
	c := core.New()
	register.Resister(c)

	var jobDispatcher worker.Dispatcher

	c.Resolve(&jobDispatcher)

	log.Print(jobDispatcher)
}
