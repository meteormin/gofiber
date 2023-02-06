package routes

import (
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/auth"
	configure "github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/internal/api_auth"
	"github.com/miniyus/gofiber/internal/groups"
	"github.com/miniyus/gofiber/internal/jobs"
	"github.com/miniyus/gofiber/internal/users"
	"github.com/miniyus/gofiber/permission"
	"github.com/miniyus/gofiber/pkg/jwt"
	rsGen "github.com/miniyus/gofiber/pkg/rs256"
	"github.com/miniyus/gofiber/pkg/worker"
	"github.com/miniyus/gofiber/utils"
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

	permissions := permission.NewPermissionsFromConfig(cfg.Permission)
	permissionCollection := permission.NewPermissionCollection(permissions...)

	authMiddlewareParam := auth.MiddlewaresParameter{
		Cfg: cfg.Auth.Jwt,
		DB:  db,
	}

	hasPermParam := permission.HasPermissionParameter{
		DB:           db,
		DefaultPerms: permissionCollection,
		FilterFunc:   nil,
	}

	apiRouter.Route(
		api_auth.Prefix,
		api_auth.Register(
			api_auth.New(
				db,
				tokenGenerator,
			),
			authMiddlewareParam,
		),
	).Name("api.auth")

	// 해당 라인 이후로는 auth middleware가 공통으로 적용된다.
	apiRouter.Middleware(auth.Middlewares(authMiddlewareParam, permission.HasPermission(hasPermParam))...)
	// job 메타 데이터에 user_id 추가
	apiRouter.Middleware(jobs.AddJobMeta(jDispatcher, db))

	apiRouter.Route(
		jobs.Prefix,
		jobs.Register(
			jobs.New(
				utils.RedisClientMaker(cfg.RedisConfig),
				jDispatcher,
			),
		),
	)

	apiRouter.Route(
		groups.Prefix,
		groups.Register(groups.New(db)),
	).Name("api.groups")

	apiRouter.Route(
		users.Prefix,
		users.Register(users.New(db)),
	).Name("api.users")

}
