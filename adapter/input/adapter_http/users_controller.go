package adapter_http

import "C"
import (
	"github.com/dexfs/go-twitter-clone/adapter/input/model/request"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/response"
	"github.com/dexfs/go-twitter-clone/config/validation"
	"github.com/dexfs/go-twitter-clone/internal/core/domain"
	"github.com/dexfs/go-twitter-clone/internal/core/port/input"
	"github.com/gin-gonic/gin"
	"net/http"
)

type usersController struct {
	getUserInfoUseCase input.GetUserInfoUseCase
	getUserFeedUseCase input.GetUserFeedUseCase
}

func NewUsersController(getUserInfoUseCase input.GetUserInfoUseCase, getUserFeedUseCase input.GetUserFeedUseCase) *usersController {
	return &usersController{
		getUserInfoUseCase,
		getUserFeedUseCase,
	}
}

func (uc *usersController) GetInfo(c *gin.Context) {
	userInfoRequest := request.UserInfoRequest{}

	if err := c.ShouldBindUri(&userInfoRequest); err != nil {
		errRest := validation.RestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
	userFeedDomain, err := uc.getUserInfoUseCase.Execute(userInfoRequest.Username)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response.GetUserInfoResponse{
		ID:        userFeedDomain.ID,
		Username:  userFeedDomain.Username,
		CreatedAt: userFeedDomain.CreatedAt.Format("2006-01-02"),
	})
}

func (uc *usersController) GetFeed(c *gin.Context) {
	userRequest := request.UserFeedRequest{}

	if err := c.ShouldBindUri(&userRequest); err != nil {
		errRest := validation.RestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	userFeedDomain, err := uc.getUserFeedUseCase.Execute(userRequest.Username)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	var items []*domain.Post
	if len(userFeedDomain) == 0 {
		items = make([]*domain.Post, 0)
	} else {
		items = userFeedDomain
	}
	c.JSON(http.StatusOK, response.GetUserFeedResponse{
		Items: items,
	})
}
