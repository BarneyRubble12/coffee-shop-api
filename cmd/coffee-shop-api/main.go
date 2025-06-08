package main

import (
	"coffee-shop-api/internal/di"
	"coffee-shop-api/internal/docs"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Coffee Shop API
// @version 1.0
// @description A RESTful API for managing coffee products
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize dependencies
	coffeeHandler, err := di.InitializeAPI()
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	// Create Gin router
	r := gin.Default()

	// Swagger documentation
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	// Coffee routes
	coffees := r.Group("/coffees")
	{
		coffees.GET("", coffeeHandler.GetAllCoffees)
		coffees.GET("/:id", coffeeHandler.GetCoffeeByID)
		coffees.POST("", coffeeHandler.CreateCoffee)
		coffees.PUT("/:id", coffeeHandler.UpdateCoffee)
		coffees.DELETE("/:id", coffeeHandler.DeleteCoffee)
		coffees.GET("/search", coffeeHandler.SearchCoffeesByOrigin)
	}

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
