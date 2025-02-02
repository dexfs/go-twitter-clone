package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/jackc/pgx/v5"
)

type postgresPostRepository struct {
	db *database.PostgresDB
}

func NewPostgresPostRepository(db *database.PostgresDB) *postgresPostRepository {
	return &postgresPostRepository{
		db: db,
	}
}

func (r *postgresPostRepository) CreatePost(ctx context.Context, aPost *domain.Post) error {
	query := `
		INSERT INTO posts (
			post_id, user_id, content, is_quote, is_repost, 
			original_post_id, original_post_content, original_post_user_id, original_post_screen_name, created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	return r.db.Insert(ctx, query,
		aPost.ID,
		aPost.User.ID,
		aPost.Content,
		aPost.IsQuote,
		aPost.IsRepost,
		aPost.OriginalPostID,
		aPost.OriginalPostContent,
		aPost.OriginalPostUserID,
		aPost.OriginalPostScreenName,
		aPost.CreatedAt,
	)
}

func (r *postgresPostRepository) HasReachedPostingLimitDay(ctx context.Context, aUserId string, aLimit uint64) bool {
	query := "select count(id) from posts where user_id = $1  AND DATE(created_at) = CURRENT_DATE"
	row := r.db.FindOne(ctx, query, aUserId)
	var count uint64

	err := row.Scan(&count)
	if err != nil {
		return false
	}

	reached := count >= aLimit

	if reached {
		return true
	}

	return false
}

func (r *postgresPostRepository) HasPostBeenRepostedByUser(ctx context.Context, aPostID string, aUserID string) bool {
	query := "select count(id) from posts where user_id = $1  AND original_post_id = $2 AND is_repost = true"
	row := r.db.FindOne(ctx, query, aUserID, aPostID)

	var count uint64

	err := row.Scan(&count)
	if err != nil {
		return false
	}

	if count > 0 {
		return true
	}

	return false
}

func (r *postgresPostRepository) AllByUserID(ctx context.Context, aUser *domain.User) []*domain.Post {
	result := make([]*domain.Post, 0)

	query := `
        SELECT 
            posts.post_id,             
            posts.content,
            posts.is_quote,
            posts.is_repost,
            posts.original_post_id,
            posts.original_post_content,
            posts.original_post_user_id,
            posts.original_post_screen_name,
            users.user_id,
            users.username            
        FROM posts 
        INNER JOIN users ON posts.user_id = users.user_id 
        WHERE posts.user_id = $1`

	rows, err := r.db.Find(ctx, query, aUser.ID)
	defer rows.Close()

	if err != nil {
		return nil
	}

	if !rows.Next() {
		return nil
	}
	for rows.Next() {
		var postSchema PostSchema
		var userSchema UserSchema
		if err := rows.Scan(
			&postSchema.ID,
			&postSchema.Content,
			&postSchema.IsQuote,
			&postSchema.IsRepost,
			&postSchema.OriginalPostID,
			&postSchema.OriginalPostContent,
			&postSchema.OriginalPostUserID,
			&postSchema.OriginalPostScreenName,
			&userSchema.ID,
			&userSchema.Username,
		); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil
			}
			return nil
		}
		user := userSchema.FromPersistence()
		post := postSchema.FromPersistence(user)
		result = append(result, post)
	}

	return result
}

func (r *postgresPostRepository) FindByID(ctx context.Context, aPostID string) (*domain.Post, error) {
	query := `
        SELECT 
            posts.post_id,             
            posts.content,
            posts.is_quote,
            posts.is_repost,
            posts.original_post_id,
            posts.original_post_content,
            posts.original_post_user_id,
            posts.original_post_screen_name,
            users.user_id,
            users.username            
        FROM posts 
        INNER JOIN users ON posts.user_id = users.user_id 
        WHERE posts.post_id = $1`

	row := r.db.FindOne(ctx, query, aPostID)

	var postSchema PostSchema
	var userSchema UserSchema

	if err := row.Scan(
		&postSchema.ID,
		&postSchema.Content,
		&postSchema.IsQuote,
		&postSchema.IsRepost,
		&postSchema.OriginalPostID,
		&postSchema.OriginalPostContent,
		&postSchema.OriginalPostUserID,
		&postSchema.OriginalPostScreenName,
		&userSchema.ID,
		&userSchema.Username,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("post not found")
		}
		return nil, fmt.Errorf("error scanning post: %w", err)

	}

	user := userSchema.FromPersistence()
	post := postSchema.FromPersistence(user)

	return post, nil

}
