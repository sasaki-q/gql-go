package config

import (
	"fmt"
	"os"

	"git/example.com/src/model"

	"github.com/google/uuid"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

func itemList() []*model.Item {
	itemList := []*model.Item{}

	for i := 0; i < 10000; i++ {
		v, _ := uuid.NewUUID()

		itemList = append(itemList, &model.Item{
			ID:    v.String(),
			Text:  fmt.Sprintf("test-item-%d", i),
			Price: 1000 * i,
		})
	}

	return itemList
}
