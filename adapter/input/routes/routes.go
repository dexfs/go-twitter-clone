package routes

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/http"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	inmemory_schema "github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory/schema"
	"github.com/dexfs/go-twitter-clone/internal/core/usecase"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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
	db := s.initDatabase()
	s.initRoutes(db)
	return s.router.Run(s.addr)
}

func (s *AppServer) initRoutes(db *database.InMemoryDB) {
	userRepo := inmemory.NewInMemoryUserRepository(db)
	postRepo := inmemory.NewInMemoryPostRepository(db)

	getUserInfoService, _ := usecase.NewGetUserInfoUseCase(userRepo)
	getUserFeedUseCase, _ := usecase.NewGetUserFeedUseCase(userRepo, postRepo)

	usersController := http.NewUsersController(getUserInfoService, getUserFeedUseCase)
	//
	s.router.GET("/users/:username/info", usersController.GetInfo)
	s.router.GET("/users/:username/feed", usersController.GetFeed)
	//
	createPostUseCase, err := usecase.NewCreatePostUseCase(postRepo, userRepo)

	if err != nil {
		log.Fatal(err)
		return
	}

	postsController := http.NewPostsController(createPostUseCase)
	s.router.POST("/posts", postsController.CreatePost)

}

func (s *AppServer) initDatabase() *database.InMemoryDB {
	db := database.NewInMemoryDB()
	if db == nil {
		panic("failed to connect to database")
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
	return db
}
