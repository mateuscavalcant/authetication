package middlewares

import (
	"authentication/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function to authenticate user sessions.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve session information from the request.
		session := utils.GetSession(c)
		userID := session.Values["id"]

		// Check if the user ID is not present in the session.
		if userID == nil {
			// If not authorized, render an HTML login page with an error message.
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Unauthorized"})
			c.Abort() // Abort further processing of the request.
			return
		}

		// If user is authenticated, proceed to the next handler.
		c.Next()
	}
}