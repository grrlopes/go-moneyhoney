package couchdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
		Get(os.Getenv("URL") + "/_design/list/_view/findall")

	if err != nil {
		return entity.Income{}, err
	}

	defer client.SetCloseConnection(true)

	var result entity.Income

	json.Unmarshal(resp.Body(), &result)

	return result, nil
}

func (db *money) Save() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://localhost:5498", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err.Error())
	}

	var result []entity.Income
	json.Unmarshal(bodyByte, &result)
}
