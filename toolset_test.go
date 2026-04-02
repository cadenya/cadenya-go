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

func TestToolSetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.New(context.TODO(), gocadenyacomcadenyago.ToolSetNewParams{
		Metadata: gocadenyacomcadenyago.F(shared.CreateResourceMetadataParam{
			Name:       gocadenyacomcadenyago.F("name"),
			ExternalID: gocadenyacomcadenyago.F("externalId"),
			Labels: gocadenyacomcadenyago.F(map[string]string{
				"foo": "string",
			}),
		}),
		Spec: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetSpecParam{
			Adapter: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetAdapterParam{
				HTTP: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetAdapterHTTPParam{
					BaseURL: gocadenyacomcadenyago.F("baseUrl"),
					Headers: gocadenyacomcadenyago.F(map[string]string{
						"foo": "string",
					}),
				}),
				Mcp: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetAdapterMcpParam{
					ExcludeTools: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterParam{
						Operator: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterOperatorOperatorUnspecified),
						Filters: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.McpToolFilterFilterParam{{
							Attribute: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersAttributeAttributeUnspecified),
							Matcher: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersMatcherParam{
								CaseSensitive: gocadenyacomcadenyago.F(true),
								Contains:      gocadenyacomcadenyago.F("contains"),
								EndsWith:      gocadenyacomcadenyago.F("endsWith"),
								Exact:         gocadenyacomcadenyago.F("exact"),
								Regex:         gocadenyacomcadenyago.F("regex"),
								StartsWith:    gocadenyacomcadenyago.F("startsWith"),
							}),
						}}),
					}),
					Headers: gocadenyacomcadenyago.F(map[string]string{
						"foo": "string",
					}),
					IncludeTools: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterParam{
						Operator: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterOperatorOperatorUnspecified),
						Filters: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.McpToolFilterFilterParam{{
							Attribute: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersAttributeAttributeUnspecified),
							Matcher: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersMatcherParam{
								CaseSensitive: gocadenyacomcadenyago.F(true),
								Contains:      gocadenyacomcadenyago.F("contains"),
								EndsWith:      gocadenyacomcadenyago.F("endsWith"),
								Exact:         gocadenyacomcadenyago.F("exact"),
								Regex:         gocadenyacomcadenyago.F("regex"),
								StartsWith:    gocadenyacomcadenyago.F("startsWith"),
							}),
						}}),
					}),
					ToolApprovals: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetAdapterMcpToolApprovalsParam{
						Always: gocadenyacomcadenyago.F(true),
						Only: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterParam{
							Operator: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterOperatorOperatorUnspecified),
							Filters: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.McpToolFilterFilterParam{{
								Attribute: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersMatcherParam{
									CaseSensitive: gocadenyacomcadenyago.F(true),
									Contains:      gocadenyacomcadenyago.F("contains"),
									EndsWith:      gocadenyacomcadenyago.F("endsWith"),
									Exact:         gocadenyacomcadenyago.F("exact"),
									Regex:         gocadenyacomcadenyago.F("regex"),
									StartsWith:    gocadenyacomcadenyago.F("startsWith"),
								}),
							}}),
						}),
					}),
					URL: gocadenyacomcadenyago.F("url"),
				}),
			}),
			Description: gocadenyacomcadenyago.F("description"),
		}),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Update(
		context.TODO(),
		"id",
		gocadenyacomcadenyago.ToolSetUpdateParams{
			Metadata: gocadenyacomcadenyago.F(shared.UpdateResourceMetadataParam{
				Name:       gocadenyacomcadenyago.F("name"),
				ExternalID: gocadenyacomcadenyago.F("externalId"),
				Labels: gocadenyacomcadenyago.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetSpecParam{
				Adapter: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetAdapterParam{
					HTTP: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetAdapterHTTPParam{
						BaseURL: gocadenyacomcadenyago.F("baseUrl"),
						Headers: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
					Mcp: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetAdapterMcpParam{
						ExcludeTools: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterParam{
							Operator: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterOperatorOperatorUnspecified),
							Filters: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.McpToolFilterFilterParam{{
								Attribute: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersMatcherParam{
									CaseSensitive: gocadenyacomcadenyago.F(true),
									Contains:      gocadenyacomcadenyago.F("contains"),
									EndsWith:      gocadenyacomcadenyago.F("endsWith"),
									Exact:         gocadenyacomcadenyago.F("exact"),
									Regex:         gocadenyacomcadenyago.F("regex"),
									StartsWith:    gocadenyacomcadenyago.F("startsWith"),
								}),
							}}),
						}),
						Headers: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
						IncludeTools: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterParam{
							Operator: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterOperatorOperatorUnspecified),
							Filters: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.McpToolFilterFilterParam{{
								Attribute: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersMatcherParam{
									CaseSensitive: gocadenyacomcadenyago.F(true),
									Contains:      gocadenyacomcadenyago.F("contains"),
									EndsWith:      gocadenyacomcadenyago.F("endsWith"),
									Exact:         gocadenyacomcadenyago.F("exact"),
									Regex:         gocadenyacomcadenyago.F("regex"),
									StartsWith:    gocadenyacomcadenyago.F("startsWith"),
								}),
							}}),
						}),
						ToolApprovals: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSetAdapterMcpToolApprovalsParam{
							Always: gocadenyacomcadenyago.F(true),
							Only: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterParam{
								Operator: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterOperatorOperatorUnspecified),
								Filters: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.McpToolFilterFilterParam{{
									Attribute: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersAttributeAttributeUnspecified),
									Matcher: gocadenyacomcadenyago.F(gocadenyacomcadenyago.McpToolFilterFiltersMatcherParam{
										CaseSensitive: gocadenyacomcadenyago.F(true),
										Contains:      gocadenyacomcadenyago.F("contains"),
										EndsWith:      gocadenyacomcadenyago.F("endsWith"),
										Exact:         gocadenyacomcadenyago.F("exact"),
										Regex:         gocadenyacomcadenyago.F("regex"),
										StartsWith:    gocadenyacomcadenyago.F("startsWith"),
									}),
								}}),
							}),
						}),
						URL: gocadenyacomcadenyago.F("url"),
					}),
				}),
				Description: gocadenyacomcadenyago.F("description"),
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

func TestToolSetListWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.List(context.TODO(), gocadenyacomcadenyago.ToolSetListParams{
		Cursor:      gocadenyacomcadenyago.F("cursor"),
		IncludeInfo: gocadenyacomcadenyago.F(true),
		Limit:       gocadenyacomcadenyago.F(int64(0)),
		Prefix:      gocadenyacomcadenyago.F("prefix"),
		SortOrder:   gocadenyacomcadenyago.F("sortOrder"),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.ToolSets.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.ListEvents(
		context.TODO(),
		"toolSetId",
		gocadenyacomcadenyago.ToolSetListEventsParams{
			Cursor:      gocadenyacomcadenyago.F("cursor"),
			IncludeInfo: gocadenyacomcadenyago.F(true),
			Limit:       gocadenyacomcadenyago.F(int64(0)),
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
