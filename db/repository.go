package db

import (
	"context"
	"log"

	"github.com/Items.Store.Api/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type iRepositorty interface {
	GetAllItems(ctx context.Context) (data.Itmes, error)
}

type Repository struct {
	db gorm.DB
}

func NewRepository() (*Repository, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=127.0.0.1 user=postgres dbname=postgres password=password123 sslmode=disable",
	}))

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&data.Item{}); err != nil {
		log.Println(err)
	}

	return &Repository{db: *db}, err
}

func (r *Repository) GetAllItems(ctx context.Context) (data.Itmes, error) {
	var users data.Itmes

	r.db.Find(&users)

	return users, nil
}
