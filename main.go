package main

import (
	"os"

	"example.com/rest-apis/events/db"
	"example.com/rest-apis/events/routes"
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
