package cart

import "clothes-shop-backend/internal/models"

type CartRepository interface {
	AddOneToCart(userID string, productID string) error
	GetCartItemsByUserID(userID string) ([]*models.CartItem, error)
	UpdateCart(userID string, products []*models.CartItemPrototype) error
}
