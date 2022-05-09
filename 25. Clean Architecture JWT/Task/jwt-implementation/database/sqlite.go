package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
}
