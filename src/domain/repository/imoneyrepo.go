package repository

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type DataMap map[string]string
type UpdateMap map[string]interface{}

type IMoneyRepo interface {
	FindAll(limit, skip int) (entity.Income, error)
	FindById(ids *entity.ById) (entity.Income, error)
	Save(data *entity.Value) (entity.Income, error)
	Update(id string, data UpdateMap) (entity.Income, error)
	// Delete(money entity.Income) (entity.Income, error)
}
