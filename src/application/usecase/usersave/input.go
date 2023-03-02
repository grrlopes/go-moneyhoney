package usersave

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.Users) (entity.Income, error)
}
