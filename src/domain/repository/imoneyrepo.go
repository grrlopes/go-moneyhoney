package repository

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type DataMap map[string]string

type IMoneyRepo interface {
	FindAll() (entity.Income, error)
	Save(data DataMap) (entity.Income, error)
	Update(money entity.Income)
	Delete(money entity.Income)
}
