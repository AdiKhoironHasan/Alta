package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	app := echo.New()
	app.GET("/users", GetAllUsers(db))
	app.POST("/users", CreateUser(db))

	return app
}

func GetAllUsers(db *gorm.DB) echo.HandlerFunc {
	users := make([]model.User, 0)
	return func(c echo.Context) error {
		users, err := service.GetAllUsersService(&users, db)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func CreateUser(db *gorm.DB) echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}
		user, err := service.CreateUserService(&user, db)
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}
