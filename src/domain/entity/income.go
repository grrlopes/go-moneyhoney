package entity

import (
	"time"
)

type Income struct {
	TotalRows int    `json:"total_rows"`
	Offset    int    `json:"offset"`
	Rows      []rows `json:"rows"`
	Error     string `json:"error"`
	Reason    string `json:"reason"`
}

type rows struct {
	ID    string `json:"id"`
	Key   key    `json:"key"`
	Value Value  `json:"value"`
}

type key struct {
	ID          string `json:"_id"`
	Rev         string `json:"_rev"`
	Author      string `json:"Author"`
	Cost        string `json:"Cost"`
	Description string `json:"Description"`
	Email       string `json:"Email"`
}

type Value struct {
	ID        string    `json:"id" bson:"_id"`
	Rev       string    `json:"rev"`
	Author    string    `json:"author" validate:"required,min=4,max=10" bson:"author"`
	Email     string    `json:"email" validate:"required,email" bson:"email"`
	Item      Items     `json:"item" bson:"item"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Items struct {
	Store       string `json:"store" validate:"required,min=4,max=10" bson:"store"`
	Description string `json:"description" validate:"required,min=4,max=30" bson:"description"`
	Cost        string `json:"cost" validate:"required,gte=1" bson:"cost"`
}

type Pagination struct {
	Limit int `json:"limit" validate:"required,gte=1,lte=50,numeric" bson:"limit"`
	Skip  int `json:"skip" validate:"numeric" bson:"skip"`
}

type Count struct {
	Total_rows int64 `json:"total_rows" validate:"numeric" bson:"total_rows"`
	Offset     int64 `json:"offset" validate:"numeric" bson:"offset"`
}

type ById struct {
	ID  string `form:"id" validate:"required"`
	Rev string `form:"rev"`
}
