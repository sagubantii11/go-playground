package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func main() {
  server := gin.Default()
  server.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  server.Run(":8081") // listen and serve on localhost:8081 (for windows "localhost:8080")
}