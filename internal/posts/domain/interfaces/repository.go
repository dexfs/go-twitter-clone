package interfaces

import postEntity "github.com/dexfs/go-twitter-clone/internal/posts"

type Limit uint64
type ID string
type Posts []*postEntity.Post
type Count uint64
type PostingLimitReached bool
type HasRepost bool

/*
*
@see https://www.codingexplorations.com/blog/mastering-the-repository-pattern-in-go-a-comprehensive-guide
*/
type PostRepository interface {
	GetAll() Posts
	CountByUser(userId ID) Count
	HasRepostByIdAndUserId(postId ID, userId string) HasRepost
	HasReachedPostingLimitDay(postId ID, userId ID, limit Limit) PostingLimitReached
}
