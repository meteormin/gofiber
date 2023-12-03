package config

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ko"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/miniyus/gofiber/pkg/validation"
)

type Validation struct {
	FallbackLocale   locales.Translator
	SupportedLocales []locales.Translator
	Validations      []validation.Tag
	Translations     []validation.TranslationTag
}

func validationConfig() Validation {
	return Validation{
		FallbackLocale: en.New(),
		SupportedLocales: []locales.Translator{
			ko.New(),
			en.New(),
		},
		Validations:  validations(),
		Translations: translations(),
	}
}

func validations() []validation.Tag {
	return []validation.Tag{}
}

func translations() []validation.TranslationTag {
	return []validation.TranslationTag{
		{
			Locale: "ko",
			Tag:    "required",
			RegisterFn: func(ut ut.Translator) error {
				return ut.Add("required", "{0} 필드는 필수 입니다.", true)
			},
			TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("required", fe.Field())
				return t
			},
		},
		{
			Locale: "ko",
			Tag:    "email",
			RegisterFn: func(ut ut.Translator) error {
				return ut.Add("email", "{0} 필드는 Email 형식이어야 합니다.", true)
			},
			TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("email", fe.Field())
				return t
			},
		},
		{
			Locale: "ko",
			Tag:    "url",
			RegisterFn: func(ut ut.Translator) error {
				return ut.Add("url", "{0} 필드는 URL 형식이어야 합니다.", true)
			},
			TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("url", fe.Field())
				return t
			},
		},
		{
			Locale: "ko",
			Tag:    "dir",
			RegisterFn: func(ut ut.Translator) error {
				return ut.Add("dir", "{0} 필드는 디렉토리 경로 형식이어야 합니다.", true)
			},
			TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("dir", fe.Field())
				return t
			},
		},
		{
			Locale: "ko",
			Tag:    "boolean",
			RegisterFn: func(ut ut.Translator) error {
				return ut.Add("boolean", "{0} 필드는 boolean 타입이어야 합니다.", true)
			},
			TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("boolean", fe.Field())
				return t
			},
		},
		{
			Locale: "ko",
			Tag:    "eqfield",
			RegisterFn: func(ut ut.Translator) error {
				return ut.Add("eqfield", "{0} 필드는 {1} 필드와 일치해야 합니다.", true)
			},
			TranslationFn: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("eqfield", fe.Field(), fe.Param())
				return t
			},
		},
	}
}
