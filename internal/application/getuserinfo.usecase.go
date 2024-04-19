package application

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/domain"
)

type GetUserInfoUseCase struct {
	userRepo domain.UserRepository
}

type GetUserInfoOutput struct {
	*domain.User
}

func NewGetUserInfoUseCase(userRepo domain.UserRepository) (*GetUserInfoUseCase, error) {
	if userRepo == nil {
		return nil, errors.New("userRepo cannot be nil")
	}
	return &GetUserInfoUseCase{userRepo: userRepo}, nil
}

func (u *GetUserInfoUseCase) Execute(username string) (GetUserInfoOutput, error) {
	user, err := u.userRepo.ByUsername(username)
	if err != nil {
		return GetUserInfoOutput{}, err
	}

	return GetUserInfoOutput{user}, nil
}
