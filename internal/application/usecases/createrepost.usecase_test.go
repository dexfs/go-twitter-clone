package app

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/infra/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestCreateRepostUseCase_WithNotFoundUser_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	mockNotFoundUser := domain.NewUser("not_found_user")
	createRepostUseCase := NewCreateRepostUseCase(mockUserRepo, postRepo)
	useCaseInput := CreateRepostUseCaseInput{
		UserID: mockNotFoundUser.ID,
		PostID: TestMocks.MockPostsSeed[0].ID,
	}

	output, err := createRepostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, CreateRepostUseCaseOutput{}) {
		t.Errorf("should report user not found, got: %v", output)
	}

	if err == nil {
		t.Errorf("should not allow create post with not found user")
	}

	if err.Error() != "user not found" {
		t.Errorf("should report 'user not found', got: %s", err.Error())
	}
}
func TestCreateRepostPostUseCase_WithNotFoundPost_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	createRepostUseCase := NewCreateRepostUseCase(mockUserRepo, postRepo)
	useCaseInput := CreateRepostUseCaseInput{
		UserID: TestMocks.MockUserSeed[0].ID,
		PostID: uuid.New().String(),
	}

	output, err := createRepostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, CreateRepostUseCaseOutput{}) {
		t.Errorf("should report user not found, got: %v", output)
	}

	if err == nil {
		t.Errorf("should not allow create quote post with not found original post")
	}

	if err.Error() != "post not found" {
		t.Errorf("should report 'post not found', got: %s", err.Error())
	}
}

func TestCreateRepostPostUseCase_WithPostOwner_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	postRepo := inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	createRepostUseCase := NewCreateRepostUseCase(mockUserRepo, postRepo)
	useCaseInput := CreateRepostUseCaseInput{
		UserID: TestMocks.MockUserSeed[0].ID,
		PostID: TestMocks.MockPostsSeed[0].ID,
	}

	output, err := createRepostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, CreateRepostUseCaseOutput{}) {
		t.Errorf("should report user not found, got: %v", output)
	}

	if err == nil {
		t.Errorf("should not allow create quote post with not found original post")
	}

	if err.Error() != "it is not possible repost your own post" {
		t.Errorf("should report 'it is not possible repost your own post', got: %s", err.Error())
	}
}

func TestCreateRepostPostUseCase_WithValidInput_ReturnsPostID(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := inmemory.NewInMemoryUserRepo(TestMocks.MockUserDB)
	mockQuoteUser := domain.NewUser("quote_user")
	TestMocks.MockUserDB.Insert(mockQuoteUser)

	postRepo := inmemory.NewInMemoryPostRepo(TestMocks.MockPostDB)
	mockOriginalPost := TestMocks.MockPostsSeed[0]

	createRepostUseCase := NewCreateRepostUseCase(mockUserRepo, postRepo)
	useCaseInput := CreateRepostUseCaseInput{
		UserID: mockQuoteUser.ID,
		PostID: mockOriginalPost.ID,
	}

	output, err := createRepostUseCase.Execute(useCaseInput)

	if reflect.DeepEqual(output, CreatePostOutput{}) {
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
