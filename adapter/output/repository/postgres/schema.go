package postgres

import (
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"time"
)

type UserSchema struct {
	ID        string    `db:"user_id"`
	Username  string    `db:"username"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (u UserSchema) FromPersistence() *domain.User {
	return &domain.User{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type PostSchema struct {
	ID                     string    `db:"post_id"`
	UserID                 string    `db:"user_id"`
	Content                string    `db:"content"`
	CreatedAt              time.Time `db:"created_at"`
	IsQuote                bool      `db:"is_quote"`
	IsRepost               bool      `db:"is_repost"`
	OriginalPostID         string    `db:"original_post_id"`
	OriginalPostContent    string    `db:"original_post_content"`
	OriginalPostUserID     string    `db:"original_post_user_id"`
	OriginalPostScreenName string    `db:"original_post_screen_name"`
}

func (p *PostSchema) FromPersistence(aUser *domain.User) *domain.Post {
	return &domain.Post{
		ID:                     p.ID,
		User:                   aUser,
		Content:                p.Content,
		CreatedAt:              p.CreatedAt,
		IsQuote:                p.IsQuote,
		IsRepost:               p.IsRepost,
		OriginalPostID:         p.OriginalPostID,
		OriginalPostContent:    p.OriginalPostContent,
		OriginalPostUserID:     p.OriginalPostUserID,
		OriginalPostScreenName: p.OriginalPostScreenName,
	}
}
