package database

import (
	"go-wishlist-api-main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect to the database
func InitDatabase() {
	var err error

	DB, err = gorm.Open(mysql.Open("go_wishlist_apidb.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Wishlist{})
}
