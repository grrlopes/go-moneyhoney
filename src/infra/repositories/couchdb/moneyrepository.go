package couchdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
)

type money struct {
	repository.IMoneyRepo
}

func NewMoneyRepository() repository.IMoneyRepo {
	return &money{}

}

func (db *money) Save() {
	fmt.Println("#####")
}

func (db *money) FindAll() []entity.Income {
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
	return result
}
