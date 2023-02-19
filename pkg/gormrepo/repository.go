package gormrepo

import "gorm.io/gorm"

type GenericRepository[T interface{}] interface {
	DB() *gorm.DB
	Debug() GenericRepository[T]
	All() ([]T, error)
	Create(ent T) (*T, error)
	Save(ent T) (*T, error)
	Find(pk uint) (*T, error)
	FindByEntity(ent T) (*T, error)
	FindByAttribute(attr string, value interface{}) (*T, error)
	Get(fn func(tx *gorm.DB) (*gorm.DB, error)) ([]T, error)
	GetByEntity(ent T) ([]T, error)
	GetByAttributes(attrs map[string]interface{}) ([]T, error)
	Delete(pk uint) (bool, error)
}

type genericRepository[T interface{}] struct {
	model T
	db    *gorm.DB
	debug bool
}

func (g *genericRepository[T]) Debug() GenericRepository[T] {
	g.debug = true
	return g
}

func (g *genericRepository[T]) DB() *gorm.DB {
	if g.debug {
		return g.DB().Debug()
	}

	return g.db
}

func (g *genericRepository[T]) getModel() T {
	return g.model
}

func (g *genericRepository[T]) All() ([]T, error) {
	models := make([]T, 0)

	err := g.DB().Find(&models).Error

	return models, err
}

func (g *genericRepository[T]) Create(ent T) (*T, error) {
	err := g.DB().Transaction(func(tx *gorm.DB) error {
		return tx.Create(&ent).Error
	})

	if err != nil {
		return nil, err
	}

	return &ent, nil
}

func (g *genericRepository[T]) Save(ent T) (*T, error) {
	err := g.DB().Transaction(func(tx *gorm.DB) error {
		return tx.Save(&ent).Error
	})

	if err != nil {
		return nil, err
	}

	return &ent, nil
}

func (g *genericRepository[T]) Find(pk uint) (*T, error) {
	model := g.getModel()
	err := g.DB().First(&model, pk).Error

	if err != nil {
		return nil, err
	}

	return &model, err
}

func (g *genericRepository[T]) FindByEntity(ent T) (*T, error) {
	model := g.getModel()
	if err := g.DB().Where(&ent).First(&model).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (g *genericRepository[T]) FindByAttribute(attr string, value interface{}) (*T, error) {
	model := g.getModel()
	err := g.DB().Where(map[string]interface{}{attr: value}).First(&model).Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (g *genericRepository[T]) Get(fn func(tx *gorm.DB) (*gorm.DB, error)) ([]T, error) {
	models := make([]T, 0)
	err := g.DB().Transaction(func(tx *gorm.DB) error {
		tx, err := fn(tx)
		if err != nil {
			return err
		}
		return tx.Find(&models).Error
	})

	return models, err
}

func (g *genericRepository[T]) GetByEntity(ent T) ([]T, error) {
	models := make([]T, 0)
	err := g.DB().Where(&ent).Find(&models).Error

	return models, err
}

func (g *genericRepository[T]) GetByAttributes(attrs map[string]interface{}) ([]T, error) {
	models := make([]T, 0)
	err := g.DB().Where(attrs).Find(&models).Error

	return models, err
}

func (g *genericRepository[T]) Delete(pk uint) (bool, error) {
	model := g.getModel()
	err := g.DB().First(&model, pk).Error
	if err != nil {
		return false, err
	}

	err = g.DB().Delete(&model).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewGenericRepository[T interface{}](db *gorm.DB, model T) GenericRepository[T] {
	return &genericRepository[T]{
		model: model,
		db:    db,
		debug: false,
	}
}
