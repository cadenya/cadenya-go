// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyasdkgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cadenya/cadenya-sdk-go"
	"github.com/cadenya/cadenya-sdk-go/internal/testutil"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/cadenya/cadenya-sdk-go/shared"
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.New(context.TODO(), gocadenyacomcadenyasdkgo.AgentNewParams{
		Metadata: gocadenyacomcadenyasdkgo.F(shared.CreateResourceMetadataParam{
			Name:       gocadenyacomcadenyasdkgo.F("name"),
			ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
			Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
				"foo": "string",
			}),
		}),
		Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentSpecParam{
			Status:                 gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentSpecStatusAgentStatusUnspecified),
			VariationSelectionMode: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
			Description:            gocadenyacomcadenyasdkgo.F("description"),
			WebhookEventsURL:       gocadenyacomcadenyasdkgo.F("webhookEventsUrl"),
		}),
		DefaultVariation: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentNewParamsDefaultVariation{
			Metadata: gocadenyacomcadenyasdkgo.F(shared.CreateResourceMetadataParam{
				Name:       gocadenyacomcadenyasdkgo.F("name"),
				ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
				Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentVariationSpecParam{
				AgentDocuments: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.AgentVariationSpecAgentDocumentParam{{
					DocumentID: gocadenyacomcadenyasdkgo.F("documentId"),
					DocumentMetadata: gocadenyacomcadenyasdkgo.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyasdkgo.F("name"),
						ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
						Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
					}),
					DocumentNamespaceID: gocadenyacomcadenyasdkgo.F("documentNamespaceId"),
					DocumentNamespaceMetadata: gocadenyacomcadenyasdkgo.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyasdkgo.F("name"),
						ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
						Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
				AgentTools: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.AgentVariationSpecAgentToolParam{{
					AgentID: gocadenyacomcadenyasdkgo.F("agentId"),
					AgentMetadata: gocadenyacomcadenyasdkgo.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyasdkgo.F("name"),
						ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
						Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
					}),
					ToolID: gocadenyacomcadenyasdkgo.F("toolId"),
					ToolMetadata: gocadenyacomcadenyasdkgo.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyasdkgo.F("name"),
						ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
						Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
					}),
					ToolSetID: gocadenyacomcadenyasdkgo.F("toolSetId"),
					ToolSetMetadata: gocadenyacomcadenyasdkgo.F(shared.ResourceMetadataParam{
						Name:       gocadenyacomcadenyasdkgo.F("name"),
						ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
						Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
							"foo": "string",
						}),
					}),
				}}),
				Constraints: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentVariationSpecConstraintsParam{
					MaxSubObjectives: gocadenyacomcadenyasdkgo.F(int64(0)),
					MaxToolCalls:     gocadenyacomcadenyasdkgo.F(int64(0)),
				}),
				Description:          gocadenyacomcadenyasdkgo.F("description"),
				EnableEpisodicMemory: gocadenyacomcadenyasdkgo.F(true),
				EpisodicMemoryTtl:    gocadenyacomcadenyasdkgo.F(int64(0)),
				ModelConfig: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentVariationSpecModelConfigParam{
					ModelID:     gocadenyacomcadenyasdkgo.F("modelId"),
					Temperature: gocadenyacomcadenyasdkgo.F(0.000000),
				}),
				Prompt: gocadenyacomcadenyasdkgo.F("prompt"),
				ToolSelection: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentVariationSpecToolSelectionParam{
					AssignedTools: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSelectionAssignedToolsParam{
						AllowDiscovery: gocadenyacomcadenyasdkgo.F(true),
					}),
					AutoDiscovery: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ToolSelectionAutoDiscoveryParam{
						Hints:    gocadenyacomcadenyasdkgo.F([]string{"string"}),
						MaxTools: gocadenyacomcadenyasdkgo.F(int64(0)),
					}),
				}),
				Weight: gocadenyacomcadenyasdkgo.F(int64(0)),
			}),
		}),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.Update(
		context.TODO(),
		"id",
		gocadenyacomcadenyasdkgo.AgentUpdateParams{
			Metadata: gocadenyacomcadenyasdkgo.F(shared.UpdateResourceMetadataParam{
				Name:       gocadenyacomcadenyasdkgo.F("name"),
				ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
				Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentSpecParam{
				Status:                 gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentSpecStatusAgentStatusUnspecified),
				VariationSelectionMode: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.AgentSpecVariationSelectionModeVariationSelectionModeUnspecified),
				Description:            gocadenyacomcadenyasdkgo.F("description"),
				WebhookEventsURL:       gocadenyacomcadenyasdkgo.F("webhookEventsUrl"),
			}),
			UpdateMask: gocadenyacomcadenyasdkgo.F("updateMask"),
		},
	)
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Agents.List(context.TODO(), gocadenyacomcadenyasdkgo.AgentListParams{
		Cursor:      gocadenyacomcadenyasdkgo.F("cursor"),
		IncludeInfo: gocadenyacomcadenyasdkgo.F(true),
		Limit:       gocadenyacomcadenyasdkgo.F(int64(0)),
		Prefix:      gocadenyacomcadenyasdkgo.F("prefix"),
		SortOrder:   gocadenyacomcadenyasdkgo.F("sortOrder"),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
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
	client := gocadenyacomcadenyasdkgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Agents.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
