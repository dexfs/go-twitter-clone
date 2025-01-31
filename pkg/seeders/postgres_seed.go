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
		ID:        "0194bd04-66e2-7cd8-b3d9-66eda709f2ee",
		Username:  "alucard",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	initialUsers = append(initialUsers, &postgres.UserSchema{
		ID:        "0194bd04-8eac-7e70-97cd-c526cdda3d6a",
		Username:  "alexander",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	initialUsers = append(initialUsers, &postgres.UserSchema{
		ID:        "0194bdb1-0588-7181-809e-a825badac714",
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
