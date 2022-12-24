package listall

import (
	"fmt"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type execute struct {
	findRepository repository.IMoneyRepo
}

func NewFindAll(repo repository.IMoneyRepo) Input {
	return execute{
		findRepository: repo,
	}

}

func (e execute) Execute(req entity.Income) {
	fmt.Println(req.Author)
	e.findRepository.Save()
}
