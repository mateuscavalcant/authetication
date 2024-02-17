package routes

import (
	"authetication/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.POST("/signup", handlers.CreateUserAccount)

}