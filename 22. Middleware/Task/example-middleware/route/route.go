package route

import (
	con "example-middleware/constants"
	"example-middleware/controller"
	m "example-middleware/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controller.GetUsersController)
	m.LogMiddleware(e)
	e.POST("/users", controller.CreateUserController)
	e.POST("/login", controller.LoginUserController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(middleware.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controller.GetUsersController)

	eJWT := e.Group("/jwt")
	eJWT.Use(middleware.JWT([]byte(con.SECRET_JWT)))
	eJWT.GET("/users", controller.GetUsersController)

	return e
}
