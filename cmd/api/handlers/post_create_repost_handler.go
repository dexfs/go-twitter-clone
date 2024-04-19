package handlers

import (
	"github.com/dexfs/go-twitter-clone/internal/application"
	"net/http"
)

type createRepostHandler struct {
	Path    string
	useCase *application.CreateRepostUseCase
}

func NewRepostHandler(useCase *application.CreateRepostUseCase) *createRepostHandler {
	return &createRepostHandler{
		Path:    "POST /posts/repost",
		useCase: useCase,
	}
}

func (h *createRepostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input application.CreateRepostUseCaseInput
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
