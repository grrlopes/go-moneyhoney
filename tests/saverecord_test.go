package tests

import (
	"testing"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/validator"
)

func TestEmailError(t *testing.T) {
	income := &entity.Income{
		Author:      "Gabriel",
		Cost:        "11",
		Description: "test",
		Email:       "gabrielgabriel.test",
		Title:       "test",
	}

	err := validator.Validate(income)

	if err != nil {
		t.Log(err)
	}

}

func TestTitleError(t *testing.T) {
	income := &entity.Income{
		Author:      "Gabriel",
		Cost:        "11",
		Description: "tttt",
		Email:       "gabriel@gabriel.test",
		Title:       "t",
	}

	err := validator.Validate(income)

	if err != nil {
		t.Log(err)
	}

}

func TestAuthorError(t *testing.T) {
	income := &entity.Income{
		Author:      "Gab",
		Cost:        "11",
		Description: "tttt",
		Email:       "gabriel@gabriel.test",
		Title:       "tttt",
	}

	err := validator.Validate(income)

	if err != nil {
		t.Log(err)
	} else {
		t.Fatal(err)
	}
}
