package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	configure "github.com/miniyus/go-fiber/config"
	"github.com/miniyus/go-fiber/database"
	"github.com/miniyus/go-fiber/internal/core/container"
	"github.com/miniyus/go-fiber/internal/core/register"
)

func Run() {
	config := configure.GetConfigs()

	wrapper := container.NewContainer(
		fiber.New(config.App),
		database.DB(config.Database),
		config,
	)

	register.Resister(wrapper)

	wrapper.Run()
}
