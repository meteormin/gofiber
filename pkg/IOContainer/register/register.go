package register

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/miniyus/gofiber/api/gofiber"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/internal/api_error"
	"github.com/miniyus/gofiber/internal/permission"
	"github.com/miniyus/gofiber/internal/resolver"
	"github.com/miniyus/gofiber/pkg/IOContainer"
	"github.com/miniyus/gofiber/pkg/jwt"
	"github.com/miniyus/gofiber/pkg/worker"
	"go.uber.org/zap"
)

// boot is High Priority
// container settings
func boot(w IOContainer.Container) {
	app := w.App()
	w.Singleton(app)

	configs := w.Config()
	w.Singleton(configs)

	db := w.Database()
	w.Singleton(db)

	redisClient := resolver.MakeRedisClient(w.Config())
	w.Singleton(redisClient)

	var tg jwt.Generator
	jwtGenerator := resolver.MakeJwtGenerator(w.Config())
	w.Bind(&tg, jwtGenerator)

	var logs *zap.SugaredLogger
	loggerStruct := resolver.MakeLogger(w.Config())
	w.Bind(&logs, loggerStruct)

	var perms permission.Collection
	permissionCollection := resolver.MakePermissionCollection(w.Config())
	w.Bind(&perms, permissionCollection)

	var dispatcher worker.Dispatcher
	jobDispatcherStruct := resolver.MakeJobDispatcher(w.Config())
	w.Bind(&dispatcher, jobDispatcherStruct)
}

// middlewares register middleware
// fiber app middleware settings
func middlewares(w IOContainer.Container) {
	w.App().Use(flogger.New(w.Config().Logger))
	w.App().Use(recover.New(recover.Config{
		EnableStackTrace: !w.IsProduction(),
	}))
	w.App().Use(api_error.ErrorHandler)
	w.App().Use(cors.New(w.Config().Cors))

	// Add Context Container
	w.App().Use(config.AddContext(config.ContainerKey, w))
	// Add Context Config
	w.App().Use(config.AddContext(config.ConfigsKey, w.Config()))
	// Add Context Logger
	var logger *zap.SugaredLogger
	w.Resolve(&logger)
	w.App().Use(config.AddContext(config.LoggerKey, logger))
	// Add Context JwtGenerator
	var jwtGen jwt.Generator
	w.Resolve(&jwtGen)
	w.App().Use(config.AddContext(config.JwtGeneratorKey, jwtGen))
	// Add Context Permissions
	var perms permission.Collection
	w.Resolve(&perms)
	w.App().Use(config.AddContext(config.PermissionsKey, perms))
	// Add Context Redis
	w.App().Use(config.AddContext(config.RedisKey, resolver.MakeRedisClient(w.Config())))
}

// routes register Routes
func routes(w IOContainer.Container) {

}

// Resister
// private 함수들 모아서 순서대로 실행 해주는 public 함수
func Resister(w IOContainer.Container) {
	boot(w)
	middlewares(w)
	routes(w)
}
