package main

import (
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/database/migrations"
)

func main() {
	configure := config.GetConfigs()

	db := database.New(configure.Database["default"])

	migrations.Migrate(db)
}
