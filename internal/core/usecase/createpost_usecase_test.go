package usecase_test

import (
	"github.com/dexfs/go-twitter-clone/adapter/output/mappers"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/internal/core/port/input"
	"github.com/dexfs/go-twitter-clone/internal/core/usecase"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"strconv"
	"testing"
)

func TestCreatePostUseCase_WithUserHasReachedLimitPostForCurrentDay_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := TestMocks.MockUserSeed
	mockUserRepo := inmemory.NewInMemoryUserRepository(TestMocks.MockDB)
	for i := 0; i < 60; i++ {
		postLoop, _ := domain.NewPost(domain.NewPostInput{
			User:    mockUser[0],
			Content: "post number" + strconv.Itoa(i),
		})
		mocks.InsertPostHelper(TestMocks.MockDB, mappers.NewPostMapper().ToPersistence(postLoop))
	}

	mockPostRepo := inmemory.NewInMemoryPostRepository(TestMocks.MockDB)

	createPostUseCase, _ := usecase.NewCreatePostUseCase(mockPostRepo, mockUserRepo)

	useCaseInput := input.CreatePostUseCaseInput{
		UserID:  mockUser[0].ID,
		Content: "Reached limit",
	}

	_, err := createPostUseCase.Execute(useCaseInput)

	if err == nil {
		t.Errorf("should not allow create post for reached limit user")
	}

	if err.Error() != "you reached your posts day limit" {
		t.Errorf("should report you reached your posts day limit, got: %s", err.Error())
	}
}
func TestCreatePostUseCase_WithNotFoundUser_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := inmemory.NewInMemoryUserRepository(TestMocks.MockDB)
	mockPostRepo := inmemory.NewInMemoryPostRepository(TestMocks.MockDB)
	mockNotFoundUser := domain.NewUser("not_found_user")
	createPostUseCase, _ := usecase.NewCreatePostUseCase(mockPostRepo, mockUserRepo)
	useCaseInput := input.CreatePostUseCaseInput{
		UserID:  mockNotFoundUser.ID,
		Content: "user not found",
	}

	_, err := createPostUseCase.Execute(useCaseInput)

	if err == nil {
		t.Errorf("should not allow create post for reached limit user")
	}

	if err.Error() != "user not found" {
		t.Errorf("should report 'user not found', got: %s", err.Error())
	}
}

func TestCreatePostUseCase_WithValidInput_ReturnsPostID(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := TestMocks.MockUserSeed
	mockUserRepo := inmemory.NewInMemoryUserRepository(TestMocks.MockDB)
	mockPostRepo := inmemory.NewInMemoryPostRepository(TestMocks.MockDB)
	useCaseInput := input.CreatePostUseCaseInput{
		UserID:  mockUser[0].ID,
		Content: "user not found",
	}
	createPostUseCase, _ := usecase.NewCreatePostUseCase(mockPostRepo, mockUserRepo)
	_, err := createPostUseCase.Execute(useCaseInput)

	if err != nil {
		t.Errorf("should allow create post")
	}
}
