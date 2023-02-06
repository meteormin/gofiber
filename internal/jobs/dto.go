package jobs

import "github.com/miniyus/gofiber/pkg/worker"

type GetJobs struct {
	Jobs []worker.Job `json:"jobs"`
}

type GetJob struct {
	worker.Job
}

type GetStatus struct {
	worker.StatusInfo
}
