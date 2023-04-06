package listbyid

import (
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
)

type InputBoundary interface {
	Execute(e *entity.ById) ([]entity.Activity, error)
}
