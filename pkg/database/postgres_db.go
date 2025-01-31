package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

type PostgresDB struct {
	conn *pgx.Conn
}

func NewPostgresDB() *PostgresDB {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error on connect database", err)
	}

	return &PostgresDB{
		conn: conn,
	}
}

func (db *PostgresDB) FindOne(ctx context.Context, query string, args ...any) pgx.Row {
	return db.conn.QueryRow(ctx, query, args...)
}

func (db *PostgresDB) Batch(ctx context.Context, batch *pgx.Batch, dataSize int) error {
	br := db.conn.SendBatch(ctx, batch)
	defer br.Close()
	for range dataSize {
		_, err := br.Exec()
		if err != nil {
			return errors.New("User seed failed: " + err.Error())
		}
	}
	return nil
}

func (db *PostgresDB) Close(ctx context.Context) {
	db.conn.Close(ctx)
}

func (db *PostgresDB) Version(ctx context.Context) {
	var version string
	err := db.conn.QueryRow(ctx, "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal("Erro ao executar query:", err)
	}
	fmt.Println("PostgreSQL version:", version)
}
