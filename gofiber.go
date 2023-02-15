package gofiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/miniyus/gofiber/api_error"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/create_admin"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/job_queue"
	cLog "github.com/miniyus/gofiber/log"
	"github.com/miniyus/gofiber/permission"
	"github.com/miniyus/gofiber/pkg/validation"
	"github.com/miniyus/gofiber/pkg/worker"
	"github.com/miniyus/gofiber/routes"
	"github.com/miniyus/gofiber/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// New gofiber application
func New(configs ...config.Configs) app.Application {
	var cfg config.Configs

	if len(configs) == 0 {
		cfg = config.GetConfigs()
	} else {
		cfg = configs[0]
	}

	appConfig := cfg.App
	appConfig.FiberConfig.ErrorHandler = api_error.OverrideDefaultErrorHandler(appConfig.Env)

	a := app.New(appConfig)

	// bindings in Container
	a.Register(bind(&cfg))
	// register middlewares
	a.Middleware(middleware)
	// register boot
	a.Register(boot)

	a.Route(routes.ApiPrefix, routes.Api, "api")
	a.Route("/", routes.External, "external")

	return a
}

// bind
// container에 객체 주입
func bind(configs *config.Configs) app.Register {
	return func(a app.Application) {
		cfg := configs
		a.Singleton(func() *config.Configs {
			return cfg
		})

		dbConfig := configs.Database["default"]

		if fiber.IsChild() {
			dbConfig.AutoMigrate = nil
		}

		db := database.New(dbConfig)
		a.Singleton(func() *gorm.DB {
			return db
		})

		var zLogger *zap.SugaredLogger
		// like singleton
		a.Bind(&zLogger, func() *zap.SugaredLogger {
			return cLog.New(configs.CustomLogger["default"])
		})

		opts := configs.JobQueueConfig
		opts.Redis = utils.RedisClientMaker(configs.RedisConfig)

		opts.WorkerOptions = utils.NewCollection(opts.WorkerOptions).Map(func(v worker.Option, i int) worker.Option {
			wLoggerCfg, ok := configs.CustomLogger["default_worker"]
			if ok {
				v.Logger = cLog.New(wLoggerCfg)
			} else {
				a.Resolve(&zLogger)
				v.Logger = zLogger
			}

			return v
		}).Items()

		dispatcher := worker.NewDispatcher(opts)

		var jDispatcher worker.Dispatcher
		// Interface Singleton
		a.Bind(&jDispatcher, func() worker.Dispatcher {
			return dispatcher
		})

	}
}

// middleware
// 미들웨어 등록
func middleware(fiberApp *fiber.App, application app.Application) {
	var cfg *config.Configs

	application.Resolve(&cfg)
	fiberApp.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals(utils.StartTime, time.Now())
		return ctx.Next()
	})

	fiberApp.Use(flogger.New(cfg.Logger))
	fiberApp.Use(recover.New(recover.Config{
		EnableStackTrace: !application.IsProduction(),
	}))
	fiberApp.Use(cors.New(cfg.Cors))
	fiberApp.Use(api_error.ErrorHandler(cfg.App.Env))

}

// boot
// 등록 과정이 끝난 후 실행되는 로직이나 사전 작업
func boot(a app.Application) {
	var dispatcher worker.Dispatcher
	a.Resolve(&dispatcher)

	var db *gorm.DB
	a.Resolve(&db)

	var cfg *config.Configs
	a.Resolve(&cfg)

	var zLogger *zap.SugaredLogger
	a.Resolve(&zLogger)

	create_admin.CreateAdmin(db, cfg)

	job_queue.New(dispatcher)
	job_queue.RecordHistory(db)

	validation.RegisterValidation(cfg.Validation.Validations)
	validation.RegisterTranslation(cfg.Validation.Translations)

	permission.CreateDefaultPermissions(db, cfg.Permission)
}
