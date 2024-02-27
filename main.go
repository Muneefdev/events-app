package main

import (
	"os"

	"github.com/muneefdev/events-app/db"
	"github.com/muneefdev/events-app/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	db.Connect()
}

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
