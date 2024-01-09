package handlers

import (
	"api-instagram/app/middlewares"
	"api-instagram/app/models"
	"api-instagram/app/services"
	"api-instagram/app/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetUsers()

	if err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, users)

}

func CreateUser(c *gin.Context) {
	var newUser models.Users

	//Parse json
	if err := c.BindJSON(&newUser); err != nil {
		utils.RespondBadRequest(c, err.Error())
		return
	}

	// Validate username and password
	if err := utils.ValidateUsername(newUser.Username); err != nil {
		utils.RespondBadRequest(c, "Invalid username")
		return
	}

	if err := utils.ValidatePassword(newUser.Password); err != nil {
		utils.RespondBadRequest(c, "Invalid password")
		return
	}

	//Hash password
	if hashPassword, err := middlewares.HashPassword(newUser.Password); err != nil {
		fmt.Println(err)
	} else {
		newUser.Password = hashPassword
	}

	// Check user is exist
	existUser, _ := services.FindUserByCriteria(newUser.Username, services.ByUsername)

	if existUser != nil {
		utils.RespondNotFound(c, "username is exist")
		return
	}

	// Save user in database
	if _, err := services.InsertUser(&newUser); err != nil {
		utils.RespondInternalServerError(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, newUser)
}
