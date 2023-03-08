package repository

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type IMongoUserRepo interface {
	FindUserByName(data *entity.Users) (entity.Users, error)
	UserSave(data *entity.Users) (entity.Income, error)
}
