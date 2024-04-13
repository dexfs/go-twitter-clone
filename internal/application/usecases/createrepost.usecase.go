package app

import (
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/domain/interfaces"
)

type CreateRepostUseCaseInput struct {
	PostID string
	UserID string
}

type CreateRepostUseCaseOutput struct {
	PostID string
}

type CreateRepostUseCase struct {
	userRepo interfaces.UserRepository
	postRepo interfaces.PostRepository
}

func NewCreateRepostUseCase(userRepo interfaces.UserRepository, postRepo interfaces.PostRepository) *CreateRepostUseCase {
	return &CreateRepostUseCase{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (uc *CreateRepostUseCase) Execute(input CreateRepostUseCaseInput) (CreateRepostUseCaseOutput, error) {
	user, err := uc.userRepo.FindByID(input.UserID)

	if err != nil {
		return CreateRepostUseCaseOutput{}, err
	}

	post, err := uc.postRepo.FindByID(input.PostID)
	if err != nil {
		return CreateRepostUseCaseOutput{}, err
	}

	aRepostInput := domain.NewRepostQuoteInput{
		User: user,
		Post: post,
	}

	aRepost, err := domain.NewRepost(aRepostInput)

	if err != nil {
		return CreateRepostUseCaseOutput{}, err
	}

	uc.postRepo.Insert(aRepost)

	return CreateRepostUseCaseOutput{PostID: aRepost.ID}, nil
}
