package main

import (
	"example-middleware/config"
	"example-middleware/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8000")
}
