package listbyid

import "github.com/grrlopes/go-moneyhoney/src/infra/presenters"

type OutputBoundary interface {
	output() []presenters.FindAllOutput
}
