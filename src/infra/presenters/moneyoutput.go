package presenters

import (
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/validator"
)

type MoneyMongoOutput struct {
	TotalRows int            `json:"total_rows"`
	Offset    int            `json:"offset"`
	Data      []entity.Value `json:"data"`
}

func MoneySuccess(data []entity.Value) MoneyMongoOutput {

	return MoneyMongoOutput{
		TotalRows: 0,
		Offset:    0,
		Data:      data,
	}
}

func MoneyError(data []entity.Value) errorOuput {
	return errorOuput{
		"Error":   data,
		"Message": data,
	}
}

func MoneyValidField(data validator.FieldValidation) errorOuput {
	mHoney := []string{}

	for _, v := range data.Message {
		mHoney = append(mHoney, v.Error())
	}

	return errorOuput{
		"Error":   data.Error,
		"Message": mHoney,
	}
}
