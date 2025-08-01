package cart

import "clothes-shop-backend/internal/models"

type CartService struct {
	CartRepository CartRepository
}

func NewCartService(cartRepository CartRepository) *CartService {
	return &CartService{CartRepository: cartRepository}
}

func (s *CartService) GetCartItemsByUserID(userID string) ([]*models.CartItem, error) {
	return s.CartRepository.GetCartItemsByUserID(userID)
}

func (s *CartService) UpdateCart(userID string, products []*models.CartItemPrototype) error {
	return s.CartRepository.UpdateCart(userID, products)
}

func (s *CartService) AddOneToCart(userID string, productID string) error {
	return s.CartRepository.AddOneToCart(userID, productID)
}
