package handlers

import (
	"github.com/dexfs/go-twitter-clone/internal/application/usecases"
	"net/http"
)

type GetUserFeedHandler struct {
	useCase *app.GetUserFeedUseCase
}

func NewGetFeedHandler(useCase *app.GetUserFeedUseCase) GetUserFeedHandler {
	return GetUserFeedHandler{
		useCase: useCase,
	}
}

func (h GetUserFeedHandler) Handle(w http.ResponseWriter, r *http.Request) {
	feed, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		JSONError(w, http.StatusInternalServerError, err)
		return
	}
	JSON(w, http.StatusOK, feed)
}
