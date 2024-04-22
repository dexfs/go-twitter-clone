package inmemory_schema

import "time"

type PostSchema struct {
	ID                     string
	UserID                 string
	Content                string
	CreatedAt              time.Time
	IsQuote                bool
	IsRepost               bool
	OriginalPostID         string
	OriginalPostContent    string
	OriginalPostUserID     string
	OriginalPostScreenName string
}

type UserSchema struct {
	ID        string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
