package post

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID                     string     `json:"id"`
	User                   *user.User `json:"user"`
	Content                string     `json:"content"`
	CreatedAt              time.Time  `json:"created_at"`
	IsQuote                bool       `json:"is_quote"`
	IsRepost               bool       `json:"is_repost"`
	OriginalPostID         string     `json:"original_post_id"`
	OriginalPostContent    string     `json:"original_post_content"`
	OriginalPostUserID     string     `json:"original_post_user_id"`
	OriginalPostScreenName string     `json:"original_post_screen_name"`
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

func NewPost(aNewPost NewPostInput) (*Post, error) {
	if aNewPost.User == nil {
		return nil, errors.New("no user provided")
	}

	if aNewPost.Content == "" {
		return nil, errors.New("no content provided")
	}

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
	}, nil
}

func NewRepost(aRepostInput NewRepostQuoteInput) (*Post, error) {
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

func NewQuote(aNewQuoteInput NewRepostQuoteInput) (*Post, error) {
	if aNewQuoteInput.Post == nil {
		return nil, errors.New("no post provided")
	}

	if aNewQuoteInput.User == nil {
		return nil, errors.New("no user provided")
	}

	if aNewQuoteInput.Content == "" {
		return nil, errors.New("no content provided")
	}

	if aNewQuoteInput.Post.IsQuote {
		return nil, errors.New("it is not possible a quote post of a quote post")
	}

	if aNewQuoteInput.Post.User.ID == aNewQuoteInput.User.ID {
		return nil, errors.New("it is not possible quote your own post")
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
