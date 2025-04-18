package middlewares

import (
	"net/http"
	"product-crud/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoggingMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		logger.Info("HTTP Request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", duration,
			"user_agent", c.Request.UserAgent(),
			"remote_addr", c.ClientIP(),
			"request_id", c.Writer.Header().Get("X-Request-ID"))
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
		}

		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Set("X-Request-ID", requestID)

		c.Next()
	}
}

func RecoveryMiddleware(logger *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Recover from panic",
					"error", err,
					"path", c.Request.URL.Path,
					"method", c.Request.Method,
				)

				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
		}()

		c.Next()
	}
}

func generateRequestID() string {
	return uuid.New().String()
}
