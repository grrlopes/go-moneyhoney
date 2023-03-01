package repository

import (
	"github.com/grrlopes/go-moneyhoney/src/domain/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataMap map[string]string
type UpdateMap map[string]interface{}

type IMoneyRepo interface {
	FindAll(limit, skip int) (entity.Income, error)
	FindById(ids *entity.ById) (entity.Income, error)
	Save(data *entity.Value) (entity.Income, error)
	Update(id primitive.ObjectID, data UpdateMap) (entity.Income, error)
	// Delete(money entity.Income) (entity.Income, error)
}

type IMongoRepo interface {
	Find(limit, skip int64) ([]entity.Activity, entity.Count, error)
}
