package input

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type GetUserFeedUseCase interface {
	Execute(ctx context.Context, username string) ([]*domain.Post, *rest_errors.RestError)
}
