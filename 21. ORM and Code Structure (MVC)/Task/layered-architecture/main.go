package main

import (
	"layered-architecture/config"
	"layered-architecture/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
