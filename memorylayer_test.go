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

func TestMemoryLayerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.MemoryLayers.New(context.TODO(), cadenya.MemoryLayerNewParams{
		WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
		Metadata: shared.CreateResourceMetadataParam{
			Name:       "name",
			ExternalID: cadenya.String("externalId"),
			Labels: map[string]string{
				"foo": "string",
			},
		},
		Spec: cadenya.MemoryLayerSpecParam{
			Type:        cadenya.MemoryLayerSpecTypeMemoryLayerTypeUnspecified,
			Description: cadenya.String("description"),
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

func TestMemoryLayerGet(t *testing.T) {
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
	_, err := client.MemoryLayers.Get(
		context.TODO(),
		"memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH",
		cadenya.MemoryLayerGetParams{
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

func TestMemoryLayerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.MemoryLayers.Update(
		context.TODO(),
		"memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH",
		cadenya.MemoryLayerUpdateParams{
			WorkspaceID: cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
			Metadata: shared.UpdateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.MemoryLayerSpecParam{
				Type:        cadenya.MemoryLayerSpecTypeMemoryLayerTypeUnspecified,
				Description: cadenya.String("description"),
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

func TestMemoryLayerListWithOptionalParams(t *testing.T) {
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
	_, err := client.MemoryLayers.List(context.TODO(), cadenya.MemoryLayerListParams{
		WorkspaceID:       cadenya.String("workspace_01HXKD2E5NQM3T9AYWCF133E3Q"),
		AgentID:           cadenya.String("agent_01HXKD2E5NQM3T9AYWCFMGWT9Y"),
		Cursor:            cadenya.String("cursor"),
		EpisodicKeyPrefix: cadenya.String("episodicKeyPrefix"),
		IncludeInfo:       cadenya.Bool(true),
		Labels:            cadenya.String("labels"),
		Limit:             cadenya.Int(0),
		Prefix:            cadenya.String("prefix"),
		Query:             cadenya.String("query"),
		SortOrder:         cadenya.String("sortOrder"),
		Type:              cadenya.MemoryLayerListParamsTypeMemoryLayerTypeUnspecified,
	})
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryLayerDelete(t *testing.T) {
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
	err := client.MemoryLayers.Delete(
		context.TODO(),
		"memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH",
		cadenya.MemoryLayerDeleteParams{
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
