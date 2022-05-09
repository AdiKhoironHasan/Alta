package config

import (
	"fmt"

	"logging-and-jwt/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DB_USER = "root"
const DB_PASS = ""
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"
const DB_NAME = "go_example"

const DB_USER_TEST = "root"
const DB_PASS_TEST = ""
const DB_HOST_TEST = "127.0.0.1"
const DB_PORT_TEST = "3306"
const DB_NAME_TEST = "go_example_test"

func InitDB() {
	var err error
	// dsn := "root:@tcp(127.0.0.1:3306)/go_example?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
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

	err = DB.AutoMigrate(&model.Book{})
	if err != nil {
		panic(err.Error())
	}
}

func InitDBTest() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USER_TEST, DB_PASS_TEST, DB_HOST_TEST, DB_PORT_TEST, DB_NAME_TEST)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if DB != nil {
		fmt.Println("succesful connect mysql")
	}

	InitMigrateTest()
}

func InitMigrateTest() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err.Error())
	}

	err = DB.AutoMigrate(&model.Book{})
	if err != nil {
		panic(err.Error())
	}
}
