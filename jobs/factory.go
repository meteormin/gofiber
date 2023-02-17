package jobs

import (
	"github.com/go-redis/redis/v9"
	"github.com/miniyus/gofiber/jobqueue"
	"github.com/miniyus/gofiber/pkg/worker"
)

func New(redis func() *redis.Client, dispatcher worker.Dispatcher, jobQueueRepository jobqueue.Repository) Handler {
	s := NewService(redis, dispatcher, jobQueueRepository)

	return NewHandler(s)
}
