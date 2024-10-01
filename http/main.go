package main

import (
	"go-demo/http/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.POST("users/new", handlers.NewUser)
	r.GET("/users", handlers.GetAllUsers)
	r.GET("/users/:userID", handlers.GetUserById)
	err := r.Run(":8080")
	if err != nil {
		log.Printf("router.Run error, %s", err.Error())
	}
}
