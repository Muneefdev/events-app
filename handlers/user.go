package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muneefdev/events-app/models"
)

func CreateUserHandler(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Invalid request body",
			"success": false,
			"data":    nil,
		})
		return
	}

	errors := user.Validate()
	if len(errors) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors,
			"success": false,
			"data":    nil,
		})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create user",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"success": true,
	})
}

func GetAllUsersHandler(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch users",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"success": true,
		"data":    users,
	})
}

func LoginHandler(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"success": false,
			"data":    nil,
		})
		return
	}

	var token string
	token, err = user.Login()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"success": true,
		"token":   token,
	})
}
