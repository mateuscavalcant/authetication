package routes

import (
	"authentication/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/signup", handlers.CreateUserAccount)
	r.POST("/login", handlers.AccessUserAccount)
	r.PUT("update-profile", handlers.UpdateUserAccount)
}