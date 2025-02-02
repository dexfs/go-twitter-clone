package input

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type GetUserInfoUseCase interface {
	Execute(ctx context.Context, username string) (*domain.User, *rest_errors.RestError)
}
