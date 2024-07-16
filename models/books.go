package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var DB *gorm.DB

func ConnectDataBase() {
	dsn := "host=localhost user=ombima password=5897 dbname=go_gin_books_api port=5432 sslmode=disable TimeZone=Africa/Nairobi"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database!", err)
	}
	database.AutoMigrate(&Book{})

	DB = database
}
