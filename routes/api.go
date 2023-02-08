package routes

import (
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/auth"
	configure "github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/database"
	groups2 "github.com/miniyus/gofiber/groups"
	jobs2 "github.com/miniyus/gofiber/jobs"
	"github.com/miniyus/gofiber/pkg/jwt"
	rsGen "github.com/miniyus/gofiber/pkg/rs256"
	"github.com/miniyus/gofiber/pkg/worker"
	users2 "github.com/miniyus/gofiber/users"
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

	authMiddlewareParam := auth.MiddlewaresParameter{
		Cfg: cfg.Auth.Jwt,
		DB:  db,
	}

	apiRouter.Route(
		auth.Prefix,
		auth.Register(
			auth.New(
				db,
				tokenGenerator,
			),
			authMiddlewareParam,
		),
	).Name("api.auth")

	// 해당 라인 이후로는 auth middleware가 공통으로 적용된다.
	apiRouter.Middleware(auth.Middlewares(authMiddlewareParam)...)
	// job 메타 데이터에 user_id 추가
	apiRouter.Middleware(jobs2.AddJobMeta(jDispatcher, db))

	apiRouter.Route(
		jobs2.Prefix,
		jobs2.Register(
			jobs2.New(
				utils.RedisClientMaker(cfg.RedisConfig),
				jDispatcher,
			),
		),
	).Name("api.jobs")

	apiRouter.Route(
		groups2.Prefix,
		groups2.Register(groups2.New(db)),
	).Name("api.groups")

	apiRouter.Route(
		users2.Prefix,
		users2.Register(users2.New(db)),
	).Name("api.users")

}
