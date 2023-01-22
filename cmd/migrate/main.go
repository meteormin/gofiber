package main

import (
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/internal/database"
	"github.com/miniyus/gofiber/internal/database/migrations"
)

func main() {
	configure := config.GetConfigs()

	db := database.DB(configure.Database)

	migrations.Migrate(db)
}
