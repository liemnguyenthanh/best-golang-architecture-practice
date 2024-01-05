package controllers

import (
	"api-instagram/auth"
	db "api-instagram/database"
	"api-instagram/models"
	"api-instagram/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Username string `json:username`
	Password string `json:password`
}

func Login(c *gin.Context) {

	var user UserLogin

	//Parse json
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Exist user
	var existUser models.Users
	db.Instance.Table("users").Where("username = ?", user.Username).First(&existUser)

	if existUser.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username is not exist!",
		})
		return
	}

	// Check password
	isCorrectPassword := auth.CheckPasswordHash(user.Password, existUser.Password)

	if !isCorrectPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password is not correct!",
		})
		return
	}

	// Create a new token with the username and set it to expire in 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": existUser.Id,
	})

	// Sign the token with your secret key
	secretKey := utils.Getenv("SECRET_JWT_KEY")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"user":         existUser,
		"access_token": tokenString,
	})
	return
}
