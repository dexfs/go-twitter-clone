package usecase

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/internal/core/port/input"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
)

type createRepostUseCase struct {
	postPort output.PostPort
	userPort output.UserPort
}

func NewCreateRepostUseCase(postPort output.PostPort, userPort output.UserPort) (*createRepostUseCase, *rest_errors.RestError) {
	if postPort == nil || userPort == nil {
		return nil, rest_errors.NewInternalServerError("postPort and userPort cannot be nil")
	}

	return &createRepostUseCase{postPort: postPort, userPort: userPort}, nil
}

func (uc *createRepostUseCase) Execute(aInput input.CreateRepostUseCaseInput) (*domain.Post, *rest_errors.RestError) {
	user, err := uc.userPort.FindByID(aInput.UserID)

	if err != nil {
		return &domain.Post{}, rest_errors.NewBadRequestError(err.Error())
	}

	isReposted := uc.postPort.HasPostBeenRepostedByUser(aInput.PostID, aInput.UserID)
	if isReposted {
		return &domain.Post{}, rest_errors.NewBadRequestError("it is not possible repost a repost post")
	}

	post, err := uc.postPort.FindByID(aInput.PostID)
	if err != nil {
		return &domain.Post{}, rest_errors.NewBadRequestError(err.Error())
	}

	aNewPostInput := domain.NewRepostQuoteInput{
		User:    user,
		Post:    post,
		Content: "",
	}

	newRepost, err := domain.NewRepost(aNewPostInput)
	if err != nil {
		return &domain.Post{}, rest_errors.NewInternalServerError(err.Error())
	}

	uc.postPort.CreatePost(newRepost)

	return newRepost, nil
}
