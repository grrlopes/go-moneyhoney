package tests

import (
	"testing"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/validator"
)

func TestEntity(t *testing.T) {
	e := &entity.Income{
		Author:      "Gabriel",
		Cost:        "11",
		Description: "test",
		Email:       "gabrielgabriel.test",
		Title:       "test",
	}

	err := validator.Validate(e)

	if err != nil {
		t.Fatal(err)
	}

}
