package main

import (
	"logging-and-jwt/config"
	"logging-and-jwt/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
