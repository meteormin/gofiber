package entity

import (
	"github.com/google/uuid"
	"github.com/miniyus/gofiber/pkg/gormhooks"
	"github.com/miniyus/gofiber/pkg/worker"
	"gorm.io/gorm"
)

type JobHistory struct {
	gorm.Model
	UserId     *uint            `json:"user_id"`
	UUID       uuid.UUID        `json:"uuid" gorm:"column:uuid;type:varchar(100);uniqueIndex"`
	WorkerName string           `json:"worker_name" gorm:"column:worker_name;type:varchar(50);"`
	JobId      string           `json:"job_id" gorm:"column:job_id;type:varchar(50)"`
	Status     worker.JobStatus `json:"status" gorm:"column:status;type:varchar(10)"`
	Error      *string          `json:"error" gorm:"column:error;type:varchar(255)"`
	User       User             `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (jh *JobHistory) BeforeSave(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).BeforeSave(tx)
}

func (jh *JobHistory) AfterSave(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).AfterSave(tx)
}

func (jh *JobHistory) BeforeCreate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).BeforeCreate(tx)
}

func (jh *JobHistory) AfterCreate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).AfterCreate(tx)
}

func (jh *JobHistory) BeforeUpdate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).BeforeUpdate(tx)
}

func (jh *JobHistory) AfterUpdate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).AfterUpdate(tx)
}

func (jh *JobHistory) BeforeDelete(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).BeforeDelete(tx)
}

func (jh *JobHistory) AfterDelete(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).AfterDelete(tx)
}

func (jh *JobHistory) AfterFind(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).AfterFind(tx)
}

func (jh *JobHistory) Before(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(jh).Before(tx)
}
