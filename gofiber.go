package gofiber

import (
	"github.com/gofiber/fiber/v2"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/miniyus/gofiber/apierrors"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/jobqueue"
	cLog "github.com/miniyus/gofiber/log"
	"github.com/miniyus/gofiber/pkg/validation"
	"github.com/miniyus/gofiber/schedule"
	"github.com/miniyus/gofiber/utils"
	"github.com/miniyus/gollection"
	"github.com/miniyus/gorm-extension/gormhooks"
	worker "github.com/miniyus/goworker"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
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
	appConfig.FiberConfig.ErrorHandler = apierrors.OverrideDefaultErrorHandler(appConfig.Env)

	a := app.New(appConfig)

	// bindings in Container
	a.Register(bind(&cfg))
	// register middlewares
	a.Middleware(middleware)
	// boot
	a.Boot(boot)

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

		db, err := database.New(dbConfig)
		if err == nil {
			a.Singleton(func() *gorm.DB {
				return db
			})
		} else {
			log.Printf("failed DB connect: %s", err)
		}

		var zLogger *zap.SugaredLogger
		a.Bind(&zLogger, func() *zap.SugaredLogger {
			return cLog.New(configs.CustomLogger["default"])
		})

		opts := configs.JobQueueConfig
		opts.Redis = utils.RedisClientMaker(configs.RedisConfig)

		opts.WorkerOptions = gollection.NewCollection(opts.WorkerOptions).Map(func(v worker.Option, i int) worker.Option {
			wLoggerCfg, ok := configs.CustomLogger["default_worker"]
			if ok {
				v.Logger = cLog.New(wLoggerCfg)
			} else {
				err = a.Resolve(&zLogger)
				if err == nil {
					v.Logger = zLogger
				}
			}

			return v
		}).Items()

		dispatcher := worker.NewDispatcher(opts)

		var jDispatcher worker.Dispatcher
		// Interface Singleton
		a.Bind(&jDispatcher, func() worker.Dispatcher {
			return dispatcher
		})

		schedulerConfig := configs.Scheduler
		schedulerLoggerCfg, ok := configs.CustomLogger["default_scheduler"]
		if ok {
			schedulerConfig.Logger = cLog.New(schedulerLoggerCfg)
		} else {
			err = a.Resolve(&zLogger)
			if err == nil {
				schedulerConfig.Logger = zLogger
			}
		}

		scheduler := schedule.NewWorker(schedulerConfig)
		a.Singleton(func() *schedule.Worker {
			return scheduler
		})
	}
}

// middleware
// 미들웨어 등록
func middleware(fiberApp *fiber.App, application app.Application) {
	var cfg *config.Configs

	err := application.Resolve(&cfg)
	if err != nil {
		log.Fatalf("failed resolve config: %v", err)
	}

	fiberApp.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals(utils.StartTime, time.Now())
		return ctx.Next()
	})

	fiberApp.Use(flogger.New(cfg.Logger))
	cfg.Recover.EnableStackTrace = !application.IsProduction()
	fiberApp.Use(recover.New(cfg.Recover))
	fiberApp.Use(apierrors.ErrorHandler(cfg.App.Env))
}

// boot
// 등록 과정이 끝난 후 실행되는 로직이나 사전 작업
func boot(a app.Application) {
	var dispatcher worker.Dispatcher
	err := a.Resolve(&dispatcher)
	if err != nil {
		log.Printf("failed resolve dispatcher: %v", err)
	}

	var db *gorm.DB
	err = a.Resolve(&db)
	if err != nil {
		log.Printf("failed resolve db: %v", err)
	}

	var cfg *config.Configs
	err = a.Resolve(&cfg)
	if err != nil {
		log.Printf("failed resolve config: %v", err)
	}

	var zLogger *zap.SugaredLogger
	err = a.Resolve(&zLogger)
	if err != nil {
		log.Printf("failed resolve logger: %v", err)
	}

	validation.NewValidator(cfg.Validation.FallbackLocale, cfg.Validation.SupportedLocales...)
	validation.RegisterValidation(cfg.Validation.Validations)
	validation.RegisterTranslation(cfg.Validation.Translations)

	jobqueue.New(dispatcher)
	jobqueue.RecordHistory(db)

	gormhooks.Register(&entity.JobHistory{})
}

func App() app.Application {
	return app.App()
}

func DB(name ...string) *gorm.DB {
	return database.GetDB(name...)
}

func Config() config.Configs {
	return config.GetConfigs()
}

func Queue() jobqueue.Container {
	return jobqueue.GetContainer()
}

func Log(name ...string) *zap.SugaredLogger {
	return cLog.GetLogger(name...)
}

func Schedule() *schedule.Worker {
	return schedule.GetWorker()
}
