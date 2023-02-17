package jobs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/auth"
	"github.com/miniyus/gofiber/jobqueue"
)

func AddJobMeta() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		meta := make(map[jobqueue.WriteableField]interface{})
		user, err := auth.GetAuthUser(ctx)
		if err != nil {
			return err
		}

		meta[jobqueue.UserId] = user.Id

		jobqueue.AddMetaOnDispatch(meta)

		return ctx.Next()
	}
}
