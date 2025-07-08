package factory

import (
	product "clothes-shop-backend/cmd/internal/repositories/product"

	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	ProductRepository *product.ProductRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		ProductRepository: product.NewProductRepository(db),
	}
}
