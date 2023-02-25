package jobs

import (
	"github.com/google/uuid"
	"github.com/miniyus/gofiber/entity"
	worker "github.com/miniyus/goworker"
	"time"
)

type GetJobs struct {
	Jobs []worker.Job `json:"jobs"`
}

type GetJob struct {
	worker.Job
}

type GetStatus struct {
	worker.StatusInfo
}

type History struct {
	UserId     *uint            `json:"user_id"`
	UUID       uuid.UUID        `json:"uuid"`
	WorkerName string           `json:"worker_name"`
	JobId      string           `json:"job_id"`
	Status     worker.JobStatus `json:"status"`
	Error      *string          `json:"error"`
	CreatedAt  time.Time        `json:"created_at"`
}

func (h History) FromEntity(history entity.JobHistory) History {
	return History{
		UserId:     history.UserId,
		UUID:       history.UUID,
		WorkerName: history.WorkerName,
		JobId:      history.JobId,
		Status:     history.Status,
		Error:      history.Error,
		CreatedAt:  history.CreatedAt,
	}
}

type HistoryQuery struct {
	UserId     *uint             `query:"user_id"`
	UUID       *uuid.UUID        `query:"uuid"`
	WorkerName *string           `query:"worker_name"`
	JobId      *string           `query:"job_id"`
	Status     *worker.JobStatus `query:"status"`
	HasError   bool              `query:"has_error"`
}

func (hq *HistoryQuery) ToEntity() entity.JobHistory {
	ent := entity.JobHistory{}

	if hq.JobId != nil {
		ent.JobId = *hq.JobId
	}

	if hq.UserId != nil {
		ent.UserId = hq.UserId
	}

	if hq.UUID != nil {
		ent.UUID = *hq.UUID
	}

	if hq.WorkerName != nil {
		ent.WorkerName = *hq.WorkerName
	}

	if hq.Status != nil {
		ent.Status = *hq.Status
	}

	return ent
}
