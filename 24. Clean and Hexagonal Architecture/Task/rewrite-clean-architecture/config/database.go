package config

import (
	"belajar-go-echo/database"
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	return database.ConnectDB()
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}

func Start() *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	err = MigrateDB(db)
	if err != nil {
		panic(err)
	}

	return db
}
