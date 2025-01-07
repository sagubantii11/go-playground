package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagubantii11/go-playground/utils"
)

func AuthenticateJWT(c *gin.Context) {
	authToken := c.Request.Header.Get("Authorization")
	if authToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized access",
		})
		return
	}
	_, err := utils.VerifyJWT(authToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized access",
		})
		return
	}

	c.Next()
}