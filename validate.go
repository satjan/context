package context

import (
	"github.com/go-playground/locales/mn"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	locale2 "github.com/satjan/validator/locale"
)

var Instance Validator

func init() {
	validate, trans := Validate()
	Instance = Validator{
		Validate: validate,
		Trans:    trans,
	}
}

type Validator struct {
	Validate *validator.Validate
	Trans    ut.Translator
}

func Validate() (*validator.Validate, ut.Translator) {
	locale := mn.New()
	uni := ut.New(locale, locale)
	trans, _ := uni.GetTranslator("mn")
	validate := validator.New()
	err := locale2.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, nil
	}

	return validate, trans
}
