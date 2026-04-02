// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

import (
	"context"
	"os"
	"testing"

	"github.com/cadenya/cadenya-sdk-go"
	"github.com/cadenya/cadenya-sdk-go/internal/testutil"
	"github.com/cadenya/cadenya-sdk-go/option"
)

func TestAutoPagination(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cadenya.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	iter := client.Agents.ListAutoPaging(context.TODO(), cadenya.AgentListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		agent := iter.Current()
		t.Logf("%+v\n", agent.Metadata)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
