package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string         `db:"id"`
	Phone     string         `db:"phone"`
	Name      string         `db:"name"`
	AvatarURL sql.NullString `db:"avatar_url"`
	Address   sql.NullString `db:"address"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}
