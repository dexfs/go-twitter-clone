package inmemory

import (
	"context"
	"errors"
	"github.com/dexfs/go-twitter-clone/adapter/output/mappers"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/pkg/shared/helpers"
	"log"
)

const POST_SCHEMA_NAME = "posts"

type inMemoryPostRepository struct {
	db *database.InMemoryDB
}

func NewInMemoryPostRepository(db *database.InMemoryDB) *inMemoryPostRepository {
	return &inMemoryPostRepository{
		db,
	}
}

func (r *inMemoryPostRepository) CreatePost(ctx context.Context, aPost *domain.Post) error {
	r.insert(&inmemory_schema.PostSchema{
		ID:                     aPost.ID,
		UserID:                 aPost.User.ID,
		Content:                aPost.Content,
		CreatedAt:              aPost.CreatedAt,
		IsQuote:                aPost.IsQuote,
		IsRepost:               aPost.IsRepost,
		OriginalPostID:         aPost.OriginalPostID,
		OriginalPostContent:    aPost.OriginalPostContent,
		OriginalPostUserID:     aPost.OriginalPostUserID,
		OriginalPostScreenName: aPost.OriginalPostScreenName,
	})
	return nil
}

func (r *inMemoryPostRepository) HasReachedPostingLimitDay(ctx context.Context, aUserId string, aLimit uint64) bool {
	var count = uint64(0)

	for _, currentData := range r.getAll() {
		matched := currentData.UserID == aUserId && helpers.IsToday(currentData.CreatedAt)

		if matched {
			count++
		}
	}

	reached := count >= aLimit
	if reached {
		return true
	} else {
		return false
	}
}

func (r *inMemoryPostRepository) AllByUserID(ctx context.Context, aUser *domain.User) []*domain.Post {
	var feed []*domain.Post
	for _, currentData := range r.getAll() {
		if currentData.UserID == aUser.ID {
			feed = append(feed, mappers.NewPostMapper().FromPersistence(currentData, aUser))
		}
	}

	return feed
}

func (r *inMemoryPostRepository) FindByID(ctx context.Context, aPostID string) (*domain.Post, error) {
	for _, currentData := range r.getAll() {
		if currentData.ID == aPostID {
			var postUser *domain.User
			for _, userData := range r.db.GetSchema(USER_SCHEMA_NAME).([]*inmemory_schema.UserSchema) {
				if currentData.UserID == userData.ID {
					postUser = mappers.NewUserMapper().FromPersistence(userData)
				}
			}
			return mappers.NewPostMapper().FromPersistence(currentData, postUser), nil
		}
	}

	return nil, errors.New("post not found")
}

func (r *inMemoryPostRepository) HasPostBeenRepostedByUser(ctx context.Context, postID string, userID string) bool {
	for _, vPost := range r.getAll() {
		if vPost.IsRepost {
			if vPost.UserID == userID && vPost.OriginalPostID == postID {
				return true
			}
		}
	}

	return false
}

func (r *inMemoryPostRepository) getAll() []*inmemory_schema.PostSchema {
	return r.db.GetSchema(POST_SCHEMA_NAME).([]*inmemory_schema.PostSchema)
}

func (r *inMemoryPostRepository) insert(newItem interface{}) {
	existing, ok := r.db.Schemas[POST_SCHEMA_NAME].([]*inmemory_schema.PostSchema)
	if !ok {
		log.Fatal("schema " + POST_SCHEMA_NAME + " not found")
		return
	}

	updateSlice := append(existing, newItem.(*inmemory_schema.PostSchema))
	r.db.Schemas[POST_SCHEMA_NAME] = updateSlice
}
