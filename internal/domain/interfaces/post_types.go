package interfaces

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
)

type ID string
type Posts []*domain.Post
type Count uint64
type PostingLimitReached bool
type HasRepost bool
type Post *domain.Post

type PostRepository interface {
	GetAll() Posts
	CountByUser(userId string) Count
	HasPostBeenRepostedByUser(postID string, userID string) HasRepost
	HasReachedPostingLimitDay(userId string, limit uint64) PostingLimitReached
	GetFeedByUserID(userID string) Posts
	Insert(item *domain.Post)
	FindByID(id string) (*domain.Post, error)
}
