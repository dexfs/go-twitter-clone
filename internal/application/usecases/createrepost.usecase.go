package app

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/domain/interfaces"
)

type CreateRepostUseCaseInput struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
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

	reposted := uc.postRepo.HasPostBeenRepostedByUser(input.PostID, input.UserID)
	if reposted {
		return CreateRepostUseCaseOutput{}, errors.New("it is not possible repost a repost post")
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
