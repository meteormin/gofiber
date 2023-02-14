package job_queue

import (
	Uuid "github.com/google/uuid"
	"github.com/miniyus/gofiber/entity"
	"gorm.io/gorm"
)

type Repository interface {
	GetByEntity(history entity.JobHistory) ([]entity.JobHistory, error)
	All(fn func(db *gorm.DB) (*gorm.DB, error)) ([]entity.JobHistory, error)
	GetByUserId(userId uint) ([]entity.JobHistory, error)
	Find(pk uint) (*entity.JobHistory, error)
	FindByUuid(uuid string) (*entity.JobHistory, error)
	Create(history entity.JobHistory) (*entity.JobHistory, error)
	Update(pk uint, history entity.JobHistory) (*entity.JobHistory, error)
	UpdateByUuid(uuid string, history entity.JobHistory) (*entity.JobHistory, error)
	Delete(pk uint) (bool, error)
	DeleteByUserId(userId uint) (int64, error)
	DeleteByUuid(uuid string) (int64, error)
}

type RepositoryStruct struct {
	db *gorm.DB
}

func (r *RepositoryStruct) All(fn func(db *gorm.DB) (*gorm.DB, error)) ([]entity.JobHistory, error) {
	histories := make([]entity.JobHistory, 0)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		tx, err := fn(tx)
		if err != nil {
			return err
		}

		return tx.Find(&histories).Error
	})

	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (r *RepositoryStruct) GetByEntity(history entity.JobHistory) ([]entity.JobHistory, error) {
	histories := make([]entity.JobHistory, 0)
	err := r.db.Where(&history).Find(&histories).Error

	return histories, err
}

func (r *RepositoryStruct) Find(pk uint) (*entity.JobHistory, error) {
	var history entity.JobHistory
	err := r.db.First(&history, pk).Error

	return &history, err
}

func (r *RepositoryStruct) GetByUserId(userId uint) ([]entity.JobHistory, error) {
	histories := make([]entity.JobHistory, 0)

	err := r.db.Where(&entity.JobHistory{UserId: &userId}).Find(&histories).Error

	return histories, err
}

func (r *RepositoryStruct) FindByUuid(uuid string) (*entity.JobHistory, error) {
	var history entity.JobHistory

	fromString, err := Uuid.FromBytes([]byte(uuid))
	if err != nil {
		return nil, err
	}

	err = r.db.Where(&entity.JobHistory{UUID: fromString}).Find(&history).Error

	return &history, err
}

func (r *RepositoryStruct) Create(history entity.JobHistory) (*entity.JobHistory, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&history).Error
	})

	if err != nil {
		return nil, err
	}

	return &history, nil
}

func (r *RepositoryStruct) Update(pk uint, history entity.JobHistory) (*entity.JobHistory, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		find, err := r.Find(pk)
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

func (r *RepositoryStruct) UpdateByUuid(uuid string, history entity.JobHistory) (*entity.JobHistory, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
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

func (r *RepositoryStruct) Delete(pk uint) (bool, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		find, err := r.Find(pk)
		if err != nil {
			return err
		}

		return tx.Delete(find).Error
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *RepositoryStruct) DeleteByUserId(userId uint) (int64, error) {
	var rowsAffected int64
	err := r.db.Transaction(func(tx *gorm.DB) error {
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
	err := r.db.Transaction(func(tx *gorm.DB) error {
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
		db: db,
	}
}
