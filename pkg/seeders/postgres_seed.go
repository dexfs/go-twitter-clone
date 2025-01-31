package seeders

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/postgres"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/jackc/pgx/v5"
	"time"
)

type PostgresSeeder struct {
	db *database.PostgresDB
}

func NewPostgresSeed(db *database.PostgresDB) *PostgresSeeder {
	return &PostgresSeeder{db: db}
}

func (s *PostgresSeeder) UsersSeed(ctx context.Context) error {
	batch := &pgx.Batch{}
	initialUsers := make([]*postgres.UserSchema, 0)
	initialUsers = append(initialUsers, &postgres.UserSchema{
		ID:        "01JJYY0V9AMD9656HT4BSV0ZEK",
		Username:  "alucard",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	initialUsers = append(initialUsers, &postgres.UserSchema{
		ID:        "01JJYY1S0JY0ERC1VQ3EEFNJC7",
		Username:  "alexander",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	initialUsers = append(initialUsers, &postgres.UserSchema{
		ID:        "01JJYY1Z0E3BMZQ0HFDH8A6NMT",
		Username:  "seras_victoria",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	batch.Queue("DELETE FROM users")
	for _, u := range initialUsers {
		batch.Queue(`INSERT INTO users (id, username, created_at, updated_at) VALUES ($1, $2, $3, $4)`,
			u.ID, u.Username, u.CreatedAt, u.UpdatedAt)
	}

	return s.db.Batch(ctx, batch, len(initialUsers))
}
