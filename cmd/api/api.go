package main

import (
	"fmt"
	"github.com/dexfs/go-twitter-clone/internal/application/handlers"
	app "github.com/dexfs/go-twitter-clone/internal/application/usecases"
	"github.com/dexfs/go-twitter-clone/internal/infra/repository/inmemory"
	"github.com/dexfs/go-twitter-clone/tests/mocks"
	"log"
	"net/http"
)

// handlers

var PostCreateHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post"))
}
var PostRepostHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post repost"))
}
var PostQuoteHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post Quote"))
}

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

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	s.initUserRoutes(router)
	//router.HandleFunc("POST /posts", PostCreateHandler)
	//router.HandleFunc("POST /posts/repost", PostRepostHandler)
	//router.HandleFunc("POST /posts/quote", PostQuoteHandler)
	// router prefix
	//router.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	server := http.Server{
		Addr:    s.addr,
		Handler: RequestLoggerMiddleware(RequestJsonContentTypeMiddleware(router)),
	}

	log.Printf("API server listening on %s\n", s.addr)

	return server.ListenAndServe()
}

func (s *APIServer) initUserRoutes(router *http.ServeMux) {
	// dabatase
	dbMocks := mocks.GetTestMocks()

	//userDb := &database.InMemoryDB[domain.User]{}
	//postDb := &database.InMemoryDB[domain.Post]{}
	// repos
	userRepo := inmemory.NewInMemoryUserRepo(dbMocks.MockUserDB)
	postRepo := inmemory.NewInMemoryPostRepo(dbMocks.MockPostDB)
	// usecases
	getUserFeed, err := app.NewGetUserFeedUseCase(userRepo, postRepo)
	if err != nil {
		log.Fatal(err)
	}

	getUserInfo, err := app.NewGetUserInfoUseCase(userRepo)
	if err != nil {
		log.Fatal(err)
	}

	// handlers

	// routes
	router.HandleFunc("GET /users/{username}/info", handlers.NewGetUserInfoHandler(getUserInfo).Handle)
	router.HandleFunc("GET /users/{username}/feed", handlers.NewGetFeedHandler(getUserFeed).Handle)
}

func main() {
	server := NewAPIServer(":8001")

	if err := server.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
