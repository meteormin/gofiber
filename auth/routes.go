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

		authMiddleware := Middleware(parameter)

		router.Get("/me", authMiddleware, handler.Me).Name("api.auth.me")

		router.Patch("/password", authMiddleware, handler.ResetPassword).Name("api.auth.password")

		router.Delete("/revoke", authMiddleware, handler.RevokeToken).Name("api.auth.revoke")
	}
}
