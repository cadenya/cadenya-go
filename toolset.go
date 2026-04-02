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

	"github.com/cadenya/cadenya-sdk-go/internal/apijson"
	"github.com/cadenya/cadenya-sdk-go/internal/apiquery"
	"github.com/cadenya/cadenya-sdk-go/internal/param"
	"github.com/cadenya/cadenya-sdk-go/internal/requestconfig"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/cadenya/cadenya-sdk-go/packages/pagination"
	"github.com/cadenya/cadenya-sdk-go/shared"
)

// ToolService manages tool sets and tools at the WORKSPACE level. Tool sets group
// related tools, and tools define specific capabilities for agents. All operations
// are implicitly scoped to the workspace determined by the JWT token.
//
// Note: When a ToolSet has managed=true, only API Key actors can modify its tools.
// Profile actors (humans) are restricted from modifying managed tool sets.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// ToolSetService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolSetService] method instead.
type ToolSetService struct {
	Options []option.RequestOption
	// ToolService manages tool sets and tools at the WORKSPACE level. Tool sets group
	// related tools, and tools define specific capabilities for agents. All operations
	// are implicitly scoped to the workspace determined by the JWT token.
	//
	// Note: When a ToolSet has managed=true, only API Key actors can modify its tools.
	// Profile actors (humans) are restricted from modifying managed tool sets.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	Tools *ToolSetToolService
}

// NewToolSetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolSetService(opts ...option.RequestOption) (r *ToolSetService) {
	r = &ToolSetService{}
	r.Options = opts
	r.Tools = NewToolSetToolService(opts...)
	return
}

// Creates a new tool set in the workspace
func (r *ToolSetService) New(ctx context.Context, body ToolSetNewParams, opts ...option.RequestOption) (res *ToolSet, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tool_sets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a tool set by ID from the workspace
func (r *ToolSetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ToolSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tool_sets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a tool set in the workspace
func (r *ToolSetService) Update(ctx context.Context, id string, body ToolSetUpdateParams, opts ...option.RequestOption) (res *ToolSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tool_sets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Lists all tool sets in the workspace
func (r *ToolSetService) List(ctx context.Context, query ToolSetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/tool_sets"
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

// Lists all tool sets in the workspace
func (r *ToolSetService) ListAutoPaging(ctx context.Context, query ToolSetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSet] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a tool set in the workspace
func (r *ToolSetService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/tool_sets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Lists all events (including sync status) for a tool set
func (r *ToolSetService) ListEvents(ctx context.Context, toolSetID string, query ToolSetListEventsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSetEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/tool_sets/%s/events", toolSetID)
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

// Lists all events (including sync status) for a tool set
func (r *ToolSetService) ListEventsAutoPaging(ctx context.Context, toolSetID string, query ToolSetListEventsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSetEvent] {
	return pagination.NewCursorPaginationAutoPager(r.ListEvents(ctx, toolSetID, query, opts...))
}

// Top-level filter with simple boolean logic (no nesting)
type McpToolFilter struct {
	Operator McpToolFilterOperator `json:"operator" api:"required"`
	Filters  []McpToolFilterFilter `json:"filters"`
	JSON     mcpToolFilterJSON     `json:"-"`
}

// mcpToolFilterJSON contains the JSON metadata for the struct [McpToolFilter]
type mcpToolFilterJSON struct {
	Operator    apijson.Field
	Filters     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpToolFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpToolFilterJSON) RawJSON() string {
	return r.raw
}

type McpToolFilterOperator string

const (
	McpToolFilterOperatorOperatorUnspecified McpToolFilterOperator = "OPERATOR_UNSPECIFIED"
	McpToolFilterOperatorOperatorAnd         McpToolFilterOperator = "OPERATOR_AND"
	McpToolFilterOperatorOperatorOr          McpToolFilterOperator = "OPERATOR_OR"
)

func (r McpToolFilterOperator) IsKnown() bool {
	switch r {
	case McpToolFilterOperatorOperatorUnspecified, McpToolFilterOperatorOperatorAnd, McpToolFilterOperatorOperatorOr:
		return true
	}
	return false
}

// Single attribute filter
type McpToolFilterFilter struct {
	Attribute McpToolFilterFiltersAttribute `json:"attribute" api:"required"`
	// String matching operations
	Matcher McpToolFilterFiltersMatcher `json:"matcher"`
	JSON    mcpToolFilterFilterJSON     `json:"-"`
}

// mcpToolFilterFilterJSON contains the JSON metadata for the struct
// [McpToolFilterFilter]
type mcpToolFilterFilterJSON struct {
	Attribute   apijson.Field
	Matcher     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McpToolFilterFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpToolFilterFilterJSON) RawJSON() string {
	return r.raw
}

type McpToolFilterFiltersAttribute string

const (
	McpToolFilterFiltersAttributeAttributeUnspecified McpToolFilterFiltersAttribute = "ATTRIBUTE_UNSPECIFIED"
	McpToolFilterFiltersAttributeAttributeName        McpToolFilterFiltersAttribute = "ATTRIBUTE_NAME"
	McpToolFilterFiltersAttributeAttributeTitle       McpToolFilterFiltersAttribute = "ATTRIBUTE_TITLE"
	McpToolFilterFiltersAttributeAttributeDescription McpToolFilterFiltersAttribute = "ATTRIBUTE_DESCRIPTION"
)

func (r McpToolFilterFiltersAttribute) IsKnown() bool {
	switch r {
	case McpToolFilterFiltersAttributeAttributeUnspecified, McpToolFilterFiltersAttributeAttributeName, McpToolFilterFiltersAttributeAttributeTitle, McpToolFilterFiltersAttributeAttributeDescription:
		return true
	}
	return false
}

// String matching operations
type McpToolFilterFiltersMatcher struct {
	CaseSensitive bool                            `json:"caseSensitive"`
	Contains      string                          `json:"contains"`
	EndsWith      string                          `json:"endsWith"`
	Exact         string                          `json:"exact"`
	Regex         string                          `json:"regex"`
	StartsWith    string                          `json:"startsWith"`
	JSON          mcpToolFilterFiltersMatcherJSON `json:"-"`
}

// mcpToolFilterFiltersMatcherJSON contains the JSON metadata for the struct
// [McpToolFilterFiltersMatcher]
type mcpToolFilterFiltersMatcherJSON struct {
	CaseSensitive apijson.Field
	Contains      apijson.Field
	EndsWith      apijson.Field
	Exact         apijson.Field
	Regex         apijson.Field
	StartsWith    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *McpToolFilterFiltersMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcpToolFilterFiltersMatcherJSON) RawJSON() string {
	return r.raw
}

// Top-level filter with simple boolean logic (no nesting)
type McpToolFilterParam struct {
	Operator param.Field[McpToolFilterOperator]      `json:"operator" api:"required"`
	Filters  param.Field[[]McpToolFilterFilterParam] `json:"filters"`
}

func (r McpToolFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Single attribute filter
type McpToolFilterFilterParam struct {
	Attribute param.Field[McpToolFilterFiltersAttribute] `json:"attribute" api:"required"`
	// String matching operations
	Matcher param.Field[McpToolFilterFiltersMatcherParam] `json:"matcher"`
}

func (r McpToolFilterFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String matching operations
type McpToolFilterFiltersMatcherParam struct {
	CaseSensitive param.Field[bool]   `json:"caseSensitive"`
	Contains      param.Field[string] `json:"contains"`
	EndsWith      param.Field[string] `json:"endsWith"`
	Exact         param.Field[string] `json:"exact"`
	Regex         param.Field[string] `json:"regex"`
	StartsWith    param.Field[string] `json:"startsWith"`
}

func (r McpToolFilterFiltersMatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SyncCompleted is emitted when a tool set sync operation completes successfully
type SyncCompleted struct {
	// Optional message with additional details
	Message string `json:"message"`
	// Number of tools synced
	ToolsSynced int64             `json:"toolsSynced"`
	JSON        syncCompletedJSON `json:"-"`
}

// syncCompletedJSON contains the JSON metadata for the struct [SyncCompleted]
type syncCompletedJSON struct {
	Message     apijson.Field
	ToolsSynced apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncCompleted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncCompletedJSON) RawJSON() string {
	return r.raw
}

// SyncFailed is emitted when a tool set sync operation fails
type SyncFailed struct {
	// Indicates this is an error event
	Error bool `json:"error"`
	// Optional error type/code for programmatic handling
	ErrorType string `json:"errorType"`
	// Error message describing what went wrong
	Message string         `json:"message"`
	JSON    syncFailedJSON `json:"-"`
}

// syncFailedJSON contains the JSON metadata for the struct [SyncFailed]
type syncFailedJSON struct {
	Error       apijson.Field
	ErrorType   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncFailedJSON) RawJSON() string {
	return r.raw
}

// SyncStarted is emitted when a tool set sync operation begins
type SyncStarted struct {
	// Timestamp when the sync was initiated
	Message string          `json:"message"`
	JSON    syncStartedJSON `json:"-"`
}

// syncStartedJSON contains the JSON metadata for the struct [SyncStarted]
type syncStartedJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SyncStarted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r syncStartedJSON) RawJSON() string {
	return r.raw
}

type ToolSet struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     ToolSetSpec             `json:"spec" api:"required"`
	// Tool set information
	Info ToolSetInfo `json:"info"`
	JSON toolSetJSON `json:"-"`
}

// toolSetJSON contains the JSON metadata for the struct [ToolSet]
type toolSetJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetJSON) RawJSON() string {
	return r.raw
}

type ToolSetAdapter struct {
	HTTP ToolSetAdapterHTTP `json:"http"`
	Mcp  ToolSetAdapterMcp  `json:"mcp"`
	JSON toolSetAdapterJSON `json:"-"`
}

// toolSetAdapterJSON contains the JSON metadata for the struct [ToolSetAdapter]
type toolSetAdapterJSON struct {
	HTTP        apijson.Field
	Mcp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetAdapter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetAdapterJSON) RawJSON() string {
	return r.raw
}

type ToolSetAdapterParam struct {
	HTTP param.Field[ToolSetAdapterHTTPParam] `json:"http"`
	Mcp  param.Field[ToolSetAdapterMcpParam]  `json:"mcp"`
}

func (r ToolSetAdapterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetAdapterHTTP struct {
	BaseURL string                 `json:"baseUrl"`
	Headers map[string]string      `json:"headers"`
	JSON    toolSetAdapterHTTPJSON `json:"-"`
}

// toolSetAdapterHTTPJSON contains the JSON metadata for the struct
// [ToolSetAdapterHTTP]
type toolSetAdapterHTTPJSON struct {
	BaseURL     apijson.Field
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetAdapterHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetAdapterHTTPJSON) RawJSON() string {
	return r.raw
}

type ToolSetAdapterHTTPParam struct {
	BaseURL param.Field[string]            `json:"baseUrl"`
	Headers param.Field[map[string]string] `json:"headers"`
}

func (r ToolSetAdapterHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetAdapterMcp struct {
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools McpToolFilter     `json:"excludeTools"`
	Headers      map[string]string `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools McpToolFilter `json:"includeTools"`
	// Approval filters that will automatically set the approval requirement on the
	// tools synced from the MCP server
	ToolApprovals ToolSetAdapterMcpToolApprovals `json:"toolApprovals"`
	URL           string                         `json:"url"`
	JSON          toolSetAdapterMcpJSON          `json:"-"`
}

// toolSetAdapterMcpJSON contains the JSON metadata for the struct
// [ToolSetAdapterMcp]
type toolSetAdapterMcpJSON struct {
	ExcludeTools  apijson.Field
	Headers       apijson.Field
	IncludeTools  apijson.Field
	ToolApprovals apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ToolSetAdapterMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetAdapterMcpJSON) RawJSON() string {
	return r.raw
}

// Approval filters that will automatically set the approval requirement on the
// tools synced from the MCP server
type ToolSetAdapterMcpToolApprovals struct {
	Always bool `json:"always"`
	// Top-level filter with simple boolean logic (no nesting)
	Only McpToolFilter                      `json:"only"`
	JSON toolSetAdapterMcpToolApprovalsJSON `json:"-"`
}

// toolSetAdapterMcpToolApprovalsJSON contains the JSON metadata for the struct
// [ToolSetAdapterMcpToolApprovals]
type toolSetAdapterMcpToolApprovalsJSON struct {
	Always      apijson.Field
	Only        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetAdapterMcpToolApprovals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetAdapterMcpToolApprovalsJSON) RawJSON() string {
	return r.raw
}

type ToolSetAdapterMcpParam struct {
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools param.Field[McpToolFilterParam] `json:"excludeTools"`
	Headers      param.Field[map[string]string]  `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools param.Field[McpToolFilterParam] `json:"includeTools"`
	// Approval filters that will automatically set the approval requirement on the
	// tools synced from the MCP server
	ToolApprovals param.Field[ToolSetAdapterMcpToolApprovalsParam] `json:"toolApprovals"`
	URL           param.Field[string]                              `json:"url"`
}

func (r ToolSetAdapterMcpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Approval filters that will automatically set the approval requirement on the
// tools synced from the MCP server
type ToolSetAdapterMcpToolApprovalsParam struct {
	Always param.Field[bool] `json:"always"`
	// Top-level filter with simple boolean logic (no nesting)
	Only param.Field[McpToolFilterParam] `json:"only"`
}

func (r ToolSetAdapterMcpToolApprovalsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ToolSetEvent represents a single event in the tool set's operation timeline
type ToolSetEvent struct {
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// ToolSetEventData represents the actual event payload for tool set operations
	Event ToolSetEventData `json:"event"`
	Info  ToolSetEventInfo `json:"info"`
	// The tool set this event is associated with
	ToolSetID string           `json:"toolSetId"`
	JSON      toolSetEventJSON `json:"-"`
}

// toolSetEventJSON contains the JSON metadata for the struct [ToolSetEvent]
type toolSetEventJSON struct {
	Metadata    apijson.Field
	Event       apijson.Field
	Info        apijson.Field
	ToolSetID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetEventJSON) RawJSON() string {
	return r.raw
}

type ToolSetEventInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy Profile `json:"createdBy"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolSet shared.ResourceMetadata `json:"toolSet"`
	JSON    toolSetEventInfoJSON    `json:"-"`
}

// toolSetEventInfoJSON contains the JSON metadata for the struct
// [ToolSetEventInfo]
type toolSetEventInfoJSON struct {
	CreatedBy   apijson.Field
	ToolSet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetEventInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetEventInfoJSON) RawJSON() string {
	return r.raw
}

// ToolSetEventData represents the actual event payload for tool set operations
type ToolSetEventData struct {
	// SyncCompleted is emitted when a tool set sync operation completes successfully
	SyncCompleted SyncCompleted `json:"syncCompleted"`
	// SyncFailed is emitted when a tool set sync operation fails
	SyncFailed SyncFailed `json:"syncFailed"`
	// SyncStarted is emitted when a tool set sync operation begins
	SyncStarted SyncStarted `json:"syncStarted"`
	// Type of the event (e.g., "sync_started", "sync_completed", "sync_failed")
	Type string               `json:"type"`
	JSON toolSetEventDataJSON `json:"-"`
}

// toolSetEventDataJSON contains the JSON metadata for the struct
// [ToolSetEventData]
type toolSetEventDataJSON struct {
	SyncCompleted apijson.Field
	SyncFailed    apijson.Field
	SyncStarted   apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ToolSetEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetEventDataJSON) RawJSON() string {
	return r.raw
}

type ToolSetInfo struct {
	AgentCount int64 `json:"agentCount"`
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy Profile         `json:"createdBy"`
	LastSync  time.Time       `json:"lastSync" format:"date-time"`
	ToolCount int64           `json:"toolCount"`
	JSON      toolSetInfoJSON `json:"-"`
}

// toolSetInfoJSON contains the JSON metadata for the struct [ToolSetInfo]
type toolSetInfoJSON struct {
	AgentCount  apijson.Field
	CreatedBy   apijson.Field
	LastSync    apijson.Field
	ToolCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetInfoJSON) RawJSON() string {
	return r.raw
}

type ToolSetSpec struct {
	Adapter     ToolSetAdapter  `json:"adapter"`
	Description string          `json:"description"`
	JSON        toolSetSpecJSON `json:"-"`
}

// toolSetSpecJSON contains the JSON metadata for the struct [ToolSetSpec]
type toolSetSpecJSON struct {
	Adapter     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetSpecJSON) RawJSON() string {
	return r.raw
}

type ToolSetSpecParam struct {
	Adapter     param.Field[ToolSetAdapterParam] `json:"adapter"`
	Description param.Field[string]              `json:"description"`
}

func (r ToolSetSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	Spec     param.Field[ToolSetSpecParam]                   `json:"spec" api:"required"`
}

func (r ToolSetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata   param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	Spec       param.Field[ToolSetSpecParam]                   `json:"spec"`
	UpdateMask param.Field[string]                             `json:"updateMask" format:"field-mask"`
}

func (r ToolSetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [ToolSetListParams]'s query parameters as `url.Values`.
func (r ToolSetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolSetListEventsParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [ToolSetListEventsParams]'s query parameters as
// `url.Values`.
func (r ToolSetListEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
