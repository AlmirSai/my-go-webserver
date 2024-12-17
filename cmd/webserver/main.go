package main

import (
	"log"
	"net/http"
	"my-go-webserver/internal/handler"
	"my-go-webserver/internal/service"
	"my-go-webserver/internal/storage"
)

func main() {
	storageLayer := storage.NewStorage()
	serviceLayer := service.NewService(storageLayer)
	handlerLayer := handler.NewHandler(serviceLayer)

	http.HandleFunc("/", handlerLayer.Home)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
