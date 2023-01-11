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
	Title       string `json:"Title"`
}

type Value struct {
	ID          string    `json:"id"`
	Rev         string    `json:"rev"`
	Author      string    `json:"author" validate:"required,min=4,max=10"`
	Title       string    `json:"title" validate:"required,min=4,max=50"`
	Description string    `json:"description" validate:"max=200"`
	Cost        string    `json:"cost" validate:"gte=1"`
	Email       string    `json:"email" validate:"required,email=200"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type Pagination struct {
	Limit int `json:"limit" validate:"required,gte=1,lte=50"`
	Skip  int `json:"skip" validate:"required,gte=1,number"`
}
