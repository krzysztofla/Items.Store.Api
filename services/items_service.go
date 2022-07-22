package services

import (
	"context"
	"errors"
	"log"

	"github.com/Items.Store.Api/data"
	"github.com/Items.Store.Api/db"
	"github.com/go-playground/validator"
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
	validate := validator.New()
	err := validate.Struct(&item)
	validationErrors := err.(validator.ValidationErrors)
	log.Fatal(validationErrors)
	is.repository.CreateItem(ctx, item)
}

func (is *ItemsService) UpdateItem(ctx context.Context, item data.Item) {

	_, err := is.GetById(ctx, item.UUID.String())
	if err != nil {
		log.Fatal(err.Error())
	}

	is.repository.UpdateItem(ctx, item)
}
