package interfaces

import (
	domain_user "github.com/dexfs/go-twitter-clone/internal/user"
)

type UserRepository interface {
	ByUsername(username string) domain_user.User
}
