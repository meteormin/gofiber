package jobs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/auth"
	"github.com/miniyus/gofiber/job_queue"
)

func AddJobMeta() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		meta := make(map[job_queue.WriteableField]interface{})
		user, err := auth.GetAuthUser(ctx)
		if err != nil {
			return err
		}

		meta[job_queue.UserId] = user.Id

		job_queue.AddMetaOnDispatch(meta)

		return ctx.Next()
	}
}
