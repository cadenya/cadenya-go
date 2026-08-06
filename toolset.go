// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// Manage tool sets and the tools they contain. Tool sets group related tools, and
// tools define specific capabilities available to agents.
//
// When a tool set is managed, only API key actors can modify its tools; human
// (profile) actors cannot.
//
// ToolSetService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolSetService] method instead.
type ToolSetService struct {
	options []option.RequestOption
	// Manage tool sets and the tools they contain. Tool sets group related tools, and
	// tools define specific capabilities available to agents.
	//
	// When a tool set is managed, only API key actors can modify its tools; human
	// (profile) actors cannot.
	Tools ToolSetToolService
	// Manage tool sets and the tools they contain. Tool sets group related tools, and
	// tools define specific capabilities available to agents.
	//
	// When a tool set is managed, only API key actors can modify its tools; human
	// (profile) actors cannot.
	Secrets ToolSetSecretService
}

// NewToolSetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolSetService(opts ...option.RequestOption) (r ToolSetService) {
	r = ToolSetService{}
	r.options = opts
	r.Tools = NewToolSetToolService(opts...)
	r.Secrets = NewToolSetSecretService(opts...)
	return
}

// Creates a new tool set in the workspace
func (r *ToolSetService) New(ctx context.Context, params ToolSetNewParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a tool set by ID from the workspace
func (r *ToolSetService) Get(ctx context.Context, id string, query ToolSetGetParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a tool set in the workspace
func (r *ToolSetService) Update(ctx context.Context, id string, params ToolSetUpdateParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all tool sets in the workspace
func (r *ToolSetService) List(ctx context.Context, params ToolSetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSet], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets", url.PathEscape(params.WorkspaceID.Value))
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

// Lists all tool sets in the workspace
func (r *ToolSetService) ListAutoPaging(ctx context.Context, params ToolSetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSet] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes a tool set in the workspace
func (r *ToolSetService) Delete(ctx context.Context, id string, body ToolSetDeleteParams, opts ...option.RequestOption) (err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Transitions a tool set to STATE_ARCHIVED. Syncing stops, the tool set is hidden
// from list results, its tools are no longer offered to objectives, and new
// variation assignments are rejected. Existing assignments are retained, and
// history is preserved — unlike delete, archiving works while the tool set is
// still assigned to agent variations.
func (r *ToolSetService) Archive(ctx context.Context, id string, body ToolSetArchiveParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s:archive", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the current OpenAPI specification JSON that has been consumed by the
// tool set. Only applicable to tool sets using the OpenAPI adapter.
func (r *ToolSetService) GetOpenAPISpec(ctx context.Context, toolSetID string, query ToolSetGetOpenAPISpecParams, opts ...option.RequestOption) (res *ToolSetGetOpenAPISpecResponse, err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/openapi_spec", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(toolSetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all events (including sync status) for a tool set
func (r *ToolSetService) ListEvents(ctx context.Context, toolSetID string, params ToolSetListEventsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSetEvent], err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/events", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID))
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

// Lists all events (including sync status) for a tool set
func (r *ToolSetService) ListEventsAutoPaging(ctx context.Context, toolSetID string, params ToolSetListEventsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSetEvent] {
	return pagination.NewCursorPaginationAutoPager(r.ListEvents(ctx, toolSetID, params, opts...))
}

// Lists the agent variations (with their parent agent) that have the tool set
// assigned. Pass tool_id to instead list variations with a direct assignment of
// that individual tool; variations that receive the tool implicitly through a
// whole-set assignment are not included in that filtered view.
func (r *ToolSetService) ListUsage(ctx context.Context, toolSetID string, params ToolSetListUsageParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSetUsage], err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/usage", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID))
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

// Lists the agent variations (with their parent agent) that have the tool set
// assigned. Pass tool_id to instead list variations with a direct assignment of
// that individual tool; variations that receive the tool implicitly through a
// whole-set assignment are not included in that filtered view.
func (r *ToolSetService) ListUsageAutoPaging(ctx context.Context, toolSetID string, params ToolSetListUsageParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSetUsage] {
	return pagination.NewCursorPaginationAutoPager(r.ListUsage(ctx, toolSetID, params, opts...))
}

// Transitions an archived tool set back to STATE_ACTIVE. Managed tool sets resume
// syncing on their next cycle and their tools become available to objectives
// again.
func (r *ToolSetService) Unarchive(ctx context.Context, id string, body ToolSetUnarchiveParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s:unarchive", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Approval filters that will automatically set the approval requirement on tools
// synced from an external source
type ApprovalRequirementFilter struct {
	Always bool `json:"always"`
	// Top-level filter with simple boolean logic (no nesting)
	Only ToolFilter `json:"only"`
	// The JSON name of the variant set in `requirement` (e.g. "always"). Required from
	// clients on writes, filled by the server on reads; drives the discriminated union
	// in the generated OpenAPI.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Always      respjson.Field
		Only        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalRequirementFilter) RawJSON() string { return r.JSON.raw }
func (r *ApprovalRequirementFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ApprovalRequirementFilter to a
// ApprovalRequirementFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ApprovalRequirementFilterParam.Overrides()
func (r ApprovalRequirementFilter) ToParam() ApprovalRequirementFilterParam {
	return param.Override[ApprovalRequirementFilterParam](json.RawMessage(r.RawJSON()))
}

// Approval filters that will automatically set the approval requirement on tools
// synced from an external source
type ApprovalRequirementFilterParam struct {
	Always param.Opt[bool] `json:"always,omitzero"`
	// The JSON name of the variant set in `requirement` (e.g. "always"). Required from
	// clients on writes, filled by the server on reads; drives the discriminated union
	// in the generated OpenAPI.
	Type param.Opt[string] `json:"type,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	Only ToolFilterParam `json:"only,omitzero"`
	paramObj
}

func (r ApprovalRequirementFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalRequirementFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalRequirementFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single attribute filter
type AttributeFilter struct {
	// Any of "ATTRIBUTE_UNSPECIFIED", "ATTRIBUTE_NAME", "ATTRIBUTE_TITLE",
	// "ATTRIBUTE_DESCRIPTION".
	Attribute AttributeFilterAttribute `json:"attribute" api:"required"`
	// String matching operations
	Matcher StringMatcher `json:"matcher"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attribute   respjson.Field
		Matcher     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributeFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AttributeFilter to a AttributeFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AttributeFilterParam.Overrides()
func (r AttributeFilter) ToParam() AttributeFilterParam {
	return param.Override[AttributeFilterParam](json.RawMessage(r.RawJSON()))
}

type AttributeFilterAttribute string

const (
	AttributeFilterAttributeAttributeUnspecified AttributeFilterAttribute = "ATTRIBUTE_UNSPECIFIED"
	AttributeFilterAttributeAttributeName        AttributeFilterAttribute = "ATTRIBUTE_NAME"
	AttributeFilterAttributeAttributeTitle       AttributeFilterAttribute = "ATTRIBUTE_TITLE"
	AttributeFilterAttributeAttributeDescription AttributeFilterAttribute = "ATTRIBUTE_DESCRIPTION"
)

// Single attribute filter
//
// The property Attribute is required.
type AttributeFilterParam struct {
	// Any of "ATTRIBUTE_UNSPECIFIED", "ATTRIBUTE_NAME", "ATTRIBUTE_TITLE",
	// "ATTRIBUTE_DESCRIPTION".
	Attribute AttributeFilterAttribute `json:"attribute,omitzero" api:"required"`
	// String matching operations
	Matcher StringMatcherParam `json:"matcher,omitzero"`
	paramObj
}

func (r AttributeFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow AttributeFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributeFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// String matching operations
type StringMatcher struct {
	CaseSensitive bool   `json:"caseSensitive"`
	Contains      string `json:"contains"`
	EndsWith      string `json:"endsWith"`
	Exact         string `json:"exact"`
	Regex         string `json:"regex"`
	StartsWith    string `json:"startsWith"`
	// The JSON name of the variant set in `match_type` (e.g. "startsWith"). Required
	// from clients on writes, filled by the server on reads; drives the discriminated
	// union in the generated OpenAPI.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CaseSensitive respjson.Field
		Contains      respjson.Field
		EndsWith      respjson.Field
		Exact         respjson.Field
		Regex         respjson.Field
		StartsWith    respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringMatcher) RawJSON() string { return r.JSON.raw }
func (r *StringMatcher) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringMatcher to a StringMatcherParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringMatcherParam.Overrides()
func (r StringMatcher) ToParam() StringMatcherParam {
	return param.Override[StringMatcherParam](json.RawMessage(r.RawJSON()))
}

// String matching operations
type StringMatcherParam struct {
	CaseSensitive param.Opt[bool]   `json:"caseSensitive,omitzero"`
	Contains      param.Opt[string] `json:"contains,omitzero"`
	EndsWith      param.Opt[string] `json:"endsWith,omitzero"`
	Exact         param.Opt[string] `json:"exact,omitzero"`
	Regex         param.Opt[string] `json:"regex,omitzero"`
	StartsWith    param.Opt[string] `json:"startsWith,omitzero"`
	// The JSON name of the variant set in `match_type` (e.g. "startsWith"). Required
	// from clients on writes, filled by the server on reads; drives the discriminated
	// union in the generated OpenAPI.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r StringMatcherParam) MarshalJSON() (data []byte, err error) {
	type shadow StringMatcherParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StringMatcherParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emitted when a tool set sync operation completes successfully.
type SyncCompleted struct {
	// Optional message with additional details.
	Message string `json:"message"`
	// Number of tools synced.
	ToolsSynced int64 `json:"toolsSynced"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ToolsSynced respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyncCompleted) RawJSON() string { return r.JSON.raw }
func (r *SyncCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emitted when a tool set sync operation fails.
type SyncFailed struct {
	// Indicates this is an error event.
	Error bool `json:"error"`
	// Optional error type/code for programmatic handling.
	ErrorType string `json:"errorType"`
	// Error message describing what went wrong.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ErrorType   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyncFailed) RawJSON() string { return r.JSON.raw }
func (r *SyncFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emitted when a tool set sync operation begins.
type SyncStarted struct {
	// Human-readable message describing the start of the sync.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyncStarted) RawJSON() string { return r.JSON.raw }
func (r *SyncStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-level filter with simple boolean logic (no nesting)
type ToolFilter struct {
	// Any of "OPERATOR_UNSPECIFIED", "OPERATOR_AND", "OPERATOR_OR".
	Operator ToolFilterOperator `json:"operator" api:"required"`
	Filters  []AttributeFilter  `json:"filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Filters     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolFilter) RawJSON() string { return r.JSON.raw }
func (r *ToolFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolFilter to a ToolFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolFilterParam.Overrides()
func (r ToolFilter) ToParam() ToolFilterParam {
	return param.Override[ToolFilterParam](json.RawMessage(r.RawJSON()))
}

type ToolFilterOperator string

const (
	ToolFilterOperatorOperatorUnspecified ToolFilterOperator = "OPERATOR_UNSPECIFIED"
	ToolFilterOperatorOperatorAnd         ToolFilterOperator = "OPERATOR_AND"
	ToolFilterOperatorOperatorOr          ToolFilterOperator = "OPERATOR_OR"
)

// Top-level filter with simple boolean logic (no nesting)
//
// The property Operator is required.
type ToolFilterParam struct {
	// Any of "OPERATOR_UNSPECIFIED", "OPERATOR_AND", "OPERATOR_OR".
	Operator ToolFilterOperator     `json:"operator,omitzero" api:"required"`
	Filters  []AttributeFilterParam `json:"filters,omitzero"`
	paramObj
}

func (r ToolFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSet struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     ToolSetSpec             `json:"spec" api:"required"`
	// The current lifecycle state of the tool set. Output only. Tool sets are created
	// STATE_ACTIVE; use the :archive and :unarchive actions to transition between
	// states.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ACTIVE", "STATE_ARCHIVED".
	State ToolSetState `json:"state" api:"required"`
	// Tool set information
	Info ToolSetInfo `json:"info"`
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
func (r ToolSet) RawJSON() string { return r.JSON.raw }
func (r *ToolSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current lifecycle state of the tool set. Output only. Tool sets are created
// STATE_ACTIVE; use the :archive and :unarchive actions to transition between
// states.
type ToolSetState string

const (
	ToolSetStateStateUnspecified ToolSetState = "STATE_UNSPECIFIED"
	ToolSetStateStateActive      ToolSetState = "STATE_ACTIVE"
	ToolSetStateStateArchived    ToolSetState = "STATE_ARCHIVED"
)

type ToolSetAdapter struct {
	// Bare tool sets define tools without an execution adapter. A bare tool call
	// doesn't fire anything: the objective's workflow pauses and waits for an external
	// API consumer to set the tool call's content (e.g. human-in-the-loop tools, or a
	// reverse harness that polls for pending tool calls, executes locally, and reports
	// results back via SetToolCallContent).
	Bare    ToolSetAdapterBare    `json:"bare"`
	HTTP    ToolSetAdapterHTTP    `json:"http"`
	MCP     ToolSetAdapterMCP     `json:"mcp"`
	OpenAPI ToolSetAdapterOpenAPI `json:"openapi"`
	// The JSON name of the variant set in `adapter` (e.g. "mcp"). Required from
	// clients on writes, filled by the server on reads; drives the discriminated union
	// in the generated OpenAPI.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bare        respjson.Field
		HTTP        respjson.Field
		MCP         respjson.Field
		OpenAPI     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapter) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapter to a ToolSetAdapterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterParam.Overrides()
func (r ToolSetAdapter) ToParam() ToolSetAdapterParam {
	return param.Override[ToolSetAdapterParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterParam struct {
	// The JSON name of the variant set in `adapter` (e.g. "mcp"). Required from
	// clients on writes, filled by the server on reads; drives the discriminated union
	// in the generated OpenAPI.
	Type param.Opt[string] `json:"type,omitzero"`
	// Bare tool sets define tools without an execution adapter. A bare tool call
	// doesn't fire anything: the objective's workflow pauses and waits for an external
	// API consumer to set the tool call's content (e.g. human-in-the-loop tools, or a
	// reverse harness that polls for pending tool calls, executes locally, and reports
	// results back via SetToolCallContent).
	Bare    ToolSetAdapterBareParam    `json:"bare,omitzero"`
	HTTP    ToolSetAdapterHTTPParam    `json:"http,omitzero"`
	MCP     ToolSetAdapterMCPParam     `json:"mcp,omitzero"`
	OpenAPI ToolSetAdapterOpenAPIParam `json:"openapi,omitzero"`
	paramObj
}

func (r ToolSetAdapterParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Bare tool sets define tools without an execution adapter. A bare tool call
// doesn't fire anything: the objective's workflow pauses and waits for an external
// API consumer to set the tool call's content (e.g. human-in-the-loop tools, or a
// reverse harness that polls for pending tool calls, executes locally, and reports
// results back via SetToolCallContent).
type ToolSetAdapterBare struct {
	// How long to wait for content to be set before the tool call errors. If unset,
	// the call waits indefinitely.
	ContentTimeout int64 `json:"contentTimeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentTimeout respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterBare) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterBare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterBare to a ToolSetAdapterBareParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterBareParam.Overrides()
func (r ToolSetAdapterBare) ToParam() ToolSetAdapterBareParam {
	return param.Override[ToolSetAdapterBareParam](json.RawMessage(r.RawJSON()))
}

// Bare tool sets define tools without an execution adapter. A bare tool call
// doesn't fire anything: the objective's workflow pauses and waits for an external
// API consumer to set the tool call's content (e.g. human-in-the-loop tools, or a
// reverse harness that polls for pending tool calls, executes locally, and reports
// results back via SetToolCallContent).
type ToolSetAdapterBareParam struct {
	// How long to wait for content to be set before the tool call errors. If unset,
	// the call waits indefinitely.
	ContentTimeout param.Opt[int64] `json:"contentTimeout,omitzero"`
	paramObj
}

func (r ToolSetAdapterBareParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterBareParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterBareParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterHTTP struct {
	BaseURL string            `json:"baseUrl"`
	Headers map[string]string `json:"headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL     respjson.Field
		Headers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterHTTP) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterHTTP to a ToolSetAdapterHTTPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterHTTPParam.Overrides()
func (r ToolSetAdapterHTTP) ToParam() ToolSetAdapterHTTPParam {
	return param.Override[ToolSetAdapterHTTPParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterHTTPParam struct {
	BaseURL param.Opt[string] `json:"baseUrl,omitzero"`
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r ToolSetAdapterHTTPParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterHTTPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterHTTPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterMCP struct {
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilter        `json:"excludeTools"`
	Headers      map[string]string `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilter `json:"includeTools"`
	// Defines behavior for just-in-time capable tool set adapters (IE: MCP).
	JustInTime ToolSetAdapterMCPJustInTime `json:"justInTime"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilter `json:"toolApprovals"`
	URL           string                    `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludeTools  respjson.Field
		Headers       respjson.Field
		IncludeTools  respjson.Field
		JustInTime    respjson.Field
		ToolApprovals respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterMCP) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterMCP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterMCP to a ToolSetAdapterMCPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterMCPParam.Overrides()
func (r ToolSetAdapterMCP) ToParam() ToolSetAdapterMCPParam {
	return param.Override[ToolSetAdapterMCPParam](json.RawMessage(r.RawJSON()))
}

// Defines behavior for just-in-time capable tool set adapters (IE: MCP).
type ToolSetAdapterMCPJustInTime struct {
	Enabled bool `json:"enabled"`
	// If set, an objective will automatically be failed if tools cannot be loaded in
	// the initial stages of an objective being created. Tools are loaded
	// asynchronously, so this setting is useful for ensuring that an objective
	// continued any further if tools are not available.
	FailObjectiveOnToolListError bool `json:"failObjectiveOnToolListError"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled                      respjson.Field
		FailObjectiveOnToolListError respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterMCPJustInTime) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterMCPJustInTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterMCPParam struct {
	URL param.Opt[string] `json:"url,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilterParam   `json:"excludeTools,omitzero"`
	Headers      map[string]string `json:"headers,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilterParam `json:"includeTools,omitzero"`
	// Defines behavior for just-in-time capable tool set adapters (IE: MCP).
	JustInTime ToolSetAdapterMCPJustInTimeParam `json:"justInTime,omitzero"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilterParam `json:"toolApprovals,omitzero"`
	paramObj
}

func (r ToolSetAdapterMCPParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterMCPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterMCPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines behavior for just-in-time capable tool set adapters (IE: MCP).
type ToolSetAdapterMCPJustInTimeParam struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// If set, an objective will automatically be failed if tools cannot be loaded in
	// the initial stages of an objective being created. Tools are loaded
	// asynchronously, so this setting is useful for ensuring that an objective
	// continued any further if tools are not available.
	FailObjectiveOnToolListError param.Opt[bool] `json:"failObjectiveOnToolListError,omitzero"`
	paramObj
}

func (r ToolSetAdapterMCPJustInTimeParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterMCPJustInTimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterMCPJustInTimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterOpenAPI struct {
	// Base URL for dispatching tool calls. If set, overrides the server resolved from
	// the spec's servers array.
	BaseURL string `json:"baseUrl"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilter `json:"excludeTools"`
	// Headers sent when fetching the spec from a URL and when dispatching tool calls.
	Headers map[string]string `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilter `json:"includeTools"`
	// Name of the server entry in the spec's servers array (OpenAPI 3.2 server.name
	// field). Used to select which server URL to dispatch to when base_url is not set.
	// If unset, the first server is used. Ignored when base_url is set.
	ServerName string `json:"serverName"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilter `json:"toolApprovals"`
	// The JSON name of the variant set in `source` (e.g. "url"). Required from clients
	// on writes, filled by the server on reads; drives the discriminated union in the
	// generated OpenAPI.
	Type string `json:"type"`
	// ID of a COMPLETE Upload containing the OpenAPI spec document.
	UploadID string `json:"uploadId"`
	// URL to fetch the OpenAPI spec from. Synced automatically every hour.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL       respjson.Field
		ExcludeTools  respjson.Field
		Headers       respjson.Field
		IncludeTools  respjson.Field
		ServerName    respjson.Field
		ToolApprovals respjson.Field
		Type          respjson.Field
		UploadID      respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterOpenAPI) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterOpenAPI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterOpenAPI to a ToolSetAdapterOpenAPIParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterOpenAPIParam.Overrides()
func (r ToolSetAdapterOpenAPI) ToParam() ToolSetAdapterOpenAPIParam {
	return param.Override[ToolSetAdapterOpenAPIParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterOpenAPIParam struct {
	// Base URL for dispatching tool calls. If set, overrides the server resolved from
	// the spec's servers array.
	BaseURL param.Opt[string] `json:"baseUrl,omitzero"`
	// Name of the server entry in the spec's servers array (OpenAPI 3.2 server.name
	// field). Used to select which server URL to dispatch to when base_url is not set.
	// If unset, the first server is used. Ignored when base_url is set.
	ServerName param.Opt[string] `json:"serverName,omitzero"`
	// The JSON name of the variant set in `source` (e.g. "url"). Required from clients
	// on writes, filled by the server on reads; drives the discriminated union in the
	// generated OpenAPI.
	Type param.Opt[string] `json:"type,omitzero"`
	// ID of a COMPLETE Upload containing the OpenAPI spec document.
	UploadID param.Opt[string] `json:"uploadId,omitzero"`
	// URL to fetch the OpenAPI spec from. Synced automatically every hour.
	URL param.Opt[string] `json:"url,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilterParam `json:"excludeTools,omitzero"`
	// Headers sent when fetching the spec from a URL and when dispatching tool calls.
	Headers map[string]string `json:"headers,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilterParam `json:"includeTools,omitzero"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilterParam `json:"toolApprovals,omitzero"`
	paramObj
}

func (r ToolSetAdapterOpenAPIParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterOpenAPIParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterOpenAPIParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single event in the tool set's operation timeline.
type ToolSetEvent struct {
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// Event payload for a tool set operation.
	Event ToolSetEventData `json:"event"`
	Info  ToolSetEventInfo `json:"info"`
	// The tool set this event is associated with.
	ToolSetID string `json:"toolSetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Event       respjson.Field
		Info        respjson.Field
		ToolSetID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetEvent) RawJSON() string { return r.JSON.raw }
func (r *ToolSetEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetEventInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolSet shared.ResourceMetadata `json:"toolSet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		ToolSet     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetEventInfo) RawJSON() string { return r.JSON.raw }
func (r *ToolSetEventInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Event payload for a tool set operation.
type ToolSetEventData struct {
	// Emitted when a tool set sync operation completes successfully.
	SyncCompleted SyncCompleted `json:"syncCompleted"`
	// Emitted when a tool set sync operation fails.
	SyncFailed SyncFailed `json:"syncFailed"`
	// Emitted when a tool set sync operation begins.
	SyncStarted SyncStarted `json:"syncStarted"`
	// The JSON name of the variant set in `data` (e.g. "syncStarted"). Filled by the
	// server; drives the discriminated union in the generated OpenAPI.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SyncCompleted respjson.Field
		SyncFailed    respjson.Field
		SyncStarted   respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetEventData) RawJSON() string { return r.JSON.raw }
func (r *ToolSetEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetInfo struct {
	AgentCount     int64 `json:"agentCount"`
	AvailableTools int64 `json:"availableTools"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy    Profile   `json:"createdBy"`
	LastSync     time.Time `json:"lastSync" format:"date-time"`
	OmittedTools int64     `json:"omittedTools"`
	ToolCount    int64     `json:"toolCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentCount     respjson.Field
		AvailableTools respjson.Field
		CreatedBy      respjson.Field
		LastSync       respjson.Field
		OmittedTools   respjson.Field
		ToolCount      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetInfo) RawJSON() string { return r.JSON.raw }
func (r *ToolSetInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetSpec struct {
	Adapter     ToolSetAdapter `json:"adapter"`
	Description string         `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Adapter     respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetSpec) RawJSON() string { return r.JSON.raw }
func (r *ToolSetSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetSpec to a ToolSetSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetSpecParam.Overrides()
func (r ToolSetSpec) ToParam() ToolSetSpecParam {
	return param.Override[ToolSetSpecParam](json.RawMessage(r.RawJSON()))
}

type ToolSetSpecParam struct {
	Description param.Opt[string]   `json:"description,omitzero"`
	Adapter     ToolSetAdapterParam `json:"adapter,omitzero"`
	paramObj
}

func (r ToolSetSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolSetUsage describes one agent variation that uses the tool set (or, when
// filtering by tool, an individual tool within it).
type ToolSetUsage struct {
	// When the assignment was created.
	AssignedAt time.Time `json:"assignedAt" api:"required" format:"date-time"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent shared.ResourceMetadata `json:"agent"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	AgentVariation shared.ResourceMetadata `json:"agentVariation"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Model shared.ResourceMetadata `json:"model"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssignedAt     respjson.Field
		Agent          respjson.Field
		AgentVariation respjson.Field
		Model          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetUsage) RawJSON() string { return r.JSON.raw }
func (r *ToolSetUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetGetOpenAPISpecResponse struct {
	// The consumed OpenAPI specification as a JSON string.
	Spec string `json:"spec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetGetOpenAPISpecResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolSetGetOpenAPISpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	Spec     ToolSetSpecParam                   `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r ToolSetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	UpdateMask  param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	Spec     ToolSetSpecParam                   `json:"spec,omitzero"`
	paramObj
}

func (r ToolSetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter expression (query param: prefix)
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Free-form search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// Filter by tool set lifecycle state. Defaults to STATE_ACTIVE when unspecified;
	// pass STATE_ARCHIVED to list archived tool sets.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ACTIVE", "STATE_ARCHIVED".
	State ToolSetListParamsState `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolSetListParams]'s query parameters as `url.Values`.
func (r ToolSetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by tool set lifecycle state. Defaults to STATE_ACTIVE when unspecified;
// pass STATE_ARCHIVED to list archived tool sets.
type ToolSetListParamsState string

const (
	ToolSetListParamsStateStateUnspecified ToolSetListParamsState = "STATE_UNSPECIFIED"
	ToolSetListParamsStateStateActive      ToolSetListParamsState = "STATE_ACTIVE"
	ToolSetListParamsStateStateArchived    ToolSetListParamsState = "STATE_ARCHIVED"
)

type ToolSetDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetArchiveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ToolSetArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetGetOpenAPISpecParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetListEventsParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolSetListEventsParams]'s query parameters as
// `url.Values`.
func (r ToolSetListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolSetListUsageParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order for results (asc or desc by assignment creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// When set, lists only variations with a direct assignment of this individual
	// tool. When unset, lists variations assigned the whole tool set. The tool must
	// belong to the tool set.
	ToolID param.Opt[string] `query:"toolId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolSetListUsageParams]'s query parameters as `url.Values`.
func (r ToolSetListUsageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolSetUnarchiveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ToolSetUnarchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetUnarchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetUnarchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
