package plaid

import (
	"net/http"

	"github.com/gin-gonic/gin"
	plaid "github.com/plaid/plaid-go/v31/plaid"
	client "github.com/q-summitLabs/StackUp/internal/plaid"
)

func CreateLinkToken(c *gin.Context) {
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: "user-id",
	}

	request := plaid.NewLinkTokenCreateRequest(
		"StackUp",
		"en",
		[]plaid.CountryCode{plaid.COUNTRYCODE_US},
		user,
	)

	request.SetProducts([]plaid.Products{plaid.PRODUCTS_TRANSACTIONS})
	linkTokenCreateResp, _, err := client.CLIENT.PlaidApi.LinkTokenCreate(c).LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		panic(err)
	}
	linkToken := linkTokenCreateResp.GetLinkToken()

	c.JSON(http.StatusOK, gin.H{
		"linkToken": linkToken,
	})
}
