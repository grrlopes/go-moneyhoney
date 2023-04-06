package listbyid

import (
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMongoRepo
}

func NewFindById(repo repository.IMongoRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}

}

func (e execute) Execute(b *entity.ById) ([]entity.Activity, error) {
	result, err := e.findRepository.FindById(b)

	if err != nil {
		return []entity.Activity{}, err
	}

	return result, nil
}
