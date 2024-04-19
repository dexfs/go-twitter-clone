package handlers

import (
	"github.com/dexfs/go-twitter-clone/internal/application"
	"net/http"
)

// Create Quote
type CreateQuoteHandler struct {
	Path    string
	useCase *application.CreateQuotePostUseCase
}

func NewCreateQuoteHandler(useCase *application.CreateQuotePostUseCase) *CreateQuoteHandler {
	return &CreateQuoteHandler{
		Path:    "POST /posts/quote",
		useCase: useCase,
	}
}

func (h *CreateQuoteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input application.CreateQuotePostUseCaseInput

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
