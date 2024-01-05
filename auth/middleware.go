package auth

import (
	"api-instagram/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	var secretKey = utils.Getenv("SECRET_JWT_KEY")

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// If the token is valid, continue processing
		c.Next()
	}
}

func ExtractUserIdFromToken(c *gin.Context) (string, error) {
	var secretKey = utils.Getenv("SECRET_JWT_KEY")

	tokenString := c.GetHeader("Authorization") // Assuming the token is sent in the Authorization header

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Replace "your-secret-key" with your actual secret key
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}

	user_id, ok := claims["user_id"].(string)
	if !ok {
		return "", err
	}

	return user_id, nil
}
