package main

import (
	"github.com/miniyus/gofiber"
)

func main() {
	app := gofiber.New()
	app.Status()
	app.Run()
}
