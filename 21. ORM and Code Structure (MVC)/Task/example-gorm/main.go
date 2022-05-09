package main

import (
	"example-gorm/config"
	"example-gorm/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":9000")
}