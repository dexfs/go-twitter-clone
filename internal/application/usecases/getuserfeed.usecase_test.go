package app

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"reflect"
	"testing"
)

func TestExecute_WithValidUsername_ReturnsFeedItems(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0], 2)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)

	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)

	userFeedUseCase, _ := NewGetUserFeedUseCase(mockUserRepo, postRepo)

	userFeed, err := userFeedUseCase.Execute(mockUser[0].Username)

	if err != nil {
		t.Errorf("want err=nil; got %v", err)
	}

	if len(userFeed.items) != 2 {
		t.Errorf("want 2 posts; got %v", len(userFeed.items))
	}
}
func TestExecute_WithEmptyUsername_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0], 2)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)
	getUserFeedUseCase, _ := NewGetUserFeedUseCase(mockUserRepo, postRepo)
	userFeedOutput, err := getUserFeedUseCase.Execute("")

	var expectedOutputItems []*domain.Post
	expectedOutputFeed := GetUserFeedUseCaseOutput{
		items: expectedOutputItems,
	}

	if !reflect.DeepEqual(userFeedOutput, expectedOutputFeed) {
		t.Errorf("want nil; got %v", userFeedOutput.items)
	}

	if err == nil {
		t.Errorf("should return an error")
	}

	if err.Error() != "username must not be empty" {
		t.Errorf("got %v want %s", err.Error(), "username must not be empty")
	}

}
func TestExecute_WithNonExistingUsername_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0], 2)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)
	getUserFeedUseCase, _ := NewGetUserFeedUseCase(mockUserRepo, postRepo)
	userFeedOutput, err := getUserFeedUseCase.Execute("non-existing-user")

	var expectedOutputItems []*domain.Post
	expectedOutputFeed := GetUserFeedUseCaseOutput{
		items: expectedOutputItems,
	}

	if !reflect.DeepEqual(userFeedOutput, expectedOutputFeed) {
		t.Errorf("want nil; got %v", userFeedOutput.items)
	}

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
	getUserFeedUseCase, err := NewGetUserFeedUseCase(nil, postRepo)

	if getUserFeedUseCase != nil {
		t.Errorf("Invalid instance of usecase")
	}

	if err == nil {
		t.Errorf("should return an error")
	}

	if err.Error() != "the dependencies should not be nil" {
		t.Errorf("got %v want %s", err.Error(), "the dependencies should not be nil")
	}
}
func TestExecute_WithNilPostRepository_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	getUserFeedUseCase, err := NewGetUserFeedUseCase(mockUserRepo, nil)

	if getUserFeedUseCase != nil {
		t.Errorf("Invalid instance of usecase")
	}

	if err == nil {
		t.Errorf("should return an error")
	}

	if err.Error() != "the dependencies should not be nil" {
		t.Errorf("got %v want %s", err.Error(), "the dependencies should not be nil")
	}
}
func TestExecute_WithPostRepositoryError_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	mockPostDB := &database.InMemoryDB[domain.Post]{}
	postRepo := mocks.MakeInMemoryPostRepo(mockPostDB)

	userFeedUseCase, _ := NewGetUserFeedUseCase(mockUserRepo, postRepo)

	userFeed, err := userFeedUseCase.Execute(mockUser[0].Username)

	if err != nil {
		t.Errorf("want err=nil; got %v", err)
	}

	if len(userFeed.items) > 0 {
		t.Errorf("want 0 posts; got %v", len(userFeed.items))
	}
}
