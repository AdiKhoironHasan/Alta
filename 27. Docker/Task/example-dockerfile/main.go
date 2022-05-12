package main

import (
	"example-gorm/route"
)

func main() {
	// config.InitDB()
	e := route.New()
	e.Start(":9000")
}
