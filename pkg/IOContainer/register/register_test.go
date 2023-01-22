package register_test

import (
	"github.com/gofiber/fiber/v2"
	configure "github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/internal/database"
	"github.com/miniyus/gofiber/pkg/IOContainer"
	"github.com/miniyus/gofiber/pkg/IOContainer/register"
	"github.com/miniyus/gofiber/pkg/worker"
	"log"
	"testing"
)

func TestRegister(t *testing.T) {
	cfg := configure.GetConfigs()
	c := IOContainer.NewContainer(fiber.New(), database.DB(cfg.Database), cfg)

	register.Resister(c)

	var jobDispatcher worker.Dispatcher

	c.Resolve(&jobDispatcher)

	log.Print(jobDispatcher)
}
