package jobqueue

import (
	"github.com/miniyus/goworker"
	"time"
)

var containerWorker = "container"
var jobContainer *JobContainer

type Container interface {
	Jobs() map[string]func(job *goworker.Job) error
	AddJob(jobId string, fn func(job *goworker.Job) error)
	RemoveJob(jobId string)
	Dispatch(jobId string) error
	SyncDispatch(jobId string) (*goworker.Job, error)
}

type JobContainer struct {
	worker   goworker.Option
	jobs     map[string]func(job *goworker.Job) error
	syncChan chan *goworker.Job
}

type WorkerOption struct {
	MaxJobCount int
	BeforeJob   func(j *goworker.Job) error
	AfterJob    func(j *goworker.Job, err error) error
	Delay       time.Duration
	Logger      goworker.Logger
}

func GetContainer() Container {
	return jobContainer
}

func NewContainer(workerOption WorkerOption) Container {

	jobContainer = &JobContainer{
		worker: goworker.Option{
			Name:        containerWorker,
			MaxJobCount: workerOption.MaxJobCount,
			Delay:       workerOption.Delay,
			Logger:      workerOption.Logger,
		},
		jobs:     make(map[string]func(job *goworker.Job) error),
		syncChan: make(chan *goworker.Job),
	}

	dispatcher.AddWorker(jobContainer.worker)
	dispatcher.Run(jobContainer.worker.Name)

	return jobContainer
}

func (jc *JobContainer) Jobs() map[string]func(job *goworker.Job) error {
	return jc.jobs
}

func (jc *JobContainer) AddJob(jobId string, fn func(job *goworker.Job) error) {
	jc.jobs[jobId] = fn
}

func (jc *JobContainer) RemoveJob(jobId string) {
	delete(jc.jobs, jobId)
}

func (jc *JobContainer) Dispatch(jobId string) error {
	return dispatcher.SelectWorker(jc.worker.Name).Dispatch(jobId, jc.jobs[jobId])
}

func (jc *JobContainer) SyncDispatch(jobId string) (*goworker.Job, error) {
	dispatcher.AfterJob(func(j *goworker.Job, err error) error {
		if jc.worker.AfterJob != nil {
			err = jc.worker.AfterJob(j, err)
			if err != nil {
				return err
			}
		}

		jc.syncChan <- j

		if err != nil {
			return err
		}

		return nil
	}, containerWorker)

	err := jc.Dispatch(jobId)
	if err != nil {
		return nil, err
	}

	return <-jc.syncChan, nil
}
