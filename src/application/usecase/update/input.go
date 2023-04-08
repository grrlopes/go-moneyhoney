package update

import "github.com/grrlopes/go-moneyhoney/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.Activity) ([]entity.Activity, error)
}
