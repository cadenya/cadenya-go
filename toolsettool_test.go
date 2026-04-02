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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Tools.New(
		context.TODO(),
		"toolSetId",
		gocadenyacomcadenyago.ToolSetToolNewParams{
			Metadata: gocadenyacomcadenyago.F(shared.CreateResourceMetadataParam{
				Name:       gocadenyacomcadenyago.F("name"),
				ExternalID: gocadenyacomcadenyago.F("externalId"),
				Labels: gocadenyacomcadenyago.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSpecParam{
				Config: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSpecConfigParam{
					HTTP: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ConfigHTTPParam{
						RequestMethod: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ConfigHTTPRequestMethodGet),
						Headers: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
						Path:                   gocadenyacomcadenyago.F("path"),
						Query:                  gocadenyacomcadenyago.F("query"),
						RequestBodyContentType: gocadenyacomcadenyago.F("requestBodyContentType"),
						RequestBodyTemplate:    gocadenyacomcadenyago.F("requestBodyTemplate"),
						ToolName:               gocadenyacomcadenyago.F("toolName"),
					}),
					Mcp: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ConfigMcpParam{
						ToolDescription: gocadenyacomcadenyago.F("toolDescription"),
						ToolName:        gocadenyacomcadenyago.F("toolName"),
						ToolTitle:       gocadenyacomcadenyago.F("toolTitle"),
					}),
				}),
				Description: gocadenyacomcadenyago.F("description"),
				Parameters: gocadenyacomcadenyago.F(map[string]interface{}{
					"foo": "bar",
				}),
				Status:           gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSpecStatusToolStatusUnspecified),
				RequiresApproval: gocadenyacomcadenyago.F(true),
			}),
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

func TestToolSetToolGet(t *testing.T) {
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
	_, err := client.ToolSets.Tools.Get(
		context.TODO(),
		"toolSetId",
		"id",
	)
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Tools.Update(
		context.TODO(),
		"toolSetId",
		"id",
		gocadenyacomcadenyago.ToolSetToolUpdateParams{
			Metadata: gocadenyacomcadenyago.F(shared.UpdateResourceMetadataParam{
				Name:       gocadenyacomcadenyago.F("name"),
				ExternalID: gocadenyacomcadenyago.F("externalId"),
				Labels: gocadenyacomcadenyago.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSpecParam{
				Config: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSpecConfigParam{
					HTTP: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ConfigHTTPParam{
						RequestMethod: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ConfigHTTPRequestMethodGet),
						Headers: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
						Path:                   gocadenyacomcadenyago.F("path"),
						Query:                  gocadenyacomcadenyago.F("query"),
						RequestBodyContentType: gocadenyacomcadenyago.F("requestBodyContentType"),
						RequestBodyTemplate:    gocadenyacomcadenyago.F("requestBodyTemplate"),
						ToolName:               gocadenyacomcadenyago.F("toolName"),
					}),
					Mcp: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ConfigMcpParam{
						ToolDescription: gocadenyacomcadenyago.F("toolDescription"),
						ToolName:        gocadenyacomcadenyago.F("toolName"),
						ToolTitle:       gocadenyacomcadenyago.F("toolTitle"),
					}),
				}),
				Description: gocadenyacomcadenyago.F("description"),
				Parameters: gocadenyacomcadenyago.F(map[string]interface{}{
					"foo": "bar",
				}),
				Status:           gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSpecStatusToolStatusUnspecified),
				RequiresApproval: gocadenyacomcadenyago.F(true),
			}),
			UpdateMask: gocadenyacomcadenyago.F("updateMask"),
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

func TestToolSetToolListWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.Tools.List(
		context.TODO(),
		"toolSetId",
		gocadenyacomcadenyago.ToolSetToolListParams{
			Cursor:      gocadenyacomcadenyago.F("cursor"),
			IncludeInfo: gocadenyacomcadenyago.F(true),
			Limit:       gocadenyacomcadenyago.F(int64(0)),
			Prefix:      gocadenyacomcadenyago.F("prefix"),
			SortOrder:   gocadenyacomcadenyago.F("sortOrder"),
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

func TestToolSetToolDelete(t *testing.T) {
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
	err := client.ToolSets.Tools.Delete(
		context.TODO(),
		"toolSetId",
		"id",
	)
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
