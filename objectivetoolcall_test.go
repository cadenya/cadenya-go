// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

import (
	"context"
	"errors"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/internal/testutil"
	"go.cadenya.com/cadenya-go/option"
	"os"
	"testing"
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
		"obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		"toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
		cadenya.ObjectiveToolCallGetParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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
		"obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		cadenya.ObjectiveToolCallListParams{
			WorkspaceID:     cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Cursor:          cadenya.String("cursor"),
			ExecutionStatus: cadenya.ObjectiveToolCallListParamsExecutionStatusToolCallExecutionStatusUnspecified,
			IncludeInfo:     cadenya.Bool(true),
			Labels:          cadenya.String("labels"),
			Limit:           cadenya.Int(0),
			Status:          cadenya.ObjectiveToolCallListParamsStatusToolCallStatusUnspecified,
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
		"obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		"toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
		cadenya.ObjectiveToolCallApproveParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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
		"obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		"toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
		cadenya.ObjectiveToolCallDenyParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Memo:        cadenya.String("memo"),
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
		"obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		"toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
		cadenya.ObjectiveToolCallSetContentParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Content: []cadenya.SetToolCallContentRequestContentBlockUnionParam{{
				OfText: &cadenya.SetToolCallContentRequestContentBlockTextParam{
					Text: cadenya.SetToolCallContentRequestTextBlockParam{
						Text: "text",
					},
					Type: cadenya.SetToolCallContentRequestContentBlockTextTypeText,
				},
			}},
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
