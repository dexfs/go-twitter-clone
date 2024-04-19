package post_usecase_test

import (
	"github.com/dexfs/go-twitter-clone/internal/infra/repository/inmemory"
	post_usecase "github.com/dexfs/go-twitter-clone/internal/post/usecase"
	"github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/mocks"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestCreateQuotePostUseCase_WithNotFoundUser_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := repo_inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := repo_inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	mockNotFoundUser := user.NewUser("not_found_user")
	createQuotePostUseCase := post_usecase.NewCreateQuotePostUseCase(mockUserRepo, postRepo)
	useCaseInput := post_usecase.CreateQuotePostUseCaseInput{
		UserID: mockNotFoundUser.ID,
		PostID: TestMocks.MockPostsSeed[0].ID,
		Quote:  "not found user",
	}

	output, err := createQuotePostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, post_usecase.CreateQuotePostUseCaseOutput{}) {
		t.Errorf("should report user not found, got: %v", output)
	}

	if err == nil {
		t.Errorf("should not allow create post with not found user")
	}

	if err.Error() != "user not found" {
		t.Errorf("should report 'user not found', got: %s", err.Error())
	}
}
func TestCreateQuotePostUseCase_WithNotFoundPost_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := repo_inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := repo_inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	createQuotePostUseCase := post_usecase.NewCreateQuotePostUseCase(mockUserRepo, postRepo)
	useCaseInput := post_usecase.CreateQuotePostUseCaseInput{
		UserID: TestMocks.MockUserSeed[0].ID,
		PostID: uuid.New().String(),
		Quote:  "not found user",
	}

	output, err := createQuotePostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, post_usecase.CreateQuotePostUseCaseOutput{}) {
		t.Errorf("should report user not found, got: %v", output)
	}

	if err == nil {
		t.Errorf("should not allow create quote post with not found original post")
	}

	if err.Error() != "post not found" {
		t.Errorf("should report 'post not found', got: %s", err.Error())
	}
}
func TestCreateQuotePostUseCase_WithValidInput_ReturnsPostID(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := repo_inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	mockQuoteUser := user.NewUser("quote_user")
	TestMocks.MockUserDB.Insert(mockQuoteUser)

	postRepo := repo_inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	createQuotePostUseCase := post_usecase.NewCreateQuotePostUseCase(mockUserRepo, postRepo)
	mockOriginalPost := TestMocks.MockPostsSeed[0]
	useCaseInput := post_usecase.CreateQuotePostUseCaseInput{
		UserID: mockQuoteUser.ID,
		PostID: mockOriginalPost.ID,
		Quote:  "New quote!",
	}

	output, err := createQuotePostUseCase.Execute(useCaseInput)

	if reflect.DeepEqual(output, post_usecase.CreatePostOutput{}) {
		t.Errorf("should return PostID, got: %v", output)
	}

	if err != nil {
		t.Errorf("should allow create quote post")
	}

	getNewQuotePost, _ := postRepo.FindByID(output.PostID)

	notExpected := getNewQuotePost.ID == mockOriginalPost.ID &&
		getNewQuotePost.OriginalPostID != mockOriginalPost.ID &&
		getNewQuotePost.OriginalPostContent != mockOriginalPost.Content

	if notExpected {
		t.Errorf("quote should be created but new quote post not found, got: %v", getNewQuotePost)
	}

}
