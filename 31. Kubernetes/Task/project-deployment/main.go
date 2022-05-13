package main

import (
	"project-deployment/route"
)

func main() {
	// config.InitDB()
	e := route.New()
	e.Start(":9000")
}
