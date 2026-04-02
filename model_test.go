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
)

func TestModelGet(t *testing.T) {
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
	_, err := client.Models.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestModelListWithOptionalParams(t *testing.T) {
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
	_, err := client.Models.List(context.TODO(), gocadenyacomcadenyasdkgo.ModelListParams{
		Cursor:    gocadenyacomcadenyasdkgo.F("cursor"),
		Limit:     gocadenyacomcadenyasdkgo.F(int64(0)),
		Prefix:    gocadenyacomcadenyasdkgo.F("prefix"),
		SortOrder: gocadenyacomcadenyasdkgo.F("sortOrder"),
		Status:    gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ModelListParamsStatusModelStatusUnspecified),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestModelSetStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Models.SetStatus(
		context.TODO(),
		"id",
		gocadenyacomcadenyasdkgo.ModelSetStatusParams{
			Status: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ModelSetStatusParamsStatusModelStatusUnspecified),
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
