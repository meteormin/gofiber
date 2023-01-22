package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber"
)

const Prefix = "/users"

func Register(handler Handler) gofiber.SubRouter {
	return func(router fiber.Router) {
		router.Get("/", handler.All).Name("api.users.all")
		router.Get("/:id", handler.Get).Name("api.users.get")
		router.Post("/", handler.Create).Name("api.users.create")
		router.Put("/:id", handler.Update).Name("api.users.update")
		router.Patch("/:id", handler.Patch).Name("api.users.patch")
		router.Delete("/:id", handler.Delete).Name("api.users.delete")
	}
}
