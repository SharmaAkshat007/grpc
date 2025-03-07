package main

import (
	"grpc/client/internal/handlers"
	"grpc/client/internal/routes"
	"grpc/client/internal/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	userService, err := services.NewGrpcUserServiceClient()

	if err != nil {
		log.Fatalf("Failed to create user service client: %v", err)
	}

	services := services.NewServices(userService)

	handlers := handlers.NewHandler(*services)

	r := mux.NewRouter()
	routes.InitRoutes(r, handlers)
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
