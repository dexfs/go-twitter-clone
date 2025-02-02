package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"time"
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

func (db *PostgresDB) Find(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return db.conn.Query(ctx, query, args...)
}

func (db *PostgresDB) Batch(ctx context.Context, batch *pgx.Batch, dataSize int) error {
	sendStart := time.Now()
	br := db.conn.SendBatch(ctx, batch)
	sendDuration := time.Since(sendStart)

	defer br.Close()
	for i := 0; i < dataSize; i++ {
		queryStart := time.Now()
		_, err := br.Exec()
		if err != nil {
			return errors.New("User seed failed: " + err.Error())
		}
		queryDuration := time.Since(queryStart)
		fmt.Printf("Query %d: %v\n", i+1, queryDuration)
	}

	fmt.Printf("Tempo de envio do batch: %v\n", sendDuration)
	fmt.Printf("Tempo total: %v\n", time.Since(sendStart))
	return nil
}

func (db *PostgresDB) Insert(ctx context.Context, query string, args ...any) error {
	_, err := db.conn.Exec(ctx, query, args...)
	if err != nil {
		return err
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
