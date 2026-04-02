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
				AgentDocuments: cadenya.F([]cadenya.AgentVariationSpecAgentDocumentParam{{
					DocumentID: cadenya.F("documentId"),
					DocumentMetadata: cadenya.F(shared.ResourceMetadataParam{
						Name:       cadenya.F("name"),
						ExternalID: cadenya.F("externalId"),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
					}),
					DocumentNamespaceID: cadenya.F("documentNamespaceId"),
					DocumentNamespaceMetadata: cadenya.F(shared.ResourceMetadataParam{
						Name:       cadenya.F("name"),
						ExternalID: cadenya.F("externalId"),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
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
				AgentDocuments: cadenya.F([]cadenya.AgentVariationSpecAgentDocumentParam{{
					DocumentID: cadenya.F("documentId"),
					DocumentMetadata: cadenya.F(shared.ResourceMetadataParam{
						Name:       cadenya.F("name"),
						ExternalID: cadenya.F("externalId"),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
					}),
					DocumentNamespaceID: cadenya.F("documentNamespaceId"),
					DocumentNamespaceMetadata: cadenya.F(shared.ResourceMetadataParam{
						Name:       cadenya.F("name"),
						ExternalID: cadenya.F("externalId"),
						Labels: cadenya.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
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
