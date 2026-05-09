package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-events/utils"
)

func Authenticate(context *gin.Context) {
	userId, err := utils.VerifyToke(context.Request.Header.Get("Authorization"))
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
	}
	context.Set("userId", userId)
}
