package resolver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/internal/core/container"
	"github.com/miniyus/gofiber/internal/core/context"
)

func Resolve[T interface{}](ctx *fiber.Ctx, dest *T) (T, error) {
	wrapper, ok := ctx.Locals(context.Container).(container.Container)
	if !ok {
		statusCode := fiber.StatusInternalServerError
		return *dest, fiber.NewError(statusCode, "Failed Get Container in Ctx")
	}

	result := wrapper.Resolve(dest)
	if result == nil {
		statusCode := fiber.StatusInternalServerError
		return *dest, fiber.NewError(statusCode, "Failed Resolve...")
	}

	return result.(T), nil
}
