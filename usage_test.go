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

func TestUsage(t *testing.T) {
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
	objective, err := client.Objectives.New(context.TODO(), cadenya.ObjectiveNewParams{
		WorkspaceID: cadenya.String("workspace_01HXKD2E5NQXAMPLE0000000"),
		AgentID:     "agent_01HXKD2E5NQXAMPLE0000000",
		SystemPromptData: map[string]any{
			"customer_name": "Ada",
		},
		FirstUserMessage: cadenya.String("Summarize the open support tickets from yesterday."),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", objective.ConfigSnapshot)
}
