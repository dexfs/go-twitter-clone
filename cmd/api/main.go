package main

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/routes"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var (
	db       *database.InMemoryDB
	pgDB     *database.PostgresDB
	userRepo output.UserPort
	postRepo output.PostPort
)

func init() {
	err := godotenv.Load()
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
	pgDB.Version()
	if db == nil {
		log.Fatal("database is nil")
	}

	initialUsers := make([]*inmemory_schema.UserSchema, 0)
	initialUsers = append(initialUsers, &inmemory_schema.UserSchema{
		ID:        "4cfe67a9-defc-42b9-8410-cb5086bec2f5",
		Username:  "alucard",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	initialUsers = append(initialUsers, &inmemory_schema.UserSchema{
		ID:        "b8903f77-5d16-4176-890f-f597594ff952",
		Username:  "alexander",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	db.RegisterSchema(inmemory.USER_SCHEMA_NAME, initialUsers)
	db.RegisterSchema(inmemory.POST_SCHEMA_NAME, []*inmemory_schema.PostSchema{})
	userRepo = inmemory.NewInMemoryUserRepository(db)
	postRepo = inmemory.NewInMemoryPostRepository(db)
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
	defer pgDB.Close()

	if err := server.Run(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
