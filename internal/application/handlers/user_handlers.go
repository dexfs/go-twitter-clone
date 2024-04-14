package handlers

import (
	"github.com/dexfs/go-twitter-clone/internal/application/usecases"
	"net/http"
)

type GetUserFeedHandler struct {
	useCase *app.GetUserFeedUseCase
}

type GetUserInfoHandler struct {
	useCase *app.GetUserInfoUseCase
}

func NewGetFeedHandler(useCase *app.GetUserFeedUseCase) GetUserFeedHandler {
	return GetUserFeedHandler{
		useCase: useCase,
	}
}

func NewGetUserInfoHandler(useCase *app.GetUserInfoUseCase) GetUserInfoHandler {
	return GetUserInfoHandler{
		useCase: useCase,
	}
}

func (h GetUserFeedHandler) Handle(w http.ResponseWriter, r *http.Request) {
	feed, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}
	JSON(w, http.StatusOK, feed)
}

func (h GetUserInfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	user, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}
	JSON(w, http.StatusCreated, user)
}
