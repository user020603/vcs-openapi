package main

import (
	"log"
	"os"

	"product-crud/api/routes"
	_ "product-crud/docs" // This is required for swagger
	"product-crud/internal/delivery/rest"
	"product-crud/internal/repository"
	"product-crud/internal/service"
	"product-crud/pkg/db"

	"github.com/joho/godotenv"
)

// @title Go Gin CRUD API
// @version 1.0
// @description A simple CRUD API built with Go Gin, PostgreSQL, and Swagger
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database connection
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Initialize database schema
	if err := db.InitSchema(database); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Initialize repository, service, and handler
	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo)
	productHandler := rest.NewProductHandler(productService)

	// Setup router with all routes and middleware
	router := routes.SetupRouter(productHandler)

	// Get port from environment variables or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}