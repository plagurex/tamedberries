package main

import (
	"log"
	"tb/internal/app"
	"tb/internal/handlers"
)

func main() {
	application := app.NewApp()
	application.AddHandlerFunc("/", handlers.HomeHandler(application))
	log.Fatal(application.Run("localhost", 8080))
}
