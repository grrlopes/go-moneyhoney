package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
)

type FindAllOutput struct {
	TotalRows int            `json:"total_rows"`
	Offset    int            `json:"offset"`
	Data      []entity.Value `json:"data"`
}

type FieldValidation struct {
	Error   string  `json:"error"`
	Message []error `json:"message"`
}

func Validate(e *entity.Value) (error FieldValidation) {
	validate := validator.New()

	eng := en.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(e)
	errs := handleError(err, trans)

	erros := FieldValidation{
		Error:   "Field not valid",
		Message: errs,
	}

	return erros
}

func handleError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
