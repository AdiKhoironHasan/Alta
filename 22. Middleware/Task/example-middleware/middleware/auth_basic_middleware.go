package middleware

import (
	"example-middleware/config"
	"example-middleware/model"

	"github.com/labstack/echo"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var user model.User
	// base64 encryption
	err := config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
