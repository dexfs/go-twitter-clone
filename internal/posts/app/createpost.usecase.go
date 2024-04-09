package app

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/internal/posts"
	postDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/posts/domain/interfaces"
	userDomainInterfaces "github.com/dexfs/go-twitter-clone/internal/user/domain/interfaces"
)

type CreatePostUseCase struct {
	userRepo userDomainInterfaces.UserRepository
	postRepo postDomainInterfaces.PostRepository
}

func NewCreatePostUseCase(userRepo userDomainInterfaces.UserRepository, postRepo postDomainInterfaces.PostRepository) *CreatePostUseCase {
	return &CreatePostUseCase{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

type CreatePostInput struct {
	UserID  string
	Content string
}

type CreatePostOutput struct {
	PostID string
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
	newPostInput := posts.NewPostInput{
		User:    user,
		Content: input.Content,
	}

	newPost, err := posts.NewPost(newPostInput)

	if err != nil {
		return CreatePostOutput{}, err
	}

	uc.postRepo.Insert(newPost)

	return CreatePostOutput{
		PostID: newPost.ID,
	}, nil
}
