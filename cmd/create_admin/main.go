package main

import (
	"github.com/miniyus/gofiber"
	"github.com/miniyus/gofiber/internal/create_admin"
)

func main() {
	create_admin.CreateAdmin(gofiber.New())
}
