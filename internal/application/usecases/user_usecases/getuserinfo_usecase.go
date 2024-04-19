package user_usecases

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/user"
)

type GetUserInfoUseCase struct {
	userRepo user.UserRepository
}

type GetUserInfoOutput struct {
	*user.User
}

func NewGetUserInfoUseCase(userRepo user.UserRepository) (*GetUserInfoUseCase, error) {
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
