package save

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type InputBoundary interface {
	Execute(id string, rev string) (entity.Income, error)
}
