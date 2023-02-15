package permission

import (
	"github.com/miniyus/gofiber/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	All() ([]entity.Permission, error)
	BatchCreate(permission []entity.Permission) ([]entity.Permission, error)
	Create(permission entity.Permission) (*entity.Permission, error)
	GetByGroupId(groupId uint) ([]entity.Permission, error)
	Update(pk uint, permission entity.Permission) (*entity.Permission, error)
	Delete(pk uint) (bool, error)
}

type RepositoryStruct struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{
		db: db,
	}
}

func (r *RepositoryStruct) All() ([]entity.Permission, error) {
	perms := make([]entity.Permission, 0)
	err := r.db.Find(&perms).Error
	return perms, err
}

func (r *RepositoryStruct) BatchCreate(permission []entity.Permission) ([]entity.Permission, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
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

func (r *RepositoryStruct) Create(permission entity.Permission) (*entity.Permission, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&permission).Error
	})

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *RepositoryStruct) GetByGroupId(groupId uint) ([]entity.Permission, error) {
	permissions := make([]entity.Permission, 0)

	err := r.db.Preload("Actions").Where(entity.Permission{GroupId: groupId}).Find(&permissions).Error

	return permissions, err
}

func (r *RepositoryStruct) Update(pk uint, permission entity.Permission) (*entity.Permission, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var ent entity.Permission
		err := tx.First(&ent, pk).Error
		if err != nil {
			return err
		}

		permission.ID = ent.ID

		return tx.Save(permission).Error
	})

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *RepositoryStruct) Delete(pk uint) (bool, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var exists entity.Permission
		err := tx.First(&exists, pk).Error
		if err != nil {
			return err
		}

		return tx.Delete(&exists).Error
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
