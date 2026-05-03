// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

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

func TestAgentScheduleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Schedules.New(
		context.TODO(),
		"workspaceId",
		"agentId",
		cadenya.AgentScheduleNewParams{
			Metadata: cadenya.F(shared.CreateResourceMetadataParam{
				Name:       cadenya.F("name"),
				BundleKey:  cadenya.F("bundleKey"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.AgentScheduleSpecParam{
				InitialMessage: cadenya.F("initialMessage"),
				Schedule: cadenya.F(cadenya.AgentScheduleSpecScheduleParam{
					Calendars: cadenya.F([]cadenya.ScheduleCalendarParam{{
						Comment: cadenya.F("comment"),
						DayOfMonth: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						DayOfWeek: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						Hour: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						Minute: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						Month: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						Second: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
					}}),
					Intervals: cadenya.F([]cadenya.ScheduleIntervalParam{{
						Every:  cadenya.F("-160513s"),
						Offset: cadenya.F("-160513s"),
					}}),
					Timezone: cadenya.F("timezone"),
				}),
				Data:          cadenya.F[any](map[string]interface{}{}),
				OverlapPolicy: cadenya.F(cadenya.AgentScheduleSpecOverlapPolicyOverlapPolicyUnspecified),
				Status:        cadenya.F(cadenya.AgentScheduleSpecStatusAgentScheduleStatusUnspecified),
				VariationID:   cadenya.F("variationId"),
			}),
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

func TestAgentScheduleGet(t *testing.T) {
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
	_, err := client.Agents.Schedules.Get(
		context.TODO(),
		"workspaceId",
		"agentId",
		"id",
	)
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentScheduleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Schedules.Update(
		context.TODO(),
		"workspaceId",
		"agentId",
		"id",
		cadenya.AgentScheduleUpdateParams{
			Metadata: cadenya.F(shared.UpdateResourceMetadataParam{
				Name:       cadenya.F("name"),
				BundleKey:  cadenya.F("bundleKey"),
				ExternalID: cadenya.F("externalId"),
				Labels: cadenya.F(map[string]string{
					"foo": "string",
				}),
			}),
			Spec: cadenya.F(cadenya.AgentScheduleSpecParam{
				InitialMessage: cadenya.F("initialMessage"),
				Schedule: cadenya.F(cadenya.AgentScheduleSpecScheduleParam{
					Calendars: cadenya.F([]cadenya.ScheduleCalendarParam{{
						Comment: cadenya.F("comment"),
						DayOfMonth: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						DayOfWeek: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						Hour: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						Minute: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						Month: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
						Second: cadenya.F([]cadenya.ScheduleRangeParam{{
							End:   cadenya.F(int64(0)),
							Start: cadenya.F(int64(0)),
							Step:  cadenya.F(int64(0)),
						}}),
					}}),
					Intervals: cadenya.F([]cadenya.ScheduleIntervalParam{{
						Every:  cadenya.F("-160513s"),
						Offset: cadenya.F("-160513s"),
					}}),
					Timezone: cadenya.F("timezone"),
				}),
				Data:          cadenya.F[any](map[string]interface{}{}),
				OverlapPolicy: cadenya.F(cadenya.AgentScheduleSpecOverlapPolicyOverlapPolicyUnspecified),
				Status:        cadenya.F(cadenya.AgentScheduleSpecStatusAgentScheduleStatusUnspecified),
				VariationID:   cadenya.F("variationId"),
			}),
			UpdateMask: cadenya.F("updateMask"),
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

func TestAgentScheduleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Schedules.List(
		context.TODO(),
		"workspaceId",
		"agentId",
		cadenya.AgentScheduleListParams{
			BundleKey:   cadenya.F("bundleKey"),
			Cursor:      cadenya.F("cursor"),
			IncludeInfo: cadenya.F(true),
			Limit:       cadenya.F(int64(0)),
			Prefix:      cadenya.F("prefix"),
			Query:       cadenya.F("query"),
			SortOrder:   cadenya.F("sortOrder"),
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

func TestAgentScheduleDelete(t *testing.T) {
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
	err := client.Agents.Schedules.Delete(
		context.TODO(),
		"workspaceId",
		"agentId",
		"id",
	)
	if err != nil {
		var apierr *cadenya.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
