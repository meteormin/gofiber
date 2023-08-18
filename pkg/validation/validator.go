package validation

import (
	"errors"
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var uni *ut.UniversalTranslator

var transErr error

func Validator() *validator.Validate {
	return validate
}

func NewValidator(fallbackLocale locales.Translator, supportedLocales ...locales.Translator) *validator.Validate {
	validate = validator.New()
	defaultTranslator := newTranslator(fallbackLocale, supportedLocales...)
	if defaultTranslator != nil {
		transErr = enTranslations.RegisterDefaultTranslations(validate, defaultTranslator)
		if transErr != nil {
			panic(transErr)
		}
	}

	return validate
}

func checkValidator() {
	if validate == nil {
		panic("must create validator, use validation.NewValidator() function")
	}
}

func newTranslator(fallbackLocale locales.Translator, supportedLocales ...locales.Translator) (t ut.Translator) {
	uni = ut.New(fallbackLocale, supportedLocales...)
	t, found := uni.GetTranslator("en")
	if found {
		return t
	}

	return nil
}

func RegisterValidation(tags []Tag) {
	checkValidator()
	for _, v := range validations(tags) {
		_ = validate.RegisterValidation(
			v.tag,
			v.fn,
		)
	}
}

func RegisterTranslation(tags []TranslationTag) {
	checkValidator()

	for _, t := range translations(tags) {
		translator, _ := uni.GetTranslator(t.locale)
		_ = validate.RegisterTranslation(
			t.tag,
			translator,
			t.registerFn,
			t.translationFn,
		)
	}
}

func Validate(data interface{}, locale ...string) map[string]string {
	checkValidator()

	translator, _ := uni.FindTranslator(locale...)

	fields := map[string]string{}
	errs := validate.Struct(data)

	if errs == nil {
		return nil
	}

	var invalidValidationError *validator.InvalidValidationError
	if errors.As(errs, &invalidValidationError) {
		errMsg := errs.Error()
		fields["InvalidValidationError"] = errMsg
		return fields
	}

	for _, err := range errs.(validator.ValidationErrors) {
		if err != nil {
			if transErr == nil && translator != nil {
				fields[err.Field()] = err.Translate(translator)
			} else {
				fields[err.Field()] = err.Error()
			}
		}
	}
	return fields
}
