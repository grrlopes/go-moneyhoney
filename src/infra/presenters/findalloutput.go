package presenters

import (
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
)

type FindAllOutput struct {
	TotalRows int            `json:"total_rows"`
	Offset    int            `json:"offset"`
	Data      []entity.Value `json:"data"`
}

func MoneySuccessResponse(data entity.Income) FindAllOutput {

	mHoney := []entity.Value{}

	for _, v := range data.Rows {
		mHoney = append(mHoney, v.Value)
	}

	return FindAllOutput{
		TotalRows: data.TotalRows,
		Offset:    data.Offset,
		Data:      mHoney,
	}

}
