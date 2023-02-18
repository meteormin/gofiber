package entity

import (
	"github.com/miniyus/gofiber/pkg/gormhooks"
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model
	Method       string `json:"method" gorm:"column:method;type:varchar(10);index:actions_unique,unique'"`
	Resource     string `json:"resource" gorm:"column:resource;type:varchar(50);index:actions_unique,unique"`
	PermissionId uint   `json:"permission_id" gorm:"column:permission_id;type:bigint;index:actions_unique,unique"`
}

func (a *Action) BeforeSave(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).BeforeSave(tx)
}

func (a *Action) AfterSave(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).AfterSave(tx)
}

func (a *Action) BeforeCreate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).BeforeCreate(tx)
}

func (a *Action) AfterCreate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).AfterCreate(tx)
}

func (a *Action) BeforeUpdate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).BeforeUpdate(tx)
}

func (a *Action) AfterUpdate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).AfterUpdate(tx)
}

func (a *Action) BeforeDelete(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).BeforeDelete(tx)
}

func (a *Action) AfterDelete(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).AfterDelete(tx)
}

func (a *Action) AfterFind(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).AfterFind(tx)
}

func (a *Action) Before(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(a).Before(tx)
}

type Permission struct {
	gorm.Model
	Group      *Group   `json:"group" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GroupId    uint     `json:"group_id" gorm:"column:group_id;index:perm_unique,unique"`
	Permission string   `json:"permission" gorm:"column:permission;type:varchar(10);index:perm_unique,unique"`
	Actions    []Action `json:"actions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Permission) BeforeSave(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).BeforeSave(tx)
}

func (p *Permission) AfterSave(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).AfterSave(tx)
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).BeforeCreate(tx)
}

func (p *Permission) AfterCreate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).AfterCreate(tx)
}

func (p *Permission) BeforeUpdate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).BeforeUpdate(tx)
}

func (p *Permission) AfterUpdate(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).AfterUpdate(tx)
}

func (p *Permission) BeforeDelete(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).BeforeDelete(tx)
}

func (p *Permission) AfterDelete(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).AfterDelete(tx)
}

func (p *Permission) AfterFind(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).AfterFind(tx)
}

func (p *Permission) Before(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(p).Before(tx)
}
