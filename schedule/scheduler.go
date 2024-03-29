package schedule

import (
	"encoding/json"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
	"log"
	"time"
)

var logTag = "schedulerWorker"
var worker *Worker

type Worker struct {
	tagsUnique bool
	jobs       map[string]*gocron.Job
	scheduler  *gocron.Scheduler
	startAt    time.Time
	quitChan   chan bool
	logger     *zap.SugaredLogger
}

type WorkerConfig struct {
	TagsUnique bool
	Location   *time.Location
	Logger     *zap.SugaredLogger
}

type JobStats struct {
	Name            string    `json:"name"`
	Error           error     `json:"error"`
	IsRunning       bool      `json:"is_running"`
	LastRun         time.Time `json:"last_run"`
	RunCount        int       `json:"run_count"`
	ScheduleAtTimes []string  `json:"schedule_at_times"`
	ScheduledAtTime string    `json:"scheduled_at_time"`
	Tags            []string  `json:"tags"`
}

func (js JobStats) Marshal() (string, error) {
	marshal, err := json.Marshal(js)
	if err != nil {
		return "null", err
	}

	return string(marshal), nil
}

func GetWorker() *Worker {
	return worker
}

func NewWorker(cfg WorkerConfig) *Worker {
	worker = &Worker{
		tagsUnique: cfg.TagsUnique,
		jobs:       make(map[string]*gocron.Job, 0),
		scheduler:  gocron.NewScheduler(cfg.Location),
		quitChan:   make(chan bool),
		logger:     cfg.Logger,
	}

	return worker
}

func (w *Worker) Jobs() map[string]*gocron.Job {
	return w.jobs
}

func (w *Worker) AddJob(name string, cronExp string, fn func(), tags ...string) (*gocron.Job, error) {
	var do *gocron.Job
	var err error

	if len(tags) == 0 {
		do, err = w.scheduler.Cron(cronExp).Do(fn)
	} else {
		do, err = w.scheduler.Tag(tags...).Cron(cronExp).Do(fn)
	}

	if err != nil {
		return nil, err
	}

	w.jobs[name] = do

	return do, err
}

func (w *Worker) RemoveJob(name string) {
	job := w.jobs[name]
	w.scheduler.Remove(job)
	delete(w.jobs, name)
}

func (w *Worker) Stats() []JobStats {
	stats := make([]JobStats, 0)
	for name, job := range w.Jobs() {
		stats = append(stats, JobStats{
			Name:            name,
			IsRunning:       job.IsRunning(),
			Error:           job.Error(),
			LastRun:         job.LastRun(),
			RunCount:        job.RunCount(),
			ScheduleAtTimes: job.ScheduledAtTimes(),
			ScheduledAtTime: job.ScheduledAtTime(),
			Tags:            job.Tags(),
		})
	}

	return stats
}

func (w *Worker) statsJob() {
	stats := w.Stats()
	runningTime := time.Now().Sub(w.startAt)

	log.Printf("[schedulerWorker] running time: %f sec", runningTime.Seconds())
	if w.logger != nil {
		w.logger.Infof("[schedulerWorker] running time: %f sec", runningTime.Seconds())
	}

	for _, s := range stats {
		marshal, err := s.Marshal()
		if err != nil {
			log.Printf("[%s] error: %s", logTag, err)
			if w.logger != nil {
				w.logger.Infof("[%s] error: %s", logTag, err)
			}
			continue
		}
		log.Printf("[%s] %s", logTag, marshal)
		if w.logger != nil {
			w.logger.Infof("[%s] %s", logTag, marshal)
		}
	}
}

func (w *Worker) Run() {
	do, err := w.scheduler.Cron("* * * * *").Do(w.statsJob)
	if err != nil {
		w.logger.Error(err)
		panic(err)
	}

	w.jobs["stats"] = do

	if w.tagsUnique {
		w.scheduler.TagsUnique()
	}

	w.scheduler.StartAsync()

	w.startAt = time.Now()
	log.Printf("[%s] Start... %s", logTag, w.startAt)
	if w.logger != nil {
		w.logger.Infof("[%s] Start... %s", logTag, w.startAt)
	}

	for {
		select {
		case <-w.quitChan:
			log.Printf("[%s] Stop", logTag)
			if w.logger != nil {
				w.logger.Infof("[%s] Stop", logTag)
			}
			return
		default:
		}
		time.Sleep(time.Second)
	}
}

func (w *Worker) Stop() {
	w.scheduler.Stop()
	w.quitChan <- true
}
