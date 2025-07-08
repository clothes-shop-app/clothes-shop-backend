package product

import (
	"clothes-shop-backend/cmd/internal/models"
)

type ProductService struct {
	ProductRepository productRepo
}

func NewProductService(productRepository productRepo) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (s *ProductService) GetPaginatedProducts(page, limit int) ([]*models.Product, error) {
	return s.ProductRepository.GetPaginatedProducts(page, limit)
}

func (s *ProductService) UploadProduct(req models.UploadProduct) (*models.Product, error) {
	return s.ProductRepository.UploadProduct(req)
}

func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.ProductRepository.GetProductByID(id)
}
