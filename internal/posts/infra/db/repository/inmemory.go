package repository

import (
	"errors"
	postEntity "github.com/dexfs/go-twitter-clone/internal/posts"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/pkg/shared/helpers"
)

type InMemoryPostRepo struct {
	db *database.InMemoryDB[postEntity.Post]
}

type Posts []*postEntity.Post

type Post *postEntity.Post

type Count uint64

func NewPostInMemory(db *database.InMemoryDB[postEntity.Post]) *InMemoryPostRepo {
	return &InMemoryPostRepo{
		db: db,
	}
}

func (r *InMemoryPostRepo) CountByUser(userId string) Count {
	count := Count(0)
	for _, currentData := range r.db.GetAll() {
		if currentData.User.ID == userId {
			count++
		}
	}

	return count
}

func (r *InMemoryPostRepo) HasPostBeenRepostedByUser(postID string, userID string) bool {
	for _, vPost := range r.db.GetAll() {
		if vPost.IsRepost {
			if vPost.User.ID == userID && vPost.OriginalPostID == postID {
				return true
			}
		}
	}
	return false
}

func (r *InMemoryPostRepo) Insert(item *postEntity.Post) {
	r.db.Insert(item)
}

func (r *InMemoryPostRepo) FindByID(id string) (*postEntity.Post, error) {
	for _, currentData := range r.db.GetAll() {
		if currentData.ID == id {
			return currentData, nil
		}
	}

	return nil, errors.New("currentUser not found")
}

func (r *InMemoryPostRepo) Remove(item *postEntity.Post) {
	r.db.Remove(item)
}

func (r *InMemoryPostRepo) GetAll() Posts {
	return r.db.GetAll()
}

func (r *InMemoryPostRepo) HasReachedPostingLimitDay(userId string, limit uint64) bool {
	var count = uint64(0)

	for _, currentData := range r.db.GetAll() {
		matched := currentData.User.ID == userId && helpers.IsToday(currentData.CreatedAt)

		if matched {
			count++
		}
	}

	reached := count == limit
	if reached {
		return true
	} else {
		return false
	}
}
