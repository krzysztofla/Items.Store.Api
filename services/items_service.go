package services

import (
	"context"
	"errors"
	"log"

	"github.com/Items.Store.Api/data"
	"github.com/Items.Store.Api/db"
	"github.com/google/uuid"
)

type ItemsService struct {
	repository *db.Repository
}

func NewItemsService() (ItemsService, error) {
	repo, err := db.NewRepository()

	if err != nil {
		log.Fatal(err.Error())
	}

	return ItemsService{
		repository: repo,
	}, err
}

func (is *ItemsService) GetAll(ctx context.Context) data.Itmes {
	resp, _ := is.repository.GetAllItems(ctx)

	return resp
}

func (is *ItemsService) GetById(ctx context.Context, id string) (*data.Item, error) {

	uid, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp, _ := is.repository.GetItemById(ctx, uid)

	if resp.UUID != uuid.Nil {
		return resp, nil
	}
	return nil, errors.New("NotFound")
}

func (is *ItemsService) CreateItem(ctx context.Context, item data.Item) {

	is.repository.CreateItem(ctx, item)
}
