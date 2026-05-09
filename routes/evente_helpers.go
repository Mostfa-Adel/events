package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-events/response"
)

func parseEventIdOrFail(context *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(context, "failed to parse id passed")
		return 0, err
	}
	return id, nil
}
