package entity

import (
	"github.com/miniyus/gofiber/pkg/gormhooks"
	"gorm.io/gorm"
	"time"
)

type AccessToken struct {
	gorm.Model
	UserId    uint      `gorm:"column:user_id"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Token     string    `gorm:"column:token;type:text;uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"column:expires_at;not null;index"`
}

func (at *AccessToken) BeforeSave(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).BeforeSave(tx)
}

func (at *AccessToken) AfterSave(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).AfterSave(tx)
}

func (at *AccessToken) BeforeCreate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).BeforeCreate(tx)
}

func (at *AccessToken) AfterCreate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).AfterCreate(tx)
}

func (at *AccessToken) BeforeUpdate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).BeforeUpdate(tx)
}

func (at *AccessToken) AfterUpdate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).AfterUpdate(tx)
}

func (at *AccessToken) BeforeDelete(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).BeforeDelete(tx)
}

func (at *AccessToken) AfterDelete(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).AfterDelete(tx)
}

func (at *AccessToken) AfterFind(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).AfterFind(tx)
}

func (at *AccessToken) Before(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(at).Before(tx)
}
