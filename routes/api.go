package routes

import (
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/auth"
	configure "github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/groups"
	"github.com/miniyus/gofiber/jobqueue"
	"github.com/miniyus/gofiber/jobs"
	"github.com/miniyus/gofiber/permission"
	"github.com/miniyus/gofiber/pkg/jwt"
	rsGen "github.com/miniyus/gofiber/pkg/rs256"
	"github.com/miniyus/gofiber/users"
	"github.com/miniyus/gofiber/utils"
	worker "github.com/miniyus/goworker"
	"gorm.io/gorm"
	"path"
)

const ApiPrefix = "/api"

func Api(apiRouter app.Router, a app.Application) {
	var cfg *configure.Configs
	a.Resolve(&cfg)

	if cfg == nil {
		configs := configure.GetConfigs()
		cfg = &configs
	}

	var db *gorm.DB
	a.Resolve(&db)

	if db == nil {
		db = database.GetDB()
	}

	var jDispatcher worker.Dispatcher
	a.Resolve(&jDispatcher)

	privateKey := rsGen.PrivatePemDecode(path.Join(cfg.Path.DataPath, "secret/private.pem"))
	tokenGenerator := jwt.NewGenerator(privateKey, privateKey.Public(), cfg.Auth.Exp)

	authHandler := auth.New(db, users.NewRepository(db), tokenGenerator)
	apiRouter.Route(
		auth.Prefix,
		auth.Register(authHandler, cfg.Auth.Jwt),
	).Name("api.auth")

	jobsHandler := jobs.New(
		utils.RedisClientMaker(cfg.RedisConfig),
		jDispatcher,
		jobqueue.NewRepository(db),
	)
	apiRouter.Route(
		jobs.Prefix,
		jobs.Register(jobsHandler),
		auth.JwtMiddleware(cfg.Auth.Jwt), auth.Middlewares(),
	).Name("api.jobs")

	hasPermission := permission.HasPermission(permission.HasPermissionParameter{
		DefaultPerms: cfg.Permission,
		DB:           db,
	})

	groupsHandler := groups.New(db)
	apiRouter.Route(
		groups.Prefix,
		groups.Register(groupsHandler),
		auth.JwtMiddleware(cfg.Auth.Jwt), auth.Middlewares(), hasPermission(),
	).Name("api.groups")

	usersHandler := users.New(db)
	apiRouter.Route(
		users.Prefix,
		users.Register(usersHandler),
		auth.JwtMiddleware(cfg.Auth.Jwt), auth.Middlewares(), hasPermission(),
	).Name("api.users")

}
