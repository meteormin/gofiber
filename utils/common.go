package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miniyus/gofiber/apierrors"
	"github.com/miniyus/gofiber/internal/base64"
	"github.com/miniyus/gofiber/internal/datetime"
	"github.com/miniyus/gofiber/internal/hash"
	"github.com/miniyus/gofiber/pkg/validation"
	"github.com/redis/go-redis/v9"
	"time"
)

type StatusResponse struct {
	Status bool `json:"status"`
}

type DataResponse[T interface{}] struct {
	Data T `json:"data"`
}

const DefaultDateLayout = datetime.DefaultDateLayout

type JsonTime datetime.JsonTime

func TimeIn(t time.Time, tz string) time.Time {
	return datetime.TimeIn(t, tz)
}

func RedisClientMaker(options *redis.Options) func() *redis.Client {
	return func() *redis.Client {
		return redis.NewClient(options)
	}
}

func HandleValidate(c *fiber.Ctx, data interface{}) *apierrors.ValidationErrorResponse {
	err := c.BodyParser(data)
	if err != nil {
		errRes := apierrors.NewValidationErrorResponse(c, map[string]string{
			"parse_error": err.Error(),
		})
		return errRes
	}

	failed := validation.Validate(data)
	if failed != nil {
		errRes := apierrors.NewValidationErrorResponse(c, failed)
		return errRes
	}

	return nil
}

func Base64UrlEncode(data string) string {
	return base64.UrlEncode(data)
}

func Base64UrlDecode(data string) (string, error) {
	return base64.UrlDecode(data)
}

func Hash() hash.BcryptWrapper {
	return hash.Bcrypt
}
