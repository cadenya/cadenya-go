// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

import (
	"context"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/internal/testutil"
	"go.cadenya.com/cadenya-go/option"
	"os"
	"testing"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.AIProviderKeys.List(context.TODO(), cadenya.AIProviderKeyListParams{
		WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, aiProviderKey := range page.Items {
		t.Logf("%+v\n", aiProviderKey.Metadata)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, aiProviderKey := range page.Items {
			t.Logf("%+v\n", aiProviderKey.Metadata)
		}
	}
}
