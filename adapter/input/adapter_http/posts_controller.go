package adapter_http

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/model/request"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/response"
	"github.com/dexfs/go-twitter-clone/config/validation"
	"github.com/dexfs/go-twitter-clone/internal/core/port/input"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postsController struct {
	createPostUseCase input.CreatePostUseCase
}

func NewPostsController(createPostUseCase input.CreatePostUseCase) *postsController {
	return &postsController{
		createPostUseCase: createPostUseCase,
	}
}

func (pc *postsController) CreatePost(c *gin.Context) {
	createPostRequest := request.CreatePostRequest{}

	if err := c.ShouldBindJSON(&createPostRequest); err != nil {
		errRest := validation.RestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	postDomain, err := pc.createPostUseCase.Execute(input.CreatePostUseCaseInput{
		UserID:  createPostRequest.UserID,
		Content: createPostRequest.Content,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, response.CreatePostResponse{
		PostID: postDomain.ID,
	})
}
