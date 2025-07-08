package product

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          string          `db:"id"`
	Name        string          `db:"name"`
	Description string          `db:"description"`
	Price       sql.NullFloat64 `db:"price"`
	CategoryID  sql.NullString  `db:"category_id"`
	CreatedAt   time.Time       `db:"created_at"`
}

type ProductImage struct {
	ID        string    `db:"id"`
	ProductID string    `db:"product_id"`
	Position  string    `db:"position"`
	Image     string    `db:"image_url"`
	CreatedAt time.Time `db:"created_at"`
}
