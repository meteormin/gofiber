package auth

import (
	"github.com/miniyus/gofiber/pkg/jwt"
	"github.com/miniyus/gofiber/users"
	"gorm.io/gorm"
)

func New(db *gorm.DB, generator jwt.Generator) Handler {
	repo := NewRepository(db)
	service := NewService(repo, users.NewRepository(db), generator)
	handler := NewHandler(service)

	return handler
}
