package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmandrew/StackUp/cmd/server/handlers"
)

func main() {
	router := gin.Default()

	// Health check route
	router.GET("/health", handlers.HealthCheck)
	router.GET("/hello", handlers.Hello)

	router.Run(":8080")
}
