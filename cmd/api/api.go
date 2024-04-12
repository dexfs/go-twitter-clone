package main

import (
	"fmt"
	"log"
	"net/http"
)

// handlers
var UserFeedHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User feed of" + r.PathValue("username")))
}

var UserInfoHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User info for " + r.PathValue("username")))
}

var PostCreateHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post"))
}
var PostRepostHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post repost"))
}
var PostQuoteHandler = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post Quote"))
}

//mux.HandleFunc("GET /users/{username}/feed", UserFeedHandler)
//mux.HandleFunc("GET /users/{username}/info", func(w http.ResponseWriter, r *http.Request) {})
//mux.HandleFunc("POST /posts", func(w http.ResponseWriter, r *http.Request) {})
//mux.HandleFunc("POST /posts/repost", func(w http.ResponseWriter, r *http.Request) {})
//mux.HandleFunc("POST /posts/quote", func(w http.ResponseWriter, r *http.Request) {})
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

	router.HandleFunc("GET /users/{username}/feed", UserFeedHandler)
	router.HandleFunc("GET /users/{username}/info", UserInfoHandler)
	router.HandleFunc("POST /posts", PostCreateHandler)
	router.HandleFunc("POST /posts/repost", PostRepostHandler)
	router.HandleFunc("POST /posts/quote", PostQuoteHandler)
	// router prefix
	//router.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	server := http.Server{
		Addr:    s.addr,
		Handler: RequestLoggerMiddleware(RequestJsonContentTypeMiddleware(router)),
	}

	log.Printf("API server listening on %s\n", s.addr)

	return server.ListenAndServe()
}

func main() {
	server := NewAPIServer(":8001")

	if err := server.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
