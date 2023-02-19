package entity

import (
	"github.com/miniyus/gofiber/pkg/gormhooks"
	"gorm.io/gorm"
	"time"
)

type UserRole string

const (
	Admin UserRole = "admin"
)

type User struct {
	gorm.Model
	Username        string     `gorm:"column:username;type:varchar(50);uniqueIndex" json:"username"`
	Email           string     `gorm:"column:email;type:varchar(100);uniqueIndex" json:"email"`
	Password        string     `gorm:"column:password;type:varchar(255)" json:"-"`
	GroupId         *uint      `gorm:"column:group_id;type:bigint" json:"group_id"`
	Role            UserRole   `gorm:"column:role;type:varchar(10)" json:"role"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at" json:"email_verified_at"`
	Group           Group
}

func (u *User) Hooks() *gormhooks.Hooks[*User] {
	return gormhooks.GetHooks(u)
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	return u.Hooks().BeforeSave(tx)
}

func (u *User) AfterSave(tx *gorm.DB) (err error) {
	return u.Hooks().AfterSave(tx)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.Hooks().BeforeCreate(tx)
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	return u.Hooks().AfterCreate(tx)
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.Hooks().BeforeUpdate(tx)
}

func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
	return u.Hooks().AfterUpdate(tx)
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	return u.Hooks().BeforeDelete(tx)
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	return u.Hooks().AfterDelete(tx)
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	return u.Hooks().AfterFind(tx)
}

func (u *User) Before(tx *gorm.DB) (err error) {
	return u.Hooks().Before(tx)
}
