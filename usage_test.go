// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyasdkgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/cadenya/cadenya-sdk-go"
	"github.com/cadenya/cadenya-sdk-go/internal/testutil"
	"github.com/cadenya/cadenya-sdk-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	account, err := client.Account.Get(context.TODO())
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", account.Metadata)
}
