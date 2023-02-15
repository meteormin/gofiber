package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtWare "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/miniyus/gofiber/app"
	"github.com/miniyus/gofiber/database"
	"github.com/miniyus/gofiber/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// 인증 관련 공통 미들웨어 작성

// User
// context에 저장될 유저 정보 구조체
type User struct {
	Id        uint   `json:"id"`
	GroupId   *uint  `json:"group_id"`
	Role      string `json:"role"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	ExpiresIn *int64 `json:"expires_in"`
}

type MiddlewaresParameter struct {
	Cfg    jwtWare.Config
	Logger *zap.SugaredLogger
	DB     *gorm.DB
}

type mw func(ctx *fiber.Ctx) (*fiber.Ctx, error)

// jwtMiddleware
// jwt 유효성 체크 미들웨어
func jwtMiddleware(jwtConfig jwtWare.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		middleware := newJwtMiddleware(jwtConfig)

		return middleware(c)
	}
}

func newJwtMiddleware(config jwtWare.Config) fiber.Handler {
	jwtConfig := config
	jwtConfig.ErrorHandler = jwtError()
	return jwtWare.New(jwtConfig)
}

// jwtError
// jwt 생성과 해독(? decode...) 관련 에러 핸들링
func jwtError() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		var status int

		if err.Error() == "Missing or malformed JWT" {
			status = fiber.StatusBadRequest
			return fiber.NewError(status, err.Error())
		}

		return fiber.NewError(status, err.Error())
	}
}

// Middlewares auth middleware
func Middlewares(parameter MiddlewaresParameter, fn ...fiber.Handler) []fiber.Handler {
	return mergeMiddlewares(parameter, fn...) // TODO: 미들웨어 추가가 좀 더 편하게 수정해야함
}

// mergeMiddlewares
// 미들웨어 슬라이스 리턴
// 인증 관련된 미들웨어 함수의 집합으로 이 함수에 등록된 순서대로 실행 가능
func mergeMiddlewares(parameter MiddlewaresParameter) fiber.Handler {
	app.App().Middleware(func(fiber *fiber.App, app app.Application) {
		fiber.Use(jwtMiddleware(parameter.Cfg))
	})
	return func(ctx *fiber.Ctx) error {
		jwtToken, ok := ctx.Locals("user").(*jwt.Token)
		if !ok {
			return fiber.ErrUnauthorized
		}

		fromJWT, err := getUserFromJWT(jwtToken)
		if err != nil {
			return err
		}

		err = checkExpired(fromJWT, parameter.DB)
		if err != nil {
			return err
		}
		logFormat := accessLogFormat{}
		parameter.Logger.Info(logFormat.ToLogFormat())
	}
}

type accessLogFormat struct {
	UserId  uint
	IP      string
	Elapsed int64
	Method  string
	Request string
	ErrMsg  string
}

func (alf accessLogFormat) ToLogFormat() string {
	return fmt.Sprintf(
		"user: %4d | IP: %15s | %6dms | %s | %s %s",
		alf.UserId, alf.IP, alf.Elapsed, alf.Method, alf.Request, alf.ErrMsg,
	)
}

// getUserFromJWT
// get user information from jwt token
func getUserFromJWT(jwtData *jwt.Token) (*User, error) {
	claims := jwtData.Claims.(jwt.MapClaims)

	userId := uint(claims["user_id"].(float64))

	var groupId uint
	if claims["group_id"] != nil {
		groupId = uint(claims["group_id"].(float64))
	}

	role := claims["role"].(string)
	username := claims["username"].(string)
	email := claims["email"].(string)
	createdAt := claims["created_at"].(string)

	var expiresIn int64
	if claims["expires_in"] != nil {
		expiresIn = int64(claims["expires_in"].(float64))
	}

	layout := "2006-01-02T15:04:05Z07:00"
	createdAtTime, err := time.Parse(layout, createdAt)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:        userId,
		GroupId:   &groupId,
		Role:      role,
		Username:  username,
		Email:     email,
		CreatedAt: createdAtTime.Format("2006-01-02 15:04:05"),
		ExpiresIn: &expiresIn,
	}, nil
}

// checkExpired
// jwt 만료 기간 체크 미들웨어
func checkExpired(user *User, gormDB ...*gorm.DB) error {
	var db *gorm.DB
	if len(gormDB) != 0 {
		db = gormDB[0]
	}

	if db == nil {
		db = database.GetDB()
	}

	tokenRepository := NewRepository(db)

	token, err := tokenRepository.FindByUserId(user.Id)
	if err != nil {
		statusCode := fiber.StatusUnauthorized
		return fiber.NewError(statusCode, "Can't Find User From Database")
	}

	if token.ExpiresAt.Unix() < time.Now().Unix() {
		statusCode := fiber.StatusUnauthorized
		return fiber.NewError(statusCode, "JWT is expired")
	}
}

func GetAuthUser(c *fiber.Ctx) (*User, error) {
	user, err := utils.GetContext[*User](c, utils.AuthUserKey)
	if err != nil {
		return nil, err
	}

	return user, nil
}
