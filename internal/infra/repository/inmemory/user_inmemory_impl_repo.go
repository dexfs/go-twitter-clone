package repo_inmemory

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

type InMemoryUserRepoImpl struct {
	db *database.InMemoryDB[user.User]
}

func NewInMemoryUserRepo(db *database.InMemoryDB[user.User]) *InMemoryUserRepoImpl {
	return &InMemoryUserRepoImpl{
		db: db,
	}
}

func (r *InMemoryUserRepoImpl) ByUsername(username string) (*user.User, error) {
	for _, currentUser := range r.db.GetAll() {
		if currentUser.Username == username {
			return currentUser, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepoImpl) Insert(item *user.User) {
	r.db.Insert(item)
}

func (r *InMemoryUserRepoImpl) GetAll() []*user.User {
	return r.db.GetAll()
}

func (r *InMemoryUserRepoImpl) Remove(item *user.User) {
	r.db.Remove(item)
}

func (r *InMemoryUserRepoImpl) FindByID(id string) (*user.User, error) {
	for _, currentUser := range r.db.GetAll() {
		if currentUser.ID == id {
			return currentUser, nil
		}
	}
	return nil, errors.New("user not found")
}
