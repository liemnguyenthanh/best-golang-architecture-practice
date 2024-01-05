// cmd/main.go
package main

import (
	"api-instagram/auth"
	"api-instagram/controllers"
	db "api-instagram/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	client := r.Group("/api")
	{
		// USERS
		client.GET("/users", auth.Auth(), controllers.GetUsers)
		client.POST("/users/create", controllers.CreateUser)

		// POSTS
		client.GET("/posts", controllers.GetPosts)
		client.POST("/posts/create", controllers.CreatePost)

		// AUTH
		client.POST("/auth/login", controllers.Login)

		// PING
		client.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world")
		})
	}
	return r
}

func main() {
	db.ConnectDb()

	r := SetupRouter()
	r.Run(":8000")
}
