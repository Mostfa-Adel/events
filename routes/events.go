package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-events/models"
	"github.com/go-events/response"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "failed get events from db",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func createEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"event":   event,
	})
}

func getEvent(context *gin.Context) {
	id, err := parseEventIdOrFail(context)
	if err != nil {
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		response.BadRequest(context, "record not exist")
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"event": event,
	})

}

func updateEvent(context *gin.Context) {
	var updatedEvent models.Event
	id, err := parseEventIdOrFail(context)
	if err != nil {
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		response.BadRequest(context, "record not exist")
		return
	}
	if context.GetInt64("userId") != event.UserID {
		response.UnAuthorized(context)
		return
	}
	updatedEvent.ID = id
	context.ShouldBindJSON(&updatedEvent)
	err = updatedEvent.UpdateEvent()
	if err != nil {
		response.BadRequest(context, "failed to update")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "updated successfully",
	})

}

func deleteEvent(context *gin.Context) {
	id, err := parseEventIdOrFail(context)
	if err != nil {
		return
	}
	event, err := models.GetEvent(id)
	if err != nil {
		response.BadRequest(context, "record not exist")
		return
	}
	if context.GetInt64("userId") != event.UserID {
		response.UnAuthorized(context)
		return
	}
	err = event.DeleteEvent()

	context.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})

}

func registerUserInEvent(context *gin.Context) {
	authUserId := context.GetInt64("userId")
	eventId, err := parseEventIdOrFail(context)
	if err != nil {
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "not found event"})
		return
	}
	err = event.RegisterUserInEvent(authUserId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "couldnt register in event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "registered successfully"})

}

func cancelUserInEvent(context *gin.Context) {
	authUserId := context.GetInt64("userId")
	eventId, err := parseEventIdOrFail(context)
	if err != nil {
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		return
	}
	err = event.DeleteUser(authUserId)
	if err != nil {
		return
	}
	response.Success(context)
}
