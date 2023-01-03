package presenters

import (
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
)

type FindAllOutput struct {
	TotalRows int            `json:"total_rows"`
	Offset    int            `json:"offset"`
	Data      []entity.Value `json:"data"`
}

type errorOuput map[string]string

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

func MoneyErrorResponse(data entity.Income) errorOuput {
	return errorOuput{
		"Error":   data.Error,
		"Message": data.Reason,
	}
}
