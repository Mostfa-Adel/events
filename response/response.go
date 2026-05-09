package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(context *gin.Context, message string) {
	context.JSON(http.StatusBadRequest, gin.H{
		message: message,
	})
}

func UnAuthorized(context *gin.Context) {
	context.JSON(http.StatusUnauthorized, gin.H{
		"message": "Not Authorizd",
	})
}

func Success(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "succcesss",
	})
}
