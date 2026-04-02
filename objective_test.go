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

func TestObjectiveNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.New(context.TODO(), gocadenyacomcadenyasdkgo.ObjectiveNewParams{
		AgentID: gocadenyacomcadenyasdkgo.F("agentId"),
		Data: gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ObjectiveDataParam{
			Data:           gocadenyacomcadenyasdkgo.F[any](map[string]interface{}{}),
			InitialMessage: gocadenyacomcadenyasdkgo.F("initialMessage"),
			Secrets: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.ObjectiveDataSecretParam{{
				Name:  gocadenyacomcadenyasdkgo.F("name"),
				Value: gocadenyacomcadenyasdkgo.F("value"),
			}}),
		}),
		Metadata: gocadenyacomcadenyasdkgo.F(shared.CreateOperationMetadataParam{
			ExternalID: gocadenyacomcadenyasdkgo.F("externalId"),
			Labels: gocadenyacomcadenyasdkgo.F(map[string]string{
				"foo": "string",
			}),
		}),
		VariationID: gocadenyacomcadenyasdkgo.F("variationId"),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectiveGet(t *testing.T) {
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
	_, err := client.Objectives.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectiveListWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.List(context.TODO(), gocadenyacomcadenyasdkgo.ObjectiveListParams{
		AgentID:           gocadenyacomcadenyasdkgo.F("agentId"),
		Cursor:            gocadenyacomcadenyasdkgo.F("cursor"),
		IncludeInfo:       gocadenyacomcadenyasdkgo.F(true),
		Limit:             gocadenyacomcadenyasdkgo.F(int64(0)),
		ParentObjectiveID: gocadenyacomcadenyasdkgo.F("parentObjectiveId"),
		ProfileID:         gocadenyacomcadenyasdkgo.F("profileId"),
		SortOrder:         gocadenyacomcadenyasdkgo.F("sortOrder"),
		State:             gocadenyacomcadenyasdkgo.F(gocadenyacomcadenyasdkgo.ObjectiveListParamsStateStateUnspecified),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyasdkgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestObjectiveCancelWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.Cancel(
		context.TODO(),
		"objectiveId",
		gocadenyacomcadenyasdkgo.ObjectiveCancelParams{
			Reason: gocadenyacomcadenyasdkgo.F("reason"),
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

func TestObjectiveContinueWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.Continue(
		context.TODO(),
		"objectiveId",
		gocadenyacomcadenyasdkgo.ObjectiveContinueParams{
			Enqueue: gocadenyacomcadenyasdkgo.F(true),
			Message: gocadenyacomcadenyasdkgo.F("message"),
			Secrets: gocadenyacomcadenyasdkgo.F([]gocadenyacomcadenyasdkgo.ObjectiveContinueParamsSecret{{
				Name:  gocadenyacomcadenyasdkgo.F("name"),
				Value: gocadenyacomcadenyasdkgo.F("value"),
			}}),
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

func TestObjectiveListContextWindowsWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.ListContextWindows(
		context.TODO(),
		"objectiveId",
		gocadenyacomcadenyasdkgo.ObjectiveListContextWindowsParams{
			Cursor:      gocadenyacomcadenyasdkgo.F("cursor"),
			IncludeInfo: gocadenyacomcadenyasdkgo.F(true),
			Limit:       gocadenyacomcadenyasdkgo.F(int64(0)),
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

func TestObjectiveListEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.ListEvents(
		context.TODO(),
		"objectiveId",
		gocadenyacomcadenyasdkgo.ObjectiveListEventsParams{
			Cursor:      gocadenyacomcadenyasdkgo.F("cursor"),
			IncludeInfo: gocadenyacomcadenyasdkgo.F(true),
			Limit:       gocadenyacomcadenyasdkgo.F(int64(0)),
			SortOrder:   gocadenyacomcadenyasdkgo.F("sortOrder"),
			WindowID:    gocadenyacomcadenyasdkgo.F("windowId"),
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
