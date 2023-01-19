package main

import (
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/internal/core/database"
	"github.com/miniyus/gofiber/internal/core/database/migrations"
)

func main() {
	configure := config.GetConfigs()

	db := database.DB(configure.Database)

	migrations.Migrate(db)
}
