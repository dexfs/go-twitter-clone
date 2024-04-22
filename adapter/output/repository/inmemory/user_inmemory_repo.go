package inmemory

import (
	"errors"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

type inMemoryUserRepository struct {
	db *database.InMemoryDB[inmemory_schema.UserSchema]
}

func NewInMemoryUserRepository(db *database.InMemoryDB[inmemory_schema.UserSchema]) *inMemoryUserRepository {
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
