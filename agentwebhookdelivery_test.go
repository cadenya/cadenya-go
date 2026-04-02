// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyago_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/internal/testutil"
	"github.com/cadenya/cadenya-go/option"
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.WebhookDeliveries.List(
		context.TODO(),
		"agentId",
		gocadenyacomcadenyago.AgentWebhookDeliveryListParams{
			Cursor:      gocadenyacomcadenyago.F("cursor"),
			EventType:   gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeUnspecified),
			Limit:       gocadenyacomcadenyago.F(int64(0)),
			ObjectiveID: gocadenyacomcadenyago.F("objectiveId"),
		},
	)
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
