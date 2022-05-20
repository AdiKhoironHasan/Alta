package service

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

func GetAllUsersService(users *[]model.User, db *gorm.DB) (*[]model.User, error) {
	err := db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func CreateUserService(user *model.User, db *gorm.DB) (*model.User, error) {
	err := db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
