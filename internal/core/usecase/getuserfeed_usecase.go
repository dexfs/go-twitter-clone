package usecase

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
)

type getUserFeedUseCase struct {
	userPort output.UserPort
	postPort output.PostPort
}

func NewGetUserFeedUseCase(userPort output.UserPort, postPort output.PostPort) (*getUserFeedUseCase, *rest_errors.RestError) {
	if userPort == nil || postPort == nil {
		return nil, rest_errors.NewInternalServerError("user port and post port is required")
	}

	return &getUserFeedUseCase{
		userPort: userPort,
		postPort: postPort,
	}, nil
}

func (uc *getUserFeedUseCase) Execute(username string) ([]*domain.Post, *rest_errors.RestError) {
	ctx := context.Background()
	user, err := uc.userPort.ByUsername(ctx, username)
	if err != nil {
		return []*domain.Post{}, rest_errors.NewNotFoundError(err.Error())
	}

	posts := uc.postPort.AllByUserID(user)
	return posts, nil
}
