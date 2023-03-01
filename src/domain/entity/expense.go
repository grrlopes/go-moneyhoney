package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Activity struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"-" bson:"user_id"`
	User      Users              `json:"user" bson:"user"`
	Item      Items              `json:"item" bson:"item"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
