package domain

// post
type ID string
type Posts []*Post
type Count uint64
type PostingLimitReached bool
type HasRepost bool

type NewPostInput struct {
	User    *User
	Content string
}

type NewRepostQuoteInput struct {
	User    *User
	Post    *Post
	Content string
}

type PostRepository interface {
	GetAll() Posts
	CountByUser(userId string) Count
	HasPostBeenRepostedByUser(postID string, userID string) HasRepost
	HasReachedPostingLimitDay(userId string, limit uint64) PostingLimitReached
	GetFeedByUserID(userID string) Posts
	Insert(item *Post)
	FindByID(id string) (*Post, error)
}

// user
type UserRepository interface {
	ByUsername(username string) (*User, error)
	FindByID(id string) (*User, error)
}
