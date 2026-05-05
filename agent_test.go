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
	_, err := client.Agents.New(
		context.TODO(),
		"workspaceId",
		cadenya.AgentNewParams{
			Metadata: cadenya.F(shared.CreateResourceMetadataParam{
				Name:       cadenya.F("name"),
				BundleKey:  cadenya.F("bundleKey"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.AgentSpecParam{
				Status:                 cadenya.F(cadenya.AgentSpecStatusAgentStatusUnspecified),
				VariationSelectionMode: cadenya.F(cadenya.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
				Description:            cadenya.F("description"),
				InputDataSchema:        cadenya.F[any](map[string]interface{}{}),
				WebhookEventsURL:       cadenya.F("webhookEventsUrl"),
			}),
			DefaultVariation: cadenya.F(cadenya.AgentNewParamsDefaultVariation{
				Metadata: cadenya.F(shared.CreateResourceMetadataParam{
					Name:       cadenya.F("name"),
					BundleKey:  cadenya.F("bundleKey"),
					ExternalID: cadenya.F("externalId"),
					Labels: cadenya.F(map[string]string{
						"foo": "string",
					}),
				}),
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
	_, err := client.Agents.Get(
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
		"workspaceId",
		"id",
		cadenya.AgentUpdateParams{
			Metadata: cadenya.F(shared.UpdateResourceMetadataParam{
				Name:       cadenya.F("name"),
				BundleKey:  cadenya.F("bundleKey"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.AgentSpecParam{
				Status:                 cadenya.F(cadenya.AgentSpecStatusAgentStatusUnspecified),
				VariationSelectionMode: cadenya.F(cadenya.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
				Description:            cadenya.F("description"),
				InputDataSchema:        cadenya.F[any](map[string]interface{}{}),
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
	_, err := client.Agents.List(
		context.TODO(),
		"workspaceId",
		cadenya.AgentListParams{
			BundleKey:              cadenya.F("bundleKey"),
			Cursor:                 cadenya.F("cursor"),
			IncludeInfo:            cadenya.F(true),
			Limit:                  cadenya.F(int64(0)),
			Prefix:                 cadenya.F("prefix"),
			Query:                  cadenya.F("query"),
			SortOrder:              cadenya.F("sortOrder"),
			Status:                 cadenya.F(cadenya.AgentListParamsStatusAgentStatusUnspecified),
			VariationSelectionMode: cadenya.F(cadenya.AgentListParamsVariationSelectionModeVariationSelectionModeUnspecified),
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
	err := client.Agents.Delete(
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
