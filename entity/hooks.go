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
	beforeSave func(s interface{}, tx *gorm.DB) (err error)
	afterSave  func(s interface{}, tx *gorm.DB) (err error)
}

func (s *saving) BeforeSave(tx *gorm.DB) (err error) {
	if s.beforeSave == nil {
		return nil
	}
	return s.beforeSave(s, tx)
}

func (s *saving) HandleBeforeSave(beforeSave func(s interface{}, tx *gorm.DB) (err error)) {
	s.beforeSave = beforeSave
}

func (s *saving) AfterSave(tx *gorm.DB) (err error) {
	if s.afterSave == nil {
		return nil
	}
	return s.afterSave(s, tx)
}

func (s *saving) HandleAfterSave(afterSave func(s interface{}, tx *gorm.DB) (err error)) {
	s.afterSave = afterSave
}

type creating struct {
	beforeCreate func(c interface{}, tx *gorm.DB) (err error)
	afterCreate  func(c interface{}, tx *gorm.DB) (err error)
}

func (c *creating) BeforeCreate(tx *gorm.DB) (err error) {
	if c.beforeCreate == nil {
		return nil
	}
	return c.beforeCreate(c, tx)
}

func (c *creating) HandleBeforeCreate(beforeCreate func(c interface{}, tx *gorm.DB) (err error)) {
	c.beforeCreate = beforeCreate
}

func (c *creating) AfterCreate(tx *gorm.DB) (err error) {
	if c.afterCreate == nil {
		return nil
	}
	return c.afterCreate(c, tx)
}

func (c *creating) HandleAfterCreate(afterCreate func(c interface{}, tx *gorm.DB) (err error)) {
	c.afterCreate = afterCreate
}

type updating struct {
	beforeUpdate func(u interface{}, tx *gorm.DB) (err error)
	afterUpdate  func(u interface{}, tx *gorm.DB) (err error)
}

func (u *updating) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.beforeUpdate == nil {
		return nil
	}
	return u.beforeUpdate(u, tx)
}

func (u *updating) HandleBeforeUpdate(beforeUpdate func(u interface{}, tx *gorm.DB) (err error)) {
	u.beforeUpdate = beforeUpdate
}

func (u *updating) AfterUpdate(tx *gorm.DB) (err error) {
	if u.afterUpdate == nil {
		return nil
	}
	return u.afterUpdate(u, tx)
}

func (u *updating) HandleAfterUpdate(afterUpdate func(u interface{}, tx *gorm.DB) (err error)) {
	u.afterUpdate = afterUpdate
}

type deleting struct {
	beforeDelete func(d interface{}, tx *gorm.DB) (err error)
	afterDelete  func(d interface{}, tx *gorm.DB) (err error)
}

func (d *deleting) BeforeDelete(tx *gorm.DB) (err error) {
	if d.beforeDelete == nil {
		return nil
	}
	return d.beforeDelete(d, tx)
}

func (d *deleting) HandleBeforeDelete(beforeDelete func(d interface{}, tx *gorm.DB) (err error)) {
	d.beforeDelete = beforeDelete
}

func (d *deleting) AfterDelete(tx *gorm.DB) (err error) {
	if d.afterDelete == nil {
		return nil
	}
	return d.afterDelete(d, tx)
}

func (d *deleting) HandleAfterDelete(afterDelete func(d interface{}, tx *gorm.DB) (err error)) {
	d.afterDelete = afterDelete
}

type querying struct {
	afterFind func(q interface{}, tx *gorm.DB) (err error)
}

func (q *querying) AfterFind(tx *gorm.DB) (err error) {
	if q.afterFind == nil {
		return nil
	}
	return q.afterFind(q, tx)
}

func (q *querying) HandleAfterFind(afterFind func(q interface{}, tx *gorm.DB) (err error)) {
	q.afterFind = afterFind
}

type modify struct {
	before func(m interface{}, tx *gorm.DB) error
}

func (m *modify) Before(tx *gorm.DB) error {
	if m.before == nil {
		return nil
	}
	return m.before(m, tx)
}

func (m *modify) HandleBefore(before func(m interface{}, tx *gorm.DB) error) {
	m.before = before
}
