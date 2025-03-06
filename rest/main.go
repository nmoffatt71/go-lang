package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example/rest/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost

}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event1"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	fmt.Println("Saving...")

	err := context.ShouldBindJSON(&event)
	fmt.Println("Saving1...")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad data"})
		return
	}
	fmt.Println("Saving2...")
	event.ID = 1
	event.UserID = 1
	fmt.Println("Saving3...")

	err = event.Save()
	fmt.Println("Saving4...")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "created!", "event": event})
}
