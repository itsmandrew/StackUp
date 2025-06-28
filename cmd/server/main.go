package main

import (
	"github.com/gin-gonic/gin"
	"github.com/q-summitLabs/StackUp/cmd/server/handlers"
	plaidHandlers "github.com/q-summitLabs/StackUp/cmd/server/handlers/plaid"
)

func main() {
	router := gin.Default()

	// Health check route
	router.GET("/health", handlers.HealthCheck)
	router.GET("/hello", handlers.Hello)

	// Plaid handler routes
	router.GET("/plaid/env", plaidHandlers.GetEnv)
	router.GET("/plaid/createLinkToken", plaidHandlers.CreateLinkToken)

	router.Run(":8080")
}
