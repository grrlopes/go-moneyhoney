package update

import (
	"errors"
	"time"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMoneyRepo
}

func NewUpdate(repo repository.IMoneyRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}

}

func (e execute) Execute(data *entity.Value) (entity.Income, error) {
	data.UpdatedAt = time.Now()
	id := data.ID

	parseData := map[string]interface{}{
		"_rev":       data.Rev,
		"author":     data.Author,
		"email":      data.Email,
		"item":       data.Item,
		"updated_at": data.UpdatedAt,
	}

	result, err := e.findRepository.Update(id, parseData)

	if err != nil {
		return entity.Income{}, err
	}

	if result.Error == "unauthorized" {
		error := errors.New(result.Error)
		return result, error
	}

	return result, nil
}
