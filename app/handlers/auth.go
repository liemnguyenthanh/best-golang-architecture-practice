package handlers

import (
	"api-instagram/app/middlewares"
	"api-instagram/app/services"
	"api-instagram/app/utils"
	"fmt"
	"net/http"

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
		utils.RespondBadRequest(c, err.Error())
		return
	}

	// Exist user
	existUser, _ := services.FindUserByCriteria(user.Username, services.ByUsername)

	fmt.Print(existUser)
	if existUser.Id == 0 {
		utils.RespondBadRequest(c, "username is not exist")
		return
	}

	// Check password
	isCorrectPassword := middlewares.CheckPasswordHash(user.Password, existUser.Password)

	if !isCorrectPassword {
		utils.RespondBadRequest(c, "password is not correct")
		return
	}

	// Create a new token with the user id
	token, err := middlewares.GenerateToken(existUser.Id)

	if err != nil {
		utils.RespondInternalServerError(c, "Failed to generate token")
	}

	var response = map[string]interface{}{
		"user":         existUser,
		"access_token": token,
	}
	// Response
	utils.RespondWithSuccess(c, http.StatusOK, response)
	return
}
