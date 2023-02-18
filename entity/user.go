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

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	return gormhooks.GetHooks(u).AfterFind(tx)
}
