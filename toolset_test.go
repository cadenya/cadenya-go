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

func TestToolSetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.New(context.TODO(), cadenya.ToolSetNewParams{
		Metadata: cadenya.F(shared.CreateResourceMetadataParam{
			Name:       cadenya.F("name"),
			BundleKey:  cadenya.F("bundleKey"),
			ExternalID: cadenya.F("externalId"),
			Labels: cadenya.F(map[string]string{
				"foo": "string",
			}),
		}),
		Spec: cadenya.F(cadenya.ToolSetSpecParam{
			Adapter: cadenya.F(cadenya.ToolSetAdapterParam{
				HTTP: cadenya.F(cadenya.ToolSetAdapterHTTPParam{
					BaseURL: cadenya.F("baseUrl"),
					Headers: cadenya.F(map[string]string{
						"foo": "string",
					}),
				}),
				Mcp: cadenya.F(cadenya.ToolSetAdapterMcpParam{
					ExcludeTools: cadenya.F(cadenya.McpToolFilterParam{
						Operator: cadenya.F(cadenya.McpToolFilterOperatorOperatorUnspecified),
						Filters: cadenya.F([]cadenya.McpToolFilterFilterParam{{
							Attribute: cadenya.F(cadenya.McpToolFilterFiltersAttributeAttributeUnspecified),
							Matcher: cadenya.F(cadenya.McpToolFilterFiltersMatcherParam{
								CaseSensitive: cadenya.F(true),
								Contains:      cadenya.F("contains"),
								EndsWith:      cadenya.F("endsWith"),
								Exact:         cadenya.F("exact"),
								Regex:         cadenya.F("regex"),
								StartsWith:    cadenya.F("startsWith"),
							}),
						}}),
					}),
					Headers: cadenya.F(map[string]string{
						"foo": "string",
					}),
					IncludeTools: cadenya.F(cadenya.McpToolFilterParam{
						Operator: cadenya.F(cadenya.McpToolFilterOperatorOperatorUnspecified),
						Filters: cadenya.F([]cadenya.McpToolFilterFilterParam{{
							Attribute: cadenya.F(cadenya.McpToolFilterFiltersAttributeAttributeUnspecified),
							Matcher: cadenya.F(cadenya.McpToolFilterFiltersMatcherParam{
								CaseSensitive: cadenya.F(true),
								Contains:      cadenya.F("contains"),
								EndsWith:      cadenya.F("endsWith"),
								Exact:         cadenya.F("exact"),
								Regex:         cadenya.F("regex"),
								StartsWith:    cadenya.F("startsWith"),
							}),
						}}),
					}),
					ToolApprovals: cadenya.F(cadenya.ToolSetAdapterMcpToolApprovalsParam{
						Always: cadenya.F(true),
						Only: cadenya.F(cadenya.McpToolFilterParam{
							Operator: cadenya.F(cadenya.McpToolFilterOperatorOperatorUnspecified),
							Filters: cadenya.F([]cadenya.McpToolFilterFilterParam{{
								Attribute: cadenya.F(cadenya.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: cadenya.F(cadenya.McpToolFilterFiltersMatcherParam{
									CaseSensitive: cadenya.F(true),
									Contains:      cadenya.F("contains"),
									EndsWith:      cadenya.F("endsWith"),
									Exact:         cadenya.F("exact"),
									Regex:         cadenya.F("regex"),
									StartsWith:    cadenya.F("startsWith"),
								}),
							}}),
						}),
					}),
					URL: cadenya.F("url"),
				}),
			}),
			Description: cadenya.F("description"),
		}),
	})
	if err != nil {
		var apierr *cadenya.Error
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
	client := cadenya.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Get(context.TODO(), "id")
	if err != nil {
		var apierr *cadenya.Error
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
	client := cadenya.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.Update(
		context.TODO(),
		"id",
		cadenya.ToolSetUpdateParams{
			Metadata: cadenya.F(shared.UpdateResourceMetadataParam{
				Name:       cadenya.F("name"),
				BundleKey:  cadenya.F("bundleKey"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.ToolSetSpecParam{
				Adapter: cadenya.F(cadenya.ToolSetAdapterParam{
					HTTP: cadenya.F(cadenya.ToolSetAdapterHTTPParam{
						BaseURL: cadenya.F("baseUrl"),
						Headers: cadenya.F(map[string]string{
							"foo": "string",
						}),
					}),
					Mcp: cadenya.F(cadenya.ToolSetAdapterMcpParam{
						ExcludeTools: cadenya.F(cadenya.McpToolFilterParam{
							Operator: cadenya.F(cadenya.McpToolFilterOperatorOperatorUnspecified),
							Filters: cadenya.F([]cadenya.McpToolFilterFilterParam{{
								Attribute: cadenya.F(cadenya.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: cadenya.F(cadenya.McpToolFilterFiltersMatcherParam{
									CaseSensitive: cadenya.F(true),
									Contains:      cadenya.F("contains"),
									EndsWith:      cadenya.F("endsWith"),
									Exact:         cadenya.F("exact"),
									Regex:         cadenya.F("regex"),
									StartsWith:    cadenya.F("startsWith"),
								}),
							}}),
						}),
						Headers: cadenya.F(map[string]string{
							"foo": "string",
						}),
						IncludeTools: cadenya.F(cadenya.McpToolFilterParam{
							Operator: cadenya.F(cadenya.McpToolFilterOperatorOperatorUnspecified),
							Filters: cadenya.F([]cadenya.McpToolFilterFilterParam{{
								Attribute: cadenya.F(cadenya.McpToolFilterFiltersAttributeAttributeUnspecified),
								Matcher: cadenya.F(cadenya.McpToolFilterFiltersMatcherParam{
									CaseSensitive: cadenya.F(true),
									Contains:      cadenya.F("contains"),
									EndsWith:      cadenya.F("endsWith"),
									Exact:         cadenya.F("exact"),
									Regex:         cadenya.F("regex"),
									StartsWith:    cadenya.F("startsWith"),
								}),
							}}),
						}),
						ToolApprovals: cadenya.F(cadenya.ToolSetAdapterMcpToolApprovalsParam{
							Always: cadenya.F(true),
							Only: cadenya.F(cadenya.McpToolFilterParam{
								Operator: cadenya.F(cadenya.McpToolFilterOperatorOperatorUnspecified),
								Filters: cadenya.F([]cadenya.McpToolFilterFilterParam{{
									Attribute: cadenya.F(cadenya.McpToolFilterFiltersAttributeAttributeUnspecified),
									Matcher: cadenya.F(cadenya.McpToolFilterFiltersMatcherParam{
										CaseSensitive: cadenya.F(true),
										Contains:      cadenya.F("contains"),
										EndsWith:      cadenya.F("endsWith"),
										Exact:         cadenya.F("exact"),
										Regex:         cadenya.F("regex"),
										StartsWith:    cadenya.F("startsWith"),
									}),
								}}),
							}),
						}),
						URL: cadenya.F("url"),
					}),
				}),
				Description: cadenya.F("description"),
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

func TestToolSetListWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.List(context.TODO(), cadenya.ToolSetListParams{
		BundleKey:   cadenya.F("bundleKey"),
		Cursor:      cadenya.F("cursor"),
		IncludeInfo: cadenya.F(true),
		Limit:       cadenya.F(int64(0)),
		Prefix:      cadenya.F("prefix"),
		Query:       cadenya.F("query"),
		SortOrder:   cadenya.F("sortOrder"),
	})
	if err != nil {
		var apierr *cadenya.Error
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
	client := cadenya.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.ToolSets.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *cadenya.Error
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
	client := cadenya.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ToolSets.ListEvents(
		context.TODO(),
		"toolSetId",
		cadenya.ToolSetListEventsParams{
			Cursor:      cadenya.F("cursor"),
			IncludeInfo: cadenya.F(true),
			Limit:       cadenya.F(int64(0)),
			SortOrder:   cadenya.F("sortOrder"),
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
