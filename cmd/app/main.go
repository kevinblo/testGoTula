package main

import (
	"log"
	"net/http"
	"os"

	"project/internal/api"
	"project/internal/storage"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	store := storage.NewMemoryStorage()
	api.SetupStorage(store)

	router := api.SetupRouter()
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
