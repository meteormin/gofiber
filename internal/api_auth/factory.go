package api_auth

import (
	"github.com/miniyus/gofiber/auth"
	"github.com/miniyus/gofiber/internal/users"
	"github.com/miniyus/gofiber/pkg/jwt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func New(db *gorm.DB, generator jwt.Generator, logger *zap.SugaredLogger) Handler {
	repo := auth.NewRepository(db)
	service := NewService(repo, users.NewRepository(db, logger), generator)
	handler := NewHandler(service)

	return handler
}
