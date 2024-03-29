package update

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.Value) (entity.Income, error)
}
