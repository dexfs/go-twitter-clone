package main

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/input/routes"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/postgres"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	db       *database.InMemoryDB
	pgDB     *database.PostgresDB
	userRepo output.UserPort
	postRepo output.PostPort
)

func init() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	m, err := migrate.New("file://migrations", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error on load migrations:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error on apply migrations:", err)
	}

	log.Println("Migrations applied")

	db = database.NewInMemoryDB()
	pgDB = database.NewPostgresDB()
	pgDB.Version(ctx)
	if db == nil {
		log.Fatal("database is nil")
	}

	userRepo = postgres.NewPostgresUserRepository(pgDB)
	postRepo = postgres.NewPostgresPostRepository(pgDB)

}

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: ":" + addr,
	}
}

func (s *APIServer) Run() error {
	router := routes.NewRouter(s.addr)
	return router.Run(userRepo, postRepo)
}

func main() {
	log.Printf("Starting Application")
	server := NewAPIServer("8001")
	defer pgDB.Close(context.Background())

	if err := server.Run(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
