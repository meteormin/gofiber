package permission

import (
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
	err := r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Save(&permission).Error
	})

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *RepositoryStruct) Get(groupId uint) ([]entity.Permission, error) {
	permissions := make([]entity.Permission, 0)

	err := r.db.Preload("Actions").Where(entity.Permission{GroupId: groupId}).Find(&permissions).Error

	return permissions, err
}
