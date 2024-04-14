package app

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/domain/interfaces"
)

type GetUserFeedUseCase struct {
	userRepo interfaces.UserRepository
	postRepo interfaces.PostRepository
}

type GetUserFeedUseCaseOutput struct {
	Items []*domain.Post `json:"items"`
}

func NewGetUserFeedUseCase(userRepo interfaces.UserRepository, postRepo interfaces.PostRepository) (*GetUserFeedUseCase, error) {
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
		Items: posts,
	}, nil
}
