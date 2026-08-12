package cadenya

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.cadenya.com/cadenya-go/internal/apijson"
	"go.cadenya.com/cadenya-go/internal/apiquery"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/packages/pagination"
	"go.cadenya.com/cadenya-go/packages/param"
	"go.cadenya.com/cadenya-go/packages/respjson"
	"go.cadenya.com/cadenya-go/shared"
	"net/http"
	"net/url"
	"slices"
	"time"
)

// Manage recurring schedules attached to agents. Schedules trigger objectives on a
// cadence defined by AgentScheduleSpec.Schedule.
//
// AgentScheduleService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentScheduleService] method instead.
type AgentScheduleService struct {
	options []option.RequestOption
}

// NewAgentScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentScheduleService(opts ...option.RequestOption) (r AgentScheduleService) {
	r = AgentScheduleService{}
	r.options = opts
	return
}

// Creates a new schedule for an agent
func (r *AgentScheduleService) New(ctx context.Context, agentID string, params AgentScheduleNewParams, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.WorkspaceID, precfg.WorkspaceID)
	if params.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/schedules", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a schedule by ID from an agent
func (r *AgentScheduleService) Get(ctx context.Context, agentID string, id string, query AgentScheduleGetParams, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&query.WorkspaceID, precfg.WorkspaceID)
	if query.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/schedules/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a schedule for an agent
func (r *AgentScheduleService) Update(ctx context.Context, agentID string, id string, params AgentScheduleUpdateParams, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.WorkspaceID, precfg.WorkspaceID)
	if params.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/schedules/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all schedules for an agent
func (r *AgentScheduleService) List(ctx context.Context, agentID string, params AgentScheduleListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[AgentSchedule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&params.WorkspaceID, precfg.WorkspaceID)
	if params.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/schedules", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID))
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all schedules for an agent
func (r *AgentScheduleService) ListAutoPaging(ctx context.Context, agentID string, params AgentScheduleListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[AgentSchedule] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, agentID, params, opts...))
}

// Deletes a schedule from an agent
func (r *AgentScheduleService) Delete(ctx context.Context, agentID string, id string, body AgentScheduleDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return err
	}
	requestconfig.UseDefaultParam(&body.WorkspaceID, precfg.WorkspaceID)
	if body.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/schedules/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Transitions a schedule to STATE_ARCHIVED and removes its underlying timer.
// Archiving is terminal: archived schedules never fire and cannot be reactivated;
// create a new schedule instead.
func (r *AgentScheduleService) Archive(ctx context.Context, agentID string, id string, body AgentScheduleArchiveParams, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.WorkspaceID, precfg.WorkspaceID)
	if body.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/schedules/%s:archive", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions a schedule to STATE_PAUSED. Paused schedules retain history but do
// not fire.
func (r *AgentScheduleService) Pause(ctx context.Context, agentID string, id string, body AgentSchedulePauseParams, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.WorkspaceID, precfg.WorkspaceID)
	if body.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/schedules/%s:pause", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions a paused schedule back to STATE_ACTIVE so it fires on its cadence
// again. Archived schedules cannot be resumed.
func (r *AgentScheduleService) Resume(ctx context.Context, agentID string, id string, body AgentScheduleResumeParams, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.WorkspaceID, precfg.WorkspaceID)
	if body.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/schedules/%s:resume", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// AgentSchedule resource — a recurring trigger attached to an agent that creates
// objectives on its cadence.
type AgentSchedule struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// AgentScheduleSpec is the user-provided configuration for a schedule.
	Spec AgentScheduleSpec `json:"spec" api:"required"`
	// The current lifecycle state of the schedule. Output only. Schedules are created
	// STATE_ACTIVE; use the :pause, :resume, and :archive actions to transition
	// between states.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ACTIVE", "STATE_PAUSED", "STATE_ARCHIVED".
	State AgentScheduleState `json:"state" api:"required"`
	// AgentScheduleInfo provides read-only runtime data about a schedule.
	Info AgentScheduleInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		State       respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSchedule) RawJSON() string { return r.JSON.raw }
func (r *AgentSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current lifecycle state of the schedule. Output only. Schedules are created
// STATE_ACTIVE; use the :pause, :resume, and :archive actions to transition
// between states.
type AgentScheduleState string

const (
	AgentScheduleStateStateUnspecified AgentScheduleState = "STATE_UNSPECIFIED"
	AgentScheduleStateStateActive      AgentScheduleState = "STATE_ACTIVE"
	AgentScheduleStateStatePaused      AgentScheduleState = "STATE_PAUSED"
	AgentScheduleStateStateArchived    AgentScheduleState = "STATE_ARCHIVED"
)

// AgentScheduleInfo provides read-only runtime data about a schedule.
type AgentScheduleInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// When the schedule last fired (regardless of objective outcome).
	LastFireAt time.Time `json:"lastFireAt" format:"date-time"`
	// ID of the most recent objective the schedule created.
	LastObjectiveID string `json:"lastObjectiveId"`
	// When the schedule most recently skipped a fire (SKIP policy + prior in flight).
	LastSkippedAt time.Time `json:"lastSkippedAt" format:"date-time"`
	// Reason for the most recent skip (e.g. "previous objective still running").
	LastSkipReason string `json:"lastSkipReason"`
	// When the schedule will next fire. Computed from the spec; absent when the
	// schedule is STATE_PAUSED/STATE_ARCHIVED or has no future fire times.
	NextFireAt time.Time `json:"nextFireAt" format:"date-time"`
	// Lifetime count of objectives created by this schedule.
	TotalFires int64 `json:"totalFires"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy       respjson.Field
		LastFireAt      respjson.Field
		LastObjectiveID respjson.Field
		LastSkippedAt   respjson.Field
		LastSkipReason  respjson.Field
		NextFireAt      respjson.Field
		TotalFires      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentScheduleInfo) RawJSON() string { return r.JSON.raw }
func (r *AgentScheduleInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentScheduleSpec is the user-provided configuration for a schedule.
type AgentScheduleSpec struct {
	// Schedule defines WHEN the schedule fires. Temporal-style structured form: a list
	// of calendar rules (wall-clock) and/or interval rules (duration), OR'd together.
	// At least one rule is required.
	Schedule AgentScheduleSpecSchedule `json:"schedule" api:"required"`
	// Optional explicit first user message passed to CreateObjective on each fire.
	// Becomes the first user message in the objective's chat history. When unset, the
	// fired objective defers to the selected variation's first_user_message_template.
	FirstUserMessage string `json:"firstUserMessage"`
	// Optional data rendered into the variation's first_user_message_template when
	// each fired objective is created. Separate from `system_prompt_data`, which
	// renders the system prompt template.
	FirstUserMessageData any `json:"firstUserMessageData"`
	// What to do when the previous run is still in flight. Defaults to SKIP.
	//
	// Any of "OVERLAP_POLICY_UNSPECIFIED", "OVERLAP_POLICY_ALLOW",
	// "OVERLAP_POLICY_SKIP".
	OverlapPolicy AgentScheduleSpecOverlapPolicy `json:"overlapPolicy"`
	// Optional data rendered into the variation's system_prompt_template when each
	// fired objective is created. If the agent has a system_prompt_data_schema, this
	// must satisfy it.
	SystemPromptData any `json:"systemPromptData"`
	// Optional explicit variation. When unset, the agent's variation_selection_mode
	// chooses per fire.
	VariationID string `json:"variationId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Schedule             respjson.Field
		FirstUserMessage     respjson.Field
		FirstUserMessageData respjson.Field
		OverlapPolicy        respjson.Field
		SystemPromptData     respjson.Field
		VariationID          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentScheduleSpec) RawJSON() string { return r.JSON.raw }
func (r *AgentScheduleSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentScheduleSpec to a AgentScheduleSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentScheduleSpecParam.Overrides()
func (r AgentScheduleSpec) ToParam() AgentScheduleSpecParam {
	return param.Override[AgentScheduleSpecParam](json.RawMessage(r.RawJSON()))
}

// What to do when the previous run is still in flight. Defaults to SKIP.
type AgentScheduleSpecOverlapPolicy string

const (
	AgentScheduleSpecOverlapPolicyOverlapPolicyUnspecified AgentScheduleSpecOverlapPolicy = "OVERLAP_POLICY_UNSPECIFIED"
	AgentScheduleSpecOverlapPolicyOverlapPolicyAllow       AgentScheduleSpecOverlapPolicy = "OVERLAP_POLICY_ALLOW"
	AgentScheduleSpecOverlapPolicyOverlapPolicySkip        AgentScheduleSpecOverlapPolicy = "OVERLAP_POLICY_SKIP"
)

// AgentScheduleSpec is the user-provided configuration for a schedule.
//
// The property Schedule is required.
type AgentScheduleSpecParam struct {
	// Schedule defines WHEN the schedule fires. Temporal-style structured form: a list
	// of calendar rules (wall-clock) and/or interval rules (duration), OR'd together.
	// At least one rule is required.
	Schedule AgentScheduleSpecScheduleParam `json:"schedule,omitzero" api:"required"`
	// Optional explicit first user message passed to CreateObjective on each fire.
	// Becomes the first user message in the objective's chat history. When unset, the
	// fired objective defers to the selected variation's first_user_message_template.
	FirstUserMessage param.Opt[string] `json:"firstUserMessage,omitzero"`
	// Optional explicit variation. When unset, the agent's variation_selection_mode
	// chooses per fire.
	VariationID param.Opt[string] `json:"variationId,omitzero"`
	// Optional data rendered into the variation's first_user_message_template when
	// each fired objective is created. Separate from `system_prompt_data`, which
	// renders the system prompt template.
	FirstUserMessageData any `json:"firstUserMessageData,omitzero"`
	// What to do when the previous run is still in flight. Defaults to SKIP.
	//
	// Any of "OVERLAP_POLICY_UNSPECIFIED", "OVERLAP_POLICY_ALLOW",
	// "OVERLAP_POLICY_SKIP".
	OverlapPolicy AgentScheduleSpecOverlapPolicy `json:"overlapPolicy,omitzero"`
	// Optional data rendered into the variation's system_prompt_template when each
	// fired objective is created. If the agent has a system_prompt_data_schema, this
	// must satisfy it.
	SystemPromptData any `json:"systemPromptData,omitzero"`
	paramObj
}

func (r AgentScheduleSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentScheduleSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentScheduleSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schedule defines WHEN the schedule fires. Temporal-style structured form: a list
// of calendar rules (wall-clock) and/or interval rules (duration), OR'd together.
// At least one rule is required.
type AgentScheduleSpecSchedule struct {
	// Wall-clock rules. May be empty if `intervals` is non-empty.
	Calendars []ScheduleCalendar `json:"calendars"`
	// Duration-based rules. May be empty if `calendars` is non-empty.
	Intervals []ScheduleInterval `json:"intervals"`
	// IANA tz name (e.g. "America/New_York"). Required. Applies to calendars;
	// intervals fire on wall-clock cadence anchored in this zone.
	Timezone string `json:"timezone"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Calendars   respjson.Field
		Intervals   respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentScheduleSpecSchedule) RawJSON() string { return r.JSON.raw }
func (r *AgentScheduleSpecSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentScheduleSpecSchedule to a
// AgentScheduleSpecScheduleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentScheduleSpecScheduleParam.Overrides()
func (r AgentScheduleSpecSchedule) ToParam() AgentScheduleSpecScheduleParam {
	return param.Override[AgentScheduleSpecScheduleParam](json.RawMessage(r.RawJSON()))
}

// Schedule defines WHEN the schedule fires. Temporal-style structured form: a list
// of calendar rules (wall-clock) and/or interval rules (duration), OR'd together.
// At least one rule is required.
type AgentScheduleSpecScheduleParam struct {
	// IANA tz name (e.g. "America/New_York"). Required. Applies to calendars;
	// intervals fire on wall-clock cadence anchored in this zone.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// Wall-clock rules. May be empty if `intervals` is non-empty.
	Calendars []ScheduleCalendarParam `json:"calendars,omitzero"`
	// Duration-based rules. May be empty if `calendars` is non-empty.
	Intervals []ScheduleIntervalParam `json:"intervals,omitzero"`
	paramObj
}

func (r AgentScheduleSpecScheduleParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentScheduleSpecScheduleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentScheduleSpecScheduleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Calendar is a wall-clock rule. Empty field-list semantics:
//
//   - second/minute/hour: empty means [{start: 0}] (top of the unit)
//   - day_of_month/month/day_of_week: empty means "any value" Fire times = cartesian
//     product across all fields.
type ScheduleCalendar struct {
	Comment    string          `json:"comment"`
	DayOfMonth []ScheduleRange `json:"dayOfMonth"`
	DayOfWeek  []ScheduleRange `json:"dayOfWeek"`
	Hour       []ScheduleRange `json:"hour"`
	Minute     []ScheduleRange `json:"minute"`
	Month      []ScheduleRange `json:"month"`
	Second     []ScheduleRange `json:"second"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comment     respjson.Field
		DayOfMonth  respjson.Field
		DayOfWeek   respjson.Field
		Hour        respjson.Field
		Minute      respjson.Field
		Month       respjson.Field
		Second      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleCalendar) RawJSON() string { return r.JSON.raw }
func (r *ScheduleCalendar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScheduleCalendar to a ScheduleCalendarParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScheduleCalendarParam.Overrides()
func (r ScheduleCalendar) ToParam() ScheduleCalendarParam {
	return param.Override[ScheduleCalendarParam](json.RawMessage(r.RawJSON()))
}

// Calendar is a wall-clock rule. Empty field-list semantics:
//
//   - second/minute/hour: empty means [{start: 0}] (top of the unit)
//   - day_of_month/month/day_of_week: empty means "any value" Fire times = cartesian
//     product across all fields.
type ScheduleCalendarParam struct {
	Comment    param.Opt[string]    `json:"comment,omitzero"`
	DayOfMonth []ScheduleRangeParam `json:"dayOfMonth,omitzero"`
	DayOfWeek  []ScheduleRangeParam `json:"dayOfWeek,omitzero"`
	Hour       []ScheduleRangeParam `json:"hour,omitzero"`
	Minute     []ScheduleRangeParam `json:"minute,omitzero"`
	Month      []ScheduleRangeParam `json:"month,omitzero"`
	Second     []ScheduleRangeParam `json:"second,omitzero"`
	paramObj
}

func (r ScheduleCalendarParam) MarshalJSON() (data []byte, err error) {
	type shadow ScheduleCalendarParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScheduleCalendarParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Interval is a duration-based rule. Fires every `every` from a stable anchor
// (workspace epoch), optionally phase-shifted by `offset`.
type ScheduleInterval struct {
	Every string `json:"every"`
	// Phase shift within `every`. Must be < `every` (enforced at runtime).
	Offset string `json:"offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Every       respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleInterval) RawJSON() string { return r.JSON.raw }
func (r *ScheduleInterval) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScheduleInterval to a ScheduleIntervalParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScheduleIntervalParam.Overrides()
func (r ScheduleInterval) ToParam() ScheduleIntervalParam {
	return param.Override[ScheduleIntervalParam](json.RawMessage(r.RawJSON()))
}

// Interval is a duration-based rule. Fires every `every` from a stable anchor
// (workspace epoch), optionally phase-shifted by `offset`.
type ScheduleIntervalParam struct {
	Every param.Opt[string] `json:"every,omitzero"`
	// Phase shift within `every`. Must be < `every` (enforced at runtime).
	Offset param.Opt[string] `json:"offset,omitzero"`
	paramObj
}

func (r ScheduleIntervalParam) MarshalJSON() (data []byte, err error) {
	type shadow ScheduleIntervalParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScheduleIntervalParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Inclusive numeric range with optional step. {start: 9} → 9 {start: 9, end: 17} →
// 9..17 {start: 0, end: 59, step: 15} → 0,15,30,45 `end` defaults to `start`;
// `step` defaults to 1.
type ScheduleRange struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
	Step  int64 `json:"step"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		Step        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleRange) RawJSON() string { return r.JSON.raw }
func (r *ScheduleRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScheduleRange to a ScheduleRangeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScheduleRangeParam.Overrides()
func (r ScheduleRange) ToParam() ScheduleRangeParam {
	return param.Override[ScheduleRangeParam](json.RawMessage(r.RawJSON()))
}

// Inclusive numeric range with optional step. {start: 9} → 9 {start: 9, end: 17} →
// 9..17 {start: 0, end: 59, step: 15} → 0,15,30,45 `end` defaults to `start`;
// `step` defaults to 1.
type ScheduleRangeParam struct {
	End   param.Opt[int64] `json:"end,omitzero"`
	Start param.Opt[int64] `json:"start,omitzero"`
	Step  param.Opt[int64] `json:"step,omitzero"`
	paramObj
}

func (r ScheduleRangeParam) MarshalJSON() (data []byte, err error) {
	type shadow ScheduleRangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScheduleRangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentScheduleNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	// AgentScheduleSpec is the user-provided configuration for a schedule.
	Spec AgentScheduleSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r AgentScheduleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentScheduleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentScheduleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentScheduleGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AgentScheduleUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Fields to update.
	UpdateMask param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	// AgentScheduleSpec is the user-provided configuration for a schedule.
	Spec AgentScheduleSpecParam `json:"spec,omitzero"`
	paramObj
}

func (r AgentScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentScheduleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentScheduleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentScheduleListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, the `info` field on each returned schedule is populated. Requests
	// with this flag count more against your rate limit.
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter expression (query param: prefix).
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Free-form search query.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time).
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentScheduleListParams]'s query parameters as
// `url.Values`.
func (r AgentScheduleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentScheduleDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AgentScheduleArchiveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r AgentScheduleArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentScheduleArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentScheduleArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentSchedulePauseParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r AgentSchedulePauseParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentSchedulePauseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSchedulePauseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentScheduleResumeParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r AgentScheduleResumeParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentScheduleResumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentScheduleResumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
