package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
)

func main() {
	config.Start()
	app := controller.New()
	app.Start(":8080")
}
