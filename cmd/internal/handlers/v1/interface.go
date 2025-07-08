package v1

import "clothes-shop-backend/cmd/internal/models"

type productService interface {
	GetPaginatedProducts(page, limit int) ([]*models.Product, error)
	UploadProduct(req models.UploadProduct) (*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
}
