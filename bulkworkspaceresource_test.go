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
)

func TestBulkWorkspaceResourceGet(t *testing.T) {
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
	_, err := client.BulkWorkspaceResources.Get(
		context.TODO(),
		"workspaceId",
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

func TestBulkWorkspaceResourceListWithOptionalParams(t *testing.T) {
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
	_, err := client.BulkWorkspaceResources.List(
		context.TODO(),
		"workspaceId",
		cadenya.BulkWorkspaceResourceListParams{
			BundleKey: cadenya.F("bundleKey"),
			Cursor:    cadenya.F("cursor"),
			Limit:     cadenya.F(int64(0)),
			SortOrder: cadenya.F("sortOrder"),
			State:     cadenya.F(cadenya.BulkWorkspaceResourceListParamsStateStateUnspecified),
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

func TestBulkWorkspaceResourceApplyWithOptionalParams(t *testing.T) {
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
	_, err := client.BulkWorkspaceResources.Apply(
		context.TODO(),
		"workspaceId",
		cadenya.BulkWorkspaceResourceApplyParams{
			Data: cadenya.F(cadenya.BulkWorkspaceApplyDataParam{
				BundleKey: cadenya.F("bundleKey"),
				Agents: cadenya.F(map[string]cadenya.AgentEntryParam{
					"foo": {
						Name: cadenya.F("name"),
						Spec: cadenya.F(cadenya.AgentSpecParam{
							Status:                 cadenya.F(cadenya.AgentSpecStatusAgentStatusUnspecified),
							VariationSelectionMode: cadenya.F(cadenya.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
							Description:            cadenya.F("description"),
							InputDataSchema:        cadenya.F[any](map[string]interface{}{}),
							WebhookEventsURL:       cadenya.F("webhookEventsUrl"),
						}),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
						Schedules: cadenya.F(map[string]cadenya.AgentScheduleEntryParam{
							"foo": {
								Name: cadenya.F("name"),
								Spec: cadenya.F(cadenya.AgentScheduleSpecParam{
									InitialMessage: cadenya.F("initialMessage"),
									Schedule: cadenya.F(cadenya.AgentScheduleSpecScheduleParam{
										Calendars: cadenya.F([]cadenya.ScheduleCalendarParam{{
											Comment: cadenya.F("comment"),
											DayOfMonth: cadenya.F([]cadenya.ScheduleRangeParam{{
												End:   cadenya.F(int64(0)),
												Start: cadenya.F(int64(0)),
												Step:  cadenya.F(int64(0)),
											}}),
											DayOfWeek: cadenya.F([]cadenya.ScheduleRangeParam{{
												End:   cadenya.F(int64(0)),
												Start: cadenya.F(int64(0)),
												Step:  cadenya.F(int64(0)),
											}}),
											Hour: cadenya.F([]cadenya.ScheduleRangeParam{{
												End:   cadenya.F(int64(0)),
												Start: cadenya.F(int64(0)),
												Step:  cadenya.F(int64(0)),
											}}),
											Minute: cadenya.F([]cadenya.ScheduleRangeParam{{
												End:   cadenya.F(int64(0)),
												Start: cadenya.F(int64(0)),
												Step:  cadenya.F(int64(0)),
											}}),
											Month: cadenya.F([]cadenya.ScheduleRangeParam{{
												End:   cadenya.F(int64(0)),
												Start: cadenya.F(int64(0)),
												Step:  cadenya.F(int64(0)),
											}}),
											Second: cadenya.F([]cadenya.ScheduleRangeParam{{
												End:   cadenya.F(int64(0)),
												Start: cadenya.F(int64(0)),
												Step:  cadenya.F(int64(0)),
											}}),
										}}),
										Intervals: cadenya.F([]cadenya.ScheduleIntervalParam{{
											Every:  cadenya.F("-160513s"),
											Offset: cadenya.F("-160513s"),
										}}),
										Timezone: cadenya.F("timezone"),
									}),
									Data:          cadenya.F[any](map[string]interface{}{}),
									OverlapPolicy: cadenya.F(cadenya.AgentScheduleSpecOverlapPolicyOverlapPolicyUnspecified),
									Status:        cadenya.F(cadenya.AgentScheduleSpecStatusAgentScheduleStatusUnspecified),
									VariationID:   cadenya.F("variationId"),
								}),
								Labels: cadenya.F(map[string]string{
									"foo": "string",
								}),
							},
						}),
						Variations: cadenya.F(map[string]cadenya.AgentVariationEntryParam{
							"foo": {
								Name: cadenya.F("name"),
								Spec: cadenya.F(cadenya.AgentVariationSpecParam{
									CompactionConfig: cadenya.F(cadenya.AgentVariationSpecCompactionConfigParam{
										Summarization: cadenya.F(cadenya.CompactionConfigSummarizationStrategyParam{
											Instructions: cadenya.F("instructions"),
										}),
										ToolResultClearing: cadenya.F(cadenya.CompactionConfigToolResultClearingStrategyParam{
											PreserveRecentResults: cadenya.F(int64(0)),
										}),
										TriggerThreshold: cadenya.F(0.000000),
									}),
									Constraints: cadenya.F(cadenya.AgentVariationSpecConstraintsParam{
										MaxSubObjectives: cadenya.F(int64(0)),
										MaxToolCalls:     cadenya.F(int64(0)),
									}),
									Description:          cadenya.F("description"),
									EnableEpisodicMemory: cadenya.F(true),
									EpisodicMemoryTtl:    cadenya.F(int64(0)),
									ModelConfig: cadenya.F(cadenya.AgentVariationSpecModelConfigParam{
										ModelID:     cadenya.F("modelId"),
										Temperature: cadenya.F(0.000000),
									}),
									ProgressiveDiscovery: cadenya.F(cadenya.AgentVariationSpecProgressiveDiscoveryParam{
										Hints:           cadenya.F([]string{"string"}),
										MaxTools:        cadenya.F(int64(0)),
										RerankThreshold: cadenya.F(0.000000),
									}),
									Prompt: cadenya.F("prompt"),
									Weight: cadenya.F(int64(0)),
								}),
								Assignments: cadenya.F([]cadenya.VariationAssignmentEntryParam{{
									SubAgentID: cadenya.F("subAgentId"),
									ToolID:     cadenya.F("toolId"),
									ToolSetID:  cadenya.F("toolSetId"),
								}}),
								Labels: cadenya.F(map[string]string{
									"foo": "string",
								}),
								MemoryLayers: cadenya.F([]cadenya.VariationMemoryLayerEntryParam{{
									MemoryLayerID: cadenya.F("memoryLayerId"),
									Position:      cadenya.F(int64(0)),
								}}),
							},
						}),
					},
				}),
				AutomaticallyPublishAgents: cadenya.F(true),
				MemoryLayers: cadenya.F(map[string]cadenya.MemoryLayerEntryParam{
					"foo": {
						Name: cadenya.F("name"),
						Spec: cadenya.F(cadenya.MemoryLayerSpecParam{
							Type:        cadenya.F(cadenya.MemoryLayerSpecTypeMemoryLayerTypeUnspecified),
							Description: cadenya.F("description"),
						}),
						Entries: cadenya.F(map[string]cadenya.MemoryEntryItemParam{
							"foo": {
								Key:         cadenya.F("key"),
								Content:     cadenya.F("content"),
								Description: cadenya.F("description"),
								UploadID:    cadenya.F("uploadId"),
							},
						}),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
					},
				}),
				SourceURL: cadenya.F("sourceUrl"),
				ToolSets: cadenya.F(map[string]cadenya.ToolSetEntryParam{
					"foo": {
						Name: cadenya.F("name"),
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
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
						Tools: cadenya.F(map[string]cadenya.ToolEntryParam{
							"foo": {
								Name: cadenya.F("name"),
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
								Labels: cadenya.F(map[string]string{
									"foo": "string",
								}),
							},
						}),
					},
				}),
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
