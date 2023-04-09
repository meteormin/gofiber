package jobqueue_test

import (
	"github.com/joho/godotenv"
	"github.com/miniyus/gofiber/config"
	"github.com/miniyus/gofiber/jobqueue"
	cLog "github.com/miniyus/gofiber/log"
	"github.com/miniyus/gofiber/utils"
	"github.com/miniyus/gollection"
	worker "github.com/miniyus/goworker"
	"log"
	"testing"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	configs := config.GetConfigs()
	opts := configs.JobQueueConfig
	opts.Redis = utils.RedisClientMaker(configs.RedisConfig)

	opts.WorkerOptions = gollection.NewCollection(opts.WorkerOptions).Map(func(v worker.Option, i int) worker.Option {
		wLoggerCfg, ok := configs.CustomLogger["default_worker"]
		if ok {
			v.Logger = cLog.New(wLoggerCfg)
		}
		return v
	}).Items()

	jobqueue.New(worker.NewDispatcher(opts))
}

func TestJobContainer_AddJob(t *testing.T) {
	jobContainer := jobqueue.NewContainer(jobqueue.WorkerOption{
		MaxJobCount: 10,
	})

	jobContainer.AddJob("TEST_1", func(job *worker.Job) error {
		log.Printf("%s - %s", job.JobId, job.Status)
		return nil
	})

	jobContainer.AddJob("TEST_2", func(job *worker.Job) error {
		log.Printf("%s - %s", job.JobId, job.Status)
		return nil
	})
}

func TestJobContainer_Jobs(t *testing.T) {
	jobContainer := jobqueue.GetContainer()

	jobs := jobContainer.Jobs()
	log.Print(jobs)
}

func TestJobContainer_Dispatch(t *testing.T) {
	jobContainer := jobqueue.GetContainer()

	err := jobContainer.Dispatch("TEST_1")
	if err != nil {
		t.Error(err)
	}

	err = jobContainer.Dispatch("TEST_2")
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Second * 3)
}

func TestJobContainer_SyncDispatch(t *testing.T) {
	jobContainer := jobqueue.GetContainer()

	jobContainer.AddJob("TEST_3", func(job *worker.Job) error {
		log.Printf("%s - %s", job.JobId, job.Status)
		time.Sleep(time.Second * 3)
		return nil
	})

	dispatch, err := jobContainer.SyncDispatch("TEST_3")
	if err != nil {
		return
	}

	if dispatch != nil {
		log.Print(dispatch)
	}
}
