package main

import (
	"log"

	"github.com/HivenOrg/Hiven/start"
)

func main() {

	app := start.BuildApp()

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
