package routes

import (
	"errors"
	"github.com/dexfs/go-twitter-clone/adapter/input/adapter_http"
	"github.com/dexfs/go-twitter-clone/internal/core/usecase"
	"github.com/fvbock/endless"
	"time"

	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/gin-gonic/gin"
)

type AppServer struct {
	router *gin.Engine
	addr   string
}

func NewRouter(addr string) *AppServer {
	//gin.SetMode(gin.DebugMode)
	//gin.ForceConsoleColor()
	return &AppServer{
		router: gin.Default(),
		addr:   addr,
	}
}

func (s *AppServer) Run() error {
	s.router.SetTrustedProxies(nil)
	db, err := s.initDatabase()
	if err != nil {
		return err
	}
	s.initRoutes(db)
	return endless.ListenAndServe(s.addr, s.router)
}

func (s *AppServer) initRoutes(db *database.InMemoryDB) {
	// repositories
	userRepo := inmemory.NewInMemoryUserRepository(db)
	postRepo := inmemory.NewInMemoryPostRepository(db)
	// usecases
	getUserInfoService, _ := usecase.NewGetUserInfoUseCase(userRepo)
	getUserFeedUseCase, _ := usecase.NewGetUserFeedUseCase(userRepo, postRepo)
	createPostUseCase, _ := usecase.NewCreatePostUseCase(postRepo, userRepo)
	createRepostUseCase, _ := usecase.NewCreateRepostUseCase(postRepo, userRepo)
	createQuoteUseCase, _ := usecase.NewCreateQuoteUseCase(postRepo, userRepo)
	//controllers
	usersController := adapter_http.NewUsersController(getUserInfoService, getUserFeedUseCase)

	postsController := adapter_http.NewPostsController(createPostUseCase, createRepostUseCase, createQuoteUseCase)
	// routes
	s.router.GET("/users/:username/info", usersController.GetInfo)
	s.router.GET("/users/:username/feed", usersController.GetFeed)
	s.router.POST("/posts", postsController.CreatePost)
	s.router.POST("/posts/repost", postsController.CreateRepost)
	s.router.POST("/posts/quote", postsController.CreateQuote)
}

func (s *AppServer) initDatabase() (*database.InMemoryDB, error) {
	db := database.NewInMemoryDB()
	if db == nil {
		return nil, errors.New("database is nil")
	}
	initialUsers := make([]*inmemory_schema.UserSchema, 0)
	initialUsers = append(initialUsers, &inmemory_schema.UserSchema{
		ID:        "4cfe67a9-defc-42b9-8410-cb5086bec2f5",
		Username:  "alucard",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	initialUsers = append(initialUsers, &inmemory_schema.UserSchema{
		ID:        "b8903f77-5d16-4176-890f-f597594ff952",
		Username:  "alexander",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	db.RegisterSchema(inmemory.USER_SCHEMA_NAME, initialUsers)
	db.RegisterSchema(inmemory.POST_SCHEMA_NAME, []*inmemory_schema.PostSchema{})
	return db, nil
}
