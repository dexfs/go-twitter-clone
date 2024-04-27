package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username string) *User {
	return &User{
		ID:        uuid.NewString(),
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
