package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username string) *User {
	return &User{
		ID:        uuid.NewString(),
		Username:  username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
