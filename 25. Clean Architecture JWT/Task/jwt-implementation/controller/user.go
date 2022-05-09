package controller

import (
	"belajar-go-echo/config"
	m "belajar-go-echo/middleware"
	"belajar-go-echo/model"
	"belajar-go-echo/service"
	"net/http"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	app := echo.New()
	app.GET("/users", GetAllUsers, m.JWTMiddleware())
	app.POST("/users", CreateUser)
	app.POST("/login", Login)

	return app
}

func GetAllUsers(c echo.Context) error {
	users := make([]model.User, 0)
	// var err ?error

	res, err := service.GetAllUsersService(&users, config.DB)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": res,
	})
}

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	res, err := service.CreateUserService(&user, config.DB)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": res,
	})
}

func Login(c echo.Context) error {
	var user model.User
	var token string

	c.Bind(&user)

	err := config.DB.Where("email = ?  AND password = ?", user.Email, user.Password).First(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"data":    err.Error(),
		})
	}

	token, err = m.CreateTokenJWT()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"data":    err.Error(),
		})
	}

	userRes := model.UserRes{
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"data":    userRes,
	})
}
