package http

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/model/request"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/response"
	"github.com/dexfs/go-twitter-clone/core/port/input"
	"github.com/gin-gonic/gin"
	"net/http"
)

type usersController struct {
	getUserInfoUseCase input.GetUserInfoUseCase
}

func NewUsersController(getUserInfoUseCase input.GetUserInfoUseCase) *usersController {
	return &usersController{
		getUserInfoUseCase,
	}
}

func (uc *usersController) GetInfo(c *gin.Context) {
	userInfoRequest := request.UserInfoRequest{}

	if err := c.ShouldBindUri(&userInfoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFeedDomain, err := uc.getUserInfoUseCase.Execute(userInfoRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetUserInfoResponse{
		ID:        userFeedDomain.ID,
		Username:  userFeedDomain.Username,
		CreatedAt: userFeedDomain.CreatedAt.Format("2006-01-02"),
	})
}
