package main

import (
	"patient/router"
)

const url = "mongodb://localhost:27017"

func main() {
	server := router.SetupRouteEcho()
	server.Start(":3000")
}
