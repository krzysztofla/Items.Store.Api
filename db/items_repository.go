package db

import (
	"context"
	"log"

	"github.com/Items.Store.Api/data"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db gorm.DB
}

func NewRepository() (*Repository, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=postgresdb user=postgres dbname=postgres password=password123 sslmode=disable",
	}))

	if err != nil {
		log.Println(err.Error())
	}

	if err = db.AutoMigrate(&data.Item{}); err != nil {
		log.Println(err)
	}

	return &Repository{db: *db}, err
}

func (r *Repository) GetAllItems(ctx context.Context) (data.Itmes, error) {
	var users []data.Item

	r.db.WithContext(ctx)
	r.db.Find(&users)

	return users, nil
}

func (r *Repository) GetItemById(ctx context.Context, uid uuid.UUID) (*data.Item, error) {
	var item data.Item

	r.db.WithContext(ctx)
	r.db.Where("UUID = ?", uid).First(&item)
	return &item, nil
}

func (r *Repository) CreateItem(ctx context.Context, item data.Item) error {

	r.db.WithContext(ctx)
	er := r.db.Create(&item).Error
	if er != nil {
		log.Println(er)
		return er
	}
	return nil
}

func (r *Repository) UpdateItem(ctx context.Context, item data.Item) error {
	r.db.WithContext(ctx)
	er := r.db.Where("UUID = ?", item.UUID).Save(&item).Error
	if er != nil {
		log.Println(er)
		return er
	}
	return nil
}

func (r *Repository) DeleteItem(ctx context.Context, uid uuid.UUID) error {
	r.db.WithContext(ctx)
	er := r.db.Delete(&data.Item{}, uid).Error
	if er != nil {
		log.Println(er)
		return er
	}
	return nil
}
