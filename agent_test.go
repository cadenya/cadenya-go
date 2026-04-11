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

func TestAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.New(context.TODO(), cadenya.AgentNewParams{
		Metadata: cadenya.F(shared.CreateResourceMetadataParam{
			Name:       cadenya.F("name"),
			ExternalID: cadenya.F("externalId"),
			Labels: cadenya.F(map[string]string{
				"foo": "string",
			}),
		}),
		Spec: cadenya.F(cadenya.AgentSpecParam{
			Status:                 cadenya.F(cadenya.AgentSpecStatusAgentStatusUnspecified),
			VariationSelectionMode: cadenya.F(cadenya.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
			Description:            cadenya.F("description"),
			WebhookEventsURL:       cadenya.F("webhookEventsUrl"),
		}),
		DefaultVariation: cadenya.F(cadenya.AgentNewParamsDefaultVariation{
			Metadata: cadenya.F(shared.CreateResourceMetadataParam{
				Name:       cadenya.F("name"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.AgentVariationSpecParam{
				AgentTools: cadenya.F([]cadenya.AgentVariationSpecAgentToolParam{{
					AgentID: cadenya.F("agentId"),
					AgentMetadata: cadenya.F(shared.ResourceMetadataParam{
						Name:       cadenya.F("name"),
						ExternalID: cadenya.F("externalId"),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
					}),
					ToolID: cadenya.F("toolId"),
					ToolMetadata: cadenya.F(shared.ResourceMetadataParam{
						Name:       cadenya.F("name"),
						ExternalID: cadenya.F("externalId"),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
					}),
					ToolSetID: cadenya.F("toolSetId"),
					ToolSetMetadata: cadenya.F(shared.ResourceMetadataParam{
						Name:       cadenya.F("name"),
						ExternalID: cadenya.F("externalId"),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
				CompactionConfig: cadenya.F(cadenya.AgentVariationSpecCompactionConfigParam{
					Summarization: cadenya.F(cadenya.CompactionConfigSummarizationStrategyParam{
						Instructions:     cadenya.F("instructions"),
						MinPreserveTurns: cadenya.F(int64(0)),
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
				Prompt: cadenya.F("prompt"),
				ToolSelection: cadenya.F(cadenya.AgentVariationSpecToolSelectionParam{
					AssignedTools: cadenya.F(cadenya.ToolSelectionAssignedToolsParam{
						AllowDiscovery: cadenya.F(true),
					}),
					AutoDiscovery: cadenya.F(cadenya.ToolSelectionAutoDiscoveryParam{
						Hints:    cadenya.F([]string{"string"}),
						MaxTools: cadenya.F(int64(0)),
					}),
				}),
				Weight: cadenya.F(int64(0)),
			}),
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

func TestAgentGet(t *testing.T) {
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
	_, err := client.Agents.Get(context.TODO(), "id")
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Update(
		context.TODO(),
		"id",
		cadenya.AgentUpdateParams{
			Metadata: cadenya.F(shared.UpdateResourceMetadataParam{
				Name:       cadenya.F("name"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.AgentSpecParam{
				Status:                 cadenya.F(cadenya.AgentSpecStatusAgentStatusUnspecified),
				VariationSelectionMode: cadenya.F(cadenya.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
				Description:            cadenya.F("description"),
				WebhookEventsURL:       cadenya.F("webhookEventsUrl"),
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

func TestAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.List(context.TODO(), cadenya.AgentListParams{
		Cursor:      cadenya.F("cursor"),
		IncludeInfo: cadenya.F(true),
		Limit:       cadenya.F(int64(0)),
		Prefix:      cadenya.F("prefix"),
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

func TestAgentDelete(t *testing.T) {
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
	err := client.Agents.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
