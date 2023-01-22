package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/pkg/IOContainer"
)

const IOContainerKey = "container"

func Resolve[T interface{}](ctx *fiber.Ctx, dest *T) (T, error) {
	wrapper, ok := ctx.Locals(IOContainerKey).(IOContainer.Container)
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