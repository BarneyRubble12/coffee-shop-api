package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"coffee-shop-api/internal/di"
	"coffee-shop-api/internal/repository"

	_ "coffee-shop-api/internal/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Coffee Shop API
// @version 1.0
// @description A RESTful API for managing a coffee shop's inventory
// @host localhost:8080
// @BasePath /
func main() {
	// Initialize dependencies
	coffeeHandler, err := di.InitializeAPI()
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	// Get the SQLite repository to handle cleanup
	sqliteRepo, ok := coffeeHandler.GetRepository().(*repository.SQLiteCoffeeRepository)
	if !ok {
		log.Fatal("Failed to get SQLite repository")
	}
	defer sqliteRepo.Close()

	// Set up Gin router
	router := gin.Default()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Coffee routes
	router.GET("/coffees", coffeeHandler.GetAllCoffees)
	router.GET("/coffees/:id", coffeeHandler.GetCoffeeByID)
	router.POST("/coffees", coffeeHandler.CreateCoffee)
	router.PUT("/coffees/:id", coffeeHandler.UpdateCoffee)
	router.DELETE("/coffees/:id", coffeeHandler.DeleteCoffee)

	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")
}
