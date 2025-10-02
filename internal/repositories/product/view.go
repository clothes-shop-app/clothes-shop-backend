package product

import "clothes-shop-backend/internal/models"

func (p *Product) FromView() *models.Product {
	var price float64
	if p.Price.Valid {
		price = p.Price.Float64
	}

	categoryID := ""
	if p.CategoryID.Valid {
		categoryID = p.CategoryID.String
	}

	return &models.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       price,
		CategoryID:  categoryID,
		CreatedAt:   p.CreatedAt,
		Images:      []string{}, // Initialize empty slice
	}
}
