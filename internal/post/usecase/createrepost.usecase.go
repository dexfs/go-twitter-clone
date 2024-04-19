package post_usecase

import (
	"errors"
	post_domain "github.com/dexfs/go-twitter-clone/internal/post"
	"github.com/dexfs/go-twitter-clone/internal/user"
)

type CreateRepostUseCaseInput struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}

type CreateRepostUseCaseOutput struct {
	PostID string
}

type CreateRepostUseCase struct {
	userRepo user.UserRepository
	postRepo post_domain.PostRepository
}

func NewCreateRepostUseCase(userRepo user.UserRepository, postRepo post_domain.PostRepository) *CreateRepostUseCase {
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

	aRepostInput := post_domain.NewRepostQuoteInput{
		User: user,
		Post: post,
	}

	aRepost, err := post_domain.NewRepost(aRepostInput)

	if err != nil {
		return CreateRepostUseCaseOutput{}, err
	}

	uc.postRepo.Insert(aRepost)

	return CreateRepostUseCaseOutput{PostID: aRepost.ID}, nil
}
