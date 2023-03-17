package server

import (
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"testing"
)

func TestServerPact_Verification(t *testing.T) {

	pact := dsl.Pact{
		Consumer: "Consumer name",
		Provider: "Provider name",
	}

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://127.0.0.1:8080",   //provider's URL
		BrokerURL:                  "https://pen.pactflow.io", //link to your remote Contract broker
		BrokerToken:                "jEQnxw7xWgYRv-3-G7Cx-g",  //your PactFlow token
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
	})

	if err != nil {
		t.Fatal(err)
	}

	// Uncomment to verify contract locally
	/*log.Println("[debug] start verification")
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		PactURLs:        []string{"../client/pactsX/den_client-den_provider.json"},
	})*/
}
