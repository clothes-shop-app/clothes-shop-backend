package user

import "clothes-shop-backend/cmd/internal/models"

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

func (s *UserService) GetUserByPhone(phone string) (*models.User, error) {
	return s.UserRepository.GetUserByPhone(phone)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepository.CreateUser(user)
}
