package users

import (
	"github.com/miniyus/gofiber/entity"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user entity.User) (*entity.User, error)
	Find(pk uint) (*entity.User, error)
	All() ([]entity.User, error)
	Update(pk uint, user entity.User) (*entity.User, error)
	Delete(pk uint) (bool, error)
	FindByUsername(username string) (*entity.User, error)
	FindByEntity(user entity.User) (*entity.User, error)
}

type RepositoryStruct struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{db}
}

func (repo *RepositoryStruct) Create(user entity.User) (*entity.User, error) {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryStruct) Find(pk uint) (*entity.User, error) {
	user := entity.User{}

	if err := repo.db.First(&user, pk).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryStruct) All() ([]entity.User, error) {
	var users []entity.User
	if err := repo.db.Find(&users).Error; err != nil {
		return make([]entity.User, 0), err
	}

	return users, nil
}

func (repo *RepositoryStruct) Update(pk uint, user entity.User) (*entity.User, error) {
	exists, err := repo.Find(pk)
	if err != nil {
		return nil, err
	}

	user.ID = exists.ID
	err = repo.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Save(&user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryStruct) Delete(pk uint) (bool, error) {

	user, err := repo.Find(pk)
	if err != nil {
		return false, err
	}

	err = repo.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(user).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *RepositoryStruct) FindByUsername(username string) (*entity.User, error) {
	var user entity.User

	if err := repo.db.Where(&entity.User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *RepositoryStruct) FindByEntity(user entity.User) (*entity.User, error) {
	var rsUser entity.User
	if err := repo.db.Where(&user).First(&rsUser).Error; err != nil {
		return nil, err
	}

	return &rsUser, nil
}
