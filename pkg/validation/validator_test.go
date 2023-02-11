package validation_test

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
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

func TestValidationTranslate(t *testing.T) {
	data := TestData{
		Id: "test",
	}
	locale := en.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator("en")
	val := validator.New()
	err := enTranslations.RegisterDefaultTranslations(val, trans)
	if err != nil {
		t.Error(err)
	}

	err = val.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		fmt.Println(errs.Translate(trans))
	}

}
