package jobqueue

import (
	Uuid "github.com/google/uuid"
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/pkg/gormrepo"
	"gorm.io/gorm"
)

type Repository interface {
	gormrepo.GenericRepository[entity.JobHistory]
	GetByUserId(userId uint) ([]entity.JobHistory, error)
	FindByUuid(uuid string) (*entity.JobHistory, error)
	UpdateByUuid(uuid string, history entity.JobHistory) (*entity.JobHistory, error)
	Delete(pk uint) (bool, error)
	DeleteByUserId(userId uint) (int64, error)
	DeleteByUuid(uuid string) (int64, error)
}

type RepositoryStruct struct {
	gormrepo.GenericRepository[entity.JobHistory]
}

func (r *RepositoryStruct) GetByUserId(userId uint) ([]entity.JobHistory, error) {
	return r.GetByEntity(entity.JobHistory{UserId: &userId})
}

func (r *RepositoryStruct) FindByUuid(uuid string) (*entity.JobHistory, error) {
	fromString, err := Uuid.FromBytes([]byte(uuid))
	if err != nil {
		return nil, err
	}

	return repo.FindByEntity(entity.JobHistory{UUID: fromString})
}

func (r *RepositoryStruct) UpdateByUuid(uuid string, history entity.JobHistory) (*entity.JobHistory, error) {
	err := r.DB().Transaction(func(tx *gorm.DB) error {
		find, err := r.FindByUuid(uuid)
		if err != nil {
			return err
		}

		history.ID = find.ID
		return tx.Save(&history).Error
	})

	if err != nil {
		return nil, err
	}

	return &history, nil
}

func (r *RepositoryStruct) DeleteByUserId(userId uint) (int64, error) {
	var rowsAffected int64
	err := r.DB().Transaction(func(tx *gorm.DB) error {
		tx = tx.Where(&entity.JobHistory{UserId: &userId}).Delete(&entity.JobHistory{})
		rowsAffected = tx.RowsAffected
		return tx.Error
	})

	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, nil
}

func (r *RepositoryStruct) DeleteByUuid(uuid string) (int64, error) {
	var rowsAffected int64 = 0
	err := r.DB().Transaction(func(tx *gorm.DB) error {
		fromString, err := Uuid.FromBytes([]byte(uuid))
		if err != nil {
			return err
		}

		tx = tx.Where(&entity.JobHistory{UUID: fromString}).Delete(&entity.JobHistory{})
		rowsAffected = tx.RowsAffected
		return tx.Error
	})

	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{
		gormrepo.NewGenericRepository(db, entity.JobHistory{}),
	}
}
