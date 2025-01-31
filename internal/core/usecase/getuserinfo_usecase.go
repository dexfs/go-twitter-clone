package usecase

import (
	"context"
	"fmt"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
)

type getUserInfoUseCase struct {
	userPort output.UserPort
}

func NewGetUserInfoUseCase(userPort output.UserPort) (*getUserInfoUseCase, *rest_errors.RestError) {
	if userPort == nil {
		return nil, rest_errors.NewInternalServerError("userPort cannot be nil")
	}
	return &getUserInfoUseCase{
		userPort: userPort,
	}, nil
}

func (s *getUserInfoUseCase) Execute(username string) (*domain.User, *rest_errors.RestError) {
	ctx := context.Background()
	fmt.Sprintf("GetUserInfoService_Execute(%s)", username)
	userInfoResponse, err := s.userPort.ByUsername(ctx, username)
	if err != nil {
		return nil, rest_errors.NewNotFoundError(err.Error())
	}
	return userInfoResponse, nil
}
