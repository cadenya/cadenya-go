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
	Options []option.RequestOption
	// Manage tool sets and the tools they contain. Tool sets group related tools, and
	// tools define specific capabilities available to agents.
	//
	// When a tool set is managed, only API key actors can modify its tools; human
	// (profile) actors cannot.
	Tools *ToolSetToolService
	// Manage tool sets and the tools they contain. Tool sets group related tools, and
	// tools define specific capabilities available to agents.
	//
	// When a tool set is managed, only API key actors can modify its tools; human
	// (profile) actors cannot.
	Secrets *ToolSetSecretService
}

// NewToolSetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolSetService(opts ...option.RequestOption) (r *ToolSetService) {
	r = &ToolSetService{}
	r.Options = opts
	r.Tools = NewToolSetToolService(opts...)
	r.Secrets = NewToolSetSecretService(opts...)
	return
}

// Creates a new tool set in the workspace
func (r *ToolSetService) New(ctx context.Context, workspaceID string, body ToolSetNewParams, opts ...option.RequestOption) (res *ToolSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a tool set by ID from the workspace
func (r *ToolSetService) Get(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (res *ToolSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a tool set in the workspace
func (r *ToolSetService) Update(ctx context.Context, workspaceID string, id string, body ToolSetUpdateParams, opts ...option.RequestOption) (res *ToolSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all tool sets in the workspace
func (r *ToolSetService) List(ctx context.Context, workspaceID string, query ToolSetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSet], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets", workspaceID)
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
func (r *ToolSetService) ListAutoPaging(ctx context.Context, workspaceID string, query ToolSetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSet] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Deletes a tool set in the workspace
func (r *ToolSetService) Delete(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Transitions a tool set to STATE_ARCHIVED. Syncing stops, the tool set is hidden
// from list results, its tools are no longer offered to objectives, and new
// variation assignments are rejected. Existing assignments are retained, and
// history is preserved — unlike delete, archiving works while the tool set is
// still assigned to agent variations.
func (r *ToolSetService) Archive(ctx context.Context, workspaceID string, id string, body ToolSetArchiveParams, opts ...option.RequestOption) (res *ToolSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s:archive", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the current OpenAPI specification JSON that has been consumed by the
// tool set. Only applicable to tool sets using the OpenAPI adapter.
func (r *ToolSetService) GetOpenAPISpec(ctx context.Context, workspaceID string, toolSetID string, opts ...option.RequestOption) (res *ToolSetGetOpenAPISpecResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/openapi_spec", workspaceID, toolSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all events (including sync status) for a tool set
func (r *ToolSetService) ListEvents(ctx context.Context, workspaceID string, toolSetID string, query ToolSetListEventsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSetEvent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/events", workspaceID, toolSetID)
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
func (r *ToolSetService) ListEventsAutoPaging(ctx context.Context, workspaceID string, toolSetID string, query ToolSetListEventsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSetEvent] {
	return pagination.NewCursorPaginationAutoPager(r.ListEvents(ctx, workspaceID, toolSetID, query, opts...))
}

// Transitions an archived tool set back to STATE_ACTIVE. Managed tool sets resume
// syncing on their next cycle and their tools become available to objectives
// again.
func (r *ToolSetService) Unarchive(ctx context.Context, workspaceID string, id string, body ToolSetUnarchiveParams, opts ...option.RequestOption) (res *ToolSet, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s:unarchive", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Approval filters that will automatically set the approval requirement on tools
// synced from an external source
type ApprovalRequirementFilter struct {
	Always bool `json:"always"`
	// Top-level filter with simple boolean logic (no nesting)
	Only ToolFilter                    `json:"only"`
	JSON approvalRequirementFilterJSON `json:"-"`
}

// approvalRequirementFilterJSON contains the JSON metadata for the struct
// [ApprovalRequirementFilter]
type approvalRequirementFilterJSON struct {
	Always      apijson.Field
	Only        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ApprovalRequirementFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r approvalRequirementFilterJSON) RawJSON() string {
	return r.raw
}

// Approval filters that will automatically set the approval requirement on tools
// synced from an external source
type ApprovalRequirementFilterParam struct {
	Always param.Field[bool] `json:"always"`
	// Top-level filter with simple boolean logic (no nesting)
	Only param.Field[ToolFilterParam] `json:"only"`
}

func (r ApprovalRequirementFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Single attribute filter
type AttributeFilter struct {
	Attribute AttributeFilterAttribute `json:"attribute" api:"required"`
	// String matching operations
	Matcher StringMatcher       `json:"matcher"`
	JSON    attributeFilterJSON `json:"-"`
}

// attributeFilterJSON contains the JSON metadata for the struct [AttributeFilter]
type attributeFilterJSON struct {
	Attribute   apijson.Field
	Matcher     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttributeFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attributeFilterJSON) RawJSON() string {
	return r.raw
}

type AttributeFilterAttribute string

const (
	AttributeFilterAttributeAttributeUnspecified AttributeFilterAttribute = "ATTRIBUTE_UNSPECIFIED"
	AttributeFilterAttributeAttributeName        AttributeFilterAttribute = "ATTRIBUTE_NAME"
	AttributeFilterAttributeAttributeTitle       AttributeFilterAttribute = "ATTRIBUTE_TITLE"
	AttributeFilterAttributeAttributeDescription AttributeFilterAttribute = "ATTRIBUTE_DESCRIPTION"
)

func (r AttributeFilterAttribute) IsKnown() bool {
	switch r {
	case AttributeFilterAttributeAttributeUnspecified, AttributeFilterAttributeAttributeName, AttributeFilterAttributeAttributeTitle, AttributeFilterAttributeAttributeDescription:
		return true
	}
	return false
}

// Single attribute filter
type AttributeFilterParam struct {
	Attribute param.Field[AttributeFilterAttribute] `json:"attribute" api:"required"`
	// String matching operations
	Matcher param.Field[StringMatcherParam] `json:"matcher"`
}

func (r AttributeFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String matching operations
type StringMatcher struct {
	CaseSensitive bool              `json:"caseSensitive"`
	Contains      string            `json:"contains"`
	EndsWith      string            `json:"endsWith"`
	Exact         string            `json:"exact"`
	Regex         string            `json:"regex"`
	StartsWith    string            `json:"startsWith"`
	JSON          stringMatcherJSON `json:"-"`
}

// stringMatcherJSON contains the JSON metadata for the struct [StringMatcher]
type stringMatcherJSON struct {
	CaseSensitive apijson.Field
	Contains      apijson.Field
	EndsWith      apijson.Field
	Exact         apijson.Field
	Regex         apijson.Field
	StartsWith    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *StringMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r stringMatcherJSON) RawJSON() string {
	return r.raw
}

// String matching operations
type StringMatcherParam struct {
	CaseSensitive param.Field[bool]   `json:"caseSensitive"`
	Contains      param.Field[string] `json:"contains"`
	EndsWith      param.Field[string] `json:"endsWith"`
	Exact         param.Field[string] `json:"exact"`
	Regex         param.Field[string] `json:"regex"`
	StartsWith    param.Field[string] `json:"startsWith"`
}

func (r StringMatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Emitted when a tool set sync operation completes successfully.
type SyncCompleted struct {
	// Optional message with additional details.
	Message string `json:"message"`
	// Number of tools synced.
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

// Emitted when a tool set sync operation fails.
type SyncFailed struct {
	// Indicates this is an error event.
	Error bool `json:"error"`
	// Optional error type/code for programmatic handling.
	ErrorType string `json:"errorType"`
	// Error message describing what went wrong.
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

// Emitted when a tool set sync operation begins.
type SyncStarted struct {
	// Human-readable message describing the start of the sync.
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

// Top-level filter with simple boolean logic (no nesting)
type ToolFilter struct {
	Operator ToolFilterOperator `json:"operator" api:"required"`
	Filters  []AttributeFilter  `json:"filters"`
	JSON     toolFilterJSON     `json:"-"`
}

// toolFilterJSON contains the JSON metadata for the struct [ToolFilter]
type toolFilterJSON struct {
	Operator    apijson.Field
	Filters     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolFilterJSON) RawJSON() string {
	return r.raw
}

type ToolFilterOperator string

const (
	ToolFilterOperatorOperatorUnspecified ToolFilterOperator = "OPERATOR_UNSPECIFIED"
	ToolFilterOperatorOperatorAnd         ToolFilterOperator = "OPERATOR_AND"
	ToolFilterOperatorOperatorOr          ToolFilterOperator = "OPERATOR_OR"
)

func (r ToolFilterOperator) IsKnown() bool {
	switch r {
	case ToolFilterOperatorOperatorUnspecified, ToolFilterOperatorOperatorAnd, ToolFilterOperatorOperatorOr:
		return true
	}
	return false
}

// Top-level filter with simple boolean logic (no nesting)
type ToolFilterParam struct {
	Operator param.Field[ToolFilterOperator]     `json:"operator" api:"required"`
	Filters  param.Field[[]AttributeFilterParam] `json:"filters"`
}

func (r ToolFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSet struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     ToolSetSpec             `json:"spec" api:"required"`
	// The current lifecycle state of the tool set. Output only. Tool sets are created
	// STATE_ACTIVE; use the :archive and :unarchive actions to transition between
	// states.
	State ToolSetState `json:"state" api:"required"`
	// Tool set information
	Info ToolSetInfo `json:"info"`
	JSON toolSetJSON `json:"-"`
}

// toolSetJSON contains the JSON metadata for the struct [ToolSet]
type toolSetJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	State       apijson.Field
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

// The current lifecycle state of the tool set. Output only. Tool sets are created
// STATE_ACTIVE; use the :archive and :unarchive actions to transition between
// states.
type ToolSetState string

const (
	ToolSetStateStateUnspecified ToolSetState = "STATE_UNSPECIFIED"
	ToolSetStateStateActive      ToolSetState = "STATE_ACTIVE"
	ToolSetStateStateArchived    ToolSetState = "STATE_ARCHIVED"
)

func (r ToolSetState) IsKnown() bool {
	switch r {
	case ToolSetStateStateUnspecified, ToolSetStateStateActive, ToolSetStateStateArchived:
		return true
	}
	return false
}

type ToolSetAdapter struct {
	HTTP    ToolSetAdapterHTTP    `json:"http"`
	Mcp     ToolSetAdapterMcp     `json:"mcp"`
	OpenAPI ToolSetAdapterOpenAPI `json:"openapi"`
	JSON    toolSetAdapterJSON    `json:"-"`
}

// toolSetAdapterJSON contains the JSON metadata for the struct [ToolSetAdapter]
type toolSetAdapterJSON struct {
	HTTP        apijson.Field
	Mcp         apijson.Field
	OpenAPI     apijson.Field
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
	HTTP    param.Field[ToolSetAdapterHTTPParam]    `json:"http"`
	Mcp     param.Field[ToolSetAdapterMcpParam]     `json:"mcp"`
	OpenAPI param.Field[ToolSetAdapterOpenAPIParam] `json:"openapi"`
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
	ExcludeTools ToolFilter        `json:"excludeTools"`
	Headers      map[string]string `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilter `json:"includeTools"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilter `json:"toolApprovals"`
	URL           string                    `json:"url"`
	JSON          toolSetAdapterMcpJSON     `json:"-"`
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

type ToolSetAdapterMcpParam struct {
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools param.Field[ToolFilterParam]   `json:"excludeTools"`
	Headers      param.Field[map[string]string] `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools param.Field[ToolFilterParam] `json:"includeTools"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals param.Field[ApprovalRequirementFilterParam] `json:"toolApprovals"`
	URL           param.Field[string]                         `json:"url"`
}

func (r ToolSetAdapterMcpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// ID of a COMPLETE Upload containing the OpenAPI spec document.
	UploadID string `json:"uploadId"`
	// URL to fetch the OpenAPI spec from. Synced automatically every hour.
	URL  string                    `json:"url"`
	JSON toolSetAdapterOpenAPIJSON `json:"-"`
}

// toolSetAdapterOpenAPIJSON contains the JSON metadata for the struct
// [ToolSetAdapterOpenAPI]
type toolSetAdapterOpenAPIJSON struct {
	BaseURL       apijson.Field
	ExcludeTools  apijson.Field
	Headers       apijson.Field
	IncludeTools  apijson.Field
	ServerName    apijson.Field
	ToolApprovals apijson.Field
	UploadID      apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ToolSetAdapterOpenAPI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetAdapterOpenAPIJSON) RawJSON() string {
	return r.raw
}

type ToolSetAdapterOpenAPIParam struct {
	// Base URL for dispatching tool calls. If set, overrides the server resolved from
	// the spec's servers array.
	BaseURL param.Field[string] `json:"baseUrl"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools param.Field[ToolFilterParam] `json:"excludeTools"`
	// Headers sent when fetching the spec from a URL and when dispatching tool calls.
	Headers param.Field[map[string]string] `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools param.Field[ToolFilterParam] `json:"includeTools"`
	// Name of the server entry in the spec's servers array (OpenAPI 3.2 server.name
	// field). Used to select which server URL to dispatch to when base_url is not set.
	// If unset, the first server is used. Ignored when base_url is set.
	ServerName param.Field[string] `json:"serverName"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals param.Field[ApprovalRequirementFilterParam] `json:"toolApprovals"`
	// ID of a COMPLETE Upload containing the OpenAPI spec document.
	UploadID param.Field[string] `json:"uploadId"`
	// URL to fetch the OpenAPI spec from. Synced automatically every hour.
	URL param.Field[string] `json:"url"`
}

func (r ToolSetAdapterOpenAPIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
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

// Event payload for a tool set operation.
type ToolSetEventData struct {
	// Emitted when a tool set sync operation completes successfully.
	SyncCompleted SyncCompleted `json:"syncCompleted"`
	// Emitted when a tool set sync operation fails.
	SyncFailed SyncFailed `json:"syncFailed"`
	// Emitted when a tool set sync operation begins.
	SyncStarted SyncStarted `json:"syncStarted"`
	// Type of the event (e.g., "sync_started", "sync_completed", "sync_failed").
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
	AgentCount     int64 `json:"agentCount"`
	AvailableTools int64 `json:"availableTools"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy    Profile         `json:"createdBy"`
	LastSync     time.Time       `json:"lastSync" format:"date-time"`
	OmittedTools int64           `json:"omittedTools"`
	ToolCount    int64           `json:"toolCount"`
	JSON         toolSetInfoJSON `json:"-"`
}

// toolSetInfoJSON contains the JSON metadata for the struct [ToolSetInfo]
type toolSetInfoJSON struct {
	AgentCount     apijson.Field
	AvailableTools apijson.Field
	CreatedBy      apijson.Field
	LastSync       apijson.Field
	OmittedTools   apijson.Field
	ToolCount      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
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

type ToolSetGetOpenAPISpecResponse struct {
	// The consumed OpenAPI specification as a JSON string.
	Spec string                            `json:"spec"`
	JSON toolSetGetOpenAPISpecResponseJSON `json:"-"`
}

// toolSetGetOpenAPISpecResponseJSON contains the JSON metadata for the struct
// [ToolSetGetOpenAPISpecResponse]
type toolSetGetOpenAPISpecResponseJSON struct {
	Spec        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetGetOpenAPISpecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetGetOpenAPISpecResponseJSON) RawJSON() string {
	return r.raw
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
	// Filter by bundle_key — return only resources owned by this bundle.
	BundleKey param.Field[string] `query:"bundleKey"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by tool set lifecycle state. Defaults to STATE_ACTIVE when unspecified;
	// pass STATE_ARCHIVED to list archived tool sets.
	State param.Field[ToolSetListParamsState] `query:"state"`
}

// URLQuery serializes [ToolSetListParams]'s query parameters as `url.Values`.
func (r ToolSetListParams) URLQuery() (v url.Values) {
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

func (r ToolSetListParamsState) IsKnown() bool {
	switch r {
	case ToolSetListParamsStateStateUnspecified, ToolSetListParamsStateStateActive, ToolSetListParamsStateStateArchived:
		return true
	}
	return false
}

type ToolSetArchiveParams struct {
}

func (r ToolSetArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type ToolSetUnarchiveParams struct {
}

func (r ToolSetUnarchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
