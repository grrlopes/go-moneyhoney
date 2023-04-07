package login

import (
	"errors"
	"fmt"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	"github.com/grrlopes/go-moneyhoney/src/helper"
)

type execute struct {
	findRepository repository.IMongoUserRepo
}

func NewLogin(repo repository.IMongoUserRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}
}

func (e execute) Execute(data *entity.Users) (OutputBoundary, error) {
	var token string
	result, err := e.findRepository.FindUserByName(data)
	if err != nil {
		return OutputBoundary{}, errors.New("Incorrect username or password.")
	}

	fmt.Println(err, "dsfsdfds")
	err = helper.ValidPassword(data, result.Password)
	if err != nil {
		return OutputBoundary{}, errors.New("Incorrect username or password.")
	}

	data.ID = result.ID
	data.CreatedAt = result.CreatedAt
	data.UpdatedAt = result.UpdatedAt

	token, err = helper.GenerateJwt(data)
	if err != nil {
		return OutputBoundary{}, err
	}

	err = helper.VerifyJwt(token)
	if err != nil {
		return OutputBoundary{}, err
	}

	output := OutputBoundary{
		"id":         result.ID,
		"author":     result.Author,
		"email":      result.Email,
		"created_at": result.CreatedAt,
		"updated_at": result.UpdatedAt,
		"token":      token,
	}

	return output, nil
}
