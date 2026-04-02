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

func TestAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.New(context.TODO(), gocadenyacomcadenyago.AgentNewParams{
		Metadata: gocadenyacomcadenyago.F(shared.CreateResourceMetadataParam{
			Name:       gocadenyacomcadenyago.F("name"),
			ExternalID: gocadenyacomcadenyago.F("externalId"),
			Labels: gocadenyacomcadenyago.F(map[string]string{
				"foo": "string",
			}),
		}),
		Spec: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentSpecParam{
			Status:                 gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentSpecStatusAgentStatusUnspecified),
			VariationSelectionMode: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
			Description:            gocadenyacomcadenyago.F("description"),
			WebhookEventsURL:       gocadenyacomcadenyago.F("webhookEventsUrl"),
		}),
		DefaultVariation: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentNewParamsDefaultVariation{
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
		}),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Update(
		context.TODO(),
		"id",
		gocadenyacomcadenyago.AgentUpdateParams{
			Metadata: gocadenyacomcadenyago.F(shared.UpdateResourceMetadataParam{
				Name:       gocadenyacomcadenyago.F("name"),
				ExternalID: gocadenyacomcadenyago.F("externalId"),
				Labels: gocadenyacomcadenyago.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentSpecParam{
				Status:                 gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentSpecStatusAgentStatusUnspecified),
				VariationSelectionMode: gocadenyacomcadenyago.F(gocadenyacomcadenyago.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
				Description:            gocadenyacomcadenyago.F("description"),
				WebhookEventsURL:       gocadenyacomcadenyago.F("webhookEventsUrl"),
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

func TestAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.List(context.TODO(), gocadenyacomcadenyago.AgentListParams{
		Cursor:      gocadenyacomcadenyago.F("cursor"),
		IncludeInfo: gocadenyacomcadenyago.F(true),
		Limit:       gocadenyacomcadenyago.F(int64(0)),
		Prefix:      gocadenyacomcadenyago.F("prefix"),
		SortOrder:   gocadenyacomcadenyago.F("sortOrder"),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Agents.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
