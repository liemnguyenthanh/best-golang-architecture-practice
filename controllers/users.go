package controllers

import (
	"api-instagram/auth"
	db "api-instagram/database"
	"api-instagram/models"
	"api-instagram/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	var users []models.Users

	if err := db.Instance.Table("Users").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func CreateUser(c *gin.Context) {

	var newUser models.Users

	//Parse json
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"newUser": newUser,
		})
		return
	}

	// Validate username and password
	if err := utils.ValidateUsername(newUser.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidatePassword(newUser.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Hash password
	newPassword, err := auth.HashPassword(newUser.Password)

	if err != nil {
		fmt.Println(err)
	}

	newUser.Password = newPassword

	// Check user is exist
	var existUser models.Users
	db.Instance.Table("users").Where("username = ?", newUser.Username).First(&existUser)

	if existUser.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username is exist!",
		})
		return
	}

	// Save user in database
	if err := db.Instance.Table("Users").Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created user",
		"user":    newUser,
	})
}
