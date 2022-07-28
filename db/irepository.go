package db

import (
	"context"

	"github.com/Items.Store.Api/data"
	"github.com/google/uuid"
)

type IRepositorty interface {
	NewRepository() (*Repository, error)
	GetAllItems(ctx context.Context) (data.Itmes, error)
	GetItemById(ctx context.Context, uid uuid.UUID) (*data.Item, error)
	CreateItem(ctx context.Context, item data.Item) error
	UpdateItem(ctx context.Context, item data.Item) error
	DeleteItem(ctx context.Context, uid uuid.UUID) error
}
