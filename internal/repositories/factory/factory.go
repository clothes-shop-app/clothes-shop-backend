package factory

import (
	cart "clothes-shop-backend/internal/repositories/cart"
	product "clothes-shop-backend/internal/repositories/product"
	user "clothes-shop-backend/internal/repositories/user"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ProductRepository *product.ProductRepository
	UserRepository    *user.UserRepository
	CartRepository    *cart.CartRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		ProductRepository: product.NewProductRepository(db),
		UserRepository:    user.NewUserRepository(db),
		CartRepository:    cart.NewCartRepository(db),
	}
}
