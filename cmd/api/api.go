package main

import (
	"fmt"
	"github.com/dexfs/go-twitter-clone/internal/infra/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/internal/post"
	post_handler "github.com/dexfs/go-twitter-clone/internal/post/handler"
	post_usecase "github.com/dexfs/go-twitter-clone/internal/post/usecase"
	"github.com/dexfs/go-twitter-clone/internal/user"
	user_handler "github.com/dexfs/go-twitter-clone/internal/user/handler"
	user_usecase "github.com/dexfs/go-twitter-clone/internal/user/usecase"
	"github.com/dexfs/go-twitter-clone/mocks"
	"log"
	"net/http"
	"time"
)

// Middlewares

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
func RequestJsonContentTypeMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
}

/** @see https://www.youtube.com/watch?v=npzXQSL4oWo **/
type APIServer struct {
	addr string
}

type Gateway struct {
	userRepo user.UserRepository
	postRepo post.PostRepository
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: ":" + addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	gateways := s.initGateways()
	s.initUserRoutes(router, gateways)
	s.initPostRoutes(router, gateways)

	// router prefix
	//router.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	server := http.Server{
		Addr:    s.addr,
		Handler: RequestLoggerMiddleware(RequestJsonContentTypeMiddleware(router)),
	}

	log.Printf("API server listening on %s\n", s.addr)

	return server.ListenAndServe()
}

func (s *APIServer) initGateways() *Gateway {
	dbMocks := mocks.GetTestMocks()
	dbMocks.MockUserDB.Insert(&user.User{
		ID:        "4cfe67a9-defc-42b9-8410-cb5086bec2f5",
		Username:  "alucard",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	dbMocks.MockUserDB.Insert(&user.User{
		ID:        "b8903f77-5d16-4176-890f-f597594ff952",
		Username:  "alexander",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	dbMocks.MockUserDB.Insert(&user.User{
		ID:        "75135a97-46be-405f-8948-0821290ca83e",
		Username:  "seras_victoria",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	userRepo := repo_inmemory.NewInMemoryUserRepo(dbMocks.MockUserDB)
	postRepo := repo_inmemory.NewInMemoryPostRepo(dbMocks.MockPostDB)

	return &Gateway{
		userRepo: userRepo,
		postRepo: postRepo,
	}
}

func (s *APIServer) initUserRoutes(router *http.ServeMux, gateways *Gateway) {

	getUserFeed, err := user_usecase.NewGetUserFeedUseCase(gateways.userRepo, gateways.postRepo)
	if err != nil {
		log.Fatal(err)
	}

	getUserInfo, err := user_usecase.NewGetUserInfoUseCase(gateways.userRepo)
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("GET /users/{username}/info", user_handler.NewGetUserInfoHandler(getUserInfo).Handle)
	router.HandleFunc("GET /users/{username}/feed", user_handler.NewGetFeedHandler(getUserFeed).Handle)
}
func (s *APIServer) initPostRoutes(router *http.ServeMux, gateways *Gateway) {

	createPostUseCase := post_usecase.NewCreatePostUseCase(gateways.userRepo, gateways.postRepo)
	createPostHandler := post_handler.NewCreatePostHandler(createPostUseCase)
	router.HandleFunc(createPostHandler.Path, createPostHandler.Handle)

	createQuotePostUseCase := post_usecase.NewCreateQuotePostUseCase(gateways.userRepo, gateways.postRepo)
	crateQuotePostHandler := post_handler.NewCreateQuoteHandler(createQuotePostUseCase)
	router.HandleFunc(crateQuotePostHandler.Path, crateQuotePostHandler.Handle)

	createRepostUseCase := post_usecase.NewCreateRepostUseCase(gateways.userRepo, gateways.postRepo)
	createRepostHandler := post_handler.NewRepostHandler(createRepostUseCase)
	router.HandleFunc(createRepostHandler.Path, createRepostHandler.Handle)
}

func main() {
	server := NewAPIServer("8001")

	if err := server.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
