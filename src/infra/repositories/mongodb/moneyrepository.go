package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type money struct {
	con *mongo.Collection
}

func NewMoneyRepository() repository.IMongoRepo {
	err := OpenDB()
	if err != nil {
		panic(err)
	}

	db := GetDBCollection("moneydata")

	return &money{
		con: db,
	}
}

func (db *money) Find(limit int, skip int) (entity.Income, error) {
	cursor, err := db.con.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
	}

	var results []entity.Value
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	var result entity.Income

	return result, nil
}
