package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuldeephsc/api/models"
)

func registerForEvent(context *gin.Context) {
	// Implement logic to register a user for an event
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not able to parse event id"})
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	err = event.RegisterUserForEvent(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered for event"})
}

func unregisterForEvent(context *gin.Context) {
	// Implement logic to unregister a user from an event

	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not able to parse event id"})
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	err = event.UnregisterUserFromEvent(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister from event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from event"})
}
