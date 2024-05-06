package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type UseCasePort interface {
	Execute()
}

type UserUseCase struct{}

func (u UserUseCase) Execute() {
	fmt.Println("UserUseCase.Execute")
}

type UserController struct {
	userUseCase UseCasePort
}

func NewUserController(useCase UseCasePort) *UserController {
	return &UserController{
		userUseCase: useCase,
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello world"})
}

func main() {
	r := gin.Default()
	userUseCase := &UserUseCase{}
	userController := NewUserController(userUseCase)

	// routes
	r.GET("/users", userController.GetUser)

	if err := r.Run(":8001"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
