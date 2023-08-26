package monitor

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/config"
	myReflect "github.com/miniyus/gofiber/internal/reflect"
	"github.com/miniyus/gofiber/schedule"
	"github.com/miniyus/goworker"
	gormLogger "gorm.io/gorm/logger"
	"path"
	"reflect"
	"time"
)

type AnalysisInfo struct {
	app           app.Application
	ContainerInfo []ContainerInfo
	Config        *config.Configs
	FiberInfo     FiberInfo
	RouterInfo    []RouterInfo
	DatabaseInfo  []DatabaseInfo
	JobQueueInfo  []JobQueueInfo
	SchedulerInfo []schedule.JobStats
}

type BindType string

const (
	Bind      BindType = "bindingInterface"
	Singleton BindType = "singleton"
)

type ContainerInfo struct {
	Key          string
	BindType     BindType
	InstanceType string
}

type FiberInfo struct {
	HandlersCount uint32
}

type RouterInfo struct {
	Name   string
	Method string
	Path   string
	Params []string
}

type DatabaseInfo struct {
	Connection        string
	Driver            string
	Host              string
	Dbname            string
	Port              string
	TimeZone          string
	SSlMode           bool
	AutoMigrate       []string
	Logger            gormLogger.Config
	MaxIdleConnection int
	MaxOpenConnection int
	MaxLifeTime       time.Duration
}

type JobQueueInfo struct {
	WorkerName  string
	IsRunning   bool
	IsPending   bool
	MaxJobCount int
}

func NewAnalysis(a app.Application) *AnalysisInfo {
	instances := a.Instances()
	containerInfos := make([]ContainerInfo, 0)

	for _, inst := range instances {
		var bt BindType
		reflectType := reflect.TypeOf(inst)
		if reflectType.Kind() == reflect.Func {
			bt = Bind
		} else {
			bt = Singleton
		}

		instType := myReflect.GetType(inst)
		ci := ContainerInfo{
			Key:          instType,
			BindType:     bt,
			InstanceType: instType,
		}
		containerInfos = append(containerInfos, ci)
	}

	var cfg *config.Configs
	_ = a.Resolve(&cfg)

	fiberInfo := FiberInfo{
		HandlersCount: a.Fiber().HandlersCount(),
	}

	routerInfo := make([]RouterInfo, 0)
	for _, r := range a.Fiber().GetRoutes() {
		routerInfo = append(routerInfo, RouterInfo{
			Name:   r.Name,
			Method: r.Method,
			Path:   r.Path,
			Params: r.Params,
		})
	}

	dbInfo := make([]DatabaseInfo, 0)
	for _, dbCfg := range cfg.Database {
		migrates := make([]string, 0)
		for _, ent := range dbCfg.AutoMigrate {
			migrates = append(migrates, myReflect.GetType(ent))
		}

		dbInfo = append(dbInfo, DatabaseInfo{
			Connection:        dbCfg.Name,
			Driver:            dbCfg.Driver,
			Host:              dbCfg.Host,
			Dbname:            dbCfg.Dbname,
			Port:              dbCfg.Port,
			TimeZone:          dbCfg.TimeZone,
			SSlMode:           dbCfg.SSLMode,
			AutoMigrate:       migrates,
			Logger:            dbCfg.Logger,
			MaxIdleConnection: dbCfg.MaxIdleConn,
			MaxOpenConnection: dbCfg.MaxOpenConn,
			MaxLifeTime:       dbCfg.MaxLifeTime,
		})
	}

	var dispatcher goworker.Dispatcher
	err := a.Resolve(&dispatcher)
	jobQueueInfo := make([]JobQueueInfo, 0)
	if err == nil {
		for _, workerStatus := range dispatcher.Status().Workers {
			jobQueueInfo = append(jobQueueInfo, JobQueueInfo{
				WorkerName:  workerStatus.Name,
				IsPending:   workerStatus.IsPending,
				IsRunning:   workerStatus.IsRunning,
				MaxJobCount: workerStatus.MaxJobCount,
			})
		}
	}

	scheduleInfo := make([]schedule.JobStats, 0)
	scheduleWorker := schedule.GetWorker()
	if scheduleWorker != nil {
		scheduleInfo = scheduleWorker.Stats()
	}

	return &AnalysisInfo{
		app:           a,
		ContainerInfo: containerInfos,
		Config:        cfg,
		FiberInfo:     fiberInfo,
		RouterInfo:    routerInfo,
		DatabaseInfo:  dbInfo,
		JobQueueInfo:  jobQueueInfo,
		SchedulerInfo: scheduleInfo,
	}
}

func (ai *AnalysisInfo) Marshal(indent bool) (string, error) {
	if indent {
		marshal, err := json.MarshalIndent(ai, "", "    ")
		if err != nil {
			return "", err
		}

		return string(marshal), nil
	}

	marshal, err := json.Marshal(ai)
	if err != nil {
		return "", err
	}

	return string(marshal), nil
}

type ResponseWrapper struct {
	ApplicationInfo *AnalysisInfo
	Links           []string `json:"_links"`
}

func makeFullUrl(c *fiber.Ctx, endPoint string) string {
	domain := c.Protocol() + "://" + c.Hostname()
	url := path.Join(c.OriginalURL(), endPoint)

	return domain + url
}

func New(application app.Application) app.SubRouter {
	analysisInfo := NewAnalysis(application)
	return func(router fiber.Router) {
		router.Get("/", func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).JSON(
				ResponseWrapper{
					ApplicationInfo: analysisInfo,
					Links: []string{
						makeFullUrl(c, "container"),
						makeFullUrl(c, "configs"),
						makeFullUrl(c, "fiber"),
						makeFullUrl(c, "routes"),
						makeFullUrl(c, "databases"),
						makeFullUrl(c, "job-queues"),
						makeFullUrl(c, "schedule"),
					},
				},
			)
		})
		router.Get("/container", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.ContainerInfo)
		})
		router.Get("/configs", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.Config)
		})
		router.Get("/fiber", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.FiberInfo)
		})
		router.Get("/routes", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.RouterInfo)
		})
		router.Get("/databases", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.DatabaseInfo)
		})
		router.Get("/job-queues", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.JobQueueInfo)
		})
		router.Get("/schedule", func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.SchedulerInfo)
		})
	}
}
