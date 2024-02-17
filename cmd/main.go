package main

import (
	"authentication/api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.InitRoutes(router.Group("/auth"))

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router.Run(":" + port)
}