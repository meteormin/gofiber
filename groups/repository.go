package groups

import (
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/utils"
	"gorm.io/gorm"
)

type Repository interface {
	Count(group entity.Group) (int64, error)
	All(page utils.Page) ([]entity.Group, int64, error)
	Create(group entity.Group) (*entity.Group, error)
	Update(pk uint, group entity.Group) (*entity.Group, error)
	Find(pk uint) (*entity.Group, error)
	FindByName(groupName string) (*entity.Group, error)
	Delete(pk uint) (bool, error)
}

type RepositoryStruct struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{db}
}

func (r *RepositoryStruct) Count(group entity.Group) (int64, error) {
	var count int64 = 0
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&group).Count(&count).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return count, err
}

func (r *RepositoryStruct) All(page utils.Page) ([]entity.Group, int64, error) {
	var groups []entity.Group

	count, err := r.Count(entity.Group{})

	if count != 0 {
		if err = r.db.Scopes(utils.Paginate(page)).Order("id desc").Find(&groups).Error; err != nil {
			return make([]entity.Group, 0), 0, err
		}
	} else {
		return make([]entity.Group, 0), 0, nil
	}

	return groups, count, err
}

func (r *RepositoryStruct) Create(group entity.Group) (*entity.Group, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&group).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *RepositoryStruct) Update(pk uint, group entity.Group) (*entity.Group, error) {
	exists, err := r.Find(pk)
	if err != nil {
		return nil, err
	}

	if exists == nil {
		return nil, gorm.ErrRecordNotFound
	}
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if group.ID == exists.ID {
			if err = tx.Save(&group).Error; err != nil {
				return err
			}
		} else {
			group.ID = exists.ID
			if err = tx.Save(&group).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *RepositoryStruct) Find(pk uint) (*entity.Group, error) {
	group := entity.Group{}
	if err := r.db.Preload("Permissions.Actions").First(&group, pk).Error; err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *RepositoryStruct) FindByName(groupName string) (*entity.Group, error) {
	group := &entity.Group{}

	if err := r.db.Preload("Permissions.Actions").Where(entity.Group{Name: groupName}).First(group).Error; err != nil {
		return nil, err
	}

	return group, nil
}

func (r *RepositoryStruct) Delete(pk uint) (bool, error) {
	exists, err := r.Find(pk)
	if err != nil {
		return false, err
	}

	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(exists).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
