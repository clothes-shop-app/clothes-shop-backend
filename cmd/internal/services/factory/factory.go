package factory

import (
	repositories "clothes-shop-backend/cmd/internal/repositories/factory"
	"clothes-shop-backend/cmd/internal/services/product"
)

type Services struct {
	ProductService *product.ProductService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		ProductService: product.NewProductService(repos.ProductRepository),
	}
}
