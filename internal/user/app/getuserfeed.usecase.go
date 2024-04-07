package app

import (
	"errors"
	postEntity "github.com/dexfs/go-twitter-clone/internal/posts"
	postDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/posts/domain/interfaces"
	userDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/user/domain/interfaces"
)

type GetUserFeedUseCase struct {
	userRepo userDomainInterfaces.UserRepository
	postRepo postDomainInterfaces.PostRepository
}

type GetUserFeedUseCaseOutput struct {
	items []*postEntity.Post
}

func NewGetUserFeedUseCase(userRepo userDomainInterfaces.UserRepository, postRepo postDomainInterfaces.PostRepository) (*GetUserFeedUseCase, error) {
	if userRepo == nil || postRepo == nil {
		return nil, errors.New("the dependencies should not be nil")
	}

	return &GetUserFeedUseCase{
		userRepo: userRepo,
		postRepo: postRepo,
	}, nil
}

func (uc *GetUserFeedUseCase) Execute(username string) (GetUserFeedUseCaseOutput, error) {
	if len(username) == 0 {
		return GetUserFeedUseCaseOutput{}, errors.New("username must not be empty")
	}
	user, err := uc.userRepo.ByUsername(username)

	if err != nil {
		return GetUserFeedUseCaseOutput{}, err
	}

	posts := uc.postRepo.GetFeedByUserID(user.ID)

	if posts == nil {
		return GetUserFeedUseCaseOutput{}, nil
	}

	return GetUserFeedUseCaseOutput{
		items: posts,
	}, nil
}
