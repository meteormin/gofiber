package permission

import (
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Save(permission entity.Permission) (*entity.Permission, error)
	Get(groupId uint) ([]entity.Permission, error)
}

type RepositoryStruct struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{
		db: db,
	}
}

func (r *RepositoryStruct) Save(permission entity.Permission) (*entity.Permission, error) {
	tx := r.db.Save(&permission)

	tx, err := database.HandleResult(tx)
	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *RepositoryStruct) Get(groupId uint) ([]entity.Permission, error) {
	permissions := make([]entity.Permission, 0)

	rs := r.db.Preload("Actions").Where(entity.Permission{GroupId: groupId}).Find(&permissions)
	_, err := database.HandleResult(rs)

	return permissions, err
}
