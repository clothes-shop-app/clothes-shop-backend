package cart

import (
	"clothes-shop-backend/internal/models"
)

func (c *CartItem) ToView() *models.CartItem {
	return &models.CartItem{
		ID:        c.ID,
		UserID:    c.UserID,
		ProductID: c.ProductID,
		Quantity:  c.Quantity,
		Name:      c.Name,
		Price:     int(c.Price.Float64),
		Image:     c.Image,
	}
}

func FromView(item *models.CartItemPrototype) *CartItemPrototype {
	return &CartItemPrototype{
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
	}
}
