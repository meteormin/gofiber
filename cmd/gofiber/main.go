package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/miniyus/gofiber"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/internal/api_error"
	"github.com/miniyus/gofiber/internal/resolver"
	"github.com/miniyus/gofiber/internal/routes"
)

// @title keyword-search-backend Swagger API Documentation
// @version 1.0.0
// @description keyword-search-backend API
// @contact.name miniyus
// @contact.url https://miniyus.github.io
// @contact.email miniyu97@gmail.com
// @license.name MIT
// @host localhost:9090
// @BasePath /
// @schemes http
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				   Bearer token type
func main() {
	app := gofiber.New()

	app.Middleware(func(fiberApp *fiber.App, application gofiber.Application) {
		configure := application.Config()

		fiberApp.Use(flogger.New(configure.Logger))
		fiberApp.Use(recover.New(recover.Config{
			EnableStackTrace: !app.IsProduction(),
		}))
		fiberApp.Use(api_error.ErrorHandler)
		fiberApp.Use(cors.New(configure.Cors))

		// Add Context Config
		fiberApp.Use(config.AddContext(config.ConfigsKey, configure))
		// Add Context Logger
		fiberApp.Use(config.AddContext(config.LoggerKey, resolver.MakeLogger(configure)))
	})

	app.Route(routes.ApiPrefix, routes.Api, "api")
	app.Route("/", routes.External, "external")

	app.Run()
}
