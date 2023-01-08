package save

import (
	"errors"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMoneyRepo
}

func NewSave(repo repository.IMoneyRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}

}

func (e execute) Execute(data entity.Value) (entity.Income, error) {
	dataMap := repository.DataMap{
		"Author":      data.Author,
		"Cost":        data.Cost,
		"Description": data.Description,
		"Email":       data.Email,
		"Title":       data.Title,
	}

	result, err := e.findRepository.Save(dataMap)

	if err != nil {
		return entity.Income{}, err
	}

	if result.Error == "unauthorized" {
		error := errors.New(result.Error)
		return result, error
	}

	return result, nil
}
