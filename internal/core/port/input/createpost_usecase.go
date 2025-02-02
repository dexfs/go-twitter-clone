package input

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type CreatePostUseCaseInput struct {
	UserID  string
	Content string
}

type CreatePostUseCase interface {
	Execute(ctx context.Context, aInput CreatePostUseCaseInput) (*domain.Post, *rest_errors.RestError)
}
