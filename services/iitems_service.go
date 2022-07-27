package services

import (
	"context"

	"github.com/Items.Store.Api/data"
)

type IItemsService interface {
	GetAll(ctx context.Context) data.Itmes
	GetById(ctx context.Context, id string) (*data.Item, error)
	CreateItem(ctx context.Context, item data.Item)
	UpdateItem(ctx context.Context, item data.Item)
	DeleteItem(ctx context.Context, uid string)
}
