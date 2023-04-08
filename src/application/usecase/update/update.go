package update

import (
	"log"
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

func (e execute) Execute(data *entity.Activity) ([]entity.Activity, error) {
	data.UpdatedAt = time.Now()
	id := data.ID

	parseData := map[string]interface{}{
		"author":     data.User,
		"item":       data.Item,
		"created_at": data.CreatedAt,
		"updated_at": data.UpdatedAt,
	}

  log.Println(parseData)

	result, err := e.findRepository.Update(id, data)

	if err != nil {
		return []entity.Activity{}, err
	}

	return result, nil
}
