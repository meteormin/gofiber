package jobs

import (
	"context"
	"fmt"
	"github.com/miniyus/gofiber/jobqueue"
	worker "github.com/miniyus/goworker"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Service interface {
	GetJobs(workerName string) ([]worker.Job, error)
	GetJob(workerName string, jobId string) (*worker.Job, error)
	Status() *worker.StatusInfo
	AllHistories(query HistoryQuery) ([]History, error)
	GetHistories(workerName string, userId uint, query HistoryQuery) ([]History, error)
	GetHistory(pk uint) (*History, error)
}

type ServiceStruct struct {
	redis      func() *redis.Client
	dispatcher worker.Dispatcher
	repo       jobqueue.Repository
}

func NewService(redis func() *redis.Client, dispatcher worker.Dispatcher, repo jobqueue.Repository) Service {
	return &ServiceStruct{
		redis:      redis,
		dispatcher: dispatcher,
		repo:       repo,
	}
}

func (s *ServiceStruct) AllHistories(query HistoryQuery) ([]History, error) {
	all, err := s.repo.Get(func(db *gorm.DB) (*gorm.DB, error) {
		ent := query.ToEntity()
		tx := db.Where(&ent)
		if query.HasError {
			tx.Where("error is not null")
		}

		return tx, nil
	})

	if err != nil {
		return make([]History, 0), err
	}

	histories := make([]History, 0)
	var history History
	for _, h := range all {
		histories = append(histories, history.FromEntity(h))
	}

	return histories, err
}

func (s *ServiceStruct) GetHistories(workerName string, userId uint, query HistoryQuery) ([]History, error) {
	all, err := s.repo.Get(func(db *gorm.DB) (*gorm.DB, error) {
		query.WorkerName = &workerName
		query.UserId = &userId
		tx := db.Where(query.ToEntity())
		if query.HasError {
			tx.Where("error is not null")
		}
		return tx, nil
	})

	if err != nil {
		return nil, err
	}

	histories := make([]History, 0)
	var history History
	for _, h := range all {
		histories = append(histories, history.FromEntity(h))
	}

	return histories, err
}

func (s *ServiceStruct) GetHistory(pk uint) (*History, error) {
	find, err := s.repo.Find(pk)
	if err != nil {
		return nil, err
	}

	var history History
	h := history.FromEntity(*find)

	return &h, nil
}

func (s *ServiceStruct) GetJobs(workerName string) ([]worker.Job, error) {
	keys := fmt.Sprintf("%s.*", workerName)
	redisClient := s.redis()
	result, err := redisClient.Keys(context.Background(), keys).Result()
	if err != nil {
		return nil, err
	}

	jobs := make([]worker.Job, 0)
	for _, r := range result {
		job := worker.Job{}

		val, err := redisClient.Get(context.Background(), r).Result()

		if err == redis.Nil {
			continue
		}

		if err != nil {
			return nil, err
		}

		err = job.UnMarshal(val)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (s *ServiceStruct) GetJob(workerName string, jobId string) (*worker.Job, error) {
	key := fmt.Sprintf("%s.%s", workerName, jobId)
	redisClient := s.redis()

	val, err := redisClient.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	job := worker.Job{}

	err = job.UnMarshal(val)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (s *ServiceStruct) Status() *worker.StatusInfo {
	return s.dispatcher.Status()
}
