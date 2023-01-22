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
	zapLogger := resolver.MakeLogger(app.Config().CustomLogger)
	tokenGenerator := resolver.MakeJwtGenerator(resolver.JwtGeneratorConfig{
		DataPath: app.Config().Path.DataPath,
		Exp:      app.Config().Auth.Exp,
	})

	authMiddlewareParam := auth.MiddlewaresParameter{
		Cfg: app.Config().Auth.Jwt,
		DB:  app.DB(),
	}

	HasPermParam := permission.HasPermissionParameter{
		DB:           app.DB(),
		DefaultPerms: resolver.MakePermissionCollection(app.Config().Permission)(),
	}

	apiRouter.Route(
		api_auth.Prefix,
		api_auth.Register(
			api_auth.New(
				app.DB(),
				tokenGenerator(),
				zapLogger(),
			),
			authMiddlewareParam,
		),
	).Name("api.auth")

	apiRouter.Route(
		groups.Prefix,
		groups.Register(groups.New(app.DB(), zapLogger())),
		auth.Middlewares(authMiddlewareParam, permission.HasPermission(HasPermParam))...,
	).Name("api.groups")

	apiRouter.Route(
		users.Prefix,
		users.Register(users.New(app.DB(), zapLogger())),
		auth.Middlewares(authMiddlewareParam)...,
	).Name("api.users")

}
