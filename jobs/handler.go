package jobs

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Status(ctx *fiber.Ctx) error
	GetJobs(ctx *fiber.Ctx) error
	GetJob(ctx *fiber.Ctx) error
	AllHistories(ctx *fiber.Ctx) error
	GetHistories(ctx *fiber.Ctx) error
	GetHistory(ctx *fiber.Ctx) error
}

type HandlerStruct struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &HandlerStruct{
		service: service,
	}
}

// AllHistories
// @Summary get all job histories
// @Description get all job histories
// @Tags Jobs
// @Success 200 {object} []History
// @Failure 403 {object} api_error.ErrorResponse
// @Failure 404 {object} api_error.ErrorResponse
// @Accept json
// @Produce json
// @Router /api/worker/histories [get]
// @Security BearerAuth
func (h *HandlerStruct) AllHistories(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

// GetHistories
// @Summary get job histories
// @Description get job histories
// @Tags Jobs
// @Param history query HistoryQuery true "query"
// @Success 200 {object} []History
// @Failure 403 {object} api_error.ErrorResponse
// @Failure 404 {object} api_error.ErrorResponse
// @Accept json
// @Produce json
// @Router /api/worker/:worker/histories [get]
// @Security BearerAuth
func (h *HandlerStruct) GetHistories(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

// GetHistory
// @Summary get job history
// @Description get job history
// @Param worker path string true "worker name"
// @Param id path int true "job history pk"
// @Tags Jobs
// @Success 200 {object} History
// @Failure 403 {object} api_error.ErrorResponse
// @Failure 404 {object} api_error.ErrorResponse
// @Accept json
// @Produce json
// @Router /api/worker/:worker/histories/:id [get]
// @Security BearerAuth
func (h *HandlerStruct) GetHistory(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

// Status
// @Summary jobs status
// @Description jobs status
// @Tags Jobs
// @Success 200 {object} GetStatus
// @Failure 403 {object} api_error.ErrorResponse
// @Failure 404 {object} api_error.ErrorResponse
// @Accept json
// @Produce json
// @Router /api/worker/status [get]
// @Security BearerAuth
func (h *HandlerStruct) Status(ctx *fiber.Ctx) error {
	return ctx.JSON(GetStatus{*h.service.Status()})
}

// GetJobs
// @Summary get jobs
// @Description get jobs
// @Tags Jobs
// @Param worker path string true "worker name"
// @Success 200 {object} GetJobs
// @Failure 403 {object} api_error.ErrorResponse
// @Failure 404 {object} api_error.ErrorResponse
// @Accept json
// @Produce json
// @Router /api/worker/{worker}/jobs [get]
// @Security BearerAuth
func (h *HandlerStruct) GetJobs(ctx *fiber.Ctx) error {
	params := ctx.AllParams()

	workerName := params["worker"]

	jobs, err := h.service.GetJobs(workerName)
	if err != nil {
		return err
	}

	return ctx.JSON(GetJobs{jobs})
}

// GetJob
// @Summary get job
// @Description get job
// @Tags Jobs
// @Param worker path string true "worker name"
// @Param job path string true "job id"
// @Success 200 {object} GetJob
// @Failure 403 {object} api_error.ErrorResponse
// @Failure 404 {object} api_error.ErrorResponse
// @Accept json
// @Produce json
// @Router /api/worker/{worker}/jobs/{job} [get]
// @Security BearerAuth
func (h *HandlerStruct) GetJob(ctx *fiber.Ctx) error {
	params := ctx.AllParams()

	workerName := params["worker"]
	jobId := params["job"]

	job, err := h.service.GetJob(workerName, jobId)
	if err != nil {
		return err
	}

	if job == nil {
		return ctx.JSON(nil)
	}

	return ctx.JSON(GetJob{*job})
}
