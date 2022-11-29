package auth

import (
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	configure "github.com/miniyus/go-fiber/config"
	"github.com/miniyus/go-fiber/internal/core/api_error"
	"log"
)

// 공통 미들웨어 작성

type User struct {
	Id        uint   `json:"id"`
	GroupId   uint   `json:"group_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	ExpiresIn int64  `json:"expires_in"`
}

func Middlewares() []fiber.Handler {
	mws := []fiber.Handler{
		JwtMiddleware,
		GetUserFromJWT,
	}

	return mws
}

func GetUserFromJWT(c *fiber.Ctx) error {
	jwtData, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		log.Print("access guest")
		return c.Next()
	}

	claims := jwtData.Claims.(jwt.MapClaims)

	userId := uint(claims["user_id"].(float64))
	groupId := uint(claims["group_id"].(float64))
	username := claims["username"].(string)
	email := claims["email"].(string)
	createdAt := claims["created_at"].(string)
	expiresIn := int64(claims["expires_in"].(float64))

	currentUser := &User{
		Id:        userId,
		GroupId:   groupId,
		Username:  username,
		Email:     email,
		CreatedAt: createdAt,
		ExpiresIn: expiresIn,
	}

	c.Locals(configure.AuthUser, currentUser)
	return c.Next()
}

func JwtMiddleware(c *fiber.Ctx) error {
	config, ok := c.Locals(configure.Config).(*configure.Configs)
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "Can not found Config Context...")
	}

	middleware := newJwtMiddleware(config.Auth.Jwt)

	return middleware(c)
}

func newJwtMiddleware(config jwtWare.Config) fiber.Handler {
	jwtConfig := config
	jwtConfig.ErrorHandler = jwtError
	return jwtWare.New(jwtConfig)
}

func jwtError(c *fiber.Ctx, err error) error {
	var errRes api_error.ErrorResponse

	if err.Error() == "Missing or malformed JWT" {
		errRes = api_error.NewErrorResponse(c, fiber.StatusBadRequest, err.Error())

		return errRes.Response()
	}

	errRes = api_error.NewErrorResponse(c, fiber.StatusBadRequest, "Invalid or expired JWT!")

	return errRes.Response()
}
