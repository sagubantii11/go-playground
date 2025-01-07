package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sagubantii11/go-playground/middlewares"
)

func RegisterRoutes(server *gin.Engine)  {
	server.GET("/ping", ping)
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventByID)
	// Grouping the routes that require authentication
	jwtAuthenticated := server.Group("/")
	jwtAuthenticated.Use(middlewares.AuthenticateJWT)
	jwtAuthenticated.POST("/events/add", addEvent)
	jwtAuthenticated.PUT("/events/:id", updateEventByID)
	jwtAuthenticated.DELETE("/events/:id", deleteEventByID)

	server.POST("/signup", addUser)
	server.POST("/login", login)
}