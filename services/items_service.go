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
		log.Println(err.Error())
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

	uid := ParseUid(id)

	resp, _ := is.repository.GetItemById(ctx, uid)

	if resp.UUID != uuid.Nil {
		return resp, nil
	}
	return nil, errors.New("NotFound")
}

func (is *ItemsService) CreateItem(ctx context.Context, item data.Item) {

	item.UUID = uuid.New()

	item.ValidateProperties()

	is.repository.CreateItem(ctx, item)
}

func (is *ItemsService) UpdateItem(ctx context.Context, item data.Item) {

	_, err := is.GetById(ctx, item.UUID.String())
	if err != nil {
		log.Println(err.Error())
	}
	item.ValidateProperties()

	is.repository.UpdateItem(ctx, item)
}

func (is *ItemsService) DeleteItem(ctx context.Context, uid string) {

	_, err := is.GetById(ctx, uid)
	if err != nil {
		log.Println(err.Error())
	}

	is.repository.DeleteItem(ctx, uuid.MustParse(uid))
}

func ParseUid(id string) uuid.UUID {
	uid, err := uuid.Parse(id)
	if err != nil {
		log.Println(err)
	}
	return uid
}
