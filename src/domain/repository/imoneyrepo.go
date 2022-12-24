package repository

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type IMoneyRepo interface {
	FindAll() []entity.Income
	Save()
	Update(money entity.Income)
	Delete(money entity.Income)
}
