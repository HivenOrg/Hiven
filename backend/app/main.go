package main

import (
	"log"

	"github.com/HivenOrg/Hiven/start"
)

func main() {

	server := start.Server()

	err := server.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
