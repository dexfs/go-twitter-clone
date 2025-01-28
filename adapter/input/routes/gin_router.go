package routes

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/adapter_http"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
	"github.com/dexfs/go-twitter-clone/internal/core/usecase"
	"github.com/fvbock/endless"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	Router *gin.Engine
	addr   string
}

func NewRouter(addr string) *AppServer {
	//gin.SetMode(gin.DebugMode)
	//gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return &AppServer{
		Router: r,
		addr:   addr,
	}
}

func (s *AppServer) Run(userRepo output.UserPort, postRepo output.PostPort) error {
	err := s.Router.SetTrustedProxies(nil)
	if err != nil {
		return err
	}

	s.InitUserResources(userRepo, postRepo)
	s.InitPostResources(userRepo, postRepo)
	return endless.ListenAndServe(s.addr, s.Router)
}
func (s *AppServer) InitUserResources(userRepo output.UserPort, postRepo output.PostPort) {
	getUserInfoService, _ := usecase.NewGetUserInfoUseCase(userRepo)
	getUserFeedUseCase, _ := usecase.NewGetUserFeedUseCase(userRepo, postRepo)

	usersController := adapter_http.NewUsersController(getUserInfoService, getUserFeedUseCase)

	s.Router.GET("/users/:username/info", usersController.GetInfo)
	s.Router.GET("/users/:username/feed", usersController.GetFeed)
}

func (s *AppServer) InitPostResources(userRepo output.UserPort, postRepo output.PostPort) {
	createPostUseCase, _ := usecase.NewCreatePostUseCase(postRepo, userRepo)
	createRepostUseCase, _ := usecase.NewCreateRepostUseCase(postRepo, userRepo)
	createQuoteUseCase, _ := usecase.NewCreateQuoteUseCase(postRepo, userRepo)

	postsController := adapter_http.NewPostsController(createPostUseCase, createRepostUseCase, createQuoteUseCase)

	s.Router.POST("/posts", postsController.CreatePost)
	s.Router.POST("/posts/repost", postsController.CreateRepost)
	s.Router.POST("/posts/quote", postsController.CreateQuote)
}
