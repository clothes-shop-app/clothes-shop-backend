package models

import "time"

type CartItem struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	ProductID string    `json:"productId"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}

type CartItemPrototype struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
