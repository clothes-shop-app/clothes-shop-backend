package factory

import (
	product "clothes-shop-backend/cmd/internal/repositories/product"
	user "clothes-shop-backend/cmd/internal/repositories/user"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ProductRepository *product.ProductRepository
	UserRepository    *user.UserRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		ProductRepository: product.NewProductRepository(db),
		UserRepository:    user.NewUserRepository(db),
	}
}
