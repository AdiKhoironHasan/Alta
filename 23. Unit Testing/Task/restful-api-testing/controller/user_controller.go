package controller

import (
	"logging-and-jwt/config"
	m "logging-and-jwt/middleware"
	"logging-and-jwt/model"
	"net/http"

	"github.com/labstack/echo"
)

type UserResponse struct {
	Message string
	Data    []model.User
}

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

func GetUserController(c echo.Context) error {
	user := model.User{}

	id := c.Param("id")
	// fmt.Println("idddd=", id)
	err := config.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by id",
		"data":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")

	err := config.DB.Delete(&model.User{}, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
		"data":    nil,
	})
}

func UpdateUserController(c echo.Context) error {
	user := model.User{}
	userUpdate := model.User{}
	id := c.Param("id")

	err := c.Bind(&userUpdate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	res := config.DB.Where("id = ?", id).First(&user)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": res.Error,
			"data":    nil,
		})
	}

	// err := DB.Model(User{}).Where("id = ?", id).Updates(userUpdate).Error
	res = res.Updates(userUpdate)
	if res.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": res.Error,
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
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
