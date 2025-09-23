package v1

import "clothes-shop-backend/internal/models"

type productService interface {
	GetPaginatedProducts(page, limit int) ([]*models.Product, error)
	UploadProduct(req models.UploadProduct) (*models.Product, error)
	GetProductByID(id string) (*models.Product, error)
}

type userService interface {
	GetUserByPhone(phone string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}

type cartService interface {
	GetCartItemsByUserID(userID string) ([]*models.CartItem, error)
	UpdateCart(userID string, products []*models.CartItemPrototype) error
	AddOneToCart(userID string, productID string) error
	Checkout(userID string) error
}
