package jobs

import (
	"github.com/miniyus/gofiber/jobqueue"
	worker "github.com/miniyus/goworker"
	"github.com/redis/go-redis/v9"
)

func New(
	redis func() *redis.Client,
	getAuthUserId GetAuthUserId,
	dispatcher worker.Dispatcher,
	jobQueueRepository jobqueue.Repository,
) Handler {
	s := NewService(redis, dispatcher, jobQueueRepository)

	return NewHandler(s, getAuthUserId)
}
