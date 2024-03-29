package save

import (
	"errors"
	"time"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMongoRepo
}

func NewSave(repo repository.IMongoRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}
}

func (e execute) Execute(data *entity.Activity) (entity.Income, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := e.findRepository.Save(data)

	if err != nil {
		return entity.Income{}, err
	}

	if result.Error == "unauthorized" {
		error := errors.New(result.Error)
		return result, error
	}

	return result, nil
}
