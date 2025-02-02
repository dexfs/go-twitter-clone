package seeders

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/postgres"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/jackc/pgx/v5"
	"strconv"
	"time"
)

type PostgresSeeder struct {
	db *database.PostgresDB
}

func NewPostgresSeed(db *database.PostgresDB) *PostgresSeeder {
	return &PostgresSeeder{db: db}
}

func (s *PostgresSeeder) UsersSeed(ctx context.Context) error {
	batch := &pgx.Batch{}
	initialUsers := make([]*postgres.UserSchema, 0)
	initialUsers = append(initialUsers, &postgres.UserSchema{
		ID:        "01JJYY0V9AMD9656HT4BSV0ZEK",
		Username:  "alucard",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	initialUsers = append(initialUsers, &postgres.UserSchema{
		ID:        "01JJYY1S0JY0ERC1VQ3EEFNJC7",
		Username:  "alexander",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	initialUsers = append(initialUsers, &postgres.UserSchema{
		ID:        "01JJYY1Z0E3BMZQ0HFDH8A6NMT",
		Username:  "seras_victoria",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	batch.Queue("DELETE FROM users")
	for _, u := range initialUsers {
		batch.Queue(`INSERT INTO users (user_id, username, created_at, updated_at) VALUES ($1, $2, $3, $4)`,
			u.ID, u.Username, u.CreatedAt, u.UpdatedAt)
	}

	return s.db.Batch(ctx, batch, len(initialUsers))
}

func (s *PostgresSeeder) PostsSeed(ctx context.Context) error {
	batch := &pgx.Batch{}

	users := make([]*domain.User, 0)
	users = append(users, &domain.User{
		ID:       "01JJYY0V9AMD9656HT4BSV0ZEK",
		Username: "alucard",
	})
	users = append(users, &domain.User{
		ID:       "01JJYY1S0JY0ERC1VQ3EEFNJC7",
		Username: "alexander",
	})
	batch.Queue("DELETE FROM posts")
	// Criar 3 posts por user
	// - alucard criar postatgem
	queryInsertPost := `
		INSERT INTO posts (
			post_id, user_id, content, is_quote, is_repost, 
			original_post_id, original_post_content, original_post_user_id, original_post_screen_name, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	var userRepost *domain.User
	for _, user := range users {
		for i := 0; i < 3; i++ {
			aPost, _ := domain.NewPost(domain.NewPostInput{User: user, Content: user.Username + " post " + strconv.Itoa(i) + " from seed"})
			batch.Queue(queryInsertPost, aPost.ID,
				aPost.User.ID,
				aPost.Content,
				aPost.IsQuote,
				aPost.IsRepost,
				aPost.OriginalPostID,
				aPost.OriginalPostContent,
				aPost.OriginalPostUserID,
				aPost.OriginalPostScreenName,
				aPost.CreatedAt)

			if user.Username == "alucard" {
				userRepost = users[1]
			} else {
				userRepost = users[0]
			}

			aRepost, _ := domain.NewRepost(domain.NewRepostQuoteInput{User: userRepost, Post: aPost})
			batch.Queue(queryInsertPost, aRepost.ID,
				aRepost.User.ID,
				aRepost.Content,
				aRepost.IsQuote,
				aRepost.IsRepost,
				aRepost.OriginalPostID,
				aRepost.OriginalPostContent,
				aRepost.OriginalPostUserID,
				aRepost.OriginalPostScreenName,
				aRepost.CreatedAt)

			aQuote, _ := domain.NewQuote(domain.NewRepostQuoteInput{User: userRepost, Post: aPost, Content: userRepost.Username + " quote " + strconv.Itoa(i)})
			batch.Queue(queryInsertPost, aQuote.ID,
				aQuote.User.ID,
				aQuote.Content,
				aQuote.IsQuote,
				aQuote.IsRepost,
				aQuote.OriginalPostID,
				aQuote.OriginalPostContent,
				aQuote.OriginalPostUserID,
				aQuote.OriginalPostScreenName,
				aQuote.CreatedAt)

		}
	}
	err := s.db.Batch(ctx, batch, len(users))
	if err != nil {
		return err
	}

	return nil
}
