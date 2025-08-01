package cart

import (
	"clothes-shop-backend/internal/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) AddOneToCart(userID string, productID string) error {
	_, err := r.db.Exec("INSERT INTO cart (user_id, product_id, quantity) VALUES (?, ?, 1) ON DUPLICATE KEY UPDATE quantity = cart.quantity + 1", userID, productID)
	return err
}

func (r *CartRepository) UpdateCart(userID string, products []*models.CartItemPrototype) error {
	err := r.ClearCart(userID)
	if err != nil {
		return err
	}

	if len(products) == 0 {
		return nil
	}

	prototypes := make([]*CartItemPrototype, len(products))
	for i, product := range products {
		prototypes[i] = FromView(product)
	}

	placeholders := strings.Repeat("(?, ?, ?),", len(products))
	placeholders = strings.TrimSuffix(placeholders, ",")

	query := `
		INSERT INTO cart (user_id, product_id, quantity) VALUES ` + placeholders + `
		ON DUPLICATE KEY UPDATE quantity = cart.quantity + VALUES(quantity)
	`

	var values []interface{}
	for _, proto := range prototypes {
		values = append(values, userID, proto.ProductID, proto.Quantity)
	}

	_, err = r.db.Exec(query, values...)

	return err
}

func (r *CartRepository) GetCartItemsByUserID(userID string) ([]*models.CartItem, error) {
	var cartItems []*CartItem
	err := r.db.Select(&cartItems, `
		SELECT 
			c.id, c.user_id, c.product_id, c.quantity, p.name, p.price,
			(
				SELECT pi.image_url
				FROM product_images pi
				WHERE pi.product_id = p.id
				ORDER BY pi.position ASC
				LIMIT 1
			) AS image_url
		FROM cart c
		LEFT JOIN products p ON c.product_id = p.id
		WHERE c.user_id = ?
	`, userID)

	items := make([]*models.CartItem, len(cartItems))
	for i, item := range cartItems {
		items[i] = item.ToView()
	}

	return items, err
}

func (r *CartRepository) ClearCart(userID string) error {
	_, err := r.db.Exec("DELETE FROM cart WHERE user_id = ?", userID)
	return err
}
