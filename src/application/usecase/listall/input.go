package listall

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type InputBoundary interface {
	Execute(p entity.Pagination) (entity.Income, error)
}
