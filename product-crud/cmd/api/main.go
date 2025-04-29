package main

import (
	"log"
	"os"
	"strconv"

	"product-crud/api/routes"
	_ "product-crud/docs"
	"product-crud/internal/delivery/rest"
	"product-crud/internal/repository"
	"product-crud/internal/service"
	"product-crud/pkg/cache"
	"product-crud/pkg/db"

	"github.com/joho/godotenv"
)

// @title Go Gin CRUD API
// @version 1.0
// @description A simple CRUD API built with Go Gin, PostgreSQL, and Swagger
// @host localhost:8080
// @BasePath /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.InitSchema(database); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Initialize Redis cache
	redisHost := getEnv("REDIS_HOST", "localhost")
	redisPort := getEnv("REDIS_PORT", "6379")
	redisPassword := getEnv("REDIS_PASSWORD", "")
	cacheTTL, _ := strconv.Atoi(getEnv("CACHE_TTL", "3600"))

	redisCache, err := cache.NewRedisCache(redisHost+":"+redisPort, redisPassword, cacheTTL)
	if err != nil {
		log.Printf("Warning: Failed to connect to Redis: %v", err)
		log.Println("Continuing without cache...")
	}

	productRepo := repository.NewProductRepository(database)
	productService := service.NewProductService(productRepo, redisCache)
	productHandler := rest.NewProductHandler(productService)

	router := routes.SetupRouter(productHandler)

	port := getEnv("PORT", "8080")

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Helper function to get environment variables with defaults
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
