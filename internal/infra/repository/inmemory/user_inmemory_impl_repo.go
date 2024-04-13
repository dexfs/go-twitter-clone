package inmemory

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

type InMemoryUserRepoImpl struct {
	db *database.InMemoryDB[domain.User]
}

func NewInMemoryUserRepo(db *database.InMemoryDB[domain.User]) *InMemoryUserRepoImpl {
	return &InMemoryUserRepoImpl{
		db: db,
	}
}

func (r *InMemoryUserRepoImpl) ByUsername(username string) (*domain.User, error) {
	for _, currentUser := range r.db.GetAll() {
		if currentUser.Username == username {
			return currentUser, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *InMemoryUserRepoImpl) Insert(item *domain.User) {
	r.db.Insert(item)
}

func (r *InMemoryUserRepoImpl) GetAll() []*domain.User {
	return r.db.GetAll()
}

func (r *InMemoryUserRepoImpl) Remove(item *domain.User) {
	r.db.Remove(item)
}

func (r *InMemoryUserRepoImpl) FindByID(id string) (*domain.User, error) {
	for _, currentUser := range r.db.GetAll() {
		if currentUser.ID == id {
			return currentUser, nil
		}
	}
	return nil, errors.New("user not found")
}
