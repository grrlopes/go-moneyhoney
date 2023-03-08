package login

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"


type InputBoundary interface {
	Execute(data *entity.Users) (entity.Users, error)
}
