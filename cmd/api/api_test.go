package main

import (
	"encoding/json"
	"fmt"
	"github.com/dexfs/go-twitter-clone/adapter/input/adapter_http"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/response"
	"github.com/dexfs/go-twitter-clone/adapter/input/model/rest_errors"
	"github.com/dexfs/go-twitter-clone/adapter/input/routes"
	"github.com/dexfs/go-twitter-clone/adapter/output/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/internal/core/port/output"
	"github.com/dexfs/go-twitter-clone/internal/core/usecase"
	"github.com/dexfs/go-twitter-clone/pkg/database"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserInfoResource_WithNoFoundUser_ReturnsErrorMessage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	dbMocks := mocks.GetTestMocks()
	userRepo, _ := InitDependencies(dbMocks.MockDB)
	wRecorder := httptest.NewRecorder()

	router := routes.NewRouter(":8002")

	getUserInfoService, _ := usecase.NewGetUserInfoUseCase(userRepo)
	usersController := adapter_http.NewUsersController(getUserInfoService, nil)
	router.Router.GET("/users/:username/info", usersController.GetInfo)

	req, _ := http.NewRequest("GET", "/users/notfound/info", nil)
	router.Router.ServeHTTP(wRecorder, req)

	var got rest_errors.RestError

	assert.Equal(t, http.StatusNotFound, wRecorder.Code)
	if err := helperDecodeJSON(wRecorder.Body, &got); err != nil {
		log.Fatal(err)
	}

	assert.EqualValues(t, got.Message, "user not found")
}

func TestUserInfoResource(t *testing.T) {
	gin.SetMode(gin.TestMode)
	dbMocks := mocks.GetTestMocks()
	userRepo, _ := InitDependencies(dbMocks.MockDB)
	wRecorder := httptest.NewRecorder()

	router := routes.NewRouter(":8002")
	getUserInfoService, _ := usecase.NewGetUserInfoUseCase(userRepo)
	usersController := adapter_http.NewUsersController(getUserInfoService, nil)
	router.Router.GET("/users/:username/info", usersController.GetInfo)

	req, _ := http.NewRequest("GET", "/users/user0/info", nil)
	router.Router.ServeHTTP(wRecorder, req)

	var got response.GetUserInfoResponse
	assert.Equal(t, http.StatusOK, wRecorder.Code)

	if err := helperDecodeJSON(wRecorder.Body, &got); err != nil {
		log.Fatal(err)
	}

	assert.EqualValues(t, "user0", got.Username)
}

func InitDependencies(db *database.InMemoryDB) (userRepo output.UserPort, port output.PostPort) {
	userRepo = inmemory.NewInMemoryUserRepository(db)
	postRepo = inmemory.NewInMemoryPostRepository(db)
	return userRepo, postRepo
}

func helperDecodeJSON(body io.Reader, v interface{}) error {
	if err := json.NewDecoder(body).Decode(v); err != nil {
		return fmt.Errorf("could not decode JSON: %v", err)
	}
	return nil
}
