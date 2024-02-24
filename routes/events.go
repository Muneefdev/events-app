package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-apis/events/models"
	"github.com/gin-gonic/gin"
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

	event.UserId = 1
	event.Save()

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

	_, err = models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	var event models.Event
	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"success": false,
			"data":    nil,
		})
		return
	}

	event.ID = id
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
