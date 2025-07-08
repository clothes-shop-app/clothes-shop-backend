package http

import (
	handlers "clothes-shop-backend/cmd/internal/handlers/factory"

	"github.com/gin-gonic/gin"
)

func InitServer(handlers *handlers.Handlers) {
	router := gin.Default()

	setupRoutes(router, handlers)

	router.Run(":8080")
}

func setupRoutes(router *gin.Engine, handlers *handlers.Handlers) {
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.GET("/products", handlers.V1Handlers.GetPaginatedProducts)
	router.POST("/products", handlers.V1Handlers.UploadProduct)
	router.GET("/products/:id", handlers.V1Handlers.GetProductByID)
}
