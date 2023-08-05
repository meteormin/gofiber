package validation_test

import (
	"github.com/go-playground/assert/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/miniyus/gofiber/pkg/validation"
	"log"
	"testing"
)

type TestData struct {
	Id          string `validate:"required"`
	Pass        string `validate:"required"`
	PassConfirm string `validate:"required,eqfield=Pass"`
	Custom      string `validate:"custom"`
}

func TestRegisterTranslation(t *testing.T) {
	tags := []validation.Tag{
		{
			Tag: "custom",
			Fn: func(fl validator.FieldLevel) bool {
				return fl.Field().String() == "custom"
			},
		},
	}

	validation.RegisterValidation(tags)

	transTags := []validation.TranslationTag{
		{
			Tag: "required",
			RegisterFn: func(ut ut.Translator) error {
				return ut.Add("required", "{0} 필드는 필수 입니다.", true)
			},
			TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("required", fe.Field())
				return t
			},
		},
		{
			Tag: "custom",
			RegisterFn: func(ut ut.Translator) error {
				return ut.Add("custom", "{0} 커스텀", true)
			},
			TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("custom", fe.Field())
				return t
			},
		},
	}

	validation.RegisterTranslation(transTags)

	data := TestData{
		Id:     "test",
		Pass:   "test",
		Custom: "cus",
	}

	validated := validation.Validate(data)

	log.Print(validated)
}

func TestValidate(t *testing.T) {
	data := TestData{
		Id:          "test",
		Pass:        "test",
		PassConfirm: "te",
	}

	validated := validation.Validate(data)

	testValidated := map[string]string{
		"Pass": "Key: 'TestData.Pass' Error:Field validation for 'Pass' failed on the 'required' tag",
	}

	t.Log(validated)
	assert.Equal(t, validated, testValidated)
}
