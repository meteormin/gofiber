package config

import (
	"github.com/go-redis/redis/v9"
	fCors "github.com/gofiber/fiber/v2/middleware/cors"
	fCsrf "github.com/gofiber/fiber/v2/middleware/csrf"
	fLoggerMiddleware "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/database"
	cLog "github.com/miniyus/gofiber/log"
	"github.com/miniyus/gofiber/permission"
	"github.com/miniyus/gofiber/pkg/worker"
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
	Permission     []permission.Config
	CreateAdmin    CreateAdminConfig
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
		Permission:     permissionConfig(),
		CreateAdmin:    createAdminConfig(),
		RedisConfig:    redisConfig(),
		JobQueueConfig: jobQueueConfig(),
		Validation:     validationConfig(),
	}
}

func GetConfigs() Configs {
	return *cfg
}
