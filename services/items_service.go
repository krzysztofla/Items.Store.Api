package services

import (
	"context"
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
	// todo: implement validation here
	return resp
}

func (is *ItemsService) GetById(ctx context.Context, id uuid.UUID) data.Item {
	return data.Item{}
}
