package postgres

import "time"

type UserSchema struct {
	ID        string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
