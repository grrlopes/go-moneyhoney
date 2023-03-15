package usersave

import (
	"errors"
	"time"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	"github.com/grrlopes/go-moneyhoney/src/helper"
)

type execute struct {
	findRepository repository.IMongoUserRepo
}

func NewUserSave(repo repository.IMongoUserRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}
}

func (e execute) Execute(data *entity.Users) (entity.Income, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.Password, _ = helper.CreatePassword(data)

	result, err := e.findRepository.UserSave(data)

	if err != nil {
		return entity.Income{}, err
	}

	if result.Error == "unauthorized" {
		error := errors.New(result.Error)
		return result, error
	}

	return result, nil
}
