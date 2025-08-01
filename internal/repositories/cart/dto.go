package cart

import "database/sql"

type CartItem struct {
	ID        string          `db:"id"`
	UserID    string          `db:"user_id"`
	ProductID string          `db:"product_id"`
	Name      string          `db:"name"`
	Image     string          `db:"image_url"`
	Price     sql.NullFloat64 `db:"price"`
	Quantity  int             `db:"quantity"`
}

type CartItemPrototype struct {
	ProductID string `db:"product_id"`
	Quantity  int    `db:"quantity"`
}
