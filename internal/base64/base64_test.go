package base64_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/miniyus/gofiber/internal/base64"
	"testing"
)

func TestB64UrlEncode(t *testing.T) {
	testString := "hello, world"

	encodeString := base64.UrlEncode(testString)
	decodeString, err := base64.UrlDecode(encodeString)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, decodeString, testString)
}
