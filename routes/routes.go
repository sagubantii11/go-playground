package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sagubantii11/go-playground/middlewares"
)

func RegisterRoutes(server *gin.Engine)  {
	server.GET("/ping", ping)
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events/add", middlewares.AuthenticateJWT, addEvent)
	server.PUT("/events/:id", middlewares.AuthenticateJWT, updateEventByID)
	server.DELETE("/events/:id", middlewares.AuthenticateJWT, deleteEventByID)

	server.POST("/signup", addUser)
	server.POST("/login", login)
}