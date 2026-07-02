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
		"workspaceId",
		"toolSetId",
		cadenya.ToolSetToolNewParams{
			Metadata: cadenya.F(shared.CreateResourceMetadataParam{
				Name:       cadenya.F("name"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.ToolSpecParam{
				Config: cadenya.F(cadenya.ToolSpecConfigParam{
					HTTP: cadenya.F(cadenya.ConfigHTTPParam{
						RequestMethod: cadenya.F(cadenya.ConfigHTTPRequestMethodHTTPMethodUnspecified),
						Headers: cadenya.F(map[string]string{
							"foo": "string",
						}),
						Path:                   cadenya.F("path"),
						Query:                  cadenya.F("query"),
						RequestBodyContentType: cadenya.F("requestBodyContentType"),
						RequestBodyTemplate:    cadenya.F("requestBodyTemplate"),
					}),
					Mcp: cadenya.F(cadenya.ConfigMcpParam{
						Annotations: cadenya.F(cadenya.McpAnnotationsParam{
							DestructiveHint: cadenya.F(true),
							IdempotentHint:  cadenya.F(true),
							OpenWorldHint:   cadenya.F(true),
							ReadOnlyHint:    cadenya.F(true),
							Title:           cadenya.F("title"),
						}),
					}),
					OpenAPI: cadenya.F(cadenya.ConfigOpenAPIParam{
						Method: cadenya.F("method"),
						Path:   cadenya.F("path"),
					}),
				}),
				Description: cadenya.F("description"),
				Parameters: cadenya.F(map[string]interface{}{
					"foo": "bar",
				}),
				RequiresApproval: cadenya.F(true),
				LlmToolName:      cadenya.F("llmToolName"),
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
		"workspaceId",
		"toolSetId",
		"id",
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
		"workspaceId",
		"toolSetId",
		"id",
		cadenya.ToolSetToolUpdateParams{
			Metadata: cadenya.F(shared.UpdateResourceMetadataParam{
				Name:       cadenya.F("name"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.ToolSpecParam{
				Config: cadenya.F(cadenya.ToolSpecConfigParam{
					HTTP: cadenya.F(cadenya.ConfigHTTPParam{
						RequestMethod: cadenya.F(cadenya.ConfigHTTPRequestMethodHTTPMethodUnspecified),
						Headers: cadenya.F(map[string]string{
							"foo": "string",
						}),
						Path:                   cadenya.F("path"),
						Query:                  cadenya.F("query"),
						RequestBodyContentType: cadenya.F("requestBodyContentType"),
						RequestBodyTemplate:    cadenya.F("requestBodyTemplate"),
					}),
					Mcp: cadenya.F(cadenya.ConfigMcpParam{
						Annotations: cadenya.F(cadenya.McpAnnotationsParam{
							DestructiveHint: cadenya.F(true),
							IdempotentHint:  cadenya.F(true),
							OpenWorldHint:   cadenya.F(true),
							ReadOnlyHint:    cadenya.F(true),
							Title:           cadenya.F("title"),
						}),
					}),
					OpenAPI: cadenya.F(cadenya.ConfigOpenAPIParam{
						Method: cadenya.F("method"),
						Path:   cadenya.F("path"),
					}),
				}),
				Description: cadenya.F("description"),
				Parameters: cadenya.F(map[string]interface{}{
					"foo": "bar",
				}),
				RequiresApproval: cadenya.F(true),
				LlmToolName:      cadenya.F("llmToolName"),
			}),
			UpdateMask: cadenya.F("updateMask"),
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
		"workspaceId",
		"toolSetId",
		cadenya.ToolSetToolListParams{
			Cursor:           cadenya.F("cursor"),
			IncludeInfo:      cadenya.F(true),
			Limit:            cadenya.F(int64(0)),
			Names:            cadenya.F([]string{"string"}),
			Prefix:           cadenya.F("prefix"),
			Query:            cadenya.F("query"),
			RequiresApproval: cadenya.F(true),
			SortOrder:        cadenya.F("sortOrder"),
			States:           cadenya.F([]cadenya.ToolSetToolListParamsState{cadenya.ToolSetToolListParamsStateStateUnspecified}),
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
		"workspaceId",
		"toolSetId",
		"id",
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
		"workspaceId",
		"toolSetId",
		"id",
		cadenya.ToolSetToolOmitParams{},
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
		"workspaceId",
		"toolSetId",
		"id",
		cadenya.ToolSetToolRestoreParams{},
	)
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
