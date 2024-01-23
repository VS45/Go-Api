package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vs45tech.com/event/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request Event Id", "Error": err.Error()})
		return
	}
	event, err := models.GetSingleEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event", "Error": err.Error()})
		return
	}
	err = event.RegisterEvent(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event", "Error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Registered !", "even": event})
}

func cancelRegistration(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id", "Error": err})
		return
	}
	event, err := models.GetSingleEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event", "Error": err})
		return
	}
	userId := context.GetInt64("userId")
	if userId != event.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorize for this task"})
		return
	}
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event", "Error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Cancelled"})
}
