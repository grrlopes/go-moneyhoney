package listbyid

import (
	"errors"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMoneyRepo
}

func NewFindById(repo repository.IMoneyRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}

}

func (e execute) Execute(b *entity.ById) (entity.Income, error) {
	result, err := e.findRepository.FindById(b)

	if err != nil {
		return entity.Income{}, err
	}

	if result.Error == "unauthorized" ||
		result.Error == "query_parse_error" {
		error := errors.New(result.Error)
		return result, error
	}

	return result, nil
}
