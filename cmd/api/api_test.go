package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dexfs/go-twitter-clone/internal/application/handlers"
	app "github.com/dexfs/go-twitter-clone/internal/application/usecases"
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/infra/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"github.com/google/uuid"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

// users
func TestUserInfoResource(t *testing.T) {
	server := http.NewServeMux()

	dbMocks := mocks.GetTestMocks()
	userRepo := inmemory.NewInMemoryUserRepo(dbMocks.MockUserDB)

	getUserInfoUseCase, err := app.NewGetUserInfoUseCase(userRepo)
	if err != nil {
		log.Fatal(err)
	}
	server.HandleFunc("/users/{username}/info", handlers.NewGetUserInfoHandler(getUserInfoUseCase).Handle)

	request, _ := http.NewRequest("GET", "/users/user0/info", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	var got app.GetUserInfoOutput

	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		fmt.Errorf("could not decode JSON: %v", err)
	}

	want := "user0"
	if got.Username != want {
		t.Errorf("got %q, want %q", got.Username, want)
	}
}
func TestUserFeedResource(t *testing.T) {}

// posts
func TestCreatePostResource(t *testing.T) {
	server := http.NewServeMux()

	dbMocks := mocks.GetTestMocks()
	userRepo := inmemory.NewInMemoryUserRepo(dbMocks.MockUserDB)
	postRepo := inmemory.NewInMemoryPostRepo(dbMocks.MockPostDB)
	createPostUseCase := app.NewCreatePostUseCase(userRepo, postRepo)
	createPostHandler := handlers.NewCreatePostHandler(createPostUseCase)
	server.HandleFunc(createPostHandler.Path, createPostHandler.Handle)

	userID := strconv.Quote(dbMocks.MockUserSeed[0].ID)
	jsonStr := `{"user_id": ` + userID + `, "content": "test content"}`

	request, _ := http.NewRequest("POST", "/posts", bytes.NewBufferString(jsonStr))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	var got app.CreatePostOutput

	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		fmt.Errorf("could not decode JSON: %v", err)
	}

	if err := uuid.Validate(got.PostID); err != nil {
		t.Errorf("got %q, want valid UUID", got.PostID)
	}
}
func TestCreateQuotePostResource(t *testing.T) {
	server := http.NewServeMux()

	dbMocks := mocks.GetTestMocks()
	newUser := &domain.User{
		ID:        "4cfe67a9-defc-42b9-8410-cb5086bec2f5",
		Username:  "alucard",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	dbMocks.MockUserDB.Insert(newUser)
	userRepo := inmemory.NewInMemoryUserRepo(dbMocks.MockUserDB)
	postRepo := inmemory.NewInMemoryPostRepo(dbMocks.MockPostDB)
	createQuotePostUseCase := app.NewCreateQuotePostUseCase(userRepo, postRepo)

	createQuotePostHandler := handlers.NewCreateQuoteHandler(createQuotePostUseCase)
	server.HandleFunc(createQuotePostHandler.Path, createQuotePostHandler.Handle)

	userID := strconv.Quote(newUser.ID)
	postID := strconv.Quote(dbMocks.MockPostsSeed[0].ID)
	jsonStr := `{"user_id": ` + userID + `, "post_id":` + postID + `, "quote": "quote post content"}`
	request, _ := http.NewRequest("POST", "/posts/quote", bytes.NewBufferString(jsonStr))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	var got app.CreateQuotePostUseCaseOutput

	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		t.Fatalf("could not decode JSON: %v", err)
	}

	if err := uuid.Validate(got.PostID); err != nil {
		t.Errorf("got %q, want valid UUID", got.PostID)
	}

}
func TestCreateRepostResource(t *testing.T) {
	server := http.NewServeMux()

	dbMocks := mocks.GetTestMocks()
	newUser := &domain.User{
		ID:        "4cfe67a9-defc-42b9-8410-cb5086bec2f5",
		Username:  "alucard",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	dbMocks.MockUserDB.Insert(newUser)
	userRepo := inmemory.NewInMemoryUserRepo(dbMocks.MockUserDB)
	postRepo := inmemory.NewInMemoryPostRepo(dbMocks.MockPostDB)
	createRepostUseCase := app.NewCreateRepostUseCase(userRepo, postRepo)

	createRepostHandler := handlers.NewRepostHandler(createRepostUseCase)
	server.HandleFunc(createRepostHandler.Path, createRepostHandler.Handle)

	userID := strconv.Quote(newUser.ID)
	postID := strconv.Quote(dbMocks.MockPostsSeed[0].ID)
	jsonStr := `{"user_id": ` + userID + `, "post_id":` + postID + `}`
	request, _ := http.NewRequest("POST", "/posts/repost", bytes.NewBufferString(jsonStr))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	var got app.CreateRepostUseCaseOutput

	if err := json.NewDecoder(response.Body).Decode(&got); err != nil {
		t.Fatalf("could not decode JSON: %v", err)
	}

	if err := uuid.Validate(got.PostID); err != nil {
		t.Errorf("got %q, want valid UUID", got.PostID)
	}
}
