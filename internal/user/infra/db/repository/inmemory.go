package repository

import (
	"errors"
	userEntity "github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

type InMemoryUserRepo struct {
	db *database.InMemoryDB[userEntity.User]
}

func NewInMemoryUserRepo(db *database.InMemoryDB[userEntity.User]) *InMemoryUserRepo {
	return &InMemoryUserRepo{
		db: db,
	}
}

func (r *InMemoryUserRepo) ByUsername(username string) (*userEntity.User, error) {
	for _, currentUser := range r.db.GetAll() {
		if currentUser.Username == username {
			return currentUser, nil
		}
	}

	return nil, errors.New("currentUser not found")
}

func (r *InMemoryUserRepo) Insert(item *userEntity.User) {
	r.db.Insert(item)
}

func (r *InMemoryUserRepo) GetAll() []*userEntity.User {
	return r.db.GetAll()
}

func (r *InMemoryUserRepo) Remove(item *userEntity.User) {
	r.db.Remove(item)
}
