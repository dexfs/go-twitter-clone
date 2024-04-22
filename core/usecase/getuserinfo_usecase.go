package usecase

import (
	"fmt"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/core/domain"
	"github.com/dexfs/go-twitter-clone/core/port/output"
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
	fmt.Sprintf("GetUserInfoService_Execute(%s)", username)
	userInfoResponse, err := s.userPort.ByUsername(username)
	if err != nil {
		return nil, rest_errors.NewNotFoundError(err.Error())
	}
	return userInfoResponse, nil
}
