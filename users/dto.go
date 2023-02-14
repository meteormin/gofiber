package users

import (
	"github.com/miniyus/gofiber/entity"
	"github.com/miniyus/gofiber/internal/datetime"
)

type CreateUser struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	Email           string `json:"email" validate:"required,email"`
}

type PatchUser struct {
	Email *string `json:"email" validate:"email"`
	Role  *string `json:"role" validate:"string"`
}

type UserResponse struct {
	Id              uint    `json:"id"`
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	EmailVerifiedAt *string `json:"email_verified_at"`
	Role            string  `json:"role"`
	GroupId         *uint   `json:"group_id"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

func ToUserEntity(user CreateUser) entity.User {
	res := entity.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}

	return res
}

func ToUserResponse(user *entity.User) UserResponse {
	createdAt := datetime.TimeIn(user.CreatedAt, "Asia/Seoul")
	updatedAt := datetime.TimeIn(user.UpdatedAt, "Asia/Seoul")

	var emailVerifiedAt *string
	if user.EmailVerifiedAt == nil {
		emailVerifiedAt = nil
	} else {
		formatString := user.EmailVerifiedAt.Format(datetime.DefaultDateLayout)
		emailVerifiedAt = &formatString
	}

	return UserResponse{
		Id:              user.ID,
		Username:        user.Username,
		Role:            string(user.Role),
		Email:           user.Email,
		EmailVerifiedAt: emailVerifiedAt,
		CreatedAt:       createdAt.Format(datetime.DefaultDateLayout),
		UpdatedAt:       updatedAt.Format(datetime.DefaultDateLayout),
	}
}
