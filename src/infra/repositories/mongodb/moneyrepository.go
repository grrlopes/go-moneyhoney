package mongodb

import (
	"context"
	"log"

	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"github.com/grrlopes/go-moneyhoney/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (db *money) Find(limit int64, skip int64, userId primitive.ObjectID) ([]entity.Activity, entity.Count, error) {
	count, err := db.con.CountDocuments(context.TODO(), bson.M{}, options.Count())
	if err != nil {
		log.Println(err)
	}

	var counts entity.Count
	counts.Total_rows = count
	counts.Offset = skip

	var results []entity.Activity

	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "user_id", Value: userId}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "users"},
					{Key: "localField", Value: "user_id"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "user"},
				},
			},
		},
		bson.D{
			{Key: "$unwind",
				Value: bson.D{
					{Key: "path", Value: "$user"},
					{Key: "preserveNullAndEmptyArrays", Value: false},
				},
			},
		},
		bson.D{{Key: "$skip", Value: skip}},
		bson.D{{Key: "$limit", Value: limit}},
	}

	cursor, err := db.con.Aggregate(context.TODO(), pipeline)
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, counts, nil
}

func (db *money) Save(data *entity.Activity) (entity.Income, error) {
	pipeline := bson.D{
		{
			Key: "user_id", Value: data.UserID,
		},
		{
			Key: "item", Value: bson.D{
				{
					Key: "name", Value: data.Item.Name,
				},
				{
					Key: "description", Value: data.Item.Description,
				},
				{
					Key: "amount", Value: data.Item.Amount,
				},
			},
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
