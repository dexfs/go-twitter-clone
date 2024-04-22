package usecase

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/core/domain"
	"github.com/dexfs/go-twitter-clone/core/port/input"
	"github.com/dexfs/go-twitter-clone/core/port/output"
)

type createPostUseCase struct {
	postPort output.PostPort
	userPort output.UserPort
}

func NewCreatePostUseCase(postPort output.PostPort, userPort output.UserPort) (*createPostUseCase, *rest_errors.RestError) {
	if postPort == nil || userPort == nil {
		return nil, rest_errors.NewInternalServerError("postPort and userPort cannot be nil")
	}

	return &createPostUseCase{postPort: postPort, userPort: userPort}, nil
}

func (uc *createPostUseCase) Execute(aInput input.CreatePostUseCaseInput) (*domain.Post, *rest_errors.RestError) {
	hasReachedLimit := uc.postPort.HasReachedPostingLimitDay(aInput.UserID, 5) // @TODO mudar isso para vir das configurações
	if hasReachedLimit {
		return &domain.Post{}, rest_errors.NewBadRequestError("you reached your posts day limit")
	}

	user, err := uc.userPort.FindByID(aInput.UserID)
	if err != nil {
		return &domain.Post{}, rest_errors.NewNotFoundError(err.Error())
	}

	aNewPost, err := domain.NewPost(domain.NewPostInput{
		User:    user,
		Content: aInput.Content,
	})
	if err != nil {
		return &domain.Post{}, rest_errors.NewBadRequestError(err.Error())
	}

	uc.postPort.CreatePost(aNewPost)

	return aNewPost, nil
}
