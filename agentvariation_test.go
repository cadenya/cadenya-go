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

func TestAgentVariationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Variations.New(
		context.TODO(),
		"agentId",
		gocadenyacomcadenyago.AgentVariationNewParams{
			Metadata: gocadenyacomcadenyago.F(shared.CreateResourceMetadataParam{
				Name:       gocadenyacomcadenyago.F("name"),
				ExternalID: gocadenyacomcadenyago.F("externalId"),
				Labels: gocadenyacomcadenyago.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentVariationSpecParam{
				AgentDocuments: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.AgentVariationSpecAgentDocumentParam{{
					DocumentID: gocadenyacomcadenyago.F("documentId"),
					DocumentMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
					DocumentNamespaceID: gocadenyacomcadenyago.F("documentNamespaceId"),
					DocumentNamespaceMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
				AgentTools: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.AgentVariationSpecAgentToolParam{{
					AgentID: gocadenyacomcadenyago.F("agentId"),
					AgentMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
					ToolID: gocadenyacomcadenyago.F("toolId"),
					ToolMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
					ToolSetID: gocadenyacomcadenyago.F("toolSetId"),
					ToolSetMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
				Constraints: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentVariationSpecConstraintsParam{
					MaxSubObjectives: gocadenyacomcadenyago.F(int64(0)),
					MaxToolCalls:     gocadenyacomcadenyago.F(int64(0)),
				}),
				Description:          gocadenyacomcadenyago.F("description"),
				EnableEpisodicMemory: gocadenyacomcadenyago.F(true),
				EpisodicMemoryTtl:    gocadenyacomcadenyago.F(int64(0)),
				ModelConfig: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentVariationSpecModelConfigParam{
					ModelID:     gocadenyacomcadenyago.F("modelId"),
					Temperature: gocadenyacomcadenyago.F(0.000000),
				}),
				Prompt: gocadenyacomcadenyago.F("prompt"),
				ToolSelection: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentVariationSpecToolSelectionParam{
					AssignedTools: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSelectionAssignedToolsParam{
						AllowDiscovery: gocadenyacomcadenyago.F(true),
					}),
					AutoDiscovery: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSelectionAutoDiscoveryParam{
						Hints:    gocadenyacomcadenyago.F([]string{"string"}),
						MaxTools: gocadenyacomcadenyago.F(int64(0)),
					}),
				}),
				Weight: gocadenyacomcadenyago.F(int64(0)),
			}),
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

func TestAgentVariationGet(t *testing.T) {
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
	_, err := client.Agents.Variations.Get(
		context.TODO(),
		"agentId",
		"id",
	)
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Variations.Update(
		context.TODO(),
		"agentId",
		"id",
		gocadenyacomcadenyago.AgentVariationUpdateParams{
			Metadata: gocadenyacomcadenyago.F(shared.UpdateResourceMetadataParam{
				Name:       gocadenyacomcadenyago.F("name"),
				ExternalID: gocadenyacomcadenyago.F("externalId"),
				Labels: gocadenyacomcadenyago.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentVariationSpecParam{
				AgentDocuments: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.AgentVariationSpecAgentDocumentParam{{
					DocumentID: gocadenyacomcadenyago.F("documentId"),
					DocumentMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
					DocumentNamespaceID: gocadenyacomcadenyago.F("documentNamespaceId"),
					DocumentNamespaceMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
				AgentTools: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.AgentVariationSpecAgentToolParam{{
					AgentID: gocadenyacomcadenyago.F("agentId"),
					AgentMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
					ToolID: gocadenyacomcadenyago.F("toolId"),
					ToolMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
					ToolSetID: gocadenyacomcadenyago.F("toolSetId"),
					ToolSetMetadata: gocadenyacomcadenyago.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyago.F("name"),
						ExternalID: gocadenyacomcadenyago.F("externalId"),
						Labels: gocadenyacomcadenyago.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
				Constraints: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentVariationSpecConstraintsParam{
					MaxSubObjectives: gocadenyacomcadenyago.F(int64(0)),
					MaxToolCalls:     gocadenyacomcadenyago.F(int64(0)),
				}),
				Description:          gocadenyacomcadenyago.F("description"),
				EnableEpisodicMemory: gocadenyacomcadenyago.F(true),
				EpisodicMemoryTtl:    gocadenyacomcadenyago.F(int64(0)),
				ModelConfig: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentVariationSpecModelConfigParam{
					ModelID:     gocadenyacomcadenyago.F("modelId"),
					Temperature: gocadenyacomcadenyago.F(0.000000),
				}),
				Prompt: gocadenyacomcadenyago.F("prompt"),
				ToolSelection: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentVariationSpecToolSelectionParam{
					AssignedTools: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSelectionAssignedToolsParam{
						AllowDiscovery: gocadenyacomcadenyago.F(true),
					}),
					AutoDiscovery: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ToolSelectionAutoDiscoveryParam{
						Hints:    gocadenyacomcadenyago.F([]string{"string"}),
						MaxTools: gocadenyacomcadenyago.F(int64(0)),
					}),
				}),
				Weight: gocadenyacomcadenyago.F(int64(0)),
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

func TestAgentVariationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Variations.List(
		context.TODO(),
		"agentId",
		gocadenyacomcadenyago.AgentVariationListParams{
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

func TestAgentVariationDelete(t *testing.T) {
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
	err := client.Agents.Variations.Delete(
		context.TODO(),
		"agentId",
		"id",
	)
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
