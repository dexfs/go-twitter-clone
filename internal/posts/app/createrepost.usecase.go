package app

import (
	"github.com/dexfs/go-twitter-clone/internal/posts"
	postDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/posts/domain/interfaces"
	userDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/user/domain/interfaces"
)

type CreateRepostUseCaseInput struct {
	PostID string
	UserID string
}

type CreateRepostUseCaseOutput struct {
	PostID string
}

type CreateRepostUseCase struct {
	userRepo userDomainInterfaces.UserRepository
	postRepo postDomainInterfaces.PostRepository
}

func NewCreateRepostUseCase(userRepo userDomainInterfaces.UserRepository, postRepo postDomainInterfaces.PostRepository) *CreateRepostUseCase {
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

	aRepostInput := posts.NewRepostQuoteInput{
		User: user,
		Post: post,
	}

	aRepost, err := posts.NewRepost(aRepostInput)

	if err != nil {
		return CreateRepostUseCaseOutput{}, err
	}

	uc.postRepo.Insert(aRepost)

	return CreateRepostUseCaseOutput{PostID: aRepost.ID}, nil
}
