package users

import (
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/utils"
	"gorm.io/gorm"
)

type Repository interface {
	utils.GenericRepository[entity.User]
	FindByUsername(username string) (*entity.User, error)
}

type RepositoryStruct struct {
	utils.GenericRepository[entity.User]
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{
		utils.NewGenericRepository[entity.User](db, entity.User{}),
	}
}

func (repo *RepositoryStruct) FindByUsername(username string) (*entity.User, error) {
	return repo.FindByEntity(entity.User{Username: username})
}
