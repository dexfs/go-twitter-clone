package inmemory

import (
	"context"
	"errors"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

const USER_SCHEMA_NAME = "users"

type inMemoryUserRepository struct {
	db *database.InMemoryDB
}

func NewInMemoryUserRepository(db *database.InMemoryDB) *inMemoryUserRepository {
	return &inMemoryUserRepository{
		db: db,
	}
}

func (r *inMemoryUserRepository) ByUsername(ctx context.Context, username string) (*domain.User, error) {
	for _, currentUser := range r.getAll() {
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

func (r *inMemoryUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	for _, currentUser := range r.getAll() {
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

func (r *inMemoryUserRepository) getAll() []*inmemory_schema.UserSchema {
	return r.db.GetSchema(USER_SCHEMA_NAME).([]*inmemory_schema.UserSchema)
}
