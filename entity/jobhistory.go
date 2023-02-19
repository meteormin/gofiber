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

func (jh *JobHistory) Hooks() *gormhooks.Hooks[*JobHistory] {
	return gormhooks.GetHooks(jh)
}

func (jh *JobHistory) BeforeSave(tx *gorm.DB) (err error) {
	return jh.Hooks().BeforeSave(tx)
}

func (jh *JobHistory) AfterSave(tx *gorm.DB) (err error) {
	return jh.Hooks().AfterSave(tx)
}

func (jh *JobHistory) BeforeCreate(tx *gorm.DB) (err error) {
	return jh.Hooks().BeforeCreate(tx)
}

func (jh *JobHistory) AfterCreate(tx *gorm.DB) (err error) {
	return jh.Hooks().AfterCreate(tx)
}

func (jh *JobHistory) BeforeUpdate(tx *gorm.DB) (err error) {
	return jh.Hooks().BeforeUpdate(tx)
}

func (jh *JobHistory) AfterUpdate(tx *gorm.DB) (err error) {
	return jh.Hooks().AfterUpdate(tx)
}

func (jh *JobHistory) BeforeDelete(tx *gorm.DB) (err error) {
	return jh.Hooks().BeforeDelete(tx)
}

func (jh *JobHistory) AfterDelete(tx *gorm.DB) (err error) {
	return jh.Hooks().AfterDelete(tx)
}

func (jh *JobHistory) AfterFind(tx *gorm.DB) (err error) {
	return jh.Hooks().AfterFind(tx)
}

func (jh *JobHistory) Before(tx *gorm.DB) (err error) {
	return jh.Hooks().Before(tx)
}
