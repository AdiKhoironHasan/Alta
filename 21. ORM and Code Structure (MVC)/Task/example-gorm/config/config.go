package config

import (
	"fmt"

	"example-gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/go_example?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if DB != nil {
		fmt.Println("succesful connect mysql")
	}

	InitMigrate()
}

func InitMigrate() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err.Error())
	}
}
