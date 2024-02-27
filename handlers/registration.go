package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muneefdev/events-app/models"
)

func RegisterEventHandler(c *gin.Context) {
	userID := c.GetInt64("userID")
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event ID.",
			"success": false,
		})
		return
	}

	isExist := models.CheckRegistration(userID, eventID)
	if isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You are already registered for this event.",
			"success": false,
		})
		return
	}

	registration := models.NewRegistration(userID, eventID)
	err = registration.RegisterEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error registering for event.",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registered for event successfully.",
		"success": true,
	})
}

func GetAllRegisteredEventsHandler(c *gin.Context) {
	registrations, err := models.GetAllRegistrations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting all registered events.",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered events fetched successfully.",
		"success": true,
		"data":    registrations,
	})
}

func DeleteRegisteredEventHandler(c *gin.Context) {
	registrationID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid registration ID.",
			"success": false,
		})
		return
	}

	registeredEvent, err := models.GetRegistrationByID(registrationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not find registration.",
			"success": false,
		})
		return
	}

	if registeredEvent.UserID != int64(c.GetInt64("userID")) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "You are not authorized to delete this registration.",
			"success": false,
		})
		return
	}

	err = models.DeleteEventRegistration(registrationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting registration.",
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration deleted successfully.",
		"success": true,
	})
}
