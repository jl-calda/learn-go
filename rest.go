package main

// event booking rest api
// GET /events - list all events
// GET /events/{id} - get event by id
// POST /events - create a new event
// PUT /events/{id} - update event by id
// DELETE /events/{id} - delete event by id
// POST /signup - sign up a new user
// POST /login - log in a user
// POST /events/{id}/register - register for an event
// DELETE /events/{id}/register - unregister for an event

import (
	"net/http"
	"time"

	"example.com/learngo/db"
	"example.com/learngo/models"
	"github.com/gin-gonic/gin"
)

func booking() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run("127.0.0.1:8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "data passed doesn't match"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.CreatedAt = time.Now()
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while saving event"})
	}

	context.JSON(http.StatusCreated, event)

}
