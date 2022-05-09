package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
)

func main() {
	db := config.Start()
	app := controller.New(db)
	app.Start(":8080")
}
