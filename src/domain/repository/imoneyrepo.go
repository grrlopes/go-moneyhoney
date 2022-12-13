package repository

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type IMoneyRepo interface {
	FindAll(data []string) []entity.Income
}
