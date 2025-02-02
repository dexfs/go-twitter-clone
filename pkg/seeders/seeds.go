package seeders

import (
	"context"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func Run() {
	ctx := context.Background()
	err := godotenv.Load()
	pgDB := database.NewPostgresDB()
	defer pgDB.Close(ctx)
	pgSeeder := NewPostgresSeed(pgDB)
	ctxWithCancel, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	err = pgSeeder.UsersSeed(ctxWithCancel)
	if err != nil {
		log.Fatal("Error on seeding users: ", err)
	}

	err = pgSeeder.PostsSeed(ctxWithCancel)
	if err != nil {
		log.Fatal("Error on seeding posts: ", err)
	}

	log.Println("seeds applied")
}
