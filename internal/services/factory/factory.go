package factory

import (
	repositories "clothes-shop-backend/internal/repositories/factory"
	"clothes-shop-backend/internal/services/cart"
	"clothes-shop-backend/internal/services/product"
	"clothes-shop-backend/internal/services/user"
)

type Services struct {
	ProductService *product.ProductService
	UserService    *user.UserService
	CartService    *cart.CartService
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		ProductService: product.NewProductService(repos.ProductRepository),
		UserService:    user.NewUserService(repos.UserRepository),
		CartService:    cart.NewCartService(repos.CartRepository),
	}
}
