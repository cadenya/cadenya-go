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
						ToolName:               cadenya.F("toolName"),
					}),
					Mcp: cadenya.F(cadenya.ConfigMcpParam{
						ToolDescription: cadenya.F("toolDescription"),
						ToolName:        cadenya.F("toolName"),
						ToolTitle:       cadenya.F("toolTitle"),
					}),
				}),
				Description: cadenya.F("description"),
				Parameters: cadenya.F(map[string]interface{}{
					"foo": "bar",
				}),
				Status:           cadenya.F(cadenya.ToolSpecStatusToolStatusUnspecified),
				RequiresApproval: cadenya.F(true),
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
						ToolName:               cadenya.F("toolName"),
					}),
					Mcp: cadenya.F(cadenya.ConfigMcpParam{
						ToolDescription: cadenya.F("toolDescription"),
						ToolName:        cadenya.F("toolName"),
						ToolTitle:       cadenya.F("toolTitle"),
					}),
				}),
				Description: cadenya.F("description"),
				Parameters: cadenya.F(map[string]interface{}{
					"foo": "bar",
				}),
				Status:           cadenya.F(cadenya.ToolSpecStatusToolStatusUnspecified),
				RequiresApproval: cadenya.F(true),
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
		"toolSetId",
		cadenya.ToolSetToolListParams{
			Cursor:      cadenya.F("cursor"),
			IncludeInfo: cadenya.F(true),
			Limit:       cadenya.F(int64(0)),
			Names:       cadenya.F([]string{"string"}),
			Prefix:      cadenya.F("prefix"),
			Query:       cadenya.F("query"),
			SortOrder:   cadenya.F("sortOrder"),
			Statuses:    cadenya.F([]cadenya.ToolSetToolListParamsStatus{cadenya.ToolSetToolListParamsStatusToolStatusUnspecified}),
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
