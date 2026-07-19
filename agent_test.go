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
		WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
		Metadata: shared.CreateResourceMetadataParam{
			Name:       "name",
			ExternalID: cadenya.String("externalId"),
			Labels: map[string]string{
				"foo": "string",
			},
		},
		Spec: cadenya.AgentSpecParam{
			VariationSelectionMode: cadenya.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified,
			Description:            cadenya.String("description"),
			EnableEpisodicMemory:   cadenya.Bool(true),
			EpisodicMemoryTtl:      cadenya.Int(0),
			OutputDefinition: map[string]any{
				"foo": "bar",
			},
			SystemPromptDataSchema: map[string]any{
				"foo": "bar",
			},
			WebhookEventsURL: cadenya.String("webhookEventsUrl"),
		},
		DefaultVariation: cadenya.AgentNewParamsDefaultVariation{
			Metadata: shared.CreateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.AgentVariationSpecParam{
				CompactionConfig: cadenya.AgentVariationSpecCompactionConfigParam{
					Summarization: cadenya.CompactionConfigSummarizationStrategyParam{
						Instructions: cadenya.String("instructions"),
					},
					ToolResultClearing: cadenya.CompactionConfigToolResultClearingStrategyParam{
						PreserveRecentResults: cadenya.Int(0),
					},
					TriggerThreshold: cadenya.Float(0),
				},
				Constraints: cadenya.AgentVariationSpecConstraintsParam{
					InactivityTimeout: cadenya.String("-160513s"),
					MaxSubObjectives:  cadenya.Int(0),
					MaxToolCalls:      cadenya.Int(0),
				},
				Description:              cadenya.String("description"),
				FirstUserMessageTemplate: cadenya.String("firstUserMessageTemplate"),
				ModelConfig: cadenya.AgentVariationSpecModelConfigParam{
					ModelID:     cadenya.String("claude/opus-4.6"),
					Temperature: cadenya.Float(0),
				},
				ProgressiveDiscovery: cadenya.AgentVariationSpecProgressiveDiscoveryParam{
					Hints:    []string{"string"},
					MaxTools: cadenya.Int(0),
				},
				SystemPromptTemplate: cadenya.String("systemPromptTemplate"),
			},
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentGetParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentUpdateParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Metadata: shared.UpdateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.AgentSpecParam{
				VariationSelectionMode: cadenya.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified,
				Description:            cadenya.String("description"),
				EnableEpisodicMemory:   cadenya.Bool(true),
				EpisodicMemoryTtl:      cadenya.Int(0),
				OutputDefinition: map[string]any{
					"foo": "bar",
				},
				SystemPromptDataSchema: map[string]any{
					"foo": "bar",
				},
				WebhookEventsURL: cadenya.String("webhookEventsUrl"),
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
		WorkspaceID:            cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
		Cursor:                 cadenya.String("cursor"),
		IncludeInfo:            cadenya.Bool(true),
		Labels:                 cadenya.String("labels"),
		Limit:                  cadenya.Int(0),
		Prefix:                 cadenya.String("prefix"),
		Query:                  cadenya.String("query"),
		SortOrder:              cadenya.String("sortOrder"),
		State:                  cadenya.AgentListParamsStateStateUnspecified,
		VariationSelectionMode: cadenya.AgentListParamsVariationSelectionModeVariationSelectionModeUnspecified,
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
	err := client.Agents.Delete(
		context.TODO(),
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentDeleteParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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

func TestAgentArchive(t *testing.T) {
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
	_, err := client.Agents.Archive(
		context.TODO(),
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentArchiveParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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

func TestAgentPublish(t *testing.T) {
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
	_, err := client.Agents.Publish(
		context.TODO(),
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentPublishParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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

func TestAgentUnarchive(t *testing.T) {
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
	_, err := client.Agents.Unarchive(
		context.TODO(),
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentUnarchiveParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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

func TestAgentUnpublish(t *testing.T) {
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
	_, err := client.Agents.Unpublish(
		context.TODO(),
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentUnpublishParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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
