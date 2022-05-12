package route

import (
	con "logging-and-jwt/constants"
	"logging-and-jwt/controller"
	m "logging-and-jwt/middleware"

	"github.com/labstack/echo"
	mEcho "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	// logging
	m.LogMiddleware(e)

	// no auth
	e.POST("/login", controller.LoginUserController)

	e.POST("/users", controller.CreateUserController)

	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)

	// auth
	eJWT := e.Group("/auth")
	eJWT.Use(mEcho.JWT([]byte(con.SECRET_JWT)))
	e.GET("/users", controller.GetUsersController)
	eJWT.GET("/users/:id", controller.GetUserController)
	eJWT.DELETE("/users/:id", controller.DeleteUserController)
	eJWT.PUT("/users/:id", controller.UpdateUserController)

	eJWT.POST("/books", controller.CreateBookController)
	eJWT.DELETE("/books/:id", controller.DeleteBookController)
	eJWT.PUT("/books/:id", controller.UpdateBookController)

	return e
}
