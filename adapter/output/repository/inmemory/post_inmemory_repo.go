package inmemory

import (
	"github.com/dexfs/go-twitter-clone/adapter/output/mappers"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/pkg/shared/helpers"
)

type inMemoryPostRepository struct {
	db *database.InMemoryDB[inmemory_schema.PostSchema]
}

func NewInMemoryPostRepository(db *database.InMemoryDB[inmemory_schema.PostSchema]) *inMemoryPostRepository {
	return &inMemoryPostRepository{
		db,
	}
}

func (r *inMemoryPostRepository) CreatePost(aPost *domain.Post) error {
	r.db.Insert(&inmemory_schema.PostSchema{
		ID:                     aPost.ID,
		UserID:                 aPost.UserID,
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

func (r *inMemoryPostRepository) HasReachedPostingLimitDay(aUserId string, aLimit uint64) bool {
	var count = uint64(0)

	for _, currentData := range r.db.GetAll() {
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

func (r *inMemoryPostRepository) AllByUserID(aUserId string) []*domain.Post {
	var feed []*domain.Post
	for _, currentData := range r.db.GetAll() {
		if currentData.UserID == aUserId {
			feed = append(feed, mappers.NewPostMapper().FromPersistence(currentData))
		}
	}

	return feed
}
