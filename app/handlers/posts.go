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

func GetPosts(c *gin.Context) {
	var filter models.PostFilter

	filter.Page = utils.GetIntQueryParam(c, "page", 1)
	filter.Limit = utils.GetIntQueryParam(c, "limit", 10)
	filter.User_id = utils.GetIntQueryParam(c, "user_id", 0)

	posts, err := services.GetPosts(&filter)

	if err != nil {
		utils.RespondBadRequest(c, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var post models.Posts

	if err := c.BindJSON(&post); err != nil {
		utils.RespondBadRequest(c, err.Error())
		return
	}

	// Extract the user ID from the request headers
	user_id, err := middlewares.ExtractUserIdFromToken(c)
	fmt.Println("user_id", user_id)
	if err != nil {
		utils.RespondUnauthorized(c, err.Error())
		return
	}

	if user_id == "" {
		utils.RespondBadRequest(c, "token is expired")
		return
	}

	// Check if the user with the given ID exists
	user, err := services.FindUserByCriteria(user_id, services.ByID)
	if err != nil {
		utils.RespondBadRequest(c, "User not found")
		return
	}

	// Save post in database
	post.User_id = user.Id

	if _, err := services.InsetPost(&post); err != nil {
		utils.RespondInternalServerError(c, err.Error())
	}

	utils.RespondWithSuccess(c, http.StatusCreated, post)
}
