package main

import (
	"bytes"
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
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setUpTests() {
	gin.SetMode(gin.TestMode)
}

func InitDependencies(db *database.InMemoryDB) (userRepo output.UserPort, port output.PostPort) {
	userRepo = inmemory.NewInMemoryUserRepository(db)
	postRepo = inmemory.NewInMemoryPostRepository(db)
	return userRepo, postRepo
}

func helperDecodeJSON(body *bytes.Buffer, v interface{}) error {
	if err := json.Unmarshal([]byte(body.String()), v); err != nil {
		return fmt.Errorf("could not decode JSON: %v", err)
	}
	return nil
}

func TestInvalidUserName(t *testing.T) {
	var tests = []struct {
		name      string
		input     string
		want      string
		causesLen int
	}{
		{" username should be alphanumeric", "not@found", "Bad Request", 1},
		{" username should be alphanumeric", "not_found", "Bad Request", 1},
		{" username should have at least 5 characters", "nf", "Bad Request", 1},
		{" username should have max 10 characters", "notfound123456789", "Bad Request", 1},
	}
	setUpTests()
	dbMocks := mocks.GetTestMocks()
	userRepo, _ := InitDependencies(dbMocks.MockDB)
	getUserInfoService, _ := usecase.NewGetUserInfoUseCase(userRepo)
	usersController := adapter_http.NewUsersController(getUserInfoService, nil)

	router := routes.NewRouter(":8002")
	router.Router.GET("/users/:username/info", usersController.GetInfo)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wRecorder := httptest.NewRecorder()
			url := fmt.Sprintf("/users/%s/info", tt.input)
			req, _ := http.NewRequest("GET", url, nil)
			router.Router.ServeHTTP(wRecorder, req)

			got := new(rest_errors.RestError)

			assert.Equal(t, http.StatusBadRequest, wRecorder.Code)
			if err := helperDecodeJSON(wRecorder.Body, got); err != nil {
				log.Fatal(err)
			}

			assert.EqualValues(t, tt.want, got.Message)
			assert.Len(t, got.Causes, tt.causesLen)

		})
	}
}

func TestUserInfoResource_WithNoFoundUser_ReturnsErrorMessage(t *testing.T) {
	setUpTests()
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
	setUpTests()
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

func TestUserFeedResource(t *testing.T) {
	setUpTests()
	dbMocks := mocks.GetTestMocks()
	userRepo, postRepo := InitDependencies(dbMocks.MockDB)
	wRecorder := httptest.NewRecorder()

	router := routes.NewRouter(":8002")
	getUserFeeUseCase, _ := usecase.NewGetUserFeedUseCase(userRepo, postRepo)
	usersController := adapter_http.NewUsersController(nil, getUserFeeUseCase)
	router.Router.GET("/users/:username/feed", usersController.GetFeed)

	req, _ := http.NewRequest("GET", "/users/user0/feed", nil)
	router.Router.ServeHTTP(wRecorder, req)

	var got response.GetUserFeedResponse
	if err := json.Unmarshal([]byte(wRecorder.Body.String()), &got); err != nil {
		log.Fatal(err)
	}

	assert.Len(t, got.Items, 2)
}

func TestUserFeedResource_WithNoFoundUser_ReturnsErrorMessage(t *testing.T) {
	setUpTests()
	dbMocks := mocks.GetTestMocks()
	userRepo, postRepo := InitDependencies(dbMocks.MockDB)
	wRecorder := httptest.NewRecorder()

	router := routes.NewRouter(":8002")
	getUserFeeUseCase, _ := usecase.NewGetUserFeedUseCase(userRepo, postRepo)
	usersController := adapter_http.NewUsersController(nil, getUserFeeUseCase)
	router.Router.GET("/users/:username/feed", usersController.GetFeed)

	req, _ := http.NewRequest("GET", "/users/notfound/feed", nil)
	router.Router.ServeHTTP(wRecorder, req)

	var got rest_errors.RestError

	assert.Equal(t, http.StatusNotFound, wRecorder.Code)

	if err := helperDecodeJSON(wRecorder.Body, &got); err != nil {
		log.Fatal(err)
	}

	assert.EqualValues(t, "user not found", got.Message)
}
