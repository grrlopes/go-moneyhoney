package couchdb

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type openDB interface {
	FindAll(limit, skip int) (entity.Income, error)
	FindById(ids *entity.ById) (entity.Income, error)
	Save(data *entity.Value) (entity.Income, error)
	Update(id string, data repository.UpdateMap) (entity.Income, error)
}

type money struct {
	con repository.IMoneyRepo
}

func NewMoneyRepository() repository.IMoneyRepo {
	var db openDB
	return &money{
		con: db,
	}
}

func (db *money) FindAll(limit, skip int) (entity.Income, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(os.Getenv("USER"), os.Getenv("PASS")).
		SetQueryParams(map[string]string{
			"limit": strconv.Itoa(limit),
			"skip":  strconv.Itoa(skip),
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

func (db *money) FindById(ids *entity.ById) (entity.Income, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBasicAuth(os.Getenv("USER"), os.Getenv("PASS")).
		SetQueryParams(map[string]string{
			"include_docs": "false",
			"key":          `"` + ids.ID + `"`,
		}).
		Get(os.Getenv("URL") + "/_design/list/_view/findbyid")

	if err != nil {
		return entity.Income{}, err
	}

	defer client.SetCloseConnection(true)

	var result entity.Income

	json.Unmarshal(resp.Body(), &result)
	return result, nil
}

func (db *money) Save(data *entity.Value) (entity.Income, error) {
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

func (db *money) Update(id string, data repository.UpdateMap) (entity.Income, error) {
	client := resty.New()
	resp, err := client.R().EnableTrace().
		SetHeader("Accept", "application/json").
		SetBasicAuth(os.Getenv("USER"), os.Getenv("PASS")).
		SetBody(data).
		Put(os.Getenv("URL") + "/" + id)

	if err != nil {
		return entity.Income{}, err
	}

	fmt.Println(id, data)

	defer client.SetCloseConnection(true)

	var result entity.Income

	json.Unmarshal(resp.Body(), &result)

	return result, nil
}
