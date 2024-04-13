package inmemory

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/domain/interfaces"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/pkg/shared/helpers"
)

type InMemoryPostRepo struct {
	db *database.InMemoryDB[domain.Post]
}

func NewInMemoryPostRepo(db *database.InMemoryDB[domain.Post]) *InMemoryPostRepo {
	return &InMemoryPostRepo{
		db: db,
	}
}

func (r *InMemoryPostRepo) CountByUser(userId string) interfaces.Count {
	count := interfaces.Count(0)
	for _, currentData := range r.db.GetAll() {
		if currentData.User.ID == userId {
			count++
		}
	}

	return count
}

func (r *InMemoryPostRepo) HasPostBeenRepostedByUser(postID string, userID string) interfaces.HasRepost {
	for _, vPost := range r.db.GetAll() {
		if vPost.IsRepost {
			if vPost.User.ID == userID && vPost.OriginalPostID == postID {
				return true
			}
		}
	}
	return false
}

func (r *InMemoryPostRepo) Insert(item *domain.Post) {
	r.db.Insert(item)
}

func (r *InMemoryPostRepo) FindByID(id string) (*domain.Post, error) {
	for _, currentData := range r.db.GetAll() {
		if currentData.ID == id {
			return currentData, nil
		}
	}

	return nil, errors.New("post not found")
}

func (r *InMemoryPostRepo) Remove(item *domain.Post) {
	r.db.Remove(item)
}

func (r *InMemoryPostRepo) GetAll() interfaces.Posts {
	return r.db.GetAll()
}

func (r *InMemoryPostRepo) HasReachedPostingLimitDay(userId string, limit uint64) interfaces.PostingLimitReached {
	var count = uint64(0)

	for _, currentData := range r.db.GetAll() {
		matched := currentData.User.ID == userId && helpers.IsToday(currentData.CreatedAt)

		if matched {
			count++
		}
	}

	reached := count >= limit
	if reached {
		return true
	} else {
		return false
	}
}

func (r *InMemoryPostRepo) GetFeedByUserID(userID string) interfaces.Posts {
	var feed []*domain.Post
	for _, currentData := range r.db.GetAll() {
		if currentData.User.ID == userID {
			feed = append(feed, currentData)
		}
	}

	return feed
}
