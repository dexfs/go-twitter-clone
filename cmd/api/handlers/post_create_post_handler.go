package handlers

import (
	"github.com/dexfs/go-twitter-clone/internal/application"
	"net/http"
)

type createPostHandler struct {
	Path    string
	useCase *application.CreatePostUseCase
}

func NewCreatePostHandler(useCase *application.CreatePostUseCase) *createPostHandler {
	return &createPostHandler{
		Path:    "POST /posts",
		useCase: useCase,
	}
}

func (h *createPostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input application.CreatePostInput
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
