package inmemory

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"time"
)

type UserSchema struct {
	ID        string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type inMemoryUserRepository struct {
	db *database.InMemoryDB[UserSchema]
}

func NewInMemoryUserRepository(db *database.InMemoryDB[UserSchema]) *inMemoryUserRepository {
	return &inMemoryUserRepository{
		db: db,
	}
}

func (r *inMemoryUserRepository) ByUsername(username string) (*domain.User, error) {
	for _, currentUser := range r.db.GetAll() {
		if currentUser.Username == username {
			return &domain.User{
				ID:        currentUser.ID,
				Username:  currentUser.Username,
				CreatedAt: currentUser.CreatedAt,
				UpdatedAt: currentUser.UpdatedAt,
			}, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *inMemoryUserRepository) FindByID(id string) (*domain.User, error) {
	for _, currentUser := range r.db.GetAll() {
		if currentUser.ID == id {
			return &domain.User{
				ID:        currentUser.ID,
				Username:  currentUser.Username,
				CreatedAt: currentUser.CreatedAt,
				UpdatedAt: currentUser.UpdatedAt,
			}, nil
		}
	}
	return nil, errors.New("user not found")
}
