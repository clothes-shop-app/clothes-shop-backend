package product

import "clothes-shop-backend/internal/models"

type productRepo interface {
	GetPaginatedProducts(page, limit int) ([]*models.Product, error)
	UploadProduct(req models.UploadProduct) (*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
}
