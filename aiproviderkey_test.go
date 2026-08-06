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

func TestAIProviderKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AIProviderKeys.New(context.TODO(), cadenya.AIProviderKeyNewParams{
		WorkspaceID: cadenya.String("workspaceId"),
		Metadata: shared.CreateResourceMetadataParam{
			Name:       "name",
			ExternalID: cadenya.String("externalId"),
			Labels: map[string]string{
				"foo": "string",
			},
		},
		Spec: cadenya.AIProviderKeySpecParam{
			Config: cadenya.AIProviderKeySpecConfigParam{
				OpenAI: cadenya.AIProviderKeySpecConfigOpenAIParam{
					OrganizationID: cadenya.String("organizationId"),
					ProjectID:      cadenya.String("projectId"),
				},
				OpenAICompatible: cadenya.AIProviderKeySpecConfigOpenAICompatibleParam{
					BaseURL: "baseUrl",
				},
				Openrouter: cadenya.AIProviderKeySpecConfigOpenrouterParam{
					Region: cadenya.String("region"),
				},
				Type: cadenya.String("type"),
			},
			Credentials: cadenya.AIProviderKeySpecCredentialsParam{
				APIKey: cadenya.AIProviderKeySpecCredentialsAPIKeyParam{
					APIKey: cadenya.String("apiKey"),
				},
				Headers: cadenya.AIProviderKeySpecCredentialsHeadersParam{
					Headers: map[string]string{
						"foo": "string",
					},
				},
				Type: cadenya.String("type"),
			},
			Provider: cadenya.AIProviderKeySpecProviderAIProviderUnspecified,
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

func TestAIProviderKeyGet(t *testing.T) {
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
	_, err := client.AIProviderKeys.Get(
		context.TODO(),
		"id",
		cadenya.AIProviderKeyGetParams{
			WorkspaceID: cadenya.String("workspaceId"),
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

func TestAIProviderKeyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.AIProviderKeys.Update(
		context.TODO(),
		"id",
		cadenya.AIProviderKeyUpdateParams{
			WorkspaceID: cadenya.String("workspaceId"),
			Metadata: shared.UpdateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.AIProviderKeySpecParam{
				Config: cadenya.AIProviderKeySpecConfigParam{
					OpenAI: cadenya.AIProviderKeySpecConfigOpenAIParam{
						OrganizationID: cadenya.String("organizationId"),
						ProjectID:      cadenya.String("projectId"),
					},
					OpenAICompatible: cadenya.AIProviderKeySpecConfigOpenAICompatibleParam{
						BaseURL: "baseUrl",
					},
					Openrouter: cadenya.AIProviderKeySpecConfigOpenrouterParam{
						Region: cadenya.String("region"),
					},
					Type: cadenya.String("type"),
				},
				Credentials: cadenya.AIProviderKeySpecCredentialsParam{
					APIKey: cadenya.AIProviderKeySpecCredentialsAPIKeyParam{
						APIKey: cadenya.String("apiKey"),
					},
					Headers: cadenya.AIProviderKeySpecCredentialsHeadersParam{
						Headers: map[string]string{
							"foo": "string",
						},
					},
					Type: cadenya.String("type"),
				},
				Provider: cadenya.AIProviderKeySpecProviderAIProviderUnspecified,
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

func TestAIProviderKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.AIProviderKeys.List(context.TODO(), cadenya.AIProviderKeyListParams{
		WorkspaceID: cadenya.String("workspaceId"),
		Cursor:      cadenya.String("cursor"),
		IncludeInfo: cadenya.Bool(true),
		Labels:      cadenya.String("labels"),
		Limit:       cadenya.Int(0),
		Prefix:      cadenya.String("prefix"),
		Promotional: cadenya.Bool(true),
		Query:       cadenya.String("query"),
		SortOrder:   cadenya.String("sortOrder"),
	})
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAIProviderKeyDelete(t *testing.T) {
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
	err := client.AIProviderKeys.Delete(
		context.TODO(),
		"id",
		cadenya.AIProviderKeyDeleteParams{
			WorkspaceID: cadenya.String("workspaceId"),
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
