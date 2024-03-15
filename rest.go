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
	"example.com/learngo/db"
	"example.com/learngo/routes"
	"github.com/gin-gonic/gin"
)

func booking() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run("127.0.0.1:8080")
}
