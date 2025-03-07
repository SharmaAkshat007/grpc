package routes

import (
	"grpc/client/internal/handlers"

	"github.com/gorilla/mux"
)

func InitRoutes(r *mux.Router, handlers *handlers.Handlers) {
	r.HandleFunc("/user/{userId}", handlers.GetUser).Methods("GET")
}
