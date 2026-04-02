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

func TestToolSetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.New(context.TODO(), gocadenyacomcadenyasdkgo.ToolSetNewParams{
		Metadata: gocadenyacomcadenyasdkgo.F(shared.CreateResourceMetadataParam{
			Name:       gocadenyacomcadenyasdkgo.F("name"),
			ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
			Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
				"foo": "string",
			}),
		}),
		Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetSpecParam{
			Adapter: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetAdapterParam{
				HTTP: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetAdapterHTTPParam{
					BaseURL: gocadenyacomcadenyasdkgo.F("baseUrl"),
					Headers: gocadenyacomcadenyasdkgo.F(map[string]string{
						"foo": "string",
					}),
				}),
				Mcp: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetAdapterMcpParam{
					ExcludeTools: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterParam{
						Operator: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterOperatorOperatorUnspecified),
						Filters: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.McpToolFilterFilterParam{{
							Attribute: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersAttributeAttributeUnspecified),
							Matcher: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersMatcherParam{
								CaseSensitive: gocadenyacomcadenyasdkgo.F(true),
								Contains:      gocadenyacomcadenyasdkgo.F("contains"),
								EndsWith:      gocadenyacomcadenyasdkgo.F("endsWith"),
								Exact:         gocadenyacomcadenyasdkgo.F("exact"),
								Regex:         gocadenyacomcadenyasdkgo.F("regex"),
								StartsWith:    gocadenyacomcadenyasdkgo.F("startsWith"),
							}),
						}}),
					}),
					Headers: gocadenyacomcadenyasdkgo.F(map[string]string{
						"foo": "string",
					}),
					IncludeTools: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterParam{
						Operator: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterOperatorOperatorUnspecified),
						Filters: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.McpToolFilterFilterParam{{
							Attribute: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersAttributeAttributeUnspecified),
							Matcher: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersMatcherParam{
								CaseSensitive: gocadenyacomcadenyasdkgo.F(true),
								Contains:      gocadenyacomcadenyasdkgo.F("contains"),
								EndsWith:      gocadenyacomcadenyasdkgo.F("endsWith"),
								Exact:         gocadenyacomcadenyasdkgo.F("exact"),
								Regex:         gocadenyacomcadenyasdkgo.F("regex"),
								StartsWith:    gocadenyacomcadenyasdkgo.F("startsWith"),
							}),
						}}),
					}),
					ToolApprovals: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetAdapterMcpToolApprovalsParam{
						Always: gocadenyacomcadenyasdkgo.F(true),
						Only: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterParam{
							Operator: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterOperatorOperatorUnspecified),
							Filters: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.McpToolFilterFilterParam{{
								Attribute: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersMatcherParam{
									CaseSensitive: gocadenyacomcadenyasdkgo.F(true),
									Contains:      gocadenyacomcadenyasdkgo.F("contains"),
									EndsWith:      gocadenyacomcadenyasdkgo.F("endsWith"),
									Exact:         gocadenyacomcadenyasdkgo.F("exact"),
									Regex:         gocadenyacomcadenyasdkgo.F("regex"),
									StartsWith:    gocadenyacomcadenyasdkgo.F("startsWith"),
								}),
							}}),
						}),
					}),
					URL: gocadenyacomcadenyasdkgo.F("url"),
				}),
			}),
			Description: gocadenyacomcadenyasdkgo.F("description"),
		}),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolSetGet(t *testing.T) {
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
	_, err := client.ToolSets.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolSetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.Update(
		context.TODO(),
		"id",
		gocadenyacomcadenyasdkgo.ToolSetUpdateParams{
			Metadata: gocadenyacomcadenyasdkgo.F(shared.UpdateResourceMetadataParam{
				Name:       gocadenyacomcadenyasdkgo.F("name"),
				ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
				Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetSpecParam{
				Adapter: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetAdapterParam{
					HTTP: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetAdapterHTTPParam{
						BaseURL: gocadenyacomcadenyasdkgo.F("baseUrl"),
						Headers: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
					}),
					Mcp: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetAdapterMcpParam{
						ExcludeTools: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterParam{
							Operator: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterOperatorOperatorUnspecified),
							Filters: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.McpToolFilterFilterParam{{
								Attribute: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersMatcherParam{
									CaseSensitive: gocadenyacomcadenyasdkgo.F(true),
									Contains:      gocadenyacomcadenyasdkgo.F("contains"),
									EndsWith:      gocadenyacomcadenyasdkgo.F("endsWith"),
									Exact:         gocadenyacomcadenyasdkgo.F("exact"),
									Regex:         gocadenyacomcadenyasdkgo.F("regex"),
									StartsWith:    gocadenyacomcadenyasdkgo.F("startsWith"),
								}),
							}}),
						}),
						Headers: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
						IncludeTools: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterParam{
							Operator: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterOperatorOperatorUnspecified),
							Filters: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.McpToolFilterFilterParam{{
								Attribute: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersMatcherParam{
									CaseSensitive: gocadenyacomcadenyasdkgo.F(true),
									Contains:      gocadenyacomcadenyasdkgo.F("contains"),
									EndsWith:      gocadenyacomcadenyasdkgo.F("endsWith"),
									Exact:         gocadenyacomcadenyasdkgo.F("exact"),
									Regex:         gocadenyacomcadenyasdkgo.F("regex"),
									StartsWith:    gocadenyacomcadenyasdkgo.F("startsWith"),
								}),
							}}),
						}),
						ToolApprovals: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSetAdapterMcpToolApprovalsParam{
							Always: gocadenyacomcadenyasdkgo.F(true),
							Only: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterParam{
								Operator: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterOperatorOperatorUnspecified),
								Filters: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.McpToolFilterFilterParam{{
									Attribute: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersAttributeAttributeUnspecified),
									Matcher: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.McpToolFilterFiltersMatcherParam{
										CaseSensitive: gocadenyacomcadenyasdkgo.F(true),
										Contains:      gocadenyacomcadenyasdkgo.F("contains"),
										EndsWith:      gocadenyacomcadenyasdkgo.F("endsWith"),
										Exact:         gocadenyacomcadenyasdkgo.F("exact"),
										Regex:         gocadenyacomcadenyasdkgo.F("regex"),
										StartsWith:    gocadenyacomcadenyasdkgo.F("startsWith"),
									}),
								}}),
							}),
						}),
						URL: gocadenyacomcadenyasdkgo.F("url"),
					}),
				}),
				Description: gocadenyacomcadenyasdkgo.F("description"),
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

func TestToolSetListWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.List(context.TODO(), gocadenyacomcadenyasdkgo.ToolSetListParams{
		Cursor:      gocadenyacomcadenyasdkgo.F("cursor"),
		IncludeInfo: gocadenyacomcadenyasdkgo.F(true),
		Limit:       gocadenyacomcadenyasdkgo.F(int64(0)),
		Prefix:      gocadenyacomcadenyasdkgo.F("prefix"),
		SortOrder:   gocadenyacomcadenyasdkgo.F("sortOrder"),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolSetDelete(t *testing.T) {
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
	err := client.ToolSets.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolSetListEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.ListEvents(
		context.TODO(),
		"toolSetId",
		gocadenyacomcadenyasdkgo.ToolSetListEventsParams{
			Cursor:      gocadenyacomcadenyasdkgo.F("cursor"),
			IncludeInfo: gocadenyacomcadenyasdkgo.F(true),
			Limit:       gocadenyacomcadenyasdkgo.F(int64(0)),
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
