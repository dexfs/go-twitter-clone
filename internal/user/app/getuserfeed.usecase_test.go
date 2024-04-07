package app

import (
	postEntity "github.com/dexfs/go-twitter-clone/internal/posts"
	userEntity "github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"reflect"
	"testing"
)

func TestExecute_WithValidUsername_ReturnsFeedItems(t *testing.T) {
	TestMocks := GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0])
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
	TestMocks := GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0])
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)
	getUserFeedUseCase, _ := NewGetUserFeedUseCase(mockUserRepo, postRepo)
	userFeedOutput, err := getUserFeedUseCase.Execute("")

	var expectedOutputItems []*postEntity.Post
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
	TestMocks := GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mocks.PostSeed(TestMocks.MockPostDB, mockUser[0])
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := mocks.MakeInMemoryPostRepo(TestMocks.MockPostDB)
	getUserFeedUseCase, _ := NewGetUserFeedUseCase(mockUserRepo, postRepo)
	userFeedOutput, err := getUserFeedUseCase.Execute("non-existing-user")

	var expectedOutputItems []*postEntity.Post
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
	TestMocks := GetTestMocks()
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
	TestMocks := GetTestMocks()
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
	TestMocks := GetTestMocks()
	mockUser := mocks.UserSeed(TestMocks.MockUserDB, 1)
	mockUserRepo := mocks.MakeInMemoryUserRepo(TestMocks.MockUserDB)
	mockPostDB := &database.InMemoryDB[postEntity.Post]{}
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

type TestMocks struct {
	MockUserDB    *database.InMemoryDB[userEntity.User]
	MockUserSeed  []*userEntity.User
	MockPostDB    *database.InMemoryDB[postEntity.Post]
	MockPostsSeed []*postEntity.Post
}

func GetTestMocks() TestMocks {
	mockUserDB := mocks.MakeDb[userEntity.User]()
	mockPostDB := mocks.MakeDb[postEntity.Post]()
	mockUserSeed := mocks.UserSeed(mockUserDB, 1)
	mockPostsSeed := mocks.PostSeed(mockPostDB, mockUserSeed[0])

	return TestMocks{
		MockUserDB:    mockUserDB,
		MockUserSeed:  mockUserSeed,
		MockPostDB:    mockPostDB,
		MockPostsSeed: mockPostsSeed,
	}
}
