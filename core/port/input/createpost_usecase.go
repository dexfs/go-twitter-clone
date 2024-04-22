package input

import "github.com/dexfs/go-twitter-clone/core/domain"

type CreatePostUseCaseInput struct {
	UserID  string
	Content string
}

type CreatePostUseCase interface {
	Execute(aInput CreatePostUseCaseInput) (*domain.Post, error)
}
