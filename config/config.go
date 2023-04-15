package config

import (
	fLoggerMiddleware "github.com/gofiber/fiber/v2/middleware/logger"
	fRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/miniyus/gofiber/app"
	_ "github.com/miniyus/gofiber/config/dotenv"
	"github.com/miniyus/gofiber/database"
	cLog "github.com/miniyus/gofiber/log"
	worker "github.com/miniyus/goworker"
	"github.com/redis/go-redis/v9"
)

type Configs struct {
	App            app.Config
	Logger         fLoggerMiddleware.Config
	Recover        fRecover.Config
	CustomLogger   map[string]cLog.Config
	Database       map[string]database.Config
	Path           Path
	Auth           Auth
	RedisConfig    *redis.Options
	JobQueueConfig worker.DispatcherOption
	Validation     Validation
}

var cfg *Configs

func init() {
	cfg = &Configs{
		App:            appConfig(),
		Logger:         flogger(),
		Recover:        recoverConfig(),
		CustomLogger:   loggerConfig(),
		Database:       databaseConfig(),
		Path:           getPath(),
		Auth:           auth(),
		RedisConfig:    redisConfig(),
		JobQueueConfig: jobQueueConfig(),
		Validation:     validationConfig(),
	}
}

func GetConfigs() Configs {
	return *cfg
}
