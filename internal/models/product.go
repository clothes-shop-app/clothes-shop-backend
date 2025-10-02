package models

import "time"

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CategoryID  string    `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	Images      []string  `json:"images"`
}
