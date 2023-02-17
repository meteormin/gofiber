package entity

import "gorm.io/gorm"

type hooks struct {
	saving
	creating
	updating
	deleting
	querying
	modify
}

func (h *hooks) Hooks() *hooks {
	return h
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

func (s *saving) HandleBeforeSave(beforeSave func(tx *gorm.DB) (err error)) {
	s.beforeSave = beforeSave
}

func (s *saving) AfterSave(tx *gorm.DB) (err error) {
	if s.afterSave == nil {
		return nil
	}
	return s.afterSave(tx)
}

func (s *saving) HandleAfterSave(afterSave func(tx *gorm.DB) (err error)) {
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

func (c *creating) HandleBeforeCreate(beforeCreate func(tx *gorm.DB) (err error)) {
	c.beforeCreate = beforeCreate
}

func (c *creating) AfterCreate(tx *gorm.DB) (err error) {
	if c.afterCreate == nil {
		return nil
	}
	return c.afterCreate(tx)
}

func (c *creating) HandleAfterCreate(afterCreate func(tx *gorm.DB) (err error)) {
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

func (u *updating) HandleBeforeUpdate(beforeUpdate func(tx *gorm.DB) (err error)) {
	u.beforeUpdate = beforeUpdate
}

func (u *updating) AfterUpdate(tx *gorm.DB) (err error) {
	if u.afterUpdate == nil {
		return nil
	}
	return u.afterUpdate(tx)
}

func (u *updating) HandleAfterUpdate(afterUpdate func(tx *gorm.DB) (err error)) {
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

func (d *deleting) HandleBeforeDelete(beforeDelete func(tx *gorm.DB) (err error)) {
	d.beforeDelete = beforeDelete
}

func (d *deleting) AfterDelete(tx *gorm.DB) (err error) {
	if d.afterDelete == nil {
		return nil
	}
	return d.afterDelete(tx)
}

func (d *deleting) HandleAfterDelete(afterDelete func(tx *gorm.DB) (err error)) {
	d.afterDelete = afterDelete
}

type querying struct {
	afterFind func(tx *gorm.DB) (err error)
}

func (q *querying) AfterFind(tx *gorm.DB) (err error) {
	if q.afterFind == nil {
		return nil
	}
	return q.afterFind(tx)
}

func (q *querying) HandleAfterFind(afterFind func(tx *gorm.DB) (err error)) {
	q.afterFind = afterFind
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

func (m *modify) HandleBefore(before func(tx *gorm.DB) error) {
	m.before = before
}
