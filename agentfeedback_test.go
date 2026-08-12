package cadenya_test

import (
	"context"
	"errors"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/internal/testutil"
	"go.cadenya.com/cadenya-go/option"
	"os"
	"testing"
	"time"
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentFeedbackListParams{
			WorkspaceID:      cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			AgentVariationID: cadenya.String("agentVariationId"),
			CreatedAfter:     cadenya.Time(time.Now()),
			CreatedBefore:    cadenya.Time(time.Now()),
			Cursor:           cadenya.String("cursor"),
			IncludeInfo:      cadenya.Bool(true),
			Labels:           cadenya.String("labels"),
			Limit:            cadenya.Int(0),
			Query:            cadenya.String("query"),
			Sentiment:        cadenya.AgentFeedbackListParamsSentimentFeedbackSentimentUnspecified,
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
