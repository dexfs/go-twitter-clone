package input

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
)

type CreateQuoteUseCase interface {
	Execute(ctx context.Context, anInput CreateQuoteUseCaseInput) (*domain.Post, *rest_errors.RestError)
}
type CreateQuoteUseCaseInput struct {
	PostID string
	UserID string
	Quote  string
}
