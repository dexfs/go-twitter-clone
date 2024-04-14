package handlers

import (
	app "github.com/dexfs/go-twitter-clone/internal/application/usecases"
	"net/http"
)

type CreatePostHandler struct {
	Path    string
	useCase *app.CreatePostUseCase
}

func NewCreatePostHandler(useCase *app.CreatePostUseCase) *CreatePostHandler {
	return &CreatePostHandler{
		Path:    "POST /posts",
		useCase: useCase,
	}
}

func (h *CreatePostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input app.CreatePostInput
	err := DecodeJSON(r, &input)
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}

	post, err := h.useCase.Execute(input)
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}

	JSON(w, http.StatusCreated, post)
}

// Create Quote
type CreateQuoteHandler struct {
	Path    string
	useCase *app.CreateQuotePostUseCase
}

func NewCreateQuoteHandler(useCase *app.CreateQuotePostUseCase) *CreateQuoteHandler {
	return &CreateQuoteHandler{
		Path:    "POST /posts/quote",
		useCase: useCase,
	}
}

func (h *CreateQuoteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input app.CreateQuotePostUseCaseInput

	err := DecodeJSON(r, &input)
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}

	quote, err := h.useCase.Execute(input)
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}

	JSON(w, http.StatusCreated, quote)
}

//

type CreateRepostHandler struct {
	Path    string
	useCase *app.CreateRepostUseCase
}

func NewRepostHandler(useCase *app.CreateRepostUseCase) *CreateRepostHandler {
	return &CreateRepostHandler{
		Path:    "POST /posts/repost",
		useCase: useCase,
	}
}

func (h *CreateRepostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input app.CreateRepostUseCaseInput
	err := DecodeJSON(r, &input)
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}

	repost, err := h.useCase.Execute(input)
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}

	JSON(w, http.StatusCreated, repost)
}
