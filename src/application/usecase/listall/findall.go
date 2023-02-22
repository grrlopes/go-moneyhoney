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

func (e execute) Execute(p entity.Pagination) ([]entity.Value, entity.Count, error) {
	result, count, err := e.findRepository.Find(int64(p.Limit), int64(p.Skip))

	if err != nil {
		return []entity.Value{}, entity.Count{}, err
	}

	return result, count, nil
}
