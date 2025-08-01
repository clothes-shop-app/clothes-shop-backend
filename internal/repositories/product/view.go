package product

import "clothes-shop-backend/internal/models"

func (p *Product) FromView() *models.Product {
	price := 0
	if p.Price.Valid {
		// Convert decimal dollars to integer cents
		price = int(p.Price.Float64 * 100)
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
