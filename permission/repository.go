package permission

import (
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	utils.GenericRepository[entity.Permission]
	BatchCreate(permission []entity.Permission) ([]entity.Permission, error)
	GetByGroupId(groupId uint) ([]entity.Permission, error)
}

type RepositoryStruct struct {
	utils.GenericRepository[entity.Permission]
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{
		utils.NewGenericRepository(db, entity.Permission{}),
	}
}

func (r *RepositoryStruct) BatchCreate(permission []entity.Permission) ([]entity.Permission, error) {
	err := r.DB().Transaction(func(tx *gorm.DB) error {
		return tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: "permission"},
				{Name: "group_id"},
			},
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).Create(&permission).Error
	})

	if err != nil {
		return make([]entity.Permission, 0), err
	}

	return permission, nil
}

func (r *RepositoryStruct) GetByGroupId(groupId uint) ([]entity.Permission, error) {
	permissions := make([]entity.Permission, 0)
	err := r.DB().Preload("Actions").Where(entity.Permission{GroupId: groupId}).Find(&permissions).Error

	return permissions, err
}
