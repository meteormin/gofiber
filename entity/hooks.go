package entity

import "gorm.io/gorm"

type hooks struct {
	saving
	creating
	updating
	deleting
	modify
}

type saving struct {
	beforeSave func(tx *gorm.DB) (err error)
	afterSave  func(tx *gorm.DB) (err error)
}

func (s *saving) BeforeSave(tx *gorm.DB) (err error) {
	if s.beforeSave == nil {
		return nil
	}
	return s.beforeSave(tx)
}

func (s *saving) SetBeforeSave(beforeSave func(tx *gorm.DB) (err error)) {
	s.beforeSave = beforeSave
}

func (s *saving) AfterSave(tx *gorm.DB) (err error) {
	if s.afterSave == nil {
		return nil
	}
	return s.afterSave(tx)
}

func (s *saving) SetAfterSave(afterSave func(tx *gorm.DB) (err error)) {
	s.afterSave = afterSave
}

type creating struct {
	beforeCreate func(tx *gorm.DB) (err error)
	afterCreate  func(tx *gorm.DB) (err error)
}

func (c *creating) BeforeCreate(tx *gorm.DB) (err error) {
	if c.beforeCreate == nil {
		return nil
	}
	return c.beforeCreate(tx)
}

func (c *creating) SetBeforeCreate(beforeCreate func(tx *gorm.DB) (err error)) {
	c.beforeCreate = beforeCreate
}

func (c *creating) AfterCreate(tx *gorm.DB) (err error) {
	if c.afterCreate == nil {
		return nil
	}
	return c.afterCreate(tx)
}

func (c *creating) SetAfterCreate(afterCreate func(tx *gorm.DB) (err error)) {
	c.afterCreate = afterCreate
}

type updating struct {
	beforeUpdate func(tx *gorm.DB) (err error)
	afterUpdate  func(tx *gorm.DB) (err error)
}

func (u *updating) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.beforeUpdate == nil {
		return nil
	}
	return u.beforeUpdate(tx)
}

func (u *updating) SetBeforeUpdate(beforeUpdate func(tx *gorm.DB) (err error)) {
	u.beforeUpdate = beforeUpdate
}

func (u *updating) AfterUpdate(tx *gorm.DB) (err error) {
	if u.afterUpdate == nil {
		return nil
	}
	return u.afterUpdate(tx)
}

func (u *updating) SetAfterUpdate(afterUpdate func(tx *gorm.DB) (err error)) {
	u.afterUpdate = afterUpdate
}

type deleting struct {
	beforeDelete func(tx *gorm.DB) (err error)
	afterDelete  func(tx *gorm.DB) (err error)
}

func (d *deleting) BeforeDelete(tx *gorm.DB) (err error) {
	if d.beforeDelete == nil {
		return nil
	}
	return d.beforeDelete(tx)
}

func (d *deleting) SetBeforeDelete(beforeDelete func(tx *gorm.DB) (err error)) {
	d.beforeDelete = beforeDelete
}

func (d *deleting) AfterDelete(tx *gorm.DB) (err error) {
	if d.afterDelete == nil {
		return nil
	}
	return d.afterDelete(tx)
}

func (d *deleting) SetAfterDelete(afterDelete func(tx *gorm.DB) (err error)) {
	d.afterDelete = afterDelete
}

type modify struct {
	before func(tx *gorm.DB) error
}

func (m *modify) Before(tx *gorm.DB) error {
	if m.before == nil {
		return nil
	}
	return m.before(tx)
}

func (m *modify) SetBefore(before func(tx *gorm.DB) error) {
	m.before = before
}
