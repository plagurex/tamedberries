package main

import (
	"log"
	"tb/internal/app"
)

func main() {
	application := app.NewApp()
	log.Fatal(application.Run("localhost", 8080))
}
