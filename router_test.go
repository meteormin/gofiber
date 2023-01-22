package gofiber_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	r := gofiber.NewRouter(fiber.New(), "/", "test")

	r.Route("/", func(router fiber.Router) {
		router.Get("/", func(ctx *fiber.Ctx) error {
			return ctx.JSON("hi")
		}).Name("tty")
	})

	for _, route := range r.GetRoutes() {
		log.Print(route)
	}
}
