package output

import "github.com/dexfs/go-twitter-clone/internal/core/domain"

type UserPort interface {
	ByUsername(username string) (*domain.User, error)
	FindByID(id string) (*domain.User, error)
}
