package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muneefdev/events-app/utails"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "No token provided",
			"success": false,
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Verify token
	claims, err := utails.VerifyJWT(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"success": false,
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if userID, ok := claims["userID"].(float64); ok {
		c.Set("userID", int64(userID))
	}
	c.Set("email", claims["email"])

	c.Next()
}
