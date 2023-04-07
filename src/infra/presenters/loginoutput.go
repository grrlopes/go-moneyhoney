package presenters

import (
	"github.com/grrlopes/go-moneyhoney/src/domain/validator"
)

type LoginMongoOutput struct {
	Data map[string]interface{} `json:"data"`
}

type data map[string]interface{}

func LoginSuccess(data map[string]interface{}) LoginMongoOutput {
	return LoginMongoOutput{
		Data: data,
	}
}

func LoginError(err error) errorOuput {
	return errorOuput{
		"message": "Unauthorized",
		"data": data{
			"error": err.Error(),
		},
	}
}

func LoginValidField(data validator.FieldValidation) errorOuput {
	mHoney := []string{}

	for _, v := range data.Message {
		mHoney = append(mHoney, v.Error())
	}

	return errorOuput{
		"Error":   data.Error,
		"Message": mHoney,
	}
}
