package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sagubantii11/go-playground/routes"
	"github.com/sagubantii11/go-playground/sqldb"
)

func main() {
	sqldb.InitDB()
	server := gin.Default()
	
	routes.RegisterRoutes(server)

	server.Run(":8081") // listen and serve on localhost:8081 (for windows "localhost:8080")
}
