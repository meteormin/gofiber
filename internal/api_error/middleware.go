package api_error

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	configure "github.com/miniyus/gofiber/config"
	"go.uber.org/zap"
	"runtime/debug"
)

func OverrideDefaultErrorHandler(ctx *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	var config *configure.Configs
	config, err = configure.GetContext[*configure.Configs](ctx, configure.ConfigsKey)

	if err != nil {
		errRes := NewFromError(ctx, err)
		return errRes.Response()
	}

	var logger *zap.SugaredLogger
	logger, err = configure.GetContext[*zap.SugaredLogger](ctx, configure.LoggerKey)

	if err != nil {
		errRes := NewFromError(ctx, err)
		return errRes.Response()
	}

	logger.Errorln(err)
	if config.AppEnv != configure.PRD {
		debug.PrintStack()
	}

	errRes := NewFromError(ctx, err)
	return errRes.Response()
}

func ErrorHandler(ctx *fiber.Ctx) error {
	err := ctx.Next()

	if err == nil {
		return nil
	}

	var config *configure.Configs
	config, err = configure.GetContext[*configure.Configs](ctx, configure.ConfigsKey)

	if err != nil {
		errRes := NewFromError(ctx, err)
		return errRes.Response()
	}

	var logger *zap.SugaredLogger
	logger, err = configure.GetContext[*zap.SugaredLogger](ctx, configure.LoggerKey)

	if err != nil {
		errRes := NewFromError(ctx, err)
		return errRes.Response()
	}

	logger.Errorln(err)
	if config.AppEnv != configure.PRD {
		debug.PrintStack()
	}

	errRes := NewFromError(ctx, err)

	b, _ := json.Marshal(errRes)

	logger.Errorln(string(b))

	return errRes.Response()
}