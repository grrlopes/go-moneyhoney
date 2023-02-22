package listall

import (
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMongoRepo
}

func NewFindAll(repo repository.IMongoRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}

}

func (e execute) Execute(p entity.Pagination) ([]entity.Value, error) {
	result, err := e.findRepository.Find(1, 1)

	if err != nil {
		return []entity.Value{}, err
	}

	return result, nil
}
