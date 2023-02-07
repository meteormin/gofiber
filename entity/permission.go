package entity

import (
	"gorm.io/gorm"
)

type Action struct {
	gorm.Model
	Method       string `json:"method" gorm:"column:method;type:varchar(10);index:actions_unique,unique'"`
	Resource     string `json:"resource" gorm:"column:resource;type:varchar(50);index:actions_unique,unique"`
	PermissionId uint   `json:"permission_id" gorm:"column:permission_id;type:bigint;index:actions_unique,unique"`
}

type Permission struct {
	gorm.Model
	Group      *Group   `json:"group" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GroupId    uint     `json:"group_id"`
	Permission string   `json:"permission" gorm:"column:permission;type:varchar(10);uniqueIndex"`
	Actions    []Action `json:"actions" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
