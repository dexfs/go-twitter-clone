package usecase

import (
	"context"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/internal/core/port/input"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
)

type createQuoteUseCase struct {
	userPort output.UserPort
	postPort output.PostPort
}

func NewCreateQuoteUseCase(postPort output.PostPort, userPort output.UserPort) (*createQuoteUseCase, *rest_errors.RestError) {
	if postPort == nil || userPort == nil {
		return nil, rest_errors.NewInternalServerError("postPort and userPort cannot be nil")
	}

	return &createQuoteUseCase{
		postPort: postPort,
		userPort: userPort,
	}, nil
}

func (uc *createQuoteUseCase) Execute(ctx context.Context, anInput input.CreateQuoteUseCaseInput) (*domain.Post, *rest_errors.RestError) {
	user, err := uc.userPort.FindByID(ctx, anInput.UserID)
	if err != nil {
		return &domain.Post{}, rest_errors.NewBadRequestError(err.Error())
	}

	post, err := uc.postPort.FindByID(ctx, anInput.PostID)
	if err != nil {
		return &domain.Post{}, rest_errors.NewBadRequestError(err.Error())
	}

	quotePostInput := domain.NewRepostQuoteInput{
		User:    user,
		Post:    post,
		Content: anInput.Quote,
	}
	newQuotePost, err := domain.NewQuote(quotePostInput)
	if err != nil {
		return &domain.Post{}, rest_errors.NewBadRequestError(err.Error())
	}

	err = uc.postPort.CreatePost(ctx, newQuotePost)

	if err != nil {
		return &domain.Post{}, rest_errors.NewBadRequestError(err.Error())
	}

	return newQuotePost, nil
}
