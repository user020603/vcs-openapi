package routes

import (
	"net/http"
	"product-crud/api/middlewares"
	"product-crud/internal/delivery/rest"
	"product-crud/pkg/logger"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(productHandler *rest.ProductHandler) *gin.Engine {
	router := gin.Default()

	logger := logger.NewLogger("info")

	router.Use(middlewares.CORSMiddleware())
	router.Use(middlewares.LoggingMiddleware(logger))
	router.Use(middlewares.RequestIDMiddleware())
	router.Use(middlewares.RecoveryMiddleware(logger))

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	v1 := router.Group("/api/v1")
	{
		setupProductRoutes(v1, productHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}

func setupProductRoutes(rg *gin.RouterGroup, handler *rest.ProductHandler) {
	products := rg.Group("/products")
	{
		products.POST("", handler.CreateProduct)
		products.GET("", handler.GetProducts)
		products.GET("/:id", handler.GetProduct)
		products.PUT("/:id", handler.UpdateProduct)
		products.DELETE("/:id", handler.DeleteProduct)
	}
}
