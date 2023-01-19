package routes

import (
	"github.com/miniyus/gofiber/internal/api/api_auth"
	"github.com/miniyus/gofiber/internal/api/groups"
	"github.com/miniyus/gofiber/internal/api/users"
	"github.com/miniyus/gofiber/internal/core/auth"
	"github.com/miniyus/gofiber/internal/core/container"
	"github.com/miniyus/gofiber/internal/core/permission"
	"github.com/miniyus/gofiber/internal/core/router"
	"github.com/miniyus/gofiber/pkg/jwt"
	"github.com/miniyus/gofiber/pkg/worker"
	"go.uber.org/zap"
)

const ApiPrefix = "/api"

func Api(c container.Container) {
	var jobDispatcher worker.Dispatcher
	c.Resolve(&jobDispatcher)

	var zapLogger *zap.SugaredLogger
	c.Resolve(&zapLogger)

	var tokenGenerator jwt.Generator
	c.Resolve(&tokenGenerator)

	apiRouter := router.New(c.App(), ApiPrefix, "api")

	apiRouter.Route(
		api_auth.Prefix,
		api_auth.Register(api_auth.New(
			c.Database(),
			tokenGenerator,
			zapLogger,
		)),
	).Name("api.auth")

	apiRouter.Route(
		groups.Prefix,
		groups.Register(groups.New(c.Database(), zapLogger)),
		auth.Middlewares(permission.HasPermission())...,
	).Name("api.groups")

	apiRouter.Route(
		users.Prefix,
		users.Register(users.New(c.Database(), zapLogger)),
		auth.Middlewares()...,
	).Name("api.users")

}
