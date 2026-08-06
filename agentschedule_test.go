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
		"agentId",
		cadenya.AgentScheduleNewParams{
			WorkspaceID: cadenya.String("workspaceId"),
			Metadata: shared.CreateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.AgentScheduleSpecParam{
				Schedule: cadenya.AgentScheduleSpecScheduleParam{
					Calendars: []cadenya.ScheduleCalendarParam{{
						Comment: cadenya.String("comment"),
						DayOfMonth: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						DayOfWeek: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						Hour: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						Minute: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						Month: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						Second: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
					}},
					Intervals: []cadenya.ScheduleIntervalParam{{
						Every:  cadenya.String("-160513s"),
						Offset: cadenya.String("-160513s"),
					}},
					Timezone: cadenya.String("timezone"),
				},
				FirstUserMessage:     cadenya.String("firstUserMessage"),
				FirstUserMessageData: map[string]any{},
				OverlapPolicy:        cadenya.AgentScheduleSpecOverlapPolicyOverlapPolicyUnspecified,
				SystemPromptData:     map[string]any{},
				VariationID:          cadenya.String("agentvar_01HXKD2E5NQM3T9AYWCF32BSPP"),
			},
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
		"agentId",
		"id",
		cadenya.AgentScheduleGetParams{
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
		"agentId",
		"id",
		cadenya.AgentScheduleUpdateParams{
			WorkspaceID: cadenya.String("workspaceId"),
			Metadata: shared.UpdateResourceMetadataParam{
				Name:       "name",
				ExternalID: cadenya.String("externalId"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: cadenya.AgentScheduleSpecParam{
				Schedule: cadenya.AgentScheduleSpecScheduleParam{
					Calendars: []cadenya.ScheduleCalendarParam{{
						Comment: cadenya.String("comment"),
						DayOfMonth: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						DayOfWeek: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						Hour: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						Minute: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						Month: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
						Second: []cadenya.ScheduleRangeParam{{
							End:   cadenya.Int(0),
							Start: cadenya.Int(0),
							Step:  cadenya.Int(0),
						}},
					}},
					Intervals: []cadenya.ScheduleIntervalParam{{
						Every:  cadenya.String("-160513s"),
						Offset: cadenya.String("-160513s"),
					}},
					Timezone: cadenya.String("timezone"),
				},
				FirstUserMessage:     cadenya.String("firstUserMessage"),
				FirstUserMessageData: map[string]any{},
				OverlapPolicy:        cadenya.AgentScheduleSpecOverlapPolicyOverlapPolicyUnspecified,
				SystemPromptData:     map[string]any{},
				VariationID:          cadenya.String("agentvar_01HXKD2E5NQM3T9AYWCF32BSPP"),
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
		"agentId",
		cadenya.AgentScheduleListParams{
			WorkspaceID: cadenya.String("workspaceId"),
			Cursor:      cadenya.String("cursor"),
			IncludeInfo: cadenya.Bool(true),
			Labels:      cadenya.String("labels"),
			Limit:       cadenya.Int(0),
			Prefix:      cadenya.String("prefix"),
			Query:       cadenya.String("query"),
			SortOrder:   cadenya.String("sortOrder"),
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
		"agentId",
		"id",
		cadenya.AgentScheduleDeleteParams{
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

func TestAgentScheduleArchive(t *testing.T) {
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
	_, err := client.Agents.Schedules.Archive(
		context.TODO(),
		"agentId",
		"id",
		cadenya.AgentScheduleArchiveParams{
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

func TestAgentSchedulePause(t *testing.T) {
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
	_, err := client.Agents.Schedules.Pause(
		context.TODO(),
		"agentId",
		"id",
		cadenya.AgentSchedulePauseParams{
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

func TestAgentScheduleResume(t *testing.T) {
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
	_, err := client.Agents.Schedules.Resume(
		context.TODO(),
		"agentId",
		"id",
		cadenya.AgentScheduleResumeParams{
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
