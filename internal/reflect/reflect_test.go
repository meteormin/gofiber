package reflect_test

import (
	"github.com/miniyus/gofiber/internal/reflect"
	"testing"
)

type TestReflect struct {
	Test string
}

func get() interface{} {
	return TestReflect{Test: "test"}
}

func TestHasField(t *testing.T) {
	testObj := TestReflect{
		Test: "test",
	}

	hasField := reflect.HasField(&testObj, "Test")

	if !hasField {
		t.Error(testObj)
	}
}

func TestGetType(t *testing.T) {
	o := get()
	t.Log(reflect.GetType(o))
}
