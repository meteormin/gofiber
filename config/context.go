package config

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ContextKey string

// context constants
// ctx.Locals() 메서드에서 주로 사용됨
const (
	ContainerKey    ContextKey = "container"
	DBKey           ContextKey = "db"
	ConfigsKey      ContextKey = "config"
	LoggerKey       ContextKey = "logger"
	AuthUserKey     ContextKey = "authUser"
	JwtGeneratorKey ContextKey = "jwtGenerator"
	PermissionsKey  ContextKey = "permissions"
	RedisKey        ContextKey = "redis"
)

func AddContext(localsKey ContextKey, value interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals(localsKey, value)

		return ctx.Next()
	}
}

func GetContext[T interface{}](ctx *fiber.Ctx, localsKey ContextKey) (T, error) {
	getCtx, ok := ctx.Locals(localsKey).(T)
	if !ok {
		return getCtx, errors.New(fmt.Sprintf("Can not get context: %s", localsKey))
	}

	return getCtx, nil
}
