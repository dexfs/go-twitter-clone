package posts

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID                     string
	User                   *user.User
	Content                string
	CreatedAt              time.Time
	IsQuote                bool
	IsRepost               bool
	OriginalPostID         string
	OriginalPostContent    string
	OriginalPostUserID     string
	OriginalPostScreenName string
}

type NewPostInput struct {
	User    *user.User
	Content string
}

type NewRepostQuoteInput struct {
	User    *user.User
	Post    *Post
	Content string
}

func NewPost(aNewPost NewPostInput) *Post {
	return &Post{
		ID:                     uuid.NewString(),
		User:                   aNewPost.User,
		Content:                aNewPost.Content,
		CreatedAt:              time.Now(),
		IsQuote:                false,
		IsRepost:               false,
		OriginalPostID:         "",
		OriginalPostContent:    "",
		OriginalPostUserID:     "",
		OriginalPostScreenName: "",
	}
}

func NewRepost(aRepostInput *NewRepostQuoteInput) (*Post, error) {
	if aRepostInput.Post.IsRepost {
		return nil, errors.New("it is not possible repost a repost post")
	}

	if aRepostInput.Post.User.ID == aRepostInput.User.ID {
		return nil, errors.New("it is not possible repost your own post")
	}

	return &Post{
		ID:                     uuid.NewString(),
		User:                   aRepostInput.User,
		CreatedAt:              time.Now(),
		IsQuote:                false,
		IsRepost:               true,
		OriginalPostID:         aRepostInput.Post.ID,
		OriginalPostContent:    aRepostInput.Post.Content,
		OriginalPostUserID:     aRepostInput.Post.User.ID,
		OriginalPostScreenName: aRepostInput.Post.User.Username,
	}, nil
}

func NewQuote(aNewQuoteInput *NewRepostQuoteInput) (*Post, error) {
	if aNewQuoteInput.Post.IsRepost {
		return nil, errors.New("it is not possible repost a repost post")
	}

	if aNewQuoteInput.Post.User.ID == aNewQuoteInput.User.ID {
		return nil, errors.New("it is not possible repost your own post")
	}

	return &Post{
		ID:                     uuid.NewString(),
		User:                   aNewQuoteInput.User,
		Content:                aNewQuoteInput.Content,
		CreatedAt:              time.Now(),
		IsQuote:                true,
		IsRepost:               false,
		OriginalPostID:         aNewQuoteInput.Post.ID,
		OriginalPostContent:    aNewQuoteInput.Post.Content,
		OriginalPostUserID:     aNewQuoteInput.Post.User.ID,
		OriginalPostScreenName: aNewQuoteInput.Post.User.Username,
	}, nil
}
