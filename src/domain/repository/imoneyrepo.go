package repository

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type IMoneyRepo interface {
	FindAll() (entity.Income, error)
	Save(id, rev string)
	Update(money entity.Income)
	Delete(money entity.Income)
}
