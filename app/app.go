package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/pkg/iocontainer"
	"github.com/miniyus/gollection"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type Env string

const (
	PRD   Env = "production"
	DEV   Env = "development"
	LOCAL Env = "local"
)

type Config struct {
	Env         Env
	Port        int
	Locale      string
	TimeZone    string
	FiberConfig fiber.Config
}

var defaultConfig = Config{
	Env:         LOCAL,
	Port:        8000,
	Locale:      "",
	TimeZone:    "Asia/Seoul",
	FiberConfig: fiber.Config{},
}

type Register func(app Application)

type Boot func(app Application)

type RouterGroup func(router Router, app Application)

type MiddlewareRegister func(fiber *fiber.App, app Application)

type Application interface {
	iocontainer.Container
	Fiber() *fiber.App
	IsProduction() bool
	Config() Config
	Middleware(fn MiddlewareRegister)
	Route(prefix string, fn RouterGroup, name ...string)
	GetRouters() gollection.Collection[Router]
	Status()
	Run()
	Register(fn Register)
	Boot(fn Boot)
	Test(req *http.Request, msTimeout ...int) (*http.Response, error)
}

var a Application

func App() Application {
	return a
}

type app struct {
	iocontainer.Container
	fiber      *fiber.App
	config     Config
	routers    gollection.Collection[Router]
	registered gollection.Collection[Register]
	boots      gollection.Collection[Boot]
	isRun      bool
}

// New
// fiber app wrapper
func New(cfgs ...Config) Application {
	var fiberConfig fiber.Config
	var cfg Config

	if len(cfgs) == 0 {
		cfg = defaultConfig
		fiberConfig = cfg.FiberConfig
	} else {
		cfg = cfgs[0]
		fiberConfig = cfg.FiberConfig
	}

	a = &app{
		Container:  iocontainer.NewContainer(),
		config:     cfg,
		fiber:      fiber.New(fiberConfig),
		isRun:      false,
		routers:    gollection.NewCollection(make([]Router, 0)),
		registered: gollection.NewCollection(make([]Register, 0)),
		boots:      gollection.NewCollection(make([]Boot, 0)),
	}

	return a
}

func (a *app) Config() Config {
	return a.config
}

func (a *app) Register(fn Register) {
	a.registered.Add(fn)
}

func (a *app) Boot(fn Boot) {
	a.boots.Add(fn)
}

// Middleware
// add middleware from closure
func (a *app) Middleware(fn MiddlewareRegister) {
	a.Register(func(this Application) {
		if _, ok := this.(*app); ok {
			fn(this.(*app).fiber, this)
		}
	})
}

// Route
// register route group
func (a *app) Route(prefix string, fn RouterGroup, name ...string) {
	a.Register(func(this Application) {
		if _, ok := this.(*app); ok {
			r := NewRouter(this.(*app).fiber, prefix, name...)
			fn(r, a)
			this.(*app).routers.Add(r)
		}
	})
}

func (a *app) GetRouters() gollection.Collection[Router] {
	return a.routers
}

// Status
// Debug 용도,
// 현재 생성된 route list
// 컨테이너가 가지고 있는 정보 콘솔 로그로 보여준다.
func (a *app) Status() {
	if a.IsProduction() {
		log.Printf("'AppEnv' is %s", PRD)
		return
	}

	log.Println("[Container Info]")
	log.Printf("ENV: %s", a.config.Env)
	log.Printf("Locale: %s", a.config.Locale)
	log.Printf("Time Zone: %s", a.config.TimeZone)

	log.Println("[Fiber App Info]")
	log.Printf("Handlers Count: %d", a.fiber.HandlersCount())
	log.Println("[Router]")

	for _, r := range a.fiber.GetRoutes() {
		log.Printf(
			"[%s] '%s' | '%s' , Params: %s",
			r.Method, r.Name, r.Path, r.Params,
		)
	}

}

func (a *app) bootstrap() {
	a.registered.For(func(v Register, i int) {
		v(a)
	})
	a.boots.For(func(v Boot, i int) {
		v(a)
	})
}

// Run
// run fiber application
func (a *app) Run() {
	if a.isRun {
		return
	}

	a.bootstrap()

	port := a.config.Port
	err := a.fiber.Listen(":" + strconv.Itoa(port))

	if err != nil {
		log.Fatalf("error start fiber app: %v", err)
	}

	a.isRun = true
}

func (a *app) Fiber() *fiber.App {
	return a.fiber
}

func (a *app) IsProduction() bool {
	return a.config.Env == PRD
}

func (a *app) Instances() map[reflect.Type]interface{} {
	return a.Container.Instances()
}

func (a *app) Bind(keyType interface{}, resolver interface{}) {
	a.Container.Bind(keyType, resolver)
}

func (a *app) Resolve(resolver interface{}) interface{} {
	return a.Container.Resolve(resolver)
}

func (a *app) Singleton(instance interface{}) {
	a.Container.Singleton(instance)
}

func (a *app) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return a.fiber.Test(req, msTimeout...)
}
