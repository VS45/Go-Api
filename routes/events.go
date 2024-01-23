package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vs45tech.com/event/models"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch Events", "Error": err})
		return
	}
	context.JSON(http.StatusOK, events)
}
func GetEvent(context *gin.Context) {
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
	context.JSON(http.StatusOK, event)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		//fmt.Println(event)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request", "Error": err})
		return
	}
	userId := context.GetInt64("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "Error": err})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created !", "even": event})
}

func updateEvent(context *gin.Context) {

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
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorize to update Event"})
		return
	}
	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request", "Error": err})
		return
	}
	err = updateEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event", "Error": err})
		return
	}
	context.JSON(http.StatusOK, updateEvent)
}

func deletEvent(context *gin.Context) {
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
	err = event.DeleteEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event", "Error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted"})
}
