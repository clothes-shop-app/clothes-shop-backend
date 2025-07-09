package factory

import (
	repositories "clothes-shop-backend/cmd/internal/repositories/factory"
	"clothes-shop-backend/cmd/internal/services/product"
	"clothes-shop-backend/cmd/internal/services/user"
)

type Services struct {
	ProductService *product.ProductService
	UserService    *user.UserService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		ProductService: product.NewProductService(repos.ProductRepository),
		UserService:    user.NewUserService(repos.UserRepository),
	}
}
