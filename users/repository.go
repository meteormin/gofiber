package users

import (
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gorm-extension/gormrepo"
	"gorm.io/gorm"
)

type Repository interface {
	gormrepo.GenericRepository[entity.User]
	FindByUsername(username string) (*entity.User, error)
}

type RepositoryStruct struct {
	gormrepo.GenericRepository[entity.User]
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{
		gormrepo.NewGenericRepository[entity.User](db, entity.User{}),
	}
}

func (repo *RepositoryStruct) FindByUsername(username string) (*entity.User, error) {
	return repo.FindByEntity(entity.User{Username: username})
}
