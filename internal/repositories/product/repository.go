package product

import (
	"clothes-shop-backend/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetPaginatedProducts(page, limit int) ([]*models.Product, error) {
	query := `
		SELECT * FROM products
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`

	var products []*Product
	err := r.db.Select(&products, query, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}

	response := make([]*models.Product, len(products))

	for i := range products {
		response[i] = products[i].FromView()
	}

	// Get product images for all products
	if len(products) > 0 {
		query = `
			SELECT * FROM product_images
			WHERE product_id IN (?)
		`

		var images []*ProductImage
		pids := make([]string, 0, len(products))
		for i := range products {
			pids = append(pids, products[i].ID)
		}

		// sqlx.In returns the expanded query and arguments
		expandedQuery, args, err := sqlx.In(query, pids)
		if err != nil {
			return nil, err
		}

		// Use the expanded query and arguments
		err = r.db.Select(&images, expandedQuery, args...)
		if err != nil {
			return nil, err
		}

		// Group images by product ID for efficient lookup
		imagesByProduct := make(map[string][]string)
		for _, img := range images {
			imagesByProduct[img.ProductID] = append(imagesByProduct[img.ProductID], img.Image)
		}

		// Assign images to products
		for i := range response {
			if productImages, exists := imagesByProduct[response[i].ID]; exists {
				response[i].Images = productImages
			}
		}
	}

	return response, nil
}

func (r *ProductRepository) GetProductByID(id string) (*models.Product, error) {
	query := `
		SELECT * FROM products
		WHERE id = ?
	`

	var product Product
	err := r.db.Get(&product, query, id)
	if err != nil {
		return nil, err
	}

	// Get product images
	imageQuery := `
		SELECT * FROM product_images
		WHERE product_id = ?
		ORDER BY position ASC
	`

	var images []*ProductImage
	err = r.db.Select(&images, imageQuery, id)
	if err != nil {
		return nil, err
	}

	// Convert to model
	result := product.FromView()

	// Add images to the result
	for _, img := range images {
		result.Images = append(result.Images, img.Image)
	}

	return result, nil
}

func (r *ProductRepository) UploadProduct(req models.UploadProduct) (*models.Product, error) {
	query := `
		INSERT INTO products (id, name, description, price, category_id)
		VALUES (?, ?, ?, ?, ?)
	`

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	newUUID := uuid.New().String()

	var categoryID interface{}
	if req.CategoryID == "" {
		categoryID = nil
	} else {
		categoryID = req.CategoryID
	}

	_, err = tx.Exec(query, newUUID, req.Name, req.Description, req.Price, categoryID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	query = `
		INSERT INTO product_images (product_id, image_url)
		VALUES (?, ?)
	`

	for _, image := range req.Images {
		tx.MustExec(query, newUUID, image)
	}

	tx.Commit()

	product, err := r.GetProductByID(newUUID)
	if err != nil {
		return nil, err
	}

	return product, nil
}
