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
		WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
		Metadata: shared.CreateResourceMetadataParam{
			Name:       "name",
			ExternalID: cadenya.String("externalId"),
			Labels: map[string]string{
				"foo": "string",
			},
		},
		Spec: cadenya.AIProviderKeySpecParam{
			Config: cadenya.AIProviderKeySpecConfigUnionParam{
				OfOpenrouter: &cadenya.AIProviderConfigOpenrouterParam{
					Openrouter: cadenya.AIProviderConfigOpenrouterOpenrouterParam{
						Region: cadenya.String("region"),
					},
					Type: cadenya.AIProviderConfigOpenrouterTypeOpenrouter,
				},
			},
			Credentials: cadenya.AIProviderKeySpecCredentialsUnionParam{
				OfAPIKey: &cadenya.AIProviderCredentialAPIKeyParam{
					APIKey: cadenya.AIProviderCredentialAPIKeyAPIKeyParam{
						APIKey: cadenya.String("apiKey"),
					},
					Type: cadenya.AIProviderCredentialAPIKeyTypeAPIKey,
				},
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
		"aipk_01HXKD2E5NQM3T9AYWCFQ41VW3",
		cadenya.AIProviderKeyGetParams{
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
		"aipk_01HXKD2E5NQM3T9AYWCFQ41VW3",
		cadenya.AIProviderKeyUpdateParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Metadata: shared.UpdateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.AIProviderKeySpecParam{
				Config: cadenya.AIProviderKeySpecConfigUnionParam{
					OfOpenrouter: &cadenya.AIProviderConfigOpenrouterParam{
						Openrouter: cadenya.AIProviderConfigOpenrouterOpenrouterParam{
							Region: cadenya.String("region"),
						},
						Type: cadenya.AIProviderConfigOpenrouterTypeOpenrouter,
					},
				},
				Credentials: cadenya.AIProviderKeySpecCredentialsUnionParam{
					OfAPIKey: &cadenya.AIProviderCredentialAPIKeyParam{
						APIKey: cadenya.AIProviderCredentialAPIKeyAPIKeyParam{
							APIKey: cadenya.String("apiKey"),
						},
						Type: cadenya.AIProviderCredentialAPIKeyTypeAPIKey,
					},
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
		WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
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
		"aipk_01HXKD2E5NQM3T9AYWCFQ41VW3",
		cadenya.AIProviderKeyDeleteParams{
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
