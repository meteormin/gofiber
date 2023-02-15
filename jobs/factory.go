package jobs

import (
	"github.com/go-redis/redis/v9"
	"github.com/miniyus/gofiber/job_queue"
	"github.com/miniyus/gofiber/pkg/worker"
)

func New(redis func() *redis.Client, dispatcher worker.Dispatcher, jobQueueRepository job_queue.Repository) Handler {
	s := NewService(redis, dispatcher, jobQueueRepository)

	return NewHandler(s)
}
