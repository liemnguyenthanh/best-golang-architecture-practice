// cmd/main.go
package main

import (
	"api-instagram/app/handlers"
	"api-instagram/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	client := r.Group("/api")
	{
		// USERS
		client.GET("/users", handlers.GetUsers)
		client.POST("/users/create", handlers.CreateUser)

		// POSTS
		client.GET("/posts", handlers.GetPosts)
		client.POST("/posts/create", handlers.CreatePost)

		// AUTH
		client.POST("/auth/login", handlers.Login)

		// PING
		client.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "server is living!!!")
		})
	}
	return r
}

func main() {
	db.ConnectDb()

	r := SetupRouter()
	r.Run(":8000")
}
