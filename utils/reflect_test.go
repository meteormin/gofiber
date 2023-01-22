package utils_test

import (
	"github.com/miniyus/gofiber/internal/utils"
	"testing"
)

type TestReflect struct {
	Test string
}

func TestHasField(t *testing.T) {
	testObj := TestReflect{
		Test: "test",
	}

	hasField := utils.HasField(&testObj, "Test")

	if !hasField {
		t.Error(testObj)
	}
}
