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
		WorkspaceID: cadenya.String("workspaceId"),
		Metadata: shared.CreateResourceMetadataParam{
			Name:       "name",
			ExternalID: cadenya.String("externalId"),
			Labels: map[string]string{
				"foo": "string",
			},
		},
		Spec: cadenya.ToolSetSpecParam{
			Adapter: cadenya.ToolSetAdapterParam{
				Bare: cadenya.ToolSetAdapterBareParam{
					ContentTimeout: cadenya.Int(0),
				},
				HTTP: cadenya.ToolSetAdapterHTTPParam{
					BaseURL: cadenya.String("baseUrl"),
					Headers: map[string]string{
						"foo": "string",
					},
				},
				MCP: cadenya.ToolSetAdapterMCPParam{
					ExcludeTools: cadenya.ToolFilterParam{
						Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
						Filters: []cadenya.AttributeFilterParam{{
							Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
							Matcher: cadenya.StringMatcherParam{
								CaseSensitive: cadenya.Bool(true),
								Contains:      cadenya.String("contains"),
								EndsWith:      cadenya.String("endsWith"),
								Exact:         cadenya.String("exact"),
								Regex:         cadenya.String("regex"),
								StartsWith:    cadenya.String("startsWith"),
								Type:          cadenya.String("type"),
							},
						}},
					},
					Headers: map[string]string{
						"foo": "string",
					},
					IncludeTools: cadenya.ToolFilterParam{
						Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
						Filters: []cadenya.AttributeFilterParam{{
							Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
							Matcher: cadenya.StringMatcherParam{
								CaseSensitive: cadenya.Bool(true),
								Contains:      cadenya.String("contains"),
								EndsWith:      cadenya.String("endsWith"),
								Exact:         cadenya.String("exact"),
								Regex:         cadenya.String("regex"),
								StartsWith:    cadenya.String("startsWith"),
								Type:          cadenya.String("type"),
							},
						}},
					},
					JustInTime: cadenya.ToolSetAdapterMCPJustInTimeParam{
						Enabled:                      cadenya.Bool(true),
						FailObjectiveOnToolListError: cadenya.Bool(true),
					},
					ToolApprovals: cadenya.ApprovalRequirementFilterParam{
						Always: cadenya.Bool(true),
						Only: cadenya.ToolFilterParam{
							Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
							Filters: []cadenya.AttributeFilterParam{{
								Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
								Matcher: cadenya.StringMatcherParam{
									CaseSensitive: cadenya.Bool(true),
									Contains:      cadenya.String("contains"),
									EndsWith:      cadenya.String("endsWith"),
									Exact:         cadenya.String("exact"),
									Regex:         cadenya.String("regex"),
									StartsWith:    cadenya.String("startsWith"),
									Type:          cadenya.String("type"),
								},
							}},
						},
						Type: cadenya.String("type"),
					},
					URL: cadenya.String("url"),
				},
				OpenAPI: cadenya.ToolSetAdapterOpenAPIParam{
					BaseURL: cadenya.String("baseUrl"),
					ExcludeTools: cadenya.ToolFilterParam{
						Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
						Filters: []cadenya.AttributeFilterParam{{
							Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
							Matcher: cadenya.StringMatcherParam{
								CaseSensitive: cadenya.Bool(true),
								Contains:      cadenya.String("contains"),
								EndsWith:      cadenya.String("endsWith"),
								Exact:         cadenya.String("exact"),
								Regex:         cadenya.String("regex"),
								StartsWith:    cadenya.String("startsWith"),
								Type:          cadenya.String("type"),
							},
						}},
					},
					Headers: map[string]string{
						"foo": "string",
					},
					IncludeTools: cadenya.ToolFilterParam{
						Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
						Filters: []cadenya.AttributeFilterParam{{
							Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
							Matcher: cadenya.StringMatcherParam{
								CaseSensitive: cadenya.Bool(true),
								Contains:      cadenya.String("contains"),
								EndsWith:      cadenya.String("endsWith"),
								Exact:         cadenya.String("exact"),
								Regex:         cadenya.String("regex"),
								StartsWith:    cadenya.String("startsWith"),
								Type:          cadenya.String("type"),
							},
						}},
					},
					ServerName: cadenya.String("serverName"),
					ToolApprovals: cadenya.ApprovalRequirementFilterParam{
						Always: cadenya.Bool(true),
						Only: cadenya.ToolFilterParam{
							Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
							Filters: []cadenya.AttributeFilterParam{{
								Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
								Matcher: cadenya.StringMatcherParam{
									CaseSensitive: cadenya.Bool(true),
									Contains:      cadenya.String("contains"),
									EndsWith:      cadenya.String("endsWith"),
									Exact:         cadenya.String("exact"),
									Regex:         cadenya.String("regex"),
									StartsWith:    cadenya.String("startsWith"),
									Type:          cadenya.String("type"),
								},
							}},
						},
						Type: cadenya.String("type"),
					},
					Type:     cadenya.String("type"),
					UploadID: cadenya.String("upload_01HXKD2E5NQM3T9AYWCFZ05DNK"),
					URL:      cadenya.String("url"),
				},
				Type: cadenya.String("type"),
			},
			Description: cadenya.String("description"),
		},
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
	_, err := client.ToolSets.Get(
		context.TODO(),
		"id",
		cadenya.ToolSetGetParams{
			WorkspaceID: cadenya.String("workspaceId"),
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
			WorkspaceID: cadenya.String("workspaceId"),
			Metadata: shared.UpdateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.ToolSetSpecParam{
				Adapter: cadenya.ToolSetAdapterParam{
					Bare: cadenya.ToolSetAdapterBareParam{
						ContentTimeout: cadenya.Int(0),
					},
					HTTP: cadenya.ToolSetAdapterHTTPParam{
						BaseURL: cadenya.String("baseUrl"),
						Headers: map[string]string{
							"foo": "string",
						},
					},
					MCP: cadenya.ToolSetAdapterMCPParam{
						ExcludeTools: cadenya.ToolFilterParam{
							Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
							Filters: []cadenya.AttributeFilterParam{{
								Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
								Matcher: cadenya.StringMatcherParam{
									CaseSensitive: cadenya.Bool(true),
									Contains:      cadenya.String("contains"),
									EndsWith:      cadenya.String("endsWith"),
									Exact:         cadenya.String("exact"),
									Regex:         cadenya.String("regex"),
									StartsWith:    cadenya.String("startsWith"),
									Type:          cadenya.String("type"),
								},
							}},
						},
						Headers: map[string]string{
							"foo": "string",
						},
						IncludeTools: cadenya.ToolFilterParam{
							Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
							Filters: []cadenya.AttributeFilterParam{{
								Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
								Matcher: cadenya.StringMatcherParam{
									CaseSensitive: cadenya.Bool(true),
									Contains:      cadenya.String("contains"),
									EndsWith:      cadenya.String("endsWith"),
									Exact:         cadenya.String("exact"),
									Regex:         cadenya.String("regex"),
									StartsWith:    cadenya.String("startsWith"),
									Type:          cadenya.String("type"),
								},
							}},
						},
						JustInTime: cadenya.ToolSetAdapterMCPJustInTimeParam{
							Enabled:                      cadenya.Bool(true),
							FailObjectiveOnToolListError: cadenya.Bool(true),
						},
						ToolApprovals: cadenya.ApprovalRequirementFilterParam{
							Always: cadenya.Bool(true),
							Only: cadenya.ToolFilterParam{
								Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
								Filters: []cadenya.AttributeFilterParam{{
									Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
									Matcher: cadenya.StringMatcherParam{
										CaseSensitive: cadenya.Bool(true),
										Contains:      cadenya.String("contains"),
										EndsWith:      cadenya.String("endsWith"),
										Exact:         cadenya.String("exact"),
										Regex:         cadenya.String("regex"),
										StartsWith:    cadenya.String("startsWith"),
										Type:          cadenya.String("type"),
									},
								}},
							},
							Type: cadenya.String("type"),
						},
						URL: cadenya.String("url"),
					},
					OpenAPI: cadenya.ToolSetAdapterOpenAPIParam{
						BaseURL: cadenya.String("baseUrl"),
						ExcludeTools: cadenya.ToolFilterParam{
							Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
							Filters: []cadenya.AttributeFilterParam{{
								Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
								Matcher: cadenya.StringMatcherParam{
									CaseSensitive: cadenya.Bool(true),
									Contains:      cadenya.String("contains"),
									EndsWith:      cadenya.String("endsWith"),
									Exact:         cadenya.String("exact"),
									Regex:         cadenya.String("regex"),
									StartsWith:    cadenya.String("startsWith"),
									Type:          cadenya.String("type"),
								},
							}},
						},
						Headers: map[string]string{
							"foo": "string",
						},
						IncludeTools: cadenya.ToolFilterParam{
							Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
							Filters: []cadenya.AttributeFilterParam{{
								Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
								Matcher: cadenya.StringMatcherParam{
									CaseSensitive: cadenya.Bool(true),
									Contains:      cadenya.String("contains"),
									EndsWith:      cadenya.String("endsWith"),
									Exact:         cadenya.String("exact"),
									Regex:         cadenya.String("regex"),
									StartsWith:    cadenya.String("startsWith"),
									Type:          cadenya.String("type"),
								},
							}},
						},
						ServerName: cadenya.String("serverName"),
						ToolApprovals: cadenya.ApprovalRequirementFilterParam{
							Always: cadenya.Bool(true),
							Only: cadenya.ToolFilterParam{
								Operator: cadenya.ToolFilterOperatorOperatorUnspecified,
								Filters: []cadenya.AttributeFilterParam{{
									Attribute: cadenya.AttributeFilterAttributeAttributeUnspecified,
									Matcher: cadenya.StringMatcherParam{
										CaseSensitive: cadenya.Bool(true),
										Contains:      cadenya.String("contains"),
										EndsWith:      cadenya.String("endsWith"),
										Exact:         cadenya.String("exact"),
										Regex:         cadenya.String("regex"),
										StartsWith:    cadenya.String("startsWith"),
										Type:          cadenya.String("type"),
									},
								}},
							},
							Type: cadenya.String("type"),
						},
						Type:     cadenya.String("type"),
						UploadID: cadenya.String("upload_01HXKD2E5NQM3T9AYWCFZ05DNK"),
						URL:      cadenya.String("url"),
					},
					Type: cadenya.String("type"),
				},
				Description: cadenya.String("description"),
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
		WorkspaceID: cadenya.String("workspaceId"),
		Cursor:      cadenya.String("cursor"),
		IncludeInfo: cadenya.Bool(true),
		Labels:      cadenya.String("labels"),
		Limit:       cadenya.Int(0),
		Prefix:      cadenya.String("prefix"),
		Query:       cadenya.String("query"),
		SortOrder:   cadenya.String("sortOrder"),
		State:       cadenya.ToolSetListParamsStateStateUnspecified,
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
	err := client.ToolSets.Delete(
		context.TODO(),
		"id",
		cadenya.ToolSetDeleteParams{
			WorkspaceID: cadenya.String("workspaceId"),
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

func TestToolSetArchive(t *testing.T) {
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
	_, err := client.ToolSets.Archive(
		context.TODO(),
		"id",
		cadenya.ToolSetArchiveParams{
			WorkspaceID: cadenya.String("workspaceId"),
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

func TestToolSetGetOpenAPISpec(t *testing.T) {
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
	_, err := client.ToolSets.GetOpenAPISpec(
		context.TODO(),
		"toolSetId",
		cadenya.ToolSetGetOpenAPISpecParams{
			WorkspaceID: cadenya.String("workspaceId"),
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
			WorkspaceID: cadenya.String("workspaceId"),
			Cursor:      cadenya.String("cursor"),
			IncludeInfo: cadenya.Bool(true),
			Labels:      cadenya.String("labels"),
			Limit:       cadenya.Int(0),
			SortOrder:   cadenya.String("sortOrder"),
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

func TestToolSetListUsageWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolSets.ListUsage(
		context.TODO(),
		"toolSetId",
		cadenya.ToolSetListUsageParams{
			WorkspaceID: cadenya.String("workspaceId"),
			Cursor:      cadenya.String("cursor"),
			Limit:       cadenya.Int(0),
			SortOrder:   cadenya.String("sortOrder"),
			ToolID:      cadenya.String("toolId"),
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

func TestToolSetUnarchive(t *testing.T) {
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
	_, err := client.ToolSets.Unarchive(
		context.TODO(),
		"id",
		cadenya.ToolSetUnarchiveParams{
			WorkspaceID: cadenya.String("workspaceId"),
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
