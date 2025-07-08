package factory

import (
	v1 "clothes-shop-backend/cmd/internal/handlers/v1"
	services "clothes-shop-backend/cmd/internal/services/factory"
)

type Handlers struct {
	V1Handlers *v1.V1Handlers
}

func NewHandlers(services *services.Services) *Handlers {
	v1Handlers := v1.NewV1Handlers(services.ProductService)

	return &Handlers{
		V1Handlers: v1Handlers,
	}
}
