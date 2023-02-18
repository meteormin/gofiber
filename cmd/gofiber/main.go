package main

import (
	"github.com/miniyus/gofiber"
	"github.com/miniyus/gofiber/routes"
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
	app.Route(routes.ApiPrefix, routes.Api, "api")
	app.Route("/", routes.External, "external")

	app.Status()
	app.Run()
}
