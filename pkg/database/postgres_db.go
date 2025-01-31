package database

import (
	"context"
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

func (db *PostgresDB) Close() {
	db.conn.Close(context.Background())
}

func (db *PostgresDB) Version() {
	var version string
	err := db.conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	defer db.conn.Close(context.Background())
	if err != nil {
		log.Fatal("Erro ao executar query:", err)
	}
	fmt.Println("PostgreSQL version:", version)
}
