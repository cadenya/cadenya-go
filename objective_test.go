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

func TestObjectiveNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.New(context.TODO(), gocadenyacomcadenyago.ObjectiveNewParams{
		AgentID: gocadenyacomcadenyago.F("agentId"),
		Data: gocadenyacomcadenyago.F(gocadenyacomcadenyago.ObjectiveDataParam{
			Data:           gocadenyacomcadenyago.F[any](map[string]interface{}{}),
			InitialMessage: gocadenyacomcadenyago.F("initialMessage"),
			Secrets: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.ObjectiveDataSecretParam{{
				Name:  gocadenyacomcadenyago.F("name"),
				Value: gocadenyacomcadenyago.F("value"),
			}}),
		}),
		Metadata: gocadenyacomcadenyago.F(shared.CreateOperationMetadataParam{
			ExternalID: gocadenyacomcadenyago.F("externalId"),
			Labels: gocadenyacomcadenyago.F(map[string]string{
				"foo": "string",
			}),
		}),
		VariationID: gocadenyacomcadenyago.F("variationId"),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Objectives.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Objectives.List(context.TODO(), gocadenyacomcadenyago.ObjectiveListParams{
		AgentID:           gocadenyacomcadenyago.F("agentId"),
		Cursor:            gocadenyacomcadenyago.F("cursor"),
		IncludeInfo:       gocadenyacomcadenyago.F(true),
		Limit:             gocadenyacomcadenyago.F(int64(0)),
		ParentObjectiveID: gocadenyacomcadenyago.F("parentObjectiveId"),
		ProfileID:         gocadenyacomcadenyago.F("profileId"),
		SortOrder:         gocadenyacomcadenyago.F("sortOrder"),
		State:             gocadenyacomcadenyago.F(gocadenyacomcadenyago.ObjectiveListParamsStateStateUnspecified),
	})
	if err != nil {
		var apierr *gocadenyacomcadenyago.Error
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
	client := gocadenyacomcadenyago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Objectives.Cancel(
		context.TODO(),
		"objectiveId",
		gocadenyacomcadenyago.ObjectiveCancelParams{
			Reason: gocadenyacomcadenyago.F("reason"),
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

func TestObjectiveContinueWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.Continue(
		context.TODO(),
		"objectiveId",
		gocadenyacomcadenyago.ObjectiveContinueParams{
			Enqueue: gocadenyacomcadenyago.F(true),
			Message: gocadenyacomcadenyago.F("message"),
			Secrets: gocadenyacomcadenyago.F([]gocadenyacomcadenyago.ObjectiveContinueParamsSecret{{
				Name:  gocadenyacomcadenyago.F("name"),
				Value: gocadenyacomcadenyago.F("value"),
			}}),
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

func TestObjectiveListContextWindowsWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.ListContextWindows(
		context.TODO(),
		"objectiveId",
		gocadenyacomcadenyago.ObjectiveListContextWindowsParams{
			Cursor:      gocadenyacomcadenyago.F("cursor"),
			IncludeInfo: gocadenyacomcadenyago.F(true),
			Limit:       gocadenyacomcadenyago.F(int64(0)),
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

func TestObjectiveListEventsWithOptionalParams(t *testing.T) {
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
	_, err := client.Objectives.ListEvents(
		context.TODO(),
		"objectiveId",
		gocadenyacomcadenyago.ObjectiveListEventsParams{
			Cursor:      gocadenyacomcadenyago.F("cursor"),
			IncludeInfo: gocadenyacomcadenyago.F(true),
			Limit:       gocadenyacomcadenyago.F(int64(0)),
			SortOrder:   gocadenyacomcadenyago.F("sortOrder"),
			WindowID:    gocadenyacomcadenyago.F("windowId"),
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
