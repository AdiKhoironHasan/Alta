package controller

import (
	"example-middleware/config"
	m "example-middleware/middleware"
	"example-middleware/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}

	c.Bind(&user)

	err := config.DB.Save(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created",
		"data":    user,
	})
}

func LoginUserController(c echo.Context) error {
	var user model.User

	c.Bind(&user)

	err := config.DB.Where("email = ?  AND password = ?", user.Email, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"data":    err.Error(),
		})
	}

	token, err := m.CreateToken(user.ID, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"data":    err.Error(),
		})
	}

	userRes := model.UserRes{user.ID, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    userRes,
	})
}
