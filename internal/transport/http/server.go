package http

import (
	handlers "clothes-shop-backend/internal/handlers/factory"
	"clothes-shop-backend/internal/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func InitServer(handlers *handlers.Handlers) {
	router := gin.Default()

	setupRoutes(router, handlers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

func setupRoutes(router *gin.Engine, handlers *handlers.Handlers) {
	router.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "http://localhost:3000" || origin == "https://clothes-shop-seven.vercel.app" {
			c.Header("Access-Control-Allow-Origin", origin)
		}
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
	router.GET("/users/phone/:phone", handlers.V1Handlers.GetUserByPhone)
	router.POST("/users", handlers.V1Handlers.CreateUser)
	router.GET("/cart", utils.AuthMiddleware(), handlers.V1Handlers.GetCartItems)
	router.PUT("/cart", utils.AuthMiddleware(), handlers.V1Handlers.UpdateCart)
	router.POST("/cart/add/:product_id", utils.AuthMiddleware(), handlers.V1Handlers.AddOneToCart)
	router.POST("/checkout/:user_id", utils.AuthMiddleware(), handlers.V1Handlers.Checkout)
}
