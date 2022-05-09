package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	var user User
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	for _, v := range users {
		if v.Id == id {
			user = v
			break
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success user by id",
		"users":    user,
	})
}

// remove index

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	users = func(s []User, index int) []User {
		ret := make([]User, 0)
		ret = append(ret, s[:index]...)
		return append(ret, s[index+1:]...)
	}(users, id-1)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user by id ",
		"users":    users,
	})
}

// // update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	var user User
	id, err := strconv.Atoi(c.Param("id"))
	c.Bind(&user)

	if err != nil {
		return err
	}

	for k, v := range users {
		if v.Id == id {
			user.Id = v.Id
			users[k].Name = user.Name
			users[k].Email = user.Email
			users[k].Password = user.Password
			break
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update user",
		"users":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
