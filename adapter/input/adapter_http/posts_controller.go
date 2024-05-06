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
	createPostUseCase   input.CreatePostUseCase
	createRepostUseCase input.CreateRepostUseCase
	createQuoteUseCase  input.CreateQuoteUseCase
}

func NewPostsController(createPostUseCase input.CreatePostUseCase, createRepostUseCase input.CreateRepostUseCase, createQuoteUseCase input.CreateQuoteUseCase) *postsController {
	return &postsController{
		createPostUseCase:   createPostUseCase,
		createRepostUseCase: createRepostUseCase,
		createQuoteUseCase:  createQuoteUseCase,
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

func (pc *postsController) CreateRepost(c *gin.Context) {
	createRequest := request.RepostRequest{}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		errRest := validation.RestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	postDomain, err := pc.createRepostUseCase.Execute(input.CreateRepostUseCaseInput{
		PostID: createRequest.PostID,
		UserID: createRequest.UserID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, response.CreatePostResponse{
		PostID: postDomain.ID,
	})
}

func (pc *postsController) CreateQuote(c *gin.Context) {
	createRequest := request.QuoteRequest{}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		errRest := validation.RestError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	postDomain, err := pc.createQuoteUseCase.Execute(input.CreateQuoteUseCaseInput{
		PostID: createRequest.PostID,
		UserID: createRequest.UserID,
		Quote:  createRequest.Quote,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, response.CreatePostResponse{
		PostID: postDomain.ID,
	})
}
