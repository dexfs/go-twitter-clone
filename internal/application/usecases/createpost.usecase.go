package app

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/domain"
	"github.com/dexfs/go-twitter-clone/internal/domain/interfaces"
)

type CreatePostUseCase struct {
	userRepo interfaces.UserRepository
	postRepo interfaces.PostRepository
}

func NewCreatePostUseCase(userRepo interfaces.UserRepository, postRepo interfaces.PostRepository) *CreatePostUseCase {
	return &CreatePostUseCase{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

type CreatePostInput struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

type CreatePostOutput struct {
	PostID string `json:"post_id"`
}

func (uc *CreatePostUseCase) Execute(input CreatePostInput) (CreatePostOutput, error) {
	// verifica se já atingiu o limite de postagens do dia retornar um erro
	hasReachedLimit := uc.postRepo.HasReachedPostingLimitDay(input.UserID, 5)
	if hasReachedLimit {
		return CreatePostOutput{}, errors.New("you reached your posts day limit")
	}

	// verifica se o usuário existe
	user, err := uc.userRepo.FindByID(input.UserID)

	if err != nil {
		return CreatePostOutput{}, err
	}
	newPostInput := domain.NewPostInput{
		User:    user,
		Content: input.Content,
	}

	newPost, err := domain.NewPost(newPostInput)

	if err != nil {
		return CreatePostOutput{}, err
	}

	uc.postRepo.Insert(newPost)

	return CreatePostOutput{
		PostID: newPost.ID,
	}, nil
}
