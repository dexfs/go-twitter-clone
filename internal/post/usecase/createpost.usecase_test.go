package post_usecase_test

import (
	repo_inmemory "github.com/dexfs/go-twitter-clone/internal/infra/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/internal/post"
	"github.com/dexfs/go-twitter-clone/internal/post/usecase"
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/mocks"
	"reflect"
	"strconv"
	"testing"
)

func TestCreatePostUseCase_WithUserHasReachedLimitPostForCurrentDay_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUser := TestMocks.MockUserSeed
	mockUserRepo := repo_inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	for i := 0; i < 60; i++ {
		postLoop, _ := post.NewPost(post.NewPostInput{
			User:    mockUser[0],
			Content: "post number" + strconv.Itoa(i),
		})
		TestMocks.MockPostDB.Insert(postLoop)
	}

	postRepo := repo_inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)

	createPostUseCase := post_usecase.NewCreatePostUseCase(mockUserRepo, postRepo)
	useCaseInput := post_usecase.CreatePostInput{
		UserID:  mockUser[0].ID,
		Content: "Reached limit",
	}
	_, err := createPostUseCase.Execute(useCaseInput)

	if err == nil {
		t.Errorf("should not allow create post for reached limit user")
	}

	if err.Error() != "you reached your post day limit" {
		t.Errorf("should report you reached your post day limit, got: %s", err.Error())
	}
}

func TestCreatePostUseCase_WithNotFoundUser_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := repo_inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := repo_inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	mockNotFoundUser := user.NewUser("not_found_user")
	createPostUseCase := post_usecase.NewCreatePostUseCase(mockUserRepo, postRepo)
	useCaseInput := post_usecase.CreatePostInput{
		UserID:  mockNotFoundUser.ID,
		Content: "user not found",
	}

	output, err := createPostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, post_usecase.CreatePostOutput{}) {
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
	mockUserRepo := repo_inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := repo_inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	createPostUseCase := post_usecase.NewCreatePostUseCase(mockUserRepo, postRepo)
	useCaseInput := post_usecase.CreatePostInput{
		UserID:  mockUser[0].ID,
		Content: "a valid post",
	}

	output, err := createPostUseCase.Execute(useCaseInput)

	if reflect.DeepEqual(output, post_usecase.CreatePostOutput{}) {
		t.Errorf("should return PostID, got: %v", output)
	}

	if err != nil {
		t.Errorf("should allow create post")
	}
}
