package routes

import (
	"github.com/miniyus/gofiber"
	"github.com/miniyus/gofiber/internal/api/api_auth"
	"github.com/miniyus/gofiber/internal/api/groups"
	"github.com/miniyus/gofiber/internal/api/users"
	"github.com/miniyus/gofiber/internal/auth"
	"github.com/miniyus/gofiber/internal/permission"
	"github.com/miniyus/gofiber/internal/resolver"
)

const ApiPrefix = "/api"

func Api(apiRouter gofiber.Router, app gofiber.Application) {
	zapLogger := resolver.MakeLogger(app.Config())

	tokenGenerator := resolver.MakeJwtGenerator(app.Config())

	apiRouter.Route(
		api_auth.Prefix,
		api_auth.Register(api_auth.New(
			app.DB(),
			tokenGenerator(),
			zapLogger(),
		)),
	).Name("api.auth")

	apiRouter.Route(
		groups.Prefix,
		groups.Register(groups.New(app.DB(), zapLogger())),
		auth.Middlewares(permission.HasPermission())...,
	).Name("api.groups")

	apiRouter.Route(
		users.Prefix,
		users.Register(users.New(app.DB(), zapLogger())),
		auth.Middlewares()...,
	).Name("api.users")

}
