package app

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/domain/interfaces"
)

type CreateQuotePostUseCaseInput struct {
	Quote  string `json:"quote"`
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}

type CreateQuotePostUseCaseOutput struct {
	PostID string
}

type CreateQuotePostUseCase struct {
	userRepo interfaces.UserRepository
	postRepo interfaces.PostRepository
}

func NewCreateQuotePostUseCase(userRepo interfaces.UserRepository, postRepo interfaces.PostRepository) *CreateQuotePostUseCase {
	return &CreateQuotePostUseCase{userRepo, postRepo}
}

func (uc CreateQuotePostUseCase) Execute(input CreateQuotePostUseCaseInput) (CreateQuotePostUseCaseOutput, error) {

	user, err := uc.userRepo.FindByID(input.UserID)

	if err != nil {
		return CreateQuotePostUseCaseOutput{}, err
	}

	post, err := uc.postRepo.FindByID(input.PostID)
	if err != nil {
		return CreateQuotePostUseCaseOutput{}, err
	}

	newQuotePostInput := domain.NewRepostQuoteInput{
		User:    user,
		Post:    post,
		Content: input.Quote,
	}
	newQuotePost, err := domain.NewQuote(newQuotePostInput)

	if err != nil {
		return CreateQuotePostUseCaseOutput{}, err
	}

	uc.postRepo.Insert(newQuotePost)

	return CreateQuotePostUseCaseOutput{PostID: newQuotePost.ID}, nil
}
