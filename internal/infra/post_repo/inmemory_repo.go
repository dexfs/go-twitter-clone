package post_repo

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/post"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/pkg/shared/helpers"
)

type InMemoryPostRepo struct {
	db *database.InMemoryDB[post.Post]
}

func NewInMemoryPostRepo(db *database.InMemoryDB[post.Post]) *InMemoryPostRepo {
	return &InMemoryPostRepo{
		db: db,
	}
}

func (r *InMemoryPostRepo) CountByUser(userId string) post.Count {
	count := post.Count(0)
	for _, currentData := range r.db.GetAll() {
		if currentData.User.ID == userId {
			count++
		}
	}

	return count
}

func (r *InMemoryPostRepo) HasPostBeenRepostedByUser(postID string, userID string) post.HasRepost {
	for _, vPost := range r.db.GetAll() {
		if vPost.IsRepost {
			if vPost.User.ID == userID && vPost.OriginalPostID == postID {
				return true
			}
		}
	}
	return false
}

func (r *InMemoryPostRepo) Insert(item *post.Post) {
	r.db.Insert(item)
}

func (r *InMemoryPostRepo) FindByID(id string) (*post.Post, error) {
	for _, currentData := range r.db.GetAll() {
		if currentData.ID == id {
			return currentData, nil
		}
	}

	return nil, errors.New("post not found")
}

func (r *InMemoryPostRepo) Remove(item *post.Post) {
	r.db.Remove(item)
}

func (r *InMemoryPostRepo) GetAll() post.Posts {
	return r.db.GetAll()
}

func (r *InMemoryPostRepo) HasReachedPostingLimitDay(userId string, limit uint64) post.PostingLimitReached {
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

func (r *InMemoryPostRepo) GetFeedByUserID(userID string) post.Posts {
	var feed []*post.Post
	for _, currentData := range r.db.GetAll() {
		if currentData.User.ID == userID {
			feed = append(feed, currentData)
		}
	}

	return feed
}
