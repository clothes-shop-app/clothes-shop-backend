package user

import "time"

type User struct {
	ID        string    `db:"id"`
	Phone     string    `db:"phone"`
	Name      string    `db:"name"`
	AvatarURL string    `db:"avatar_url"`
	Address   string    `db:"address"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
