package entity

import (
	"fmt"
	"time"
)

type Income struct {
	TotalRows int    `json:"total_rows"`
	Offset    int    `json:"offset"`
	Rows      []Rows `json:"rows"`
}

type Rows struct {
	ID    string `json:"id"`
	Key   key    `json:"key"`
	Value value  `json:"value"`
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

type value struct {
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

func (i *Income) SetAuthor(author string) error {
	fmt.Println(author)

	// i.Author = author

	return nil
}
