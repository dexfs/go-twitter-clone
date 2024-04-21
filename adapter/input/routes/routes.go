package routes

import (
	"github.com/dexfs/go-twitter-clone/adapter/input/http"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/core/port/output"
	"github.com/dexfs/go-twitter-clone/core/usecase"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	userRepo output.UserPort
)

func InitRoutes(r *gin.Engine) {
	initAdapters()
	getUserInfoService := usecase.NewGetUserInfoService(userRepo)
	usersController := http.NewUsersController(getUserInfoService)

	r.GET("/users/:username/info", usersController.GetInfo)
}

func initAdapters() {
	userDb := &database.InMemoryDB[inmemory.UserSchema]{}
	userDb.Insert(&inmemory.UserSchema{
		ID:        "userid",
		Username:  "seeduser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	userRepo = inmemory.NewInMemoryUserRepository(userDb)
}
