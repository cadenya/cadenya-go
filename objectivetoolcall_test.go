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
)

func TestObjectiveToolCallGet(t *testing.T) {
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
	_, err := client.Objectives.ToolCalls.Get(
		context.TODO(),
		"workspaceId",
		"objectiveId",
		"toolCallId",
	)
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectiveToolCallListWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.ToolCalls.List(
		context.TODO(),
		"workspaceId",
		"objectiveId",
		cadenya.ObjectiveToolCallListParams{
			Cursor:          cadenya.F("cursor"),
			ExecutionStatus: cadenya.F(cadenya.ObjectiveToolCallListParamsExecutionStatusToolCallExecutionStatusUnspecified),
			IncludeInfo:     cadenya.F(true),
			Limit:           cadenya.F(int64(0)),
			Status:          cadenya.F(cadenya.ObjectiveToolCallListParamsStatusToolCallStatusUnspecified),
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

func TestObjectiveToolCallApprove(t *testing.T) {
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
	_, err := client.Objectives.ToolCalls.Approve(
		context.TODO(),
		"workspaceId",
		"objectiveId",
		"toolCallId",
		cadenya.ObjectiveToolCallApproveParams{},
	)
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectiveToolCallDenyWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.ToolCalls.Deny(
		context.TODO(),
		"workspaceId",
		"objectiveId",
		"toolCallId",
		cadenya.ObjectiveToolCallDenyParams{
			Memo: cadenya.F("memo"),
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

func TestObjectiveToolCallSetContent(t *testing.T) {
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
	_, err := client.Objectives.ToolCalls.SetContent(
		context.TODO(),
		"workspaceId",
		"objectiveId",
		"toolCallId",
		cadenya.ObjectiveToolCallSetContentParams{
			Content: cadenya.F([]cadenya.SetToolCallContentRequestContentBlockParam{{
				Audio: cadenya.F(cadenya.SetToolCallContentRequestAudioBlockParam{
					Data:     cadenya.F("data"),
					MimeType: cadenya.F("mimeType"),
				}),
				Image: cadenya.F(cadenya.SetToolCallContentRequestImageBlockParam{
					Data:     cadenya.F("data"),
					MimeType: cadenya.F("mimeType"),
				}),
				Text: cadenya.F(cadenya.SetToolCallContentRequestTextBlockParam{
					Text: cadenya.F("text"),
				}),
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
