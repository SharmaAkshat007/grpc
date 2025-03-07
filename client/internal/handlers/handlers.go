package handlers

import "grpc/client/internal/services"

type Handlers struct {
	services services.Services
}

func NewHandler(services services.Services) *Handlers {
	return &Handlers{services: services}
}
