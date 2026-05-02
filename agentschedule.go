// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cadenya/cadenya-go/internal/apijson"
	"github.com/cadenya/cadenya-go/internal/apiquery"
	"github.com/cadenya/cadenya-go/internal/param"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/packages/pagination"
	"github.com/cadenya/cadenya-go/shared"
)

// AgentScheduleService manages recurring schedules attached to agents. Schedules
// trigger objectives on a cadence defined by AgentScheduleSpec.Schedule. All
// operations are implicitly scoped to the workspace determined by the JWT token.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// AgentScheduleService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentScheduleService] method instead.
type AgentScheduleService struct {
	Options []option.RequestOption
}

// NewAgentScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentScheduleService(opts ...option.RequestOption) (r *AgentScheduleService) {
	r = &AgentScheduleService{}
	r.Options = opts
	return
}

// Creates a new schedule for an agent
func (r *AgentScheduleService) New(ctx context.Context, agentID string, body AgentScheduleNewParams, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/schedules", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a schedule by ID from an agent
func (r *AgentScheduleService) Get(ctx context.Context, agentID string, id string, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/schedules/%s", agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a schedule for an agent
func (r *AgentScheduleService) Update(ctx context.Context, agentID string, id string, body AgentScheduleUpdateParams, opts ...option.RequestOption) (res *AgentSchedule, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/schedules/%s", agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all schedules for an agent
func (r *AgentScheduleService) List(ctx context.Context, agentID string, query AgentScheduleListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[AgentSchedule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/schedules", agentID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *AgentScheduleService) ListAutoPaging(ctx context.Context, agentID string, query AgentScheduleListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[AgentSchedule] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, agentID, query, opts...))
}

// Deletes a schedule from an agent
func (r *AgentScheduleService) Delete(ctx context.Context, agentID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/agents/%s/schedules/%s", agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// AgentSchedule resource — a recurring trigger attached to an agent that creates
// objectives on its cadence.
type AgentSchedule struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// AgentScheduleSpec is the user-provided configuration for a schedule.
	Spec AgentScheduleSpec `json:"spec" api:"required"`
	// AgentScheduleInfo provides read-only runtime data about a schedule.
	Info AgentScheduleInfo `json:"info"`
	JSON agentScheduleJSON `json:"-"`
}

// agentScheduleJSON contains the JSON metadata for the struct [AgentSchedule]
type agentScheduleJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentScheduleJSON) RawJSON() string {
	return r.raw
}

// AgentScheduleInfo provides read-only runtime data about a schedule.
type AgentScheduleInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
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
	// schedule is PAUSED/ARCHIVED or has no future fire times.
	NextFireAt time.Time `json:"nextFireAt" format:"date-time"`
	// Lifetime count of objectives created by this schedule.
	TotalFires int64                 `json:"totalFires"`
	JSON       agentScheduleInfoJSON `json:"-"`
}

// agentScheduleInfoJSON contains the JSON metadata for the struct
// [AgentScheduleInfo]
type agentScheduleInfoJSON struct {
	CreatedBy       apijson.Field
	LastFireAt      apijson.Field
	LastObjectiveID apijson.Field
	LastSkippedAt   apijson.Field
	LastSkipReason  apijson.Field
	NextFireAt      apijson.Field
	TotalFires      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AgentScheduleInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentScheduleInfoJSON) RawJSON() string {
	return r.raw
}

// AgentScheduleSpec is the user-provided configuration for a schedule.
type AgentScheduleSpec struct {
	// The initial message passed to CreateObjective on each fire. Becomes the first
	// user message in the objective's chat history.
	InitialMessage string `json:"initialMessage" api:"required"`
	// Schedule defines WHEN the schedule fires. Temporal-style structured form: a list
	// of calendar rules (wall-clock) and/or interval rules (duration), OR'd together.
	// At least one rule is required.
	Schedule AgentScheduleSpecSchedule `json:"schedule" api:"required"`
	// Optional input data passed to the objective. If the agent has an
	// input_data_schema, this must satisfy it.
	Data interface{} `json:"data"`
	// What to do when the previous run is still in flight. Defaults to SKIP.
	OverlapPolicy AgentScheduleSpecOverlapPolicy `json:"overlapPolicy"`
	// Lifecycle. Defaults to ACTIVE on create when unspecified.
	Status AgentScheduleSpecStatus `json:"status"`
	// Optional explicit variation. When unset, the agent's variation_selection_mode
	// chooses per fire.
	VariationID string                `json:"variationId"`
	JSON        agentScheduleSpecJSON `json:"-"`
}

// agentScheduleSpecJSON contains the JSON metadata for the struct
// [AgentScheduleSpec]
type agentScheduleSpecJSON struct {
	InitialMessage apijson.Field
	Schedule       apijson.Field
	Data           apijson.Field
	OverlapPolicy  apijson.Field
	Status         apijson.Field
	VariationID    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AgentScheduleSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentScheduleSpecJSON) RawJSON() string {
	return r.raw
}

// What to do when the previous run is still in flight. Defaults to SKIP.
type AgentScheduleSpecOverlapPolicy string

const (
	AgentScheduleSpecOverlapPolicyOverlapPolicyUnspecified AgentScheduleSpecOverlapPolicy = "OVERLAP_POLICY_UNSPECIFIED"
	AgentScheduleSpecOverlapPolicyOverlapPolicyAllow       AgentScheduleSpecOverlapPolicy = "OVERLAP_POLICY_ALLOW"
	AgentScheduleSpecOverlapPolicyOverlapPolicySkip        AgentScheduleSpecOverlapPolicy = "OVERLAP_POLICY_SKIP"
)

func (r AgentScheduleSpecOverlapPolicy) IsKnown() bool {
	switch r {
	case AgentScheduleSpecOverlapPolicyOverlapPolicyUnspecified, AgentScheduleSpecOverlapPolicyOverlapPolicyAllow, AgentScheduleSpecOverlapPolicyOverlapPolicySkip:
		return true
	}
	return false
}

// Lifecycle. Defaults to ACTIVE on create when unspecified.
type AgentScheduleSpecStatus string

const (
	AgentScheduleSpecStatusAgentScheduleStatusUnspecified AgentScheduleSpecStatus = "AGENT_SCHEDULE_STATUS_UNSPECIFIED"
	AgentScheduleSpecStatusAgentScheduleStatusActive      AgentScheduleSpecStatus = "AGENT_SCHEDULE_STATUS_ACTIVE"
	AgentScheduleSpecStatusAgentScheduleStatusPaused      AgentScheduleSpecStatus = "AGENT_SCHEDULE_STATUS_PAUSED"
	AgentScheduleSpecStatusAgentScheduleStatusArchived    AgentScheduleSpecStatus = "AGENT_SCHEDULE_STATUS_ARCHIVED"
)

func (r AgentScheduleSpecStatus) IsKnown() bool {
	switch r {
	case AgentScheduleSpecStatusAgentScheduleStatusUnspecified, AgentScheduleSpecStatusAgentScheduleStatusActive, AgentScheduleSpecStatusAgentScheduleStatusPaused, AgentScheduleSpecStatusAgentScheduleStatusArchived:
		return true
	}
	return false
}

// AgentScheduleSpec is the user-provided configuration for a schedule.
type AgentScheduleSpecParam struct {
	// The initial message passed to CreateObjective on each fire. Becomes the first
	// user message in the objective's chat history.
	InitialMessage param.Field[string] `json:"initialMessage" api:"required"`
	// Schedule defines WHEN the schedule fires. Temporal-style structured form: a list
	// of calendar rules (wall-clock) and/or interval rules (duration), OR'd together.
	// At least one rule is required.
	Schedule param.Field[AgentScheduleSpecScheduleParam] `json:"schedule" api:"required"`
	// Optional input data passed to the objective. If the agent has an
	// input_data_schema, this must satisfy it.
	Data param.Field[interface{}] `json:"data"`
	// What to do when the previous run is still in flight. Defaults to SKIP.
	OverlapPolicy param.Field[AgentScheduleSpecOverlapPolicy] `json:"overlapPolicy"`
	// Lifecycle. Defaults to ACTIVE on create when unspecified.
	Status param.Field[AgentScheduleSpecStatus] `json:"status"`
	// Optional explicit variation. When unset, the agent's variation_selection_mode
	// chooses per fire.
	VariationID param.Field[string] `json:"variationId"`
}

func (r AgentScheduleSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Timezone string                        `json:"timezone"`
	JSON     agentScheduleSpecScheduleJSON `json:"-"`
}

// agentScheduleSpecScheduleJSON contains the JSON metadata for the struct
// [AgentScheduleSpecSchedule]
type agentScheduleSpecScheduleJSON struct {
	Calendars   apijson.Field
	Intervals   apijson.Field
	Timezone    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentScheduleSpecSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentScheduleSpecScheduleJSON) RawJSON() string {
	return r.raw
}

// Schedule defines WHEN the schedule fires. Temporal-style structured form: a list
// of calendar rules (wall-clock) and/or interval rules (duration), OR'd together.
// At least one rule is required.
type AgentScheduleSpecScheduleParam struct {
	// Wall-clock rules. May be empty if `intervals` is non-empty.
	Calendars param.Field[[]ScheduleCalendarParam] `json:"calendars"`
	// Duration-based rules. May be empty if `calendars` is non-empty.
	Intervals param.Field[[]ScheduleIntervalParam] `json:"intervals"`
	// IANA tz name (e.g. "America/New_York"). Required. Applies to calendars;
	// intervals fire on wall-clock cadence anchored in this zone.
	Timezone param.Field[string] `json:"timezone"`
}

func (r AgentScheduleSpecScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Calendar is a wall-clock rule. Empty field-list semantics:
//
//   - second/minute/hour: empty means [{start: 0}] (top of the unit)
//   - day_of_month/month/day_of_week: empty means "any value" Fire times = cartesian
//     product across all fields.
type ScheduleCalendar struct {
	Comment    string               `json:"comment"`
	DayOfMonth []ScheduleRange      `json:"dayOfMonth"`
	DayOfWeek  []ScheduleRange      `json:"dayOfWeek"`
	Hour       []ScheduleRange      `json:"hour"`
	Minute     []ScheduleRange      `json:"minute"`
	Month      []ScheduleRange      `json:"month"`
	Second     []ScheduleRange      `json:"second"`
	JSON       scheduleCalendarJSON `json:"-"`
}

// scheduleCalendarJSON contains the JSON metadata for the struct
// [ScheduleCalendar]
type scheduleCalendarJSON struct {
	Comment     apijson.Field
	DayOfMonth  apijson.Field
	DayOfWeek   apijson.Field
	Hour        apijson.Field
	Minute      apijson.Field
	Month       apijson.Field
	Second      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleCalendar) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleCalendarJSON) RawJSON() string {
	return r.raw
}

// Calendar is a wall-clock rule. Empty field-list semantics:
//
//   - second/minute/hour: empty means [{start: 0}] (top of the unit)
//   - day_of_month/month/day_of_week: empty means "any value" Fire times = cartesian
//     product across all fields.
type ScheduleCalendarParam struct {
	Comment    param.Field[string]               `json:"comment"`
	DayOfMonth param.Field[[]ScheduleRangeParam] `json:"dayOfMonth"`
	DayOfWeek  param.Field[[]ScheduleRangeParam] `json:"dayOfWeek"`
	Hour       param.Field[[]ScheduleRangeParam] `json:"hour"`
	Minute     param.Field[[]ScheduleRangeParam] `json:"minute"`
	Month      param.Field[[]ScheduleRangeParam] `json:"month"`
	Second     param.Field[[]ScheduleRangeParam] `json:"second"`
}

func (r ScheduleCalendarParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Interval is a duration-based rule. Fires every `every` from a stable anchor
// (workspace epoch), optionally phase-shifted by `offset`.
type ScheduleInterval struct {
	Every string `json:"every"`
	// Phase shift within `every`. Must be < `every` (enforced at runtime).
	Offset string               `json:"offset"`
	JSON   scheduleIntervalJSON `json:"-"`
}

// scheduleIntervalJSON contains the JSON metadata for the struct
// [ScheduleInterval]
type scheduleIntervalJSON struct {
	Every       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleInterval) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleIntervalJSON) RawJSON() string {
	return r.raw
}

// Interval is a duration-based rule. Fires every `every` from a stable anchor
// (workspace epoch), optionally phase-shifted by `offset`.
type ScheduleIntervalParam struct {
	Every param.Field[string] `json:"every"`
	// Phase shift within `every`. Must be < `every` (enforced at runtime).
	Offset param.Field[string] `json:"offset"`
}

func (r ScheduleIntervalParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Inclusive numeric range with optional step. {start: 9} → 9 {start: 9, end: 17} →
// 9..17 {start: 0, end: 59, step: 15} → 0,15,30,45 `end` defaults to `start`;
// `step` defaults to 1.
type ScheduleRange struct {
	End   int64             `json:"end"`
	Start int64             `json:"start"`
	Step  int64             `json:"step"`
	JSON  scheduleRangeJSON `json:"-"`
}

// scheduleRangeJSON contains the JSON metadata for the struct [ScheduleRange]
type scheduleRangeJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Step        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScheduleRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleRangeJSON) RawJSON() string {
	return r.raw
}

// Inclusive numeric range with optional step. {start: 9} → 9 {start: 9, end: 17} →
// 9..17 {start: 0, end: 59, step: 15} → 0,15,30,45 `end` defaults to `start`;
// `step` defaults to 1.
type ScheduleRangeParam struct {
	End   param.Field[int64] `json:"end"`
	Start param.Field[int64] `json:"start"`
	Step  param.Field[int64] `json:"step"`
}

func (r ScheduleRangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentScheduleNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	// AgentScheduleSpec is the user-provided configuration for a schedule.
	Spec param.Field[AgentScheduleSpecParam] `json:"spec" api:"required"`
}

func (r AgentScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentScheduleUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	// AgentScheduleSpec is the user-provided configuration for a schedule.
	Spec param.Field[AgentScheduleSpecParam] `json:"spec"`
	// Fields to update.
	UpdateMask param.Field[string] `json:"updateMask" format:"field-mask"`
}

func (r AgentScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentScheduleListParams struct {
	// Filter by bundle_key — return only resources owned by this bundle.
	BundleKey param.Field[string] `query:"bundleKey"`
	// Pagination cursor from previous response.
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit.
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return.
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix).
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query.
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time).
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [AgentScheduleListParams]'s query parameters as
// `url.Values`.
func (r AgentScheduleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
