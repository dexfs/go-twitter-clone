package usecase_test

import (
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/core/usecase"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"testing"
)

func TestExecute_WithValidUsername_ReturnsFeedItems(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0], 2)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)

	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)

	userFeedUseCase, _ := usecase.NewGetUserFeedUseCase(mockUserRepo, postRepo)

	userFeed, err := userFeedUseCase.Execute(mockUser[0].Username)

	if err != nil {
		t.Errorf("want err=nil; got %v", err)
	}

	if len(userFeed) != 2 {
		t.Errorf("want 2 posts; got %v", len(userFeed))
	}
}

func TestExecute_WithEmptyUsername_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0], 2)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)
	getUserFeedUseCase, _ := usecase.NewGetUserFeedUseCase(mockUserRepo, postRepo)
	_, err := getUserFeedUseCase.Execute("")

	if err == nil {
		t.Errorf("should return an error")
	}

	if err.Error() != "user not found" {
		t.Errorf("got %v want %s", err.Error(), "username must not be empty")
	}

}
func TestExecute_WithNonExistingUsername_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0], 2)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)
	getUserFeedUseCase, _ := usecase.NewGetUserFeedUseCase(mockUserRepo, postRepo)
	_, err := getUserFeedUseCase.Execute("non-existing-user")

	if err == nil {
		t.Errorf("should return an error")
	}

	if err.Error() != "user not found" {
		t.Errorf("got %v want %s", err.Error(), "user not found")
	}
}
func TestExecute_WithNilUserRepository_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)
	getUserFeedUseCase, err := usecase.NewGetUserFeedUseCase(nil, postRepo)

	if getUserFeedUseCase != nil {
		t.Errorf("Invalid instance of usecase")
	}

	if err == nil {
		t.Errorf("should return an error")
	}

	if err.Error() != "user port and post port is required" {
		t.Errorf("got %v want %s", err.Error(), "the dependencies should not be nil")
	}
}
func TestExecute_WithNilPostRepository_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	getUserFeedUseCase, err := usecase.NewGetUserFeedUseCase(mockUserRepo, nil)

	if getUserFeedUseCase != nil {
		t.Errorf("Invalid instance of usecase")
	}

	if err == nil {
		t.Errorf("should return an error")
	}

	if err.Error() != "user port and post port is required" {
		t.Errorf("got %v want %s", err.Error(), "the dependencies should not be nil")
	}
}
func TestExecute_WithPostRepositoryError_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	mockPostDB := &database.InMemoryDB[inmemory_schema.PostSchema]{}
	postRepo := mocks.MakeInMemoryPostRepo(mockPostDB)

	userFeedUseCase, _ := usecase.NewGetUserFeedUseCase(mockUserRepo, postRepo)

	userFeed, err := userFeedUseCase.Execute(mockUser[0].Username)

	if err != nil {
		t.Errorf("want err=nil; got %v", err)
	}

	if len(userFeed) > 0 {
		t.Errorf("want 0 posts; got %v", len(userFeed))
	}
}
