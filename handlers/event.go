package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muneefdev/events-app/models"
)

func GetAllEventsHandler(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch events",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Events fetched successfully",
		"success": true,
		"data":    events,
	})
}

func CreateEventsHandler(c *gin.Context) {

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"success": false,
			"data":    nil,
		})
		return
	}

	event.UserId = c.GetInt64("userID")
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event created successfully",
		"success": true,
		"data":    event,
	})
}

func GetEventByIdHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
			"success": false,
			"data":    nil,
		})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch event",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event fetched successfully",
		"success": true,
		"data":    event,
	})
}

func UpdateEventByIdHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
			"success": false,
			"data":    nil,
		})
		return
	}

  var event *models.Event
	event, err = models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	// Check if the user is the owner of the event
	userId := c.GetInt64("userID")
	if userId != event.UserId {
    c.JSON(http.StatusForbidden, gin.H{
      "message": "You are not authorized to update this event",
      "success": false,
      "data":    nil,
    })
    return
	}

	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"success": false,
			"data":    nil,
		})
		return
	}

	err = event.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not update event",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"success": true,
		"data":    event,
	})
}

func DeleteEventByIdHandler(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
			"success": false,
			"data":    nil,
		})
		return
	}

	var event *models.Event
	event, err = models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	// Check if the user is the owner of the event
	userId := c.GetInt64("userID")
	if userId != event.UserId {
    c.JSON(http.StatusForbidden, gin.H{
      "message": "You are not authorized to delete this event",
      "success": false,
      "data":    nil,
    })
    return
	}

	err = models.DeleteEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not delete event",
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
		"success": true,
		"data":    event,
	})
}
