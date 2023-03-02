package mongodb

import (
	"context"
	"log"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type users struct {
	con *mongo.Collection
}

func NewUserRepository() repository.IMongoRepo {
	err := OpenDB()
	if err != nil {
		panic(err)
	}

	db := GetDBCollection("users")

	return &money{
		con: db,
	}
}

func (db *money) UserSave(data *entity.Users) (entity.Income, error) {
	pipeline := bson.D{
		{
			Key: "author", Value: data.Author,
		},
		{
			Key: "email", Value: data.Email,
		},
		{
			Key: "created_at", Value: data.CreatedAt,
		},
		{
			Key: "updated_at", Value: data.UpdatedAt,
		},
	}

	_, err := db.con.InsertOne(context.TODO(), pipeline)
	if err != nil {
		log.Println(err)
		return entity.Income{}, err
	}

	var result entity.Income
	result.Reason = "created!"

	return result, err
}
