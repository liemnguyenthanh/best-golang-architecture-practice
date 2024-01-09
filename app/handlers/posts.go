package handlers

import (
	"api-instagram/app/middlewares"
	"api-instagram/app/models"
	"api-instagram/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewPost struct {
	User_id int
	Content string
}

func GetPosts(c *gin.Context) {
	// db := database.ConnectDb()

	c.JSON(http.StatusOK, gin.H{
		"message": "post List",
		"posts":   "posts",
	})
}

func CreatePost(c *gin.Context) {

	var post models.Posts

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	// Extract the user ID from the request headers
	user_id, err := middlewares.ExtractUserIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":    "Unauthorized",
			"username": user_id,
		})
		return
	}

	// Check if the user with the given ID exists
	var user models.Users
	result := db.Instance.Table("users").First(&user, user_id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Save post in database
	// Create a new post
	newPost := NewPost{
		User_id: user.Id,
		Content: post.Content,
	}

	if err := db.Instance.Table("Posts").Create(&newPost).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created post",
		"post":    "json",
	})
}
