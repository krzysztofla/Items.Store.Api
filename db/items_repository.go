package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Items.Store.Api/data"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db gorm.DB
}

func NewRepository() (*Repository, error) {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: psqlInfo,
	}))

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if err = db.AutoMigrate(&data.Item{}); err != nil {
		log.Println(err)
		return nil, err
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
