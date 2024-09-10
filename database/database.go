package database

import (
	"github.com/LENOVO/go-wishlist-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite" // This is the driver we are using
)

var DB *gorm.DB

// connect to the database
func InitDatabase() {
	var err error

	DB, err = gorm.Open(sqlite.Open("go_wishlist_apidb.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Wishlist{})
}
