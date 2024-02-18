package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func LimitLoginAttempts() gin.HandlerFunc {
	return func(c *gin.Context) {

		loginAttempts := cache.New(5*time.Minute, 10*time.Minute)
		username := c.PostForm("username") 

		if attempts, found := loginAttempts.Get(username); found && attempts.(int) >= 3 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many login attempts. Please try again later."})
			c.Abort()
			return
		}

		c.Next()
	}
}