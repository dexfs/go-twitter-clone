package input

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type CreateRepostUseCaseInput struct {
	PostID string
	UserID string
}

type CreateRepostUseCase interface {
	Execute(ctx context.Context, input CreateRepostUseCaseInput) (*domain.Post, *rest_errors.RestError)
}
