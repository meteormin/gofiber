package jobs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/app"
)

const Prefix = "worker"

func Register(handler Handler) app.SubRouter {
	return func(router fiber.Router) {
		router.Get("/status", handler.Status)
		router.Get("/:worker/jobs", handler.GetJobs)
		router.Get("/:worker/jobs/:job", handler.GetJob)
		router.Get("/:worker/histories", handler.GetHistories)
		router.Get("/:worker/histories/:id", handler.GetHistory)
		router.Get("/histories", handler.AllHistories)
	}
}
