package routes

import (
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/auth"
	api_auth2 "github.com/miniyus/gofiber/internal/api_auth"
	groups2 "github.com/miniyus/gofiber/internal/groups"
	users2 "github.com/miniyus/gofiber/internal/users"
	"github.com/miniyus/gofiber/permission"
	"github.com/miniyus/gofiber/resolver"
)

const ApiPrefix = "/api"

func Api(apiRouter app.Router, app app.Application) {
	zapLogger := resolver.MakeLogger(app.Config().CustomLogger)
	tokenGenerator := resolver.MakeJwtGenerator(resolver.JwtGeneratorConfig{
		DataPath: app.Config().Path.DataPath,
		Exp:      app.Config().Auth.Exp,
	})

	authMiddlewareParam := auth.MiddlewaresParameter{
		Cfg: app.Config().Auth.Jwt,
		DB:  app.DB(),
	}

	hasPermParam := permission.HasPermissionParameter{
		DB:           app.DB(),
		DefaultPerms: resolver.MakePermissionCollection(app.Config().Permission)(),
	}

	apiRouter.Route(
		api_auth2.Prefix,
		api_auth2.Register(
			api_auth2.New(
				app.DB(),
				tokenGenerator(),
				zapLogger(),
			),
			authMiddlewareParam,
		),
	).Name("api.auth")

	apiRouter.Route(
		groups2.Prefix,
		groups2.Register(groups2.New(app.DB(), zapLogger())),
		auth.Middlewares(authMiddlewareParam, permission.HasPermission(hasPermParam))...,
	).Name("api.groups")

	apiRouter.Route(
		users2.Prefix,
		users2.Register(users2.New(app.DB(), zapLogger())),
		auth.Middlewares(authMiddlewareParam)...,
	).Name("api.users")

}
