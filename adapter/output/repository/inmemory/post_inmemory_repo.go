package inmemory

import (
	"github.com/dexfs/go-twitter-clone/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/pkg/shared/helpers"
	"time"
)

type PostSchema struct {
	ID                     string
	UserId                 string
	Content                string
	CreatedAt              time.Time
	IsQuote                bool
	IsRepost               bool
	OriginalPostID         string
	OriginalPostContent    string
	OriginalPostUserID     string
	OriginalPostScreenName string
}

type inMemoryPostRepository struct {
	db *database.InMemoryDB[PostSchema]
}

func NewInMemoryPostRepository(db *database.InMemoryDB[PostSchema]) *inMemoryPostRepository {
	return &inMemoryPostRepository{
		db,
	}
}

func (r *inMemoryPostRepository) CreatePost(aPost *domain.Post) error {
	r.db.Insert(&PostSchema{
		ID:                     aPost.ID,
		UserId:                 aPost.UserID,
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
		matched := currentData.UserId == aUserId && helpers.IsToday(currentData.CreatedAt)

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
