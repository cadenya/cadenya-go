// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyasdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cadenya/cadenya-sdk-go"
	"github.com/cadenya/cadenya-sdk-go/internal/testutil"
	"github.com/cadenya/cadenya-sdk-go/option"
)

func TestAgentWebhookDeliveryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.WebhookDeliveries.List(
		context.TODO(),
		"agentId",
		gocadenyacomcadenyasdkgo.AgentWebhookDeliveryListParams{
			Cursor:      gocadenyacomcadenyasdkgo.F("cursor"),
			EventType:   gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeUnspecified),
			Limit:       gocadenyacomcadenyasdkgo.F(int64(0)),
			ObjectiveID: gocadenyacomcadenyasdkgo.F("objectiveId"),
		},
	)
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
