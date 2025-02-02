package output

import (
	"context"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type PostPort interface {
	CreatePost(ctx context.Context, aPost *domain.Post) error
	HasReachedPostingLimitDay(ctx context.Context, aUserId string, aLimit uint64) bool
	HasPostBeenRepostedByUser(ctx context.Context, aPostID string, aUserID string) bool
	AllByUserID(ctx context.Context, aUser *domain.User) []*domain.Post
	FindByID(ctx context.Context, aPostID string) (*domain.Post, error)
}
