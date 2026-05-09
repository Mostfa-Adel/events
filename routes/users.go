package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-events/models"
	"github.com/go-events/response"
	"github.com/go-events/utils"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		response.BadRequest(context, "passed data not correct")
		return
	}
	err = user.SaveUser()
	if err != nil {
		response.BadRequest(context, "failed to save user")
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
		"user":    user,
	})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		response.BadRequest(context, "passed data not correct")
		return
	}

	userId, err := models.ValidateUserCredintials(user.Email, user.Password)
	if err != nil {
		response.BadRequest(context, err.Error())
		return
	}
	token, err := utils.GenerateJwtToken(user.Email, userId)
	if err != nil {
		response.BadRequest(context, "filed to authenticate")
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "loggedin successfully",
		"user":    user,
		"token":   token,
	})

}
