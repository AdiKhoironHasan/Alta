package route

import (
	"project-deployment/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", controller.Home)
	return e
}
