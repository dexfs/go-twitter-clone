package output

import (
	"context"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type UserPort interface {
	ByUsername(ctx context.Context, username string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
}
