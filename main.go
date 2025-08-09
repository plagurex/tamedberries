package main

import (
	"log"
	"tb/internal/app"
	"tb/internal/handlers"
)

func main() {
	application := app.NewApp()
	application.AddHandlerFunc("/", handlers.HomeHandler(application))
	application.AddHandlerFunc("/about-us", handlers.AboutUsHandler(application))
	application.AddHandlerFunc("/catalog", handlers.CatalogHandler(application))
	application.AddHandlerFunc("/search", handlers.SearchHandler(application))
	application.AddHandlerFunc("/product/", handlers.ProductHandler(application))
	log.Fatal(application.Run("localhost", 8080))
}
