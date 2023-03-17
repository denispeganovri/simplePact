package client

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"net/http"
	"testing"
	"transPact/pkg/server"
	//"transPact/pkg/server"
)

func Test_ContractTest(t *testing.T) {

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "Consumer name",
		Provider: "Provider name",
	}

	defer pact.Teardown()

	pact.
		AddInteraction().
		Given("Test name").
		UponReceiving("A request to provider").
		WithRequest(dsl.Request{
			Method: "GET",
			Path:   dsl.Term("/endpoint", "/endpoint"),
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.Term("application/json; charset=utf-8", `application\/json`),
			},
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Body: dsl.Match(
				//here you can define a struct which you want to receive from you provider
				server.TransformedSku{},
			),
			Headers: dsl.MapMatcher{
				"Content-Type": dsl.Term("application/json; charset=utf-8", `application\/json`),
			},
		})

	var test = func() (err error) {
		url := fmt.Sprintf("http://localhost:%d/endpoint", pact.Server.Port)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")

		_, err = http.DefaultClient.Do(req)
		if err != nil {
			return
		}

		return
	}

	err := pact.Verify(test)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}

	// specify PACT publisher
	publisher := dsl.Publisher{}
	err = publisher.Publish(types.PublishRequest{
		PactURLs:        []string{"../client/pacts/consumer_name-provider_name.json"},
		PactBroker:      "https://pen.pactflow.io/", //link to your remote Contract broker
		BrokerToken:     "jEQnxw7xWgYRv-3-G7Cx-g",   //your PactFlow token
		ConsumerVersion: "2.0.2",
		Tags:            []string{"2.0.2", "latest"},
	})
	if err != nil {
		t.Fatal(err)
	}
}
