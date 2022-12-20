package entity

import (
	"fmt"
	"time"
)

type Income struct {
	Author      string    `json:"author" validate:"required,min=4,max=10"`
	Cost        string    `json:"cost" validate:"gte=1"`
	Description string    `json:"description" validate:"max=200"`
	Email       string    `json:"email" validate:"required,email"`
	Title       string    `json:"title" validate:"required,min=4,max=100"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (i *Income) SetAuthor(author string) error {
	fmt.Println(author)

	i.Author = author

	return nil
}
