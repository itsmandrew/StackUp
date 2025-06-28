package plaid

import (
	"github.com/gin-gonic/gin"
	client "github.com/q-summitLabs/StackUp/internal/plaid"
)

func GetEnv(c *gin.Context) {
	env := client.PLAID_ENV
	c.JSON(200, gin.H{"env": env})
}
