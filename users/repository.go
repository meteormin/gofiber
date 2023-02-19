package users

import (
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/pkg/gormrepo"
	"gorm.io/gorm"
)

type Repository interface {
	gormrepo.GenericRepository[entity.User]
	Update(pk uint, ent entity.User) (*entity.User, error)
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

func (repo *RepositoryStruct) Update(pk uint, ent entity.User) (*entity.User, error) {
	find, err := repo.Find(pk)
	if err != nil {
		return nil, err
	}

	ent.ID = find.ID

	return repo.Save(ent)
}
