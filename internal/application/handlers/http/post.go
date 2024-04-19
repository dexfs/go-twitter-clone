package http

import (
	"github.com/dexfs/go-twitter-clone/internal/application/handlers"
	"github.com/dexfs/go-twitter-clone/internal/application/usecases/post_usecases"
	"net/http"
)

type CreatePostHandler struct {
	Path    string
	useCase *post_usecases.CreatePostUseCase
}

func NewCreatePostHandler(useCase *post_usecases.CreatePostUseCase) *CreatePostHandler {
	return &CreatePostHandler{
		Path:    "POST /posts",
		useCase: useCase,
	}
}

func (h *CreatePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input post_usecases.CreatePostInput
	err := handlers.DecodeJSON(r, &input)
	if err != nil {
		handlers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	post, err := h.useCase.Execute(input)
	if err != nil {
		handlers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	handlers.JSON(w, http.StatusCreated, post)
}

// Create Quote
type CreateQuoteHandler struct {
	Path    string
	useCase *post_usecases.CreateQuotePostUseCase
}

func NewCreateQuoteHandler(useCase *post_usecases.CreateQuotePostUseCase) *CreateQuoteHandler {
	return &CreateQuoteHandler{
		Path:    "POST /posts/quote",
		useCase: useCase,
	}
}

func (h *CreateQuoteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input post_usecases.CreateQuotePostUseCaseInput

	err := handlers.DecodeJSON(r, &input)
	if err != nil {
		handlers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	quote, err := h.useCase.Execute(input)
	if err != nil {
		handlers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	handlers.JSON(w, http.StatusCreated, quote)
}

//

type CreateRepostHandler struct {
	Path    string
	useCase *post_usecases.CreateRepostUseCase
}

func NewRepostHandler(useCase *post_usecases.CreateRepostUseCase) *CreateRepostHandler {
	return &CreateRepostHandler{
		Path:    "POST /posts/repost",
		useCase: useCase,
	}
}

func (h *CreateRepostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input post_usecases.CreateRepostUseCaseInput
	err := handlers.DecodeJSON(r, &input)
	if err != nil {
		handlers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	repost, err := h.useCase.Execute(input)
	if err != nil {
		handlers.JSONError(w, http.StatusBadRequest, err)
		return
	}

	handlers.JSON(w, http.StatusCreated, repost)
}
