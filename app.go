package gofiber

import (
	"github.com/gofiber/fiber/v2"
	configure "github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/internal/api_error"
	"github.com/miniyus/gofiber/internal/database"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type Application interface {
	IsProduction() bool
	Middleware(fn func(fiber *fiber.App, app Application))
	Route(prefix string, fn func(router Router, app Application), name ...string)
	Run()
	Fiber() *fiber.App
	Config() *configure.Configs
	DB() *gorm.DB
}

type app struct {
	fiber  *fiber.App
	config *configure.Configs
	db     *gorm.DB
}

func (a *app) Fiber() *fiber.App {
	return a.fiber
}

func (a *app) Config() *configure.Configs {
	return a.config
}

func (a *app) DB() *gorm.DB {
	return a.db
}

func (a *app) Middleware(fn func(fiber *fiber.App, app Application)) {
	fn(a.Fiber(), a)
}

func (a *app) Route(prefix string, fn func(r Router, app Application), name ...string) {
	r := NewRouter(a.Fiber(), prefix, name...)

	fn(r, a)
}

func (a *app) Run() {
	port := a.config.AppPort
	err := a.fiber.Listen(":" + strconv.Itoa(port))

	if err != nil {
		log.Fatalf("error start fiber app: %v", err)
	}
}

func (a *app) IsProduction() bool {
	return a.Config().AppEnv == configure.PRD
}

// New
// fiber app wrapper
func New(configs ...*configure.Configs) Application {
	var config *configure.Configs

	if len(configs) == 0 {
		config = configure.GetConfigs()
	} else {
		config = configs[0]
	}

	config.App.ErrorHandler = api_error.OverrideDefaultErrorHandler

	return &app{
		fiber.New(config.App),
		config,
		database.DB(config.Database),
	}
}
