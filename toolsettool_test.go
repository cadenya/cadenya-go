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

func TestToolSetToolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.Tools.New(
		context.TODO(),
		"toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		cadenya.ToolSetToolNewParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Metadata: shared.CreateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.ToolSpecParam{
				Config: cadenya.ToolSpecConfigUnionParam{
					OfHTTP: &cadenya.ToolSpecConfigHTTPParam{
						HTTP: cadenya.ConfigHTTPParam{
							RequestMethod: cadenya.ConfigHTTPRequestMethodHTTPMethodUnspecified,
							Headers: map[string]string{
								"foo": "string",
							},
							Path:                   cadenya.String("path"),
							Query:                  cadenya.String("query"),
							RequestBodyContentType: cadenya.String("requestBodyContentType"),
							RequestBodyTemplate:    cadenya.String("requestBodyTemplate"),
						},
						Type: cadenya.ToolSpecConfigHTTPTypeHTTP,
					},
				},
				Description: "description",
				Parameters: map[string]any{
					"foo": "bar",
				},
				RequiresApproval: true,
				LlmToolName:      cadenya.String("llmToolName"),
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

func TestToolSetToolGet(t *testing.T) {
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
	_, err := client.ToolSets.Tools.Get(
		context.TODO(),
		"toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		"tool_01HXKD2E5NQM3T9AYWCFWVYY9K",
		cadenya.ToolSetToolGetParams{
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

func TestToolSetToolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.Tools.Update(
		context.TODO(),
		"toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		"tool_01HXKD2E5NQM3T9AYWCFWVYY9K",
		cadenya.ToolSetToolUpdateParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Metadata: shared.UpdateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.ToolSpecParam{
				Config: cadenya.ToolSpecConfigUnionParam{
					OfHTTP: &cadenya.ToolSpecConfigHTTPParam{
						HTTP: cadenya.ConfigHTTPParam{
							RequestMethod: cadenya.ConfigHTTPRequestMethodHTTPMethodUnspecified,
							Headers: map[string]string{
								"foo": "string",
							},
							Path:                   cadenya.String("path"),
							Query:                  cadenya.String("query"),
							RequestBodyContentType: cadenya.String("requestBodyContentType"),
							RequestBodyTemplate:    cadenya.String("requestBodyTemplate"),
						},
						Type: cadenya.ToolSpecConfigHTTPTypeHTTP,
					},
				},
				Description: "description",
				Parameters: map[string]any{
					"foo": "bar",
				},
				RequiresApproval: true,
				LlmToolName:      cadenya.String("llmToolName"),
			},
			UpdateMask: cadenya.String("updateMask"),
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

func TestToolSetToolListWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.Tools.List(
		context.TODO(),
		"toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		cadenya.ToolSetToolListParams{
			WorkspaceID:      cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Cursor:           cadenya.String("cursor"),
			IncludeInfo:      cadenya.Bool(true),
			Labels:           cadenya.String("labels"),
			Limit:            cadenya.Int(0),
			Names:            []string{"string"},
			Prefix:           cadenya.String("prefix"),
			Query:            cadenya.String("query"),
			RequiresApproval: cadenya.Bool(true),
			SortOrder:        cadenya.String("sortOrder"),
			States:           []string{"STATE_UNSPECIFIED"},
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

func TestToolSetToolDelete(t *testing.T) {
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
	err := client.ToolSets.Tools.Delete(
		context.TODO(),
		"toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		"tool_01HXKD2E5NQM3T9AYWCFWVYY9K",
		cadenya.ToolSetToolDeleteParams{
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

func TestToolSetToolOmit(t *testing.T) {
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
	_, err := client.ToolSets.Tools.Omit(
		context.TODO(),
		"toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		"tool_01HXKD2E5NQM3T9AYWCFWVYY9K",
		cadenya.ToolSetToolOmitParams{
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

func TestToolSetToolRestore(t *testing.T) {
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
	_, err := client.ToolSets.Tools.Restore(
		context.TODO(),
		"toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		"tool_01HXKD2E5NQM3T9AYWCFWVYY9K",
		cadenya.ToolSetToolRestoreParams{
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
