package main

import (
	"fmt"
	"github.com/dexfs/go-twitter-clone/adapter/input/routes"
	"log"
	"runtime/debug"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: ":" + addr,
	}
}

func (s *APIServer) Run() error {
	router := routes.NewRouter(s.addr)
	return router.Run()
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	log.Printf("Starting Application")
	server := NewAPIServer("8001")
	//
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
