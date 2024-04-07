package app

import (
	"errors"
	userEntity "github.com/dexfs/go-twitter-clone/internal/user"
	"github.com/dexfs/go-twitter-clone/internal/user/domain/interfaces"
)

type GetUserInfoUseCase struct {
	userRepo interfaces.UserRepository
}

type GetUserInfoOutput struct {
	info *userEntity.User
}

func NewGetUserFeedUseCase(userRepo interfaces.UserRepository) (*GetUserInfoUseCase, error) {
	if userRepo == nil {
		return nil, errors.New("userRepo cannot be nil")
	}
	return &GetUserInfoUseCase{userRepo: userRepo}, nil
}

func (u *GetUserInfoUseCase) Execute(username string) (GetUserInfoOutput, error) {
	result, err := u.userRepo.ByUsername(username)
	if err != nil {
		return GetUserInfoOutput{}, err
	}

	return GetUserInfoOutput{info: result}, nil
}
