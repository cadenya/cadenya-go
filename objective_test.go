// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

import (
	"context"
	"errors"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/internal/testutil"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/shared"
	"os"
	"testing"
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
		WorkspaceID: cadenya.String("workspaceId"),
		AgentID:     "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		SystemPromptData: map[string]any{
			"foo": "bar",
		},
		EpisodicMemory: cadenya.ObjectiveNewParamsEpisodicMemory{
			Key: "key",
		},
		FirstUserMessage: cadenya.String("firstUserMessage"),
		FirstUserMessageData: map[string]any{
			"foo": "bar",
		},
		MemoryCascade: []cadenya.MemoryReferenceParam{{
			MemoryLayerID: "memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH",
			MemoryEntryID: cadenya.String("mementry_01HXKD2E5NQM3T9AYWCF5E52Z0"),
		}},
		Metadata: shared.CreateOperationMetadataParam{
			ExternalID: cadenya.String("externalId"),
			Labels: map[string]string{
				"foo": "string",
			},
		},
		PinnedParameters: map[string]string{
			"foo": "string",
		},
		Secrets: []cadenya.ObjectiveNewParamsSecret{{
			Name:  cadenya.String("name"),
			Value: cadenya.String("value"),
		}},
		Subject: cadenya.SubjectAssertionParam{
			ID:   "customer-user-42",
			Name: cadenya.String("Jane Doe"),
		},
		Tenant: cadenya.TenantAssertionParam{
			ID:   "acme-corp",
			Name: cadenya.String("Acme Corp"),
		},
		VariationID: cadenya.String("agentvar_01HXKD2E5NQM3T9AYWCF32BSPP"),
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
	_, err := client.Objectives.Get(
		context.TODO(),
		"id",
		cadenya.ObjectiveGetParams{
			WorkspaceID: cadenya.String("workspaceId"),
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
		WorkspaceID:       cadenya.String("workspaceId"),
		AgentID:           cadenya.String("agentId"),
		AgentScheduleID:   cadenya.String("agentScheduleId"),
		Cursor:            cadenya.String("cursor"),
		IncludeInfo:       cadenya.Bool(true),
		Labels:            cadenya.String("labels"),
		Limit:             cadenya.Int(0),
		ParentObjectiveID: cadenya.String("parentObjectiveId"),
		ProfileID:         cadenya.String("profileId"),
		SortOrder:         cadenya.String("sortOrder"),
		State:             cadenya.ObjectiveListParamsStateStateUnspecified,
		SubjectID:         cadenya.String("subjectId"),
		TenantID:          cadenya.String("tenantId"),
		WidgetID:          cadenya.String("widgetId"),
		WidgetSessionID:   cadenya.String("widgetSessionId"),
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
			WorkspaceID: cadenya.String("workspaceId"),
			Reason:      cadenya.String("reason"),
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
			WorkspaceID: cadenya.String("workspaceId"),
			CompactionConfig: cadenya.AgentVariationSpecCompactionConfigParam{
				Summarization: cadenya.CompactionConfigSummarizationStrategyParam{
					Instructions: cadenya.String("instructions"),
				},
				ToolResultClearing: cadenya.CompactionConfigToolResultClearingStrategyParam{
					PreserveRecentResults: cadenya.Int(0),
				},
				TriggerThreshold: cadenya.Float(0),
			},
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
			WorkspaceID: cadenya.String("workspaceId"),
			Message:     "message",
			Enqueue:     cadenya.Bool(true),
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
			WorkspaceID: cadenya.String("workspaceId"),
			Cursor:      cadenya.String("cursor"),
			IncludeInfo: cadenya.Bool(true),
			Labels:      cadenya.String("labels"),
			Limit:       cadenya.Int(0),
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
			WorkspaceID:  cadenya.String("workspaceId"),
			Cursor:       cadenya.String("cursor"),
			IncludeInfo:  cadenya.Bool(true),
			Labels:       cadenya.String("labels"),
			Limit:        cadenya.Int(0),
			SinceEventID: cadenya.String("sinceEventId"),
			SortOrder:    cadenya.String("sortOrder"),
			WindowID:     cadenya.String("windowId"),
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

func TestObjectiveGetDiagnostics(t *testing.T) {
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
	_, err := client.Objectives.GetDiagnostics(
		context.TODO(),
		"objectiveId",
		cadenya.ObjectiveGetDiagnosticsParams{
			WorkspaceID: cadenya.String("workspaceId"),
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
