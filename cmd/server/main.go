package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsmandrew/StackUp/cmd/server/handlers"
	client "github.com/itsmandrew/StackUp/internal/plaid"
)

func main() {
	router := gin.Default()

	// Health check route
	router.GET("/health", handlers.HealthCheck)
	router.GET("/hello", handlers.Hello)

	// Plaid routes
	router.GET("/plaid/env", func(c *gin.Context) {
		env := client.PLAID_ENV
		c.JSON(200, gin.H{"env": env})
	})

	router.Run(":8080")
}
