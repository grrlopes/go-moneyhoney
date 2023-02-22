package mongodb

import (
	"context"
	"log"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (db *money) Find(limit int64, skip int64) ([]entity.Value, error) {
	cursor, err := db.con.Find(context.TODO(), bson.M{},
		options.Find().SetSkip(skip).SetLimit(limit))

	if err != nil {
		log.Println(err)
	}

	var results []entity.Value
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
	}

	return results, nil
}
