package auth

import (
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/utils"
	"gorm.io/gorm"
)

// Repository
// JWT 관련 토큰 Repository
type Repository interface {
	utils.GenericRepository[entity.AccessToken]
	FindByToken(token string) (*entity.AccessToken, error)
	FindByUserId(userId uint) (*entity.AccessToken, error)
	Delete(pk uint) (bool, error)
}

// RepositoryStruct
// Repository 인터페이스 구현 구조체
type RepositoryStruct struct {
	utils.GenericRepository[entity.AccessToken]
}

// NewRepository
// Repository 생성 함수
func NewRepository(db *gorm.DB) Repository {
	return &RepositoryStruct{utils.NewGenericRepository(db, entity.AccessToken{})}
}

// FindByToken
// find token by token(string)
func (repo *RepositoryStruct) FindByToken(token string) (*entity.AccessToken, error) {
	return repo.FindByEntity(entity.AccessToken{Token: token})
}

// FindByUserId
// find token by user id
func (repo *RepositoryStruct) FindByUserId(userId uint) (*entity.AccessToken, error) {
	return repo.FindByEntity(entity.AccessToken{UserId: userId})
}
