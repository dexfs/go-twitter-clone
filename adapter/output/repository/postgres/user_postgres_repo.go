package postgres

import (
	"context"
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
)

type PostgresUserRepository struct {
	db *database.PostgresDB
}

func NewPostgresUserRepository(db *database.PostgresDB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) ByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := "SELECT * FROM users WHERE username = $1"
	row := r.db.FindOne(ctx, query, username)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := r.db.FindOne(ctx, query, id)

	var user domain.User
	if err := row.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
