package resolver

import (
	goContext "context"
	"github.com/go-redis/redis/v9"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/internal/logger"
	"github.com/miniyus/gofiber/internal/permission"
	"github.com/miniyus/gofiber/pkg/jwt"
	rsGen "github.com/miniyus/gofiber/pkg/rs256"
	"github.com/miniyus/gofiber/pkg/worker"
	"go.uber.org/zap"
	"log"
	"path"
)

func MakeJwtGenerator(cfg *config.Configs) func() jwt.Generator {
	dataPath := cfg.Path.DataPath

	privateKey := rsGen.PrivatePemDecode(path.Join(dataPath, "secret/private.pem"))

	return func() jwt.Generator {
		return &jwt.GeneratorStruct{
			PrivateKey: privateKey,
			PublicKey:  privateKey.Public(),
			Exp:        cfg.Auth.Exp,
		}
	}
}

func MakeLogger(cfg *config.Configs) func() *zap.SugaredLogger {
	loggerConfig := cfg.CustomLogger
	return func() *zap.SugaredLogger {
		return logger.New(parseLoggerConfig(loggerConfig))
	}
}

func parseLoggerConfig(loggerConfig config.LoggerConfig) logger.Config {
	return logger.Config{
		TimeFormat: loggerConfig.TimeFormat,
		FilePath:   loggerConfig.FilePath,
		Filename:   loggerConfig.Filename,
		MaxAge:     loggerConfig.MaxAge,
		MaxBackups: loggerConfig.MaxBackups,
		MaxSize:    loggerConfig.MaxSize,
		Compress:   loggerConfig.Compress,
		TimeKey:    loggerConfig.TimeKey,
		TimeZone:   loggerConfig.TimeZone,
		LogLevel:   loggerConfig.LogLevel,
	}
}

func MakePermissionCollection(cfg *config.Configs) func() permission.Collection {
	permCfg := permission.NewPermissionsFromConfig(parsePermissionConfig(cfg.Permission))

	return func() permission.Collection {
		return permission.NewPermissionCollection(permCfg...)
	}
}

func parsePermissionConfig(permissionConfig []config.PermissionConfig) []permission.Config {
	var permCfg []permission.Config
	for _, cfg := range permissionConfig {
		permCfg = append(permCfg, permission.Config{
			Name:      cfg.Name,
			GroupId:   cfg.GroupId,
			Methods:   parseMethodConstants(cfg.Methods),
			Resources: cfg.Resources,
		})
	}

	return permCfg
}

func parseMethodConstants(methods []config.PermissionMethod) []permission.Method {
	var authMethods []permission.Method
	for _, method := range methods {
		authMethods = append(authMethods, permission.Method(method))
	}

	return authMethods
}

func MakeJobDispatcher(cfg *config.Configs) func() worker.Dispatcher {
	opts := cfg.QueueConfig

	opts.Redis = MakeRedisClient(cfg)

	return func() worker.Dispatcher {
		return worker.NewDispatcher(opts)
	}
}

func MakeRedisClient(cfg *config.Configs) func() *redis.Client {
	return func() *redis.Client {
		client := redis.NewClient(cfg.RedisConfig)
		pong, err := client.Ping(goContext.Background()).Result()
		log.Print(pong, err)
		return client
	}
}
