package toolexchange_test

import (
	"fmt"
	"testing"
	"toolexchange"
)

func TestExchange(t *testing.T) {
	// Tool1 wants to send data to Tool2.
	// Tool1 requests a token by adding data.
	ex := toolexchange.NewExchanger()
	req := toolexchange.Item{
		Data: map[string]string{
			"query": "(a AND b)",
			"seeds": "1234,54321,2134,4312",
		},
		Referrer: "Tool1",
	}
	token := ex.PutItem(req)

	// Now, a request is sent to Tool2 with the token.
	// Tool2 uses the token to request the saved data.
	resp, _ := ex.GetItem(token)
	// Tool2 can double check the token is from Tool1 using the referrer.
	fmt.Println(resp.Referrer)
	// Now Tool2 can use the data received in whatever way it pleases.
	fmt.Println(resp.Data)
	// Tool2 can continue to use the data up until it expires.
	fmt.Println(resp.Expiration)
}
