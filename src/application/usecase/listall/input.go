package listall

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type Input interface {
	Execute(entity.Income)
}
