package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/api/events", GetAllEventsHandler)
	r.GET("/api/events/:id", GetEventByIdHandler)
	r.POST("/api/events", CreateEventsHandler)
  r.PATCH("/api/events/:id", UpdateEventByIdHandler)
  r.DELETE("/api/events/:id", DeleteEventByIdHandler)
}
