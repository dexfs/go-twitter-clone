package post

type ID string
type Posts []*Post
type Count uint64
type PostingLimitReached bool
type HasRepost bool

type PostRepository interface {
	GetAll() Posts
	CountByUser(userId string) Count
	HasPostBeenRepostedByUser(postID string, userID string) HasRepost
	HasReachedPostingLimitDay(userId string, limit uint64) PostingLimitReached
	GetFeedByUserID(userID string) Posts
	Insert(item *Post)
	FindByID(id string) (*Post, error)
}
