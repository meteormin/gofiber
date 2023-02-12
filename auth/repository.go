package auth

import (
	"github.com/miniyus/gofiber/entity"
	"gorm.io/gorm"
)

// Repository
// JWT 관련 토큰 Repository
type Repository interface {
	All() ([]*entity.AccessToken, error)
	Create(token entity.AccessToken) (*entity.AccessToken, error)
	Find(pk uint) (*entity.AccessToken, error)
	FindByToken(token string) (*entity.AccessToken, error)
	FindByUserId(userId uint) (*entity.AccessToken, error)
	Update(token entity.AccessToken) (*entity.AccessToken, error)
	Delete(pk uint) (bool, error)
}

// RepositoryStruct
// Repository 인터페이스 구현 구조체
type RepositoryStruct struct {
	db *gorm.DB
}

// NewRepository
// Repository 생성 함수
func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{db}
}

// All
// 모든 데이터 조회
func (repo *RepositoryStruct) All() ([]*entity.AccessToken, error) {
	var tokens []*entity.AccessToken
	if err := repo.db.Find(&tokens).Error; err != nil {
		return nil, err
	}

	return tokens, nil
}

// Create
// 토큰 생성
func (repo *RepositoryStruct) Create(token entity.AccessToken) (*entity.AccessToken, error) {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&token).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// Find
//
//	token by pk
func (repo *RepositoryStruct) Find(pk uint) (*entity.AccessToken, error) {
	token := entity.AccessToken{}
	if err := repo.db.Find(&token, pk).Error; err != nil {
		return nil, err
	}

	return &token, nil
}

// Update
// update token
func (repo *RepositoryStruct) Update(token entity.AccessToken) (*entity.AccessToken, error) {
	err := repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&token).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// Delete
// delete token by pk
func (repo *RepositoryStruct) Delete(pk uint) (bool, error) {
	token, err := repo.Find(pk)
	if token == nil && err != nil {
		return false, err
	}

	err = repo.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Delete(&token).Error; err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		return false, err
	}

	return true, err

}

// FindByToken
// find token by token(string)
func (repo *RepositoryStruct) FindByToken(token string) (*entity.AccessToken, error) {
	ent := entity.AccessToken{}

	if err := repo.db.Where(&entity.AccessToken{Token: token}).First(&ent).Error; err != nil {
		return nil, err
	}

	return &ent, nil
}

// FindByUserId
// find token by user id
func (repo *RepositoryStruct) FindByUserId(userId uint) (*entity.AccessToken, error) {
	ent := entity.AccessToken{}

	err := repo.db.Where(&entity.AccessToken{
		UserId: userId,
	}).Last(&ent).Error

	if err != nil {
		return nil, err
	}

	return &ent, nil
}
