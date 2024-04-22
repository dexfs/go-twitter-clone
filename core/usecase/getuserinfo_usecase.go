package usecase

import (
	"errors"
	"fmt"
	"github.com/dexfs/go-twitter-clone/core/domain"
	"github.com/dexfs/go-twitter-clone/core/port/output"
)

type getUserInfoUseCase struct {
	userPort output.UserPort
}

func NewGetUserInfoUseCase(userPort output.UserPort) (*getUserInfoUseCase, error) {
	if userPort == nil {
		return nil, errors.New("userPort cannot be nil")
	}
	return &getUserInfoUseCase{
		userPort: userPort,
	}, nil
}

func (s *getUserInfoUseCase) Execute(username string) (*domain.User, error) {
	fmt.Sprintf("GetUserInfoService_Execute(%s)", username)
	userInfoResponse, err := s.userPort.ByUsername(username)

	return userInfoResponse, err
}
