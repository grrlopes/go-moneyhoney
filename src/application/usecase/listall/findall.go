package listall

import (
	"errors"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMoneyRepo
}

func NewFindAll(repo repository.IMoneyRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}

}

func (e execute) Execute(p entity.Pagination) (entity.Income, error) {
	result, err := e.findRepository.FindAll(p.Limit, p.Skip)

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
