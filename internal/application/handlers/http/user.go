package http

import (
	"github.com/dexfs/go-twitter-clone/internal/application/handlers"
	"github.com/dexfs/go-twitter-clone/internal/application/usecases/user_usecases"
	"net/http"
)

type GetUserFeedHandler struct {
	useCase *user_usecases.GetUserFeedUseCase
}

type GetUserInfoHandler struct {
	useCase *user_usecases.GetUserInfoUseCase
}

func NewGetFeedHandler(useCase *user_usecases.GetUserFeedUseCase) GetUserFeedHandler {
	return GetUserFeedHandler{
		useCase: useCase,
	}
}

func NewGetUserInfoHandler(useCase *user_usecases.GetUserInfoUseCase) GetUserInfoHandler {
	return GetUserInfoHandler{
		useCase: useCase,
	}
}

func (h GetUserFeedHandler) Handle(w http.ResponseWriter, r *http.Request) {
	feed, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		handlers.JSONError(w, http.StatusBadRequest, err)
		return
	}
	handlers.JSON(w, http.StatusOK, feed)
}

func (h GetUserInfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	user, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		handlers.JSONError(w, http.StatusBadRequest, err)
		return
	}
	handlers.JSON(w, http.StatusCreated, user)
}
