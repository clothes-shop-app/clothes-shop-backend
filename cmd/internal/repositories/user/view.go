package user

import "clothes-shop-backend/cmd/internal/models"

func (u *User) FromView() *models.User {
	return &models.User{
		ID:        u.ID,
		Phone:     u.Phone,
		Name:      u.Name,
		AvatarURL: u.AvatarURL.String,
		Address:   u.Address.String,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
