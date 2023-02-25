package entity

import (
	"github.com/miniyus/gorm-extension/gormhooks"
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model
	Method       string `json:"method" gorm:"column:method;type:varchar(10);index:actions_unique,unique'"`
	Resource     string `json:"resource" gorm:"column:resource;type:varchar(50);index:actions_unique,unique"`
	PermissionId uint   `json:"permission_id" gorm:"column:permission_id;type:bigint;index:actions_unique,unique"`
}

func (a *Action) Hooks() *gormhooks.Hooks[*Action] {
	return gormhooks.GetHooks(a)
}

func (a *Action) BeforeSave(tx *gorm.DB) (err error) {
	return a.Hooks().BeforeSave(tx)
}

func (a *Action) AfterSave(tx *gorm.DB) (err error) {
	return a.Hooks().AfterSave(tx)
}

func (a *Action) BeforeCreate(tx *gorm.DB) (err error) {
	return a.Hooks().BeforeCreate(tx)
}

func (a *Action) AfterCreate(tx *gorm.DB) (err error) {
	return a.Hooks().AfterCreate(tx)
}

func (a *Action) BeforeUpdate(tx *gorm.DB) (err error) {
	return a.Hooks().BeforeUpdate(tx)
}

func (a *Action) AfterUpdate(tx *gorm.DB) (err error) {
	return a.Hooks().AfterUpdate(tx)
}

func (a *Action) BeforeDelete(tx *gorm.DB) (err error) {
	return a.Hooks().BeforeDelete(tx)
}

func (a *Action) AfterDelete(tx *gorm.DB) (err error) {
	return a.Hooks().AfterDelete(tx)
}

func (a *Action) AfterFind(tx *gorm.DB) (err error) {
	return a.Hooks().AfterFind(tx)
}

func (a *Action) Before(tx *gorm.DB) (err error) {
	return a.Hooks().Before(tx)
}

type Permission struct {
	gorm.Model
	Group      *Group   `json:"group" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GroupId    uint     `json:"group_id" gorm:"column:group_id;index:perm_unique,unique"`
	Permission string   `json:"permission" gorm:"column:permission;type:varchar(10);index:perm_unique,unique"`
	Actions    []Action `json:"actions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Permission) Hooks() *gormhooks.Hooks[*Permission] {
	return gormhooks.GetHooks(p)
}

func (p *Permission) BeforeSave(tx *gorm.DB) (err error) {
	return p.Hooks().BeforeSave(tx)
}

func (p *Permission) AfterSave(tx *gorm.DB) (err error) {
	return p.Hooks().AfterSave(tx)
}

func (p *Permission) BeforeCreate(tx *gorm.DB) (err error) {
	return p.Hooks().BeforeCreate(tx)
}

func (p *Permission) AfterCreate(tx *gorm.DB) (err error) {
	return p.Hooks().AfterCreate(tx)
}

func (p *Permission) BeforeUpdate(tx *gorm.DB) (err error) {
	return p.Hooks().BeforeUpdate(tx)
}

func (p *Permission) AfterUpdate(tx *gorm.DB) (err error) {
	return p.Hooks().AfterUpdate(tx)
}

func (p *Permission) BeforeDelete(tx *gorm.DB) (err error) {
	return p.Hooks().BeforeDelete(tx)
}

func (p *Permission) AfterDelete(tx *gorm.DB) (err error) {
	return p.Hooks().AfterDelete(tx)
}

func (p *Permission) AfterFind(tx *gorm.DB) (err error) {
	return p.Hooks().AfterFind(tx)
}

func (p *Permission) Before(tx *gorm.DB) (err error) {
	return p.Hooks().Before(tx)
}
