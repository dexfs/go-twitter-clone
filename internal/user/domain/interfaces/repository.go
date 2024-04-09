package interfaces

import (
	userEntity "github.com/dexfs/go-twitter-clone/internal/user"
)

type UserRepository interface {
	ByUsername(username string) (*userEntity.User, error)
	FindByID(id string) (*userEntity.User, error)
}
