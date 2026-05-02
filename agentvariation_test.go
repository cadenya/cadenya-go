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

func TestAgentVariationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Variations.New(
		context.TODO(),
		"agentId",
		cadenya.AgentVariationNewParams{
			Metadata: cadenya.F(shared.CreateResourceMetadataParam{
				Name:       cadenya.F("name"),
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

func TestAgentVariationGet(t *testing.T) {
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
	_, err := client.Agents.Variations.Get(
		context.TODO(),
		"agentId",
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

func TestAgentVariationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Variations.Update(
		context.TODO(),
		"agentId",
		"id",
		cadenya.AgentVariationUpdateParams{
			Metadata: cadenya.F(shared.UpdateResourceMetadataParam{
				Name:       cadenya.F("name"),
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

func TestAgentVariationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Variations.List(
		context.TODO(),
		"agentId",
		cadenya.AgentVariationListParams{
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

func TestAgentVariationDelete(t *testing.T) {
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
	err := client.Agents.Variations.Delete(
		context.TODO(),
		"agentId",
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

func TestAgentVariationAddAssignmentWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Variations.AddAssignment(
		context.TODO(),
		"agentId",
		"variationId",
		cadenya.AgentVariationAddAssignmentParams{
			SubAgentID: cadenya.F("subAgentId"),
			ToolID:     cadenya.F("toolId"),
			ToolSetID:  cadenya.F("toolSetId"),
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

func TestAgentVariationAddMemoryLayerWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Variations.AddMemoryLayer(
		context.TODO(),
		"agentId",
		"variationId",
		cadenya.AgentVariationAddMemoryLayerParams{
			MemoryLayerID: cadenya.F("memoryLayerId"),
			Position:      cadenya.F(int64(0)),
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

func TestAgentVariationRemoveAssignment(t *testing.T) {
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
	err := client.Agents.Variations.RemoveAssignment(
		context.TODO(),
		"agentId",
		"variationId",
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

func TestAgentVariationRemoveMemoryLayer(t *testing.T) {
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
	err := client.Agents.Variations.RemoveMemoryLayer(
		context.TODO(),
		"agentId",
		"variationId",
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

func TestAgentVariationUpdateMemoryLayerWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Variations.UpdateMemoryLayer(
		context.TODO(),
		"agentId",
		"variationId",
		"id",
		cadenya.AgentVariationUpdateMemoryLayerParams{
			Position: cadenya.F(int64(0)),
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
