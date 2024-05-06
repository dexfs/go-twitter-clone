package usecase_test

import (
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/internal/core/port/input"
	"github.com/dexfs/go-twitter-clone/internal/core/usecase"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestCreateQuotePostUseCase_WithNotFoundUser_ReturnsError(t *testing.T) {
	TestMocks := mocks.GetTestMocks()
	mockUserRepo := inmemory.NewInMemoryUserRepository(TestMocks.MockDB)
	postRepo := inmemory.NewInMemoryPostRepository(TestMocks.MockDB)
	mockNotFoundUser := domain.NewUser("not_found_user")
	createQuotePostUseCase, _ := usecase.NewCreateQuoteUseCase(postRepo, mockUserRepo)
	useCaseInput := input.CreateQuoteUseCaseInput{
		UserID: mockNotFoundUser.ID,
		PostID: TestMocks.MockPostsSeed[0].ID,
		Quote:  "not found user",
	}

	output, err := createQuotePostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, &domain.Post{}) {
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
	mockUserRepo := inmemory.NewInMemoryUserRepository(TestMocks.MockDB)
	postRepo := inmemory.NewInMemoryPostRepository(TestMocks.MockDB)
	createQuotePostUseCase, _ := usecase.NewCreateQuoteUseCase(postRepo, mockUserRepo)
	useCaseInput := input.CreateQuoteUseCaseInput{
		UserID: TestMocks.MockUserSeed[0].ID,
		PostID: uuid.New().String(),
		Quote:  "not found user",
	}

	output, err := createQuotePostUseCase.Execute(useCaseInput)

	if !reflect.DeepEqual(output, &domain.Post{}) {
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
	mockUserRepo := inmemory.NewInMemoryUserRepository(TestMocks.MockDB)
	mockQuoteUser := domain.NewUser("quote_user")

	mocks.InsertUserHelper(TestMocks.MockDB, &inmemory_schema.UserSchema{
		ID:        mockQuoteUser.ID,
		Username:  mockQuoteUser.Username,
		CreatedAt: mockQuoteUser.CreatedAt,
		UpdatedAt: mockQuoteUser.UpdatedAt,
	})

	postRepo := inmemory.NewInMemoryPostRepository(TestMocks.MockDB)
	createQuotePostUseCase, _ := usecase.NewCreateQuoteUseCase(postRepo, mockUserRepo)
	mockOriginalPost := TestMocks.MockPostsSeed[0]
	useCaseInput := input.CreateQuoteUseCaseInput{
		UserID: mockQuoteUser.ID,
		PostID: mockOriginalPost.ID,
		Quote:  "New quote!",
	}

	output, err := createQuotePostUseCase.Execute(useCaseInput)

	if reflect.DeepEqual(output, &domain.Post{}) {
		t.Errorf("should return PostID, got: %v", output)
	}

	if err != nil {
		t.Errorf("should allow create quote post")
	}

	notExpected := output.ID == mockOriginalPost.ID &&
		output.OriginalPostID != mockOriginalPost.ID &&
		output.OriginalPostContent != mockOriginalPost.Content

	if notExpected {
		t.Errorf("quote should be created but new quote post not found, got: %v", output)
	}
}
