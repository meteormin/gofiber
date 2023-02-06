package jobs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/auth"
	"github.com/miniyus/gofiber/job_queue"
	"github.com/miniyus/gofiber/pkg/worker"
	"gorm.io/gorm"
)

func AddJobMeta(jDispatcher worker.Dispatcher, db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		meta := make(map[job_queue.WriteableField]interface{})
		user, err := auth.GetAuthUser(ctx)
		if err != nil {
			return err
		}

		meta[job_queue.UserId] = user.Id

		job_queue.AddMetaOnDispatch(jDispatcher, db, meta)

		return ctx.Next()
	}
}
