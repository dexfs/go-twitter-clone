package input

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type GetUserInfoUseCase interface {
	Execute(username string) (*domain.User, *rest_errors.RestError)
}
