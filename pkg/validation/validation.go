package validation

import "github.com/go-playground/validator/v10"

type Tag struct {
	Tag string
	Fn  validator.Func
}

type registerValidationStruct struct {
	tag string
	fn  validator.Func
}

func validations(cfg []Tag) []registerValidationStruct {
	registerValidations := make([]registerValidationStruct, 0)
	for _, c := range cfg {
		register := registerValidationStruct{
			tag: c.Tag,
			fn:  c.Fn,
		}

		registerValidations = append(registerValidations, register)
	}

	return registerValidations
}
