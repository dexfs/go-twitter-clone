package user_handler

import (
	user_usecase "github.com/dexfs/go-twitter-clone/internal/user/usecase"
	"github.com/dexfs/go-twitter-clone/pkg/helpers"
	"net/http"
)

type GetUserFeedHandler struct {
	useCase *user_usecase.GetUserFeedUseCase
}

type GetUserInfoHandler struct {
	useCase *user_usecase.GetUserInfoUseCase
}

func NewGetFeedHandler(useCase *user_usecase.GetUserFeedUseCase) GetUserFeedHandler {
	return GetUserFeedHandler{
		useCase: useCase,
	}
}

func NewGetUserInfoHandler(useCase *user_usecase.GetUserInfoUseCase) GetUserInfoHandler {
	return GetUserInfoHandler{
		useCase: useCase,
	}
}

func (h GetUserFeedHandler) Handle(w http.ResponseWriter, r *http.Request) {
	feed, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err)
		return
	}
	helpers.JSON(w, http.StatusOK, feed)
}

func (h GetUserInfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	user, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		helpers.JSONError(w, http.StatusBadRequest, err)
		return
	}
	helpers.JSON(w, http.StatusCreated, user)
}
