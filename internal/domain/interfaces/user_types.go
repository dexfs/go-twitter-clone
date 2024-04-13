package interfaces

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
)

type UserRepository interface {
	ByUsername(username string) (domain.User, error)
	FindByID(id string) (domain.User, error)
}
