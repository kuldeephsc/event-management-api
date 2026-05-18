package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kuldeephsc/api/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data, Try again later"})
		return
	}
	// context.JSON(http.StatusOK, gin.H{"message": "Hello From Server"})
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// token := context.Request.Header.Get("Authorization")

	// if token == "" {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }

	// userId, err := utils.VerifyToken(token)
	// if err != nil {
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }

	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could Not parse request data."})
		return
	}
	// event.ID = 1
	userId, _ := context.Get("userId")
	event.UserID = (userId).(int64)
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event, Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "Event": event})
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id. Try again later ."})
		return
	}

	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, event)

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not able to fetch event id"})
		return
	}

	UserId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if event.UserID != UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to update this event"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not able to find event"})
		return
	}
	var updateEvent models.Event

	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse Request data"})
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not updated the event"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated Successfully"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	UserId := context.GetInt64("userId")
	// event, err := models.GetEventById(eventId)

	if event.UserID != UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized to delete this event"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event not Present in the database"})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted"})
}
