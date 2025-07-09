package user

import (
	"clothes-shop-backend/cmd/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByPhone(phone string) (*models.User, error) {
	query := `
		SELECT * FROM users
		WHERE phone = ?
	`

	var user User
	err := r.db.Get(&user, query, phone)
	if err != nil {
		return nil, err
	}

	return user.FromView(), nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (id, phone, name, avatar_url, address)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, user.ID, user.Phone, user.Name, user.AvatarURL, user.Address)
	return err
}
