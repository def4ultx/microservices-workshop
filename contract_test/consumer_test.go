// +build consumer

package pact

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestConsumer(t *testing.T) {

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "Order-API",
		Provider: "Inventory-API",
		Host:     "localhost",
	}
	defer pact.Teardown()

	// Pass in test case. This is the component that makes the external HTTP call
	var test = func() (err error) {
		u := fmt.Sprintf("http://localhost:%d/product/668337293889241011", pact.Server.Port)
		req, err := http.NewRequest(http.MethodGet, u, nil)
		if err != nil {
			return
		}

		req.Header.Set("Content-Type", "application/json")
		_, err = http.DefaultClient.Do(req)
		return
	}

	product := struct {
		ID     int64  `json:"id"`
		Name   string `json:"name"`
		Price  int    `json:"price"`
		Amount int    `json:"amount"`
	}{}

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given("Product test1 exist").
		UponReceiving("A request to get product test").
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String("/product/668337293889241011"),
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    nil,
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Match(&product),
		})

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}
