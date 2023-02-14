package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/app"
)

const Prefix = "/auth"

func Register(handler Handler, parameter MiddlewaresParameter) app.SubRouter {
	return func(router fiber.Router) {
		router.Post("/register", handler.SignUp).Name("api.auth.register")
		router.Post("/token", handler.SignIn).Name("api.auth.token")

		router.Get("/me", Middlewares(parameter, handler.Me)...).Name("api.auth.me")

		router.Patch("/password", Middlewares(parameter, handler.Me)...).Name("api.auth.password")

		router.Delete("/revoke", Middlewares(parameter, handler.Me)...).Name("api.auth.revoke")
	}
}
