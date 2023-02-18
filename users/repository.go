package users

import (
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/pkg/gormrepo"
	"gorm.io/gorm"
)

type Repository interface {
	gormrepo.GenericRepository[entity.User]
	FindByUsername(username string) (*entity.User, error)
}

type RepositoryStruct struct {
	gormrepo.GenericRepository[entity.User]
}

func NewRepository(db *gorm.DB, ent ...entity.User) Repository {
	var u entity.User
	if len(ent) == 0 {
		u = entity.User{}
	} else {
		u = ent[0]
	}

	return &RepositoryStruct{
		gormrepo.NewGenericRepository[entity.User](db, u),
	}
}

func (repo *RepositoryStruct) FindByUsername(username string) (*entity.User, error) {
	return repo.FindByEntity(entity.User{Username: username})
}
