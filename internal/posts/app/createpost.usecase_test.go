package app

import (
	"github.com/dexfs/go-twitter-clone/internal/posts"
	postInfraImpl "github.com/dexfs/go-twitter-clone/internal/posts/infra/db/repository"
	"github.com/dexfs/go-twitter-clone/internal/user"
	userInfraImpl "github.com/dexfs/go-twitter-clone/internal/user/infra/db/repository"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"reflect"
	"strconv"
	"testing"
)

func TestCreatePostUseCase_WithUserHasReachedLimitPostForCurrentDay_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := TestMocks.MockUserSeed
	mockUserRepo := userInfraImpl.NewInMemoryUserRepo(TestMocks.MockUserDB)
	for i := 0; i < 60; i++ {
		postLoop, _ := posts.NewPost(posts.NewPostInput{
			User:    mockUser[0],
			Content: "post number" + strconv.Itoa(i),
		})
		TestMocks.MockPostDB.Insert(postLoop)
	}

	postRepo := postInfraImpl.NewInMemoryPostRepo(TestMocks.MockPostDB)

	createPostUseCase := NewCreatePostUseCase(mockUserRepo, postRepo)
	useCaseInput := CreatePostInput{
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
	mockUserRepo := userInfraImpl.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := postInfraImpl.NewInMemoryPostRepo(TestMocks.MockPostDB)
	mockNotFoundUser := user.NewUser("not_found_user")
	createPostUseCase := NewCreatePostUseCase(mockUserRepo, postRepo)
	useCaseInput := CreatePostInput{
		UserID:  mockNotFoundUser.ID,
		Content: "user not found",
	}

	output, err := createPostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, CreatePostOutput{}) {
		t.Errorf("should report user not found, got: %v", output)
	}

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
	mockUserRepo := userInfraImpl.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := postInfraImpl.NewInMemoryPostRepo(TestMocks.MockPostDB)
	createPostUseCase := NewCreatePostUseCase(mockUserRepo, postRepo)
	useCaseInput := CreatePostInput{
		UserID:  mockUser[0].ID,
		Content: "a valid post",
	}

	output, err := createPostUseCase.Execute(useCaseInput)

	if reflect.DeepEqual(output, CreatePostOutput{}) {
		t.Errorf("should return PostID, got: %v", output)
	}

	if err != nil {
		t.Errorf("should allow create post")
	}
}
