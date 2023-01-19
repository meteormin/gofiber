package core

import (
	"github.com/gofiber/fiber/v2"
	configure "github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/internal/core/api_error"
	"github.com/miniyus/gofiber/internal/core/container"
	"github.com/miniyus/gofiber/internal/core/database"
	"github.com/miniyus/gofiber/internal/core/register"
)

// New
// IoC 컨테이너 생성
func New() container.Container {
	config := configure.GetConfigs()
	config.App.ErrorHandler = api_error.OverrideDefaultErrorHandler

	wrapper := container.New(
		fiber.New(config.App),
		database.DB(config.Database),
		config,
	)

	register.Resister(wrapper)

	return wrapper
}
