package user

import "clothes-shop-backend/internal/models"

type UserRepository interface {
	GetUserByPhone(phone string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
}
