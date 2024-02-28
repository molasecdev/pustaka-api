package config

import (
	"os"
	"pustaka-api/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConfig() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Auth{})
	db.AutoMigrate(&models.User{}, &models.Role{})
	db.AutoMigrate(&models.Author{}, &models.Publisher{}, &models.Bookshelfs{}, &models.Category{}, &models.Language{}, &models.Book{})
	db.AutoMigrate(&models.Loan{})
	db.AutoMigrate(&models.Notification{})

	return db
}
