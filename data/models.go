package data

import (
	"github.com/google/uuid"
)

type Item struct {
	UUID        uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required,min=3,max=12"`
	Price       float32   `json:"price" validate:"required,gt=0"`
	Description string    `json:"description"`
	SKU         string    `json:"sku"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

type Itmes []Item
