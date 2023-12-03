package config

import (
	"github.com/miniyus/gofiber/schedule"
	"time"
)

func schedulerConfig() schedule.WorkerConfig {
	return schedule.WorkerConfig{
		TagsUnique: true,
		Location:   time.Local,
	}
}
