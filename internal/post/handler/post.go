package post_handler

import (
	post_usecase "github.com/dexfs/go-twitter-clone/internal/post/usecase"
	"github.com/dexfs/go-twitter-clone/pkg/helpers"
	"net/http"
)

type CreatePostHandler struct {
	Path    string
	useCase *post_usecase.CreatePostUseCase
}

func NewCreatePostHandler(useCase *post_usecase.CreatePostUseCase) *CreatePostHandler {
	return &CreatePostHandler{
		Path:    "POST /posts",
		useCase: useCase,
	}
}

func (h *CreatePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input post_usecase.CreatePostInput
	err := helpers.DecodeJSON(r, &input)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	post, err := h.useCase.Execute(input)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, post)
}

// Create Quote
type CreateQuoteHandler struct {
	Path    string
	useCase *post_usecase.CreateQuotePostUseCase
}

func NewCreateQuoteHandler(useCase *post_usecase.CreateQuotePostUseCase) *CreateQuoteHandler {
	return &CreateQuoteHandler{
		Path:    "POST /posts/quote",
		useCase: useCase,
	}
}

func (h *CreateQuoteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input post_usecase.CreateQuotePostUseCaseInput

	err := helpers.DecodeJSON(r, &input)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	quote, err := h.useCase.Execute(input)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, quote)
}

//

type CreateRepostHandler struct {
	Path    string
	useCase *post_usecase.CreateRepostUseCase
}

func NewRepostHandler(useCase *post_usecase.CreateRepostUseCase) *CreateRepostHandler {
	return &CreateRepostHandler{
		Path:    "POST /posts/repost",
		useCase: useCase,
	}
}

func (h *CreateRepostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input post_usecase.CreateRepostUseCaseInput
	err := helpers.DecodeJSON(r, &input)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	repost, err := h.useCase.Execute(input)
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	helpers.JSON(w, http.StatusCreated, repost)
}
