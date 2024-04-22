package input

import "github.com/dexfs/go-twitter-clone/core/domain"

type GetUserInfoUseCase interface {
	Execute(username string) (*domain.User, error)
}
