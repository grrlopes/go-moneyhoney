package couchdb

import (
	"encoding/json"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type money struct {
	repository.IMoneyRepo
}

func NewMoneyRepository() repository.IMoneyRepo {
	return &money{}
}

func (db *money) FindAll() (entity.Income, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(os.Getenv("USER"), os.Getenv("PASS")).
		SetQueryParams(map[string]string{
			"limit": "1",
			"skip":  "dfsf",
		}).
		Get(os.Getenv("URL") + "/_design/list/_view/findall")
	if err != nil {
		return entity.Income{}, err
	}

	defer client.SetCloseConnection(true)

	var result entity.Income

	json.Unmarshal(resp.Body(), &result)

	return result, nil
}

func (db *money) Save(data repository.DataMap) (entity.Income, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(os.Getenv("USER"), os.Getenv("PASS")).
		SetBody(data).
		Post(os.Getenv("URL"))

	if err != nil {
		return entity.Income{}, err
	}

	defer client.SetCloseConnection(true)

	var result entity.Income

	json.Unmarshal(resp.Body(), &result)

	return result, nil
}
