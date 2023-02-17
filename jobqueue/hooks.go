package jobqueue

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v9"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/pkg/worker"
	"gorm.io/gorm"
	"log"
)

type WriteableField string

const (
	UserId WriteableField = "UserId"
)

var dispatcher worker.Dispatcher
var repo Repository

func New(workerDispatcher worker.Dispatcher, workers ...string) {
	dispatcher = workerDispatcher
	dispatcher.Run(workers...)
}

func GetDispatcher() worker.Dispatcher {
	return dispatcher
}

func GetRepository() Repository {
	return repo
}

func convJobHistory(job *worker.Job, err error) entity.JobHistory {
	jh := entity.JobHistory{}
	var errMsg *string
	if err != nil {
		errMessage := err.Error()
		errMsg = &errMessage
	}

	jh.JobId = job.JobId
	jh.Status = job.Status
	jh.UUID = job.UUID
	jh.WorkerName = job.WorkerName
	jh.Error = errMsg

	return jh
}

func parseMeta(jh entity.JobHistory, m map[string]interface{}) entity.JobHistory {
	if len(m) != 0 {
		for key, val := range m {
			switch key {
			case string(UserId):
				if v, ok := val.(uint); ok {
					jh.UserId = &v
				}
				break
			default:
				break
			}
		}
	}

	return jh
}

func createJobHistory(j *worker.Job) error {
	log.Print("create job history")
	if j == nil {
		return errors.New("job is nil")
	}

	jh := convJobHistory(j, nil)
	jh = parseMeta(jh, j.Meta)

	_, err := repo.Create(jh)

	if err != nil {
		return err
	}

	return nil
}

func updateJobHistory(j *worker.Job, err error) error {
	if j == nil {
		return errors.New("job is nil")
	}

	jh := convJobHistory(j, err)

	_, err = repo.UpdateByUuid(jh.UUID.String(), jh)
	if err != nil {
		return err
	}

	return nil
}

var jobMeta = make(map[string]interface{})

func AddMetaOnDispatch(meta map[WriteableField]interface{}) {
	for key, val := range meta {
		jobMeta[string(key)] = val
	}
}

func RecordHistory(db *gorm.DB) {
	if dispatcher == nil {
		panic("you need call New() method in job_queue package")
	}

	if db == nil {
		db = database.GetDB()
	}

	repo = NewRepository(db)

	dispatcher.OnDispatch(func(j *worker.Job) error {
		j.Meta = jobMeta
		return createJobHistory(j)
	})

	dispatcher.BeforeJob(func(j *worker.Job) error {
		return updateJobHistory(j, nil)
	})

	dispatcher.AfterJob(func(j *worker.Job, err error) error {
		return updateJobHistory(j, err)
	})

}

func FindJob(jobId string, workerName ...string) (*worker.Job, error) {
	name := worker.DefaultWorker
	if len(workerName) != 0 {
		name = workerName[0]
	}

	dispatcher.SelectWorker(name)

	redisClient := dispatcher.GetRedis()()

	var convJob *worker.Job
	value, err := redisClient.Get(context.Background(), jobId).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {

		bytes := []byte(value)
		err = json.Unmarshal(bytes, &convJob)
		if err != nil {
			return nil, err
		}
	}

	return convJob, nil

}
