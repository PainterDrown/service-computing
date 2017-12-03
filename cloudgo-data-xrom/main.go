package main

import (
	"github.com/painterdrown/service-computing/cloudgo-data/services"
)

func main() {
	server := services.NewServer()
	server.Run(":8080")
}
