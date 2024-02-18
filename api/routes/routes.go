package routes

import (
	"authentication/api/handlers"
	"authentication/api/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/signup", handlers.CreateUserAccount)
	r.POST("/login", middlewares.LimitLoginAttempts(),handlers.AccessUserAccount)
	r.PUT("/update-profile", handlers.UpdateUserAccount)
	r.POST("/delete-account", handlers.DeleteUserAccount)
}