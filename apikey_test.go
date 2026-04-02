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

func TestAPIKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.APIKeys.New(context.TODO(), gocadenyacomcadenyasdkgo.APIKeyNewParams{
		Metadata: gocadenyacomcadenyasdkgo.F(shared.CreateResourceMetadataParam{
			Name:       gocadenyacomcadenyasdkgo.F("name"),
			ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
			Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
				"foo": "string",
			}),
		}),
		Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.APIKeySpecParam{
			Description: gocadenyacomcadenyasdkgo.F("description"),
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

func TestAPIKeyGet(t *testing.T) {
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
	_, err := client.APIKeys.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIKeyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.APIKeys.Update(
		context.TODO(),
		"id",
		gocadenyacomcadenyasdkgo.APIKeyUpdateParams{
			Metadata: gocadenyacomcadenyasdkgo.F(shared.UpdateResourceMetadataParam{
				Name:       gocadenyacomcadenyasdkgo.F("name"),
				ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
				Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.APIKeySpecParam{
				Description: gocadenyacomcadenyasdkgo.F("description"),
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

func TestAPIKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIKeys.List(context.TODO(), gocadenyacomcadenyasdkgo.APIKeyListParams{
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

func TestAPIKeyDelete(t *testing.T) {
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
	err := client.APIKeys.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIKeyRotate(t *testing.T) {
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
	_, err := client.APIKeys.Rotate(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
