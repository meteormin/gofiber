package main

import (
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/create_admin"
)

func main() {
	create_admin.CreateAdmin(app.New())
}
