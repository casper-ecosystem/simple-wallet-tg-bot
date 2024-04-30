package mwauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(use bool, token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !use {
			c.Next()
			return
		}
		// Get token from the 'x-token' header
		headerToken := c.GetHeader("X-Token")

		// Check if token is present and valid
		if headerToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort() // Stop processing the request
			return
		}

		// Check if the provided token matches the one from the header
		if headerToken != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort() // Stop processing the request
			return
		}

		// If the token is valid, proceed with the request
		c.Next()
	}
}
