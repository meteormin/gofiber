package entity

import (
	"github.com/miniyus/gofiber/pkg/gormhooks"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name        string       `json:"name" gorm:"column:name;type:varchar(50);uniqueIndex"`
	Permissions []Permission `json:"permissions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Users       []User       `json:"users" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (g *Group) Hooks() *gormhooks.Hooks[*Group] {
	return gormhooks.GetHooks(g)
}

func (g *Group) BeforeSave(tx *gorm.DB) (err error) {
	return g.Hooks().BeforeSave(tx)
}

func (g *Group) AfterSave(tx *gorm.DB) (err error) {
	return g.Hooks().AfterSave(tx)
}

func (g *Group) BeforeCreate(tx *gorm.DB) (err error) {
	return g.Hooks().BeforeCreate(tx)
}

func (g *Group) AfterCreate(tx *gorm.DB) (err error) {
	return g.Hooks().AfterCreate(tx)
}

func (g *Group) BeforeUpdate(tx *gorm.DB) (err error) {
	return g.Hooks().BeforeUpdate(tx)
}

func (g *Group) AfterUpdate(tx *gorm.DB) (err error) {
	return g.Hooks().AfterUpdate(tx)
}

func (g *Group) BeforeDelete(tx *gorm.DB) (err error) {
	return g.Hooks().BeforeDelete(tx)
}

func (g *Group) AfterDelete(tx *gorm.DB) (err error) {
	return g.Hooks().AfterDelete(tx)
}

func (g *Group) AfterFind(tx *gorm.DB) (err error) {
	return g.Hooks().AfterFind(tx)
}

func (g *Group) Before(tx *gorm.DB) (err error) {
	return g.Hooks().Before(tx)
}
