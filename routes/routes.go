package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine)  {
	server.GET("/ping", ping)
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events/add", addEvent)
	server.PUT("/events/:id", updateEventByID)
	server.DELETE("/events/:id", deleteEventByID)
}