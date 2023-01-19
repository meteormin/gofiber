package register

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	flogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/miniyus/gofiber/api/gofiber"
	"github.com/miniyus/gofiber/internal/core/api_error"
	"github.com/miniyus/gofiber/internal/core/container"
	"github.com/miniyus/gofiber/internal/core/context"
	"github.com/miniyus/gofiber/internal/core/permission"
	"github.com/miniyus/gofiber/internal/core/register/resolver"
	router "github.com/miniyus/gofiber/internal/routes"
	"github.com/miniyus/gofiber/pkg/jwt"
	"github.com/miniyus/gofiber/pkg/worker"
	"go.uber.org/zap"
	"log"
)

// boot is High Priority
// container settings
func boot(w container.Container) {
	app := w.App()
	w.Singleton(app)

	config := w.Config()
	w.Singleton(config)

	db := w.Database()
	w.Singleton(db)

	redisClient := resolver.MakeRedisClient(w)
	w.Singleton(redisClient)

	var tg jwt.Generator
	jwtGenerator := resolver.MakeJwtGenerator(w)
	w.Bind(&tg, jwtGenerator)

	var logs *zap.SugaredLogger
	loggerStruct := resolver.MakeLogger(w)
	w.Bind(&logs, loggerStruct)

	var perms permission.Collection
	permissionCollection := resolver.MakePermissionCollection(w)
	w.Bind(&perms, permissionCollection)

	var dispatcher worker.Dispatcher
	jobDispatcherStruct := resolver.MakeJobDispatcher(w)
	w.Bind(&dispatcher, jobDispatcherStruct)
}

// middlewares register middleware
// fiber app middleware settings
func middlewares(w container.Container) {
	log.Print(w.Instances())
	w.App().Use(flogger.New(w.Config().Logger))
	w.App().Use(recover.New(recover.Config{
		EnableStackTrace: !w.IsProduction(),
	}))
	w.App().Use(api_error.ErrorHandler)
	w.App().Use(cors.New(w.Config().Cors))

	// Add Context Container
	w.App().Use(context.AddContext(context.Container, w))
	// Add Context Config
	w.App().Use(context.AddContext(context.Config, w.Config()))
	// Add Context Logger
	var logger *zap.SugaredLogger
	w.Resolve(&logger)
	w.App().Use(context.AddContext(context.Logger, logger))
	// Add Context Permissions
	var perms permission.Collection
	w.Resolve(&perms)
	w.App().Use(context.AddContext(context.Permissions, perms))
	// Add Context Redis
	w.App().Use(context.AddContext(context.Redis, resolver.MakeRedisClient(w)))
}

// routes register Routes
func routes(w container.Container) {
	router.Api(w)
	router.External(w)

}

// Resister
// private 함수들 모아서 순서대로 실행 해주는 public 함수
func Resister(w container.Container) {
	boot(w)
	middlewares(w)
	routes(w)
}
