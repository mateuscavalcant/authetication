package main

import (
	"authentication/api/routes"
	"authentication/pkg/database"
	"os"

	"github.com/gin-gonic/gin"
)


func main() {
	database.InitializeDB()
	r := gin.Default()

	// Configuração do CORS
	r.Use(corsMiddleware())

	// Inicialize suas rotas
	routes.InitRoutes(r.Group("/"))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r.Run(":" + port)
}

// Defina sua função de middleware CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

	