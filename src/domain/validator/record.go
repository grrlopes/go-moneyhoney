package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
)

func Validate(e *entity.Income) error {
	validate := validator.New()
	err := validate.Struct(e)

	errs := err.(validator.ValidationErrors)

	for _, e := range errs {
		fmt.Println(e.Error(), "dsdfs")
	}
	return err
}
