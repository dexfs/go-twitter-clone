package output

import "github.com/dexfs/go-twitter-clone/internal/core/domain"

type PostPort interface {
	CreatePost(aPost *domain.Post) error
	HasReachedPostingLimitDay(aUserId string, aLimit uint64) bool
	HasPostBeenRepostedByUser(postID string, userID string) bool
	AllByUserID(aUser *domain.User) []*domain.Post
	FindByID(aPostID string) (*domain.Post, error)
}
