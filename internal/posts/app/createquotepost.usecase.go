package app

import (
	"github.com/dexfs/go-twitter-clone/internal/posts"
	postDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/posts/domain/interfaces"
	userDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/user/domain/interfaces"
)

type CreateQuotePostUseCaseInput struct {
	Quote  string
	PostID string
	UserID string
}

type CreateQuotePostUseCaseOutput struct {
	PostID string
}

type CreateQuotePostUseCase struct {
	userRepo userDomainInterfaces.UserRepository
	postRepo postDomainInterfaces.PostRepository
}

func NewCreateQuotePostUseCase(userRepo userDomainInterfaces.UserRepository, postRepo postDomainInterfaces.PostRepository) CreateQuotePostUseCase {
	return CreateQuotePostUseCase{userRepo, postRepo}
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

	newQuotePostInput := posts.NewRepostQuoteInput{
		User:    user,
		Post:    post,
		Content: input.Quote,
	}
	newQuotePost, err := posts.NewQuote(newQuotePostInput)

	if err != nil {
		return CreateQuotePostUseCaseOutput{}, err
	}

	uc.postRepo.Insert(newQuotePost)

	return CreateQuotePostUseCaseOutput{PostID: newQuotePost.ID}, nil
}
