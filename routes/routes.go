package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/muneefdev/events-app/handlers"
	"github.com/muneefdev/events-app/middlewares"
)

func RegisterRoutes(r *gin.Engine) {
	// Public routes
	r.GET("/api/events", handlers.GetAllEventsHandler)
	r.GET("/api/events/:id", handlers.GetEventByIdHandler)

	// Authenticated routes
	auth := r.Group("/api", middlewares.AuthMiddleware)
	auth.POST("/events", handlers.CreateEventsHandler)
	auth.PATCH("/events/:id", handlers.UpdateEventByIdHandler)
	auth.DELETE("/events/:id", handlers.DeleteEventByIdHandler)

  // Registering for an event
	auth.GET("/events/registered", handlers.GetAllRegisteredEventsHandler)
	auth.POST("/events/:id/register", handlers.RegisterEventHandler)
	auth.DELETE("/events/:id/registered", handlers.DeleteRegisteredEventHandler)

	// User routes
	r.POST("/api/register", handlers.CreateUserHandler)
	r.POST("/api/login", handlers.LoginHandler)
	r.GET("/api/users", handlers.GetAllUsersHandler)

}
