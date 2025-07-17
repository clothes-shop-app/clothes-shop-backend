package user

import (
	"clothes-shop-backend/cmd/internal/models"
	"database/sql"

	"github.com/google/uuid"
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

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user.FromView(), nil
}

func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	newUUID := uuid.New()

	query := `
		INSERT INTO users (id, phone, name, address)
		VALUES (?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, newUUID, user.Phone, user.Name, user.Address)
	if err != nil {
		return nil, err
	}

	user.ID = newUUID.String()

	return user, err
}
