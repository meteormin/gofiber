package main

import (
	"github.com/miniyus/gofiber"
	App "github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/monitor"
)

// @title gofiber Swagger API Documentation
// @version 1.0,0
// @description gofiber API
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
	app.Route("/analysis", func(router App.Router, app App.Application) {
		router.Route("/", monitor.New(app))
	})
	app.Status()
	app.Run()
}
