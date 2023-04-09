package update

import (
	"time"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMongoRepo
}

func NewUpdate(repo repository.IMongoRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}

}

func (e execute) Execute(data *entity.Activity) (map[string]interface{}, error) {
	var status string = "Data was updated!"
	data.UpdatedAt = time.Now()
	id := data.ID

	result, err := e.findRepository.Update(id, data)

	if err != nil {
		return map[string]interface{}{}, err
	}

	parseResult := map[string]interface{}{
		"message": &status,
	}

	if result <= 0 {
		status = "Data was not updated!"
		return parseResult, err
	}

	return parseResult, nil
}
