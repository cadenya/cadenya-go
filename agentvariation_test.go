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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentVariationNewParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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
					ModelID:         "claude/opus-4.6",
					CachingEnabled:  cadenya.Bool(true),
					MaxOutputTokens: cadenya.Int(0),
					ReasoningEffort: cadenya.AgentVariationSpecModelConfigReasoningEffortReasoningEffortUnspecified,
					StopSequences:   []string{"string"},
					Temperature:     cadenya.Float(0),
					TopK:            cadenya.Int(0),
					TopP:            cadenya.Float(0),
				},
				ProgressiveDiscovery: cadenya.AgentVariationSpecProgressiveDiscoveryParam{
					Hints:    []string{"string"},
					MaxTools: cadenya.Int(0),
				},
				SystemPromptTemplate: cadenya.String("systemPromptTemplate"),
			},
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		"agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		cadenya.AgentVariationGetParams{
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		"agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		cadenya.AgentVariationUpdateParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Metadata: shared.UpdateResourceMetadataParam{
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
					ModelID:         "claude/opus-4.6",
					CachingEnabled:  cadenya.Bool(true),
					MaxOutputTokens: cadenya.Int(0),
					ReasoningEffort: cadenya.AgentVariationSpecModelConfigReasoningEffortReasoningEffortUnspecified,
					StopSequences:   []string{"string"},
					Temperature:     cadenya.Float(0),
					TopK:            cadenya.Int(0),
					TopP:            cadenya.Float(0),
				},
				ProgressiveDiscovery: cadenya.AgentVariationSpecProgressiveDiscoveryParam{
					Hints:    []string{"string"},
					MaxTools: cadenya.Int(0),
				},
				SystemPromptTemplate: cadenya.String("systemPromptTemplate"),
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		cadenya.AgentVariationListParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		"agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		cadenya.AgentVariationDeleteParams{
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		"agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		cadenya.AgentVariationAddAssignmentParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			OfToolID: &cadenya.AddAgentVariationAssignmentRequestToolIDParam{
				ToolID: "tool_01HXKD2E5NQM3T9AYWCFWVYY9K",
				Type:   cadenya.AddAgentVariationAssignmentRequestToolIDTypeToolID,
			},
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		"agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		cadenya.AgentVariationAddMemoryLayerParams{
			WorkspaceID:   cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			MemoryLayerID: "memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH",
			Position:      cadenya.Int(0),
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		"agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		"avt_01HXKD2E5NQM3T9AYWCFJE6K89",
		cadenya.AgentVariationRemoveAssignmentParams{
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		"agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		"avml_01HXKD2E5NQM3T9AYWCFX8AF59",
		cadenya.AgentVariationRemoveMemoryLayerParams{
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
		"agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		"agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		"avml_01HXKD2E5NQM3T9AYWCFX8AF59",
		cadenya.AgentVariationUpdateMemoryLayerParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Position:    cadenya.Int(0),
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
