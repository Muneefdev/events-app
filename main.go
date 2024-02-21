package main

import (
	"net/http"

	"example.com/rest-apis/events/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/events", GetAllEventsHandler)
	r.POST("/api/events", CreateEventHandler)

	r.Run(":8080")
}

func GetAllEventsHandler(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    events,
	})
}

func CreateEventHandler(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"data":    nil,
			"message": "invalid data",
		})
		return
	}

	event.ID = 1
	event.UserId = 1
	event.Save()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    event,
	})
}
