package config

import "github.com/miniyus/gofiber/pkg/worker"

type WorkerName string

const (
	DefaultWorker WorkerName = "default"
)

func queueConfig() worker.DispatcherOption {
	return worker.DispatcherOption{
		WorkerOptions: []worker.Option{
			{
				Name:        string(DefaultWorker),
				MaxJobCount: 12,
			},
		},
	}
}