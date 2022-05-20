package config

import (
	"belajar-go-echo/database"
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	return database.ConnectDB()
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}

func Start() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	DB = db

	err = MigrateDB(DB)
	if err != nil {
		panic(err)
	}
}
