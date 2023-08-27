package monitor

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/schedule"
	"github.com/miniyus/goworker"
	"gorm.io/gorm"
	"path"
)

type AnalysisInfo struct {
	app           app.Application
	ContainerInfo []containerInfo
	Config        *config.Configs
	FiberInfo     fiberInfo
	RouterInfo    []routerInfo
	DatabaseInfo  databaseInfo
	JobQueueInfo  []JobQueueInfo
	SchedulerInfo []schedule.JobStats
}

type JobQueueInfo struct {
	WorkerName  string
	IsRunning   bool
	IsPending   bool
	MaxJobCount int
}

func NewAnalysis(a app.Application) *AnalysisInfo {
	ci := newContainerInfo(a)
	fi := newFiberInfo(a.Fiber())
	ri := newRouterInfo(a.Fiber())

	var cfg *config.Configs
	_ = a.Resolve(&cfg)

	var db *gorm.DB
	_ = a.Resolve(&db)

	dbInfo := newDbInfo(cfg.Database, db)

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
		ContainerInfo: ci,
		Config:        cfg,
		FiberInfo:     fi,
		RouterInfo:    ri,
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

type hasLinks struct {
	Links []string `json:"_links"`
}

type ResponseWrapper struct {
	hasLinks
	ApplicationInfo *AnalysisInfo
}

func makeFullUrl(c *fiber.Ctx, endPoint string) string {
	domain := c.Protocol() + "://" + c.Hostname()
	url := path.Join(c.OriginalURL(), endPoint)

	return domain + url
}

func New(application app.Application) app.SubRouter {
	return func(router fiber.Router) {
		var cfg config.Configs
		_ = application.Resolve(&cfg)
		var dispatcher goworker.Dispatcher
		_ = application.Resolve(&dispatcher)

		router.Get("/", func(c *fiber.Ctx) error {
			analysisInfo := NewAnalysis(application)
			return c.Status(fiber.StatusOK).JSON(
				ResponseWrapper{
					ApplicationInfo: analysisInfo,
					hasLinks: hasLinks{
						[]string{
							makeFullUrl(c, "container"),
							makeFullUrl(c, "configs"),
							makeFullUrl(c, "fiber"),
							makeFullUrl(c, "routes"),
							makeFullUrl(c, "databases"),
							makeFullUrl(c, "job-queues"),
							makeFullUrl(c, "schedule"),
						},
					},
				},
			)
		})
		router.Get("/container", func(ctx *fiber.Ctx) error {
			analysisInfo := NewAnalysis(application)
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.ContainerInfo)
		})
		router.Get("/configs", func(ctx *fiber.Ctx) error {
			analysisInfo := NewAnalysis(application)
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.Config)
		})
		router.Get("/fiber", func(ctx *fiber.Ctx) error {
			analysisInfo := NewAnalysis(application)
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.FiberInfo)
		})
		router.Get("/routes", func(ctx *fiber.Ctx) error {
			analysisInfo := NewAnalysis(application)
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.RouterInfo)
		})
		router.Get("/databases", func(ctx *fiber.Ctx) error {
			analysisInfo := NewAnalysis(application)
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.DatabaseInfo)
		})
		router.Get("/job-queues", func(ctx *fiber.Ctx) error {
			analysisInfo := NewAnalysis(application)
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.JobQueueInfo)
		})
		router.Get("/schedule", func(ctx *fiber.Ctx) error {
			analysisInfo := NewAnalysis(application)
			return ctx.Status(fiber.StatusOK).JSON(analysisInfo.SchedulerInfo)
		})
	}
}
