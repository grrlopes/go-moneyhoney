package listall

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type InputBoundary interface {
	Execute(entity.Income) (entity.Income, error)
}
