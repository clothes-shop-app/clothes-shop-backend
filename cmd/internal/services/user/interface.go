package user

import "clothes-shop-backend/cmd/internal/models"

type UserRepository interface {
	GetUserByPhone(phone string) (*models.User, error)
	CreateUser(user *models.User) error
}
