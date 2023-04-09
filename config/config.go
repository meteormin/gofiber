package config

import (
	fCors "github.com/gofiber/fiber/v2/middleware/cors"
	fCsrf "github.com/gofiber/fiber/v2/middleware/csrf"
	fLoggerMiddleware "github.com/gofiber/fiber/v2/middleware/logger"
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
	CustomLogger   map[string]cLog.Config
	Database       map[string]database.Config
	Path           Path
	Auth           Auth
	Cors           fCors.Config
	Csrf           fCsrf.Config
	RedisConfig    *redis.Options
	JobQueueConfig worker.DispatcherOption
	Validation     Validation
}

var cfg *Configs

func init() {
	cfg = &Configs{
		App:            appConfig(),
		Logger:         flogger(),
		CustomLogger:   loggerConfig(),
		Database:       databaseConfig(),
		Path:           getPath(),
		Auth:           auth(),
		Cors:           cors(),
		Csrf:           csrf(),
		RedisConfig:    redisConfig(),
		JobQueueConfig: jobQueueConfig(),
		Validation:     validationConfig(),
	}
}

func GetConfigs() Configs {
	return *cfg
}
