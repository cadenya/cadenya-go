// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyasdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cadenya/cadenya-sdk-go"
	"github.com/cadenya/cadenya-sdk-go/internal/testutil"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/cadenya/cadenya-sdk-go/shared"
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Tools.New(
		context.TODO(),
		"toolSetId",
		gocadenyacomcadenyasdkgo.ToolSetToolNewParams{
			Metadata: gocadenyacomcadenyasdkgo.F(shared.CreateResourceMetadataParam{
				Name:       gocadenyacomcadenyasdkgo.F("name"),
				ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
				Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSpecParam{
				Config: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSpecConfigParam{
					HTTP: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ConfigHTTPParam{
						RequestMethod: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ConfigHTTPRequestMethodGet),
						Headers: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
						Path:                   gocadenyacomcadenyasdkgo.F("path"),
						Query:                  gocadenyacomcadenyasdkgo.F("query"),
						RequestBodyContentType: gocadenyacomcadenyasdkgo.F("requestBodyContentType"),
						RequestBodyTemplate:    gocadenyacomcadenyasdkgo.F("requestBodyTemplate"),
						ToolName:               gocadenyacomcadenyasdkgo.F("toolName"),
					}),
					Mcp: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ConfigMcpParam{
						ToolDescription: gocadenyacomcadenyasdkgo.F("toolDescription"),
						ToolName:        gocadenyacomcadenyasdkgo.F("toolName"),
						ToolTitle:       gocadenyacomcadenyasdkgo.F("toolTitle"),
					}),
				}),
				Description: gocadenyacomcadenyasdkgo.F("description"),
				Parameters: gocadenyacomcadenyasdkgo.F(map[string]interface{}{
					"foo": "bar",
				}),
				Status:           gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSpecStatusToolStatusUnspecified),
				RequiresApproval: gocadenyacomcadenyasdkgo.F(true),
			}),
		},
	)
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Tools.Get(
		context.TODO(),
		"toolSetId",
		"id",
	)
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Tools.Update(
		context.TODO(),
		"toolSetId",
		"id",
		gocadenyacomcadenyasdkgo.ToolSetToolUpdateParams{
			Metadata: gocadenyacomcadenyasdkgo.F(shared.UpdateResourceMetadataParam{
				Name:       gocadenyacomcadenyasdkgo.F("name"),
				ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
				Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSpecParam{
				Config: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSpecConfigParam{
					HTTP: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ConfigHTTPParam{
						RequestMethod: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ConfigHTTPRequestMethodGet),
						Headers: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
						Path:                   gocadenyacomcadenyasdkgo.F("path"),
						Query:                  gocadenyacomcadenyasdkgo.F("query"),
						RequestBodyContentType: gocadenyacomcadenyasdkgo.F("requestBodyContentType"),
						RequestBodyTemplate:    gocadenyacomcadenyasdkgo.F("requestBodyTemplate"),
						ToolName:               gocadenyacomcadenyasdkgo.F("toolName"),
					}),
					Mcp: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ConfigMcpParam{
						ToolDescription: gocadenyacomcadenyasdkgo.F("toolDescription"),
						ToolName:        gocadenyacomcadenyasdkgo.F("toolName"),
						ToolTitle:       gocadenyacomcadenyasdkgo.F("toolTitle"),
					}),
				}),
				Description: gocadenyacomcadenyasdkgo.F("description"),
				Parameters: gocadenyacomcadenyasdkgo.F(map[string]interface{}{
					"foo": "bar",
				}),
				Status:           gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSpecStatusToolStatusUnspecified),
				RequiresApproval: gocadenyacomcadenyasdkgo.F(true),
			}),
			UpdateMask: gocadenyacomcadenyasdkgo.F("updateMask"),
		},
	)
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Tools.List(
		context.TODO(),
		"toolSetId",
		gocadenyacomcadenyasdkgo.ToolSetToolListParams{
			Cursor:      gocadenyacomcadenyasdkgo.F("cursor"),
			IncludeInfo: gocadenyacomcadenyasdkgo.F(true),
			Limit:       gocadenyacomcadenyasdkgo.F(int64(0)),
			Prefix:      gocadenyacomcadenyasdkgo.F("prefix"),
			SortOrder:   gocadenyacomcadenyasdkgo.F("sortOrder"),
		},
	)
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.ToolSets.Tools.Delete(
		context.TODO(),
		"toolSetId",
		"id",
	)
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
