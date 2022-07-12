package data

import "github.com/google/uuid"

type Item struct {
	UUID        uuid.UUID `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Price       float32   `json:"price" validate:"required,gt=0"`
	Description string    `json:"description"`
	SKU         string    `json:"sku" validate:"sku"`
	CreatedAt   string    `json:"-"`
	UpdatedAt   string    `json:"-"`
	DeletedAt   string    `json:"-"`
}

type Items []*Items
