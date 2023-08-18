package validation

import (
	"github.com/go-playground/validator/v10"
)

// TranslationTag 언어 변환 설정 구조체
type TranslationTag struct {
	// Locale 언어
	Locale string
	// Tag 유효성 검사 태그이름
	Tag string
	// RegisterFn 유효성 검사 규칙 등록 함수(클로저)
	RegisterFn validator.RegisterTranslationsFunc
	// TranslationFn
	TranslationFn validator.TranslationFunc
}

// registerTranslationStruct
type registerTranslationStruct struct {
	tag           string
	locale        string
	registerFn    validator.RegisterTranslationsFunc
	translationFn validator.TranslationFunc
}

func translations(cfg []TranslationTag) []registerTranslationStruct {
	translationSlice := make([]registerTranslationStruct, 0)
	for _, c := range cfg {
		register := registerTranslationStruct{
			tag:           c.Tag,
			locale:        c.Locale,
			registerFn:    c.RegisterFn,
			translationFn: c.TranslationFn,
		}
		translationSlice = append(translationSlice, register)
	}

	return translationSlice
}
