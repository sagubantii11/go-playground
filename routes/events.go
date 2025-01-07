package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sagubantii11/go-playground/models"
)

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

func getEventByID(c *gin.Context) {
	parsedID := c.Param("id")
	id, err := strconv.ParseInt(parsedID, 10, 64)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"event": models.GetEventByID(id),
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

func updateEventByID(c *gin.Context) {
	parsedID := c.Param("id")
	id, err := strconv.ParseInt(parsedID, 10, 64)
	if err != nil {
		panic(err)
	}
	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	event = models.UpdateEventByID(id, event)
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   event,
	})
}

func deleteEventByID(c *gin.Context) {
	parsedID := c.Param("id")
	id, err := strconv.ParseInt(parsedID, 10, 64)
	if err != nil {
		panic(err)
	}
	models.DeleteEventByID(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}