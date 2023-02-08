package validation

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type TranslationTag struct {
	Tag           string
	RegisterFn    validator.RegisterTranslationsFunc
	TranslationFn validator.TranslationFunc
}

type registerTranslationStruct struct {
	tag           string
	trans         ut.Translator
	registerFn    validator.RegisterTranslationsFunc
	translationFn validator.TranslationFunc
}

func translations(translator ut.Translator, cfg []TranslationTag) []registerTranslationStruct {
	trans := make([]registerTranslationStruct, 0)
	for _, c := range cfg {
		register := registerTranslationStruct{
			tag:           c.Tag,
			trans:         translator,
			registerFn:    c.RegisterFn,
			translationFn: c.TranslationFn,
		}
		trans = append(trans, register)
	}

	return trans
}
