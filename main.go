package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagubantii11/go-playground/models"
)

func main() {
	server := gin.Default()
	server.GET("/ping", ping)
	server.GET("/events", getAllEvents)
	server.POST("/events/add", addEvent)

	server.Run(":8081") // listen and serve on localhost:8081 (for windows "localhost:8080")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getAllEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"events": models.GetAllEvents(),
	})
}

func addEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	models.AddEvent(event)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event added successfully",
		"event":   event,
	})
}
