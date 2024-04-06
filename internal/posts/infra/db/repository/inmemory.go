package repository

import (
	"errors"
	postEntity "github.com/dexfs/go-twitter-clone/internal/posts"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/pkg/shared/helpers"
)

type PostDb struct {
	db *database.InMemoryDB[postEntity.Post]
}

type Posts []*postEntity.Post

type Post *postEntity.Post

type Count uint64

func NewPostInMemory(db *database.InMemoryDB[postEntity.Post]) *PostDb {
	return &PostDb{
		db: db,
	}
}

func (p *PostDb) CountByUser(userId string) Count {
	count := Count(0)
	for _, currentData := range p.db.GetAll() {
		if currentData.User.ID == userId {
			count++
		}
	}

	return count
}

func (p *PostDb) Insert(item *postEntity.Post) {
	p.db.Insert(item)
}

func (p *PostDb) FindByID(id string) (*postEntity.Post, error) {
	for _, currentData := range p.db.GetAll() {
		if currentData.ID == id {
			return currentData, nil
		}
	}

	return nil, errors.New("currentUser not found")
}

func (p *PostDb) Remove(item *postEntity.Post) {
	p.db.Remove(item)
}

func (p *PostDb) GetAll() Posts {
	return p.db.GetAll()
}

func (p *PostDb) HasReachedPostingLimitDay(userId string, limit uint64) bool {
	var count = uint64(0)

	for _, currentData := range p.db.GetAll() {
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
