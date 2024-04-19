package handlers

import (
	"github.com/dexfs/go-twitter-clone/internal/application"
	"net/http"
)

type getUserInfoHandler struct {
	useCase *application.GetUserInfoUseCase
}

func NewGetUserInfoHandler(useCase *application.GetUserInfoUseCase) getUserInfoHandler {
	return getUserInfoHandler{
		useCase: useCase,
	}
}

func (h getUserInfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	user, err := h.useCase.Execute(r.PathValue("username"))
	if err != nil {
		JSONError(w, http.StatusBadRequest, err)
		return
	}
	JSON(w, http.StatusCreated, user)
}
