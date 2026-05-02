// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/internal/testutil"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/shared"
)

func TestObjectiveNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.New(context.TODO(), cadenya.ObjectiveNewParams{
		AgentID: cadenya.F("agentId"),
		Data: cadenya.F(cadenya.ObjectiveDataParam{
			Data:           cadenya.F[any](map[string]interface{}{}),
			InitialMessage: cadenya.F("initialMessage"),
			MemoryStack: cadenya.F([]cadenya.MemoryReferenceParam{{
				MemoryEntryID: cadenya.F("memoryEntryId"),
				MemoryLayerID: cadenya.F("memoryLayerId"),
			}}),
			Secrets: cadenya.F([]cadenya.ObjectiveDataSecretParam{{
				Name:  cadenya.F("name"),
				Value: cadenya.F("value"),
			}}),
		}),
		Metadata: cadenya.F(shared.CreateOperationMetadataParam{
			ExternalID: cadenya.F("externalId"),
			Labels: cadenya.F(map[string]string{
				"foo": "string",
			}),
		}),
		VariationID: cadenya.F("variationId"),
	})
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectiveGet(t *testing.T) {
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
	_, err := client.Objectives.Get(context.TODO(), "id")
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectiveListWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.List(context.TODO(), cadenya.ObjectiveListParams{
		AgentID:           cadenya.F("agentId"),
		AgentScheduleID:   cadenya.F("agentScheduleId"),
		Cursor:            cadenya.F("cursor"),
		IncludeInfo:       cadenya.F(true),
		Limit:             cadenya.F(int64(0)),
		ParentObjectiveID: cadenya.F("parentObjectiveId"),
		ProfileID:         cadenya.F("profileId"),
		SortOrder:         cadenya.F("sortOrder"),
		State:             cadenya.F(cadenya.ObjectiveListParamsStateStateUnspecified),
	})
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectiveCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.Cancel(
		context.TODO(),
		"objectiveId",
		cadenya.ObjectiveCancelParams{
			Reason: cadenya.F("reason"),
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

func TestObjectiveCompactWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.Compact(
		context.TODO(),
		"objectiveId",
		cadenya.ObjectiveCompactParams{
			CompactionConfig: cadenya.F(cadenya.AgentVariationSpecCompactionConfigParam{
				Summarization: cadenya.F(cadenya.CompactionConfigSummarizationStrategyParam{
					Instructions: cadenya.F("instructions"),
				}),
				ToolResultClearing: cadenya.F(cadenya.CompactionConfigToolResultClearingStrategyParam{
					PreserveRecentResults: cadenya.F(int64(0)),
				}),
				TriggerThreshold: cadenya.F(0.000000),
			}),
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

func TestObjectiveContinueWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.Continue(
		context.TODO(),
		"objectiveId",
		cadenya.ObjectiveContinueParams{
			Enqueue: cadenya.F(true),
			Message: cadenya.F("message"),
			Secrets: cadenya.F([]cadenya.ObjectiveContinueParamsSecret{{
				Name:  cadenya.F("name"),
				Value: cadenya.F("value"),
			}}),
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

func TestObjectiveListContextWindowsWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.ListContextWindows(
		context.TODO(),
		"objectiveId",
		cadenya.ObjectiveListContextWindowsParams{
			Cursor:      cadenya.F("cursor"),
			IncludeInfo: cadenya.F(true),
			Limit:       cadenya.F(int64(0)),
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

func TestObjectiveListEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.ListEvents(
		context.TODO(),
		"objectiveId",
		cadenya.ObjectiveListEventsParams{
			Cursor:      cadenya.F("cursor"),
			IncludeInfo: cadenya.F(true),
			Limit:       cadenya.F(int64(0)),
			SortOrder:   cadenya.F("sortOrder"),
			WindowID:    cadenya.F("windowId"),
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
