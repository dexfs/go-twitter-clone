package interfaces

import postEntity "github.com/dexfs/go-twitter-clone/internal/posts"

type ID string
type Posts []*postEntity.Post
type Count uint64
type PostingLimitReached bool
type HasRepost bool
type Post *postEntity.Post

/*
*
@see https://www.codingexplorations.com/blog/mastering-the-repository-pattern-in-go-a-comprehensive-guide
*/
type PostRepository interface {
	GetAll() Posts
	CountByUser(userId string) Count
	HasPostBeenRepostedByUser(postID string, userID string) HasRepost
	HasReachedPostingLimitDay(userId string, limit uint64) PostingLimitReached
	GetFeedByUserID(userID string) Posts
	Insert(item *postEntity.Post)
}
