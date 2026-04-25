// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/internal/testutil"
	"github.com/cadenya/cadenya-go/option"
)

func TestAgentFeedbackListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Feedback.List(
		context.TODO(),
		"agentId",
		cadenya.AgentFeedbackListParams{
			AgentVariationID: cadenya.F("agentVariationId"),
			CreatedAfter:     cadenya.F(time.Now()),
			CreatedBefore:    cadenya.F(time.Now()),
			Cursor:           cadenya.F("cursor"),
			IncludeInfo:      cadenya.F(true),
			Limit:            cadenya.F(int64(0)),
			Query:            cadenya.F("query"),
			Sentiment:        cadenya.F(cadenya.AgentFeedbackListParamsSentimentFeedbackSentimentUnspecified),
		},
	)
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
