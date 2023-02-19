package repository

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type DataMap map[string]string

type IMoneyRepo interface {
	FindAll(limit, skip int) (entity.Income, error)
	FindById(ids *entity.ById) (entity.Income, error)
	Save(data *entity.Value) (entity.Income, error)
	Update(money entity.Income)
	Delete(money entity.Income)
}
