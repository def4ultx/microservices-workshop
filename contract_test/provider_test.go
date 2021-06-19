// +build provider

package pact

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Provider: "Inventory-API",
	}

	// Verify the Provider using the locally saved Pact Files
	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:8080",
		PactURLs:        []string{filepath.ToSlash(fmt.Sprintf("./pacts/order-api-inventory-api.json"))},
		StateHandlers: types.StateHandlers{

			"Product test1 exist": func() error {
				return nil
			},
		},
	})
}
