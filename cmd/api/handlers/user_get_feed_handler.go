package handlers

import (
	"github.com/dexfs/go-twitter-clone/internal/application"
	"net/http"
)

type getUserFeedHandler struct {
	useCase *application.GetUserFeedUseCase
}

func NewGetFeedHandler(useCase *application.GetUserFeedUseCase) getUserFeedHandler {
	return getUserFeedHandler{
		useCase: useCase,
	}
}

func (h getUserFeedHandler) Handle(w http.ResponseWriter, r *http.Request) {
	feed, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}
	JSON(w, http.StatusOK, feed)
}
