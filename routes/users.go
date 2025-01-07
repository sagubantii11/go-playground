package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagubantii11/go-playground/models"
	"github.com/sagubantii11/go-playground/utils"
)

func addUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	user.RegisterUser()
	c.JSON(http.StatusCreated, gin.H{
		"message": "User added successfully",
		"user":    user,
	})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}
	if user.VerifyUserLogin() {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	} else {
		validJWT, err := utils.GenerateJWT(user.EmailID, user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"token":   validJWT,
		})
	}
}
