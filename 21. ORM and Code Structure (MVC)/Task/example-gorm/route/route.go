package route

import (
	"example-gorm/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUserController)

	return e
}
