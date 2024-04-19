package post_usecases

import (
	"errors"
	PostModule "github.com/dexfs/go-twitter-clone/internal/post"
	UserModule "github.com/dexfs/go-twitter-clone/internal/user"
)

type CreateRepostUseCaseInput struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}

type CreateRepostUseCaseOutput struct {
	PostID string
}

type CreateRepostUseCase struct {
	userRepo UserModule.UserRepository
	postRepo PostModule.PostRepository
}

func NewCreateRepostUseCase(userRepo UserModule.UserRepository, postRepo PostModule.PostRepository) *CreateRepostUseCase {
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

	aRepostInput := PostModule.NewRepostQuoteInput{
		User: user,
		Post: post,
	}

	aRepost, err := PostModule.NewRepost(aRepostInput)

	if err != nil {
		return CreateRepostUseCaseOutput{}, err
	}

	uc.postRepo.Insert(aRepost)

	return CreateRepostUseCaseOutput{PostID: aRepost.ID}, nil
}
