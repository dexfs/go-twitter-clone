package routes

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/http"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/core/port/output"
	"github.com/dexfs/go-twitter-clone/core/usecase"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	userRepo output.UserPort
	postRepo output.PostPort
)

type AppServer struct {
	router *gin.Engine
	addr   string
}

func NewRouter(addr string) *AppServer {
	gin.SetMode(gin.DebugMode)
	gin.ForceConsoleColor()
	return &AppServer{
		router: gin.Default(),
		addr:   addr,
	}
}

func (s *AppServer) Run() error {
	s.router.SetTrustedProxies(nil)
	s.initAdapters()
	s.initUserRoutes()
	s.initPostRoutes()
	return s.router.Run(s.addr)
}

func (s *AppServer) initPostRoutes() {
	createPostUseCase, _ := usecase.NewCreatePostUseCase(postRepo, userRepo)
	postsController := http.NewPostsController(createPostUseCase)

	s.router.POST("/posts", postsController.CreatePost)
}

func (s *AppServer) initUserRoutes() {
	getUserInfoService, _ := usecase.NewGetUserInfoUseCase(userRepo)
	getUserFeedUseCase, _ := usecase.NewGetUserFeedUseCase(userRepo, postRepo)
	usersController := http.NewUsersController(getUserInfoService, getUserFeedUseCase)

	s.router.GET("/users/:username/info", usersController.GetInfo)
	s.router.GET("/users/:username/feed", usersController.GetFeed)
}

func (s *AppServer) initAdapters() {
	userDb := &database.InMemoryDB[inmemory_schema.UserSchema]{}
	userDb.Insert(&inmemory_schema.UserSchema{
		ID:        "4cfe67a9-defc-42b9-8410-cb5086bec2f5",
		Username:  "alucard",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	userDb.Insert(&inmemory_schema.UserSchema{
		ID:        "b8903f77-5d16-4176-890f-f597594ff952",
		Username:  "alexander",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	userRepo = inmemory.NewInMemoryUserRepository(userDb)

	postDB := &database.InMemoryDB[inmemory_schema.PostSchema]{}
	postRepo = inmemory.NewInMemoryPostRepository(postDB)
}
