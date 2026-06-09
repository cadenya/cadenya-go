// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

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
// ToolSetToolService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolSetToolService] method instead.
type ToolSetToolService struct {
	Options []option.RequestOption
}

// NewToolSetToolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolSetToolService(opts ...option.RequestOption) (r *ToolSetToolService) {
	r = &ToolSetToolService{}
	r.Options = opts
	return
}

// Creates a new tool in the tool set
func (r *ToolSetToolService) New(ctx context.Context, workspaceID string, toolSetID string, body ToolSetToolNewParams, opts ...option.RequestOption) (res *Tool, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools", workspaceID, toolSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a tool by ID from the workspace
func (r *ToolSetToolService) Get(ctx context.Context, workspaceID string, toolSetID string, id string, opts ...option.RequestOption) (res *Tool, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s", workspaceID, toolSetID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a tool in the tool set
func (r *ToolSetToolService) Update(ctx context.Context, workspaceID string, toolSetID string, id string, body ToolSetToolUpdateParams, opts ...option.RequestOption) (res *Tool, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s", workspaceID, toolSetID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all tools in the tool set
func (r *ToolSetToolService) List(ctx context.Context, workspaceID string, toolSetID string, query ToolSetToolListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Tool], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools", workspaceID, toolSetID)
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

// Lists all tools in the tool set
func (r *ToolSetToolService) ListAutoPaging(ctx context.Context, workspaceID string, toolSetID string, query ToolSetToolListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Tool] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, toolSetID, query, opts...))
}

// Deletes a tool in the tool set
func (r *ToolSetToolService) Delete(ctx context.Context, workspaceID string, toolSetID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s", workspaceID, toolSetID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Transitions a tool to STATE_OMITTED, excluding it from agent use. Fails if the
// tool is currently assigned to agent variations.
func (r *ToolSetToolService) Omit(ctx context.Context, workspaceID string, toolSetID string, id string, body ToolSetToolOmitParams, opts ...option.RequestOption) (res *Tool, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s:omit", workspaceID, toolSetID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions an omitted tool back to STATE_AVAILABLE. For managed tool sets, the
// next sync may omit the tool again if its filters still exclude it.
func (r *ToolSetToolService) Restore(ctx context.Context, workspaceID string, toolSetID string, id string, body ToolSetToolRestoreParams, opts ...option.RequestOption) (res *Tool, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s:restore", workspaceID, toolSetID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ConfigHTTP struct {
	RequestMethod          ConfigHTTPRequestMethod `json:"requestMethod" api:"required"`
	Headers                map[string]string       `json:"headers"`
	Path                   string                  `json:"path"`
	Query                  string                  `json:"query"`
	RequestBodyContentType string                  `json:"requestBodyContentType"`
	// These are only used when the request method is a POST, PUT, or PATCH
	RequestBodyTemplate string `json:"requestBodyTemplate"`
	// The tool name (commonly an "operation id" in OpenAPI specs) to call on the HTTP
	// adapter. This is used to match the tool spec to the correct endpoint on the HTTP
	// adapter. it will be derived from the name of the tool if not provided.
	ToolName string         `json:"toolName"`
	JSON     configHTTPJSON `json:"-"`
}

// configHTTPJSON contains the JSON metadata for the struct [ConfigHTTP]
type configHTTPJSON struct {
	RequestMethod          apijson.Field
	Headers                apijson.Field
	Path                   apijson.Field
	Query                  apijson.Field
	RequestBodyContentType apijson.Field
	RequestBodyTemplate    apijson.Field
	ToolName               apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ConfigHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configHTTPJSON) RawJSON() string {
	return r.raw
}

type ConfigHTTPRequestMethod string

const (
	ConfigHTTPRequestMethodHTTPMethodUnspecified ConfigHTTPRequestMethod = "HTTP_METHOD_UNSPECIFIED"
	ConfigHTTPRequestMethodGet                   ConfigHTTPRequestMethod = "GET"
	ConfigHTTPRequestMethodPost                  ConfigHTTPRequestMethod = "POST"
	ConfigHTTPRequestMethodPut                   ConfigHTTPRequestMethod = "PUT"
	ConfigHTTPRequestMethodPatch                 ConfigHTTPRequestMethod = "PATCH"
	ConfigHTTPRequestMethodDelete                ConfigHTTPRequestMethod = "DELETE"
)

func (r ConfigHTTPRequestMethod) IsKnown() bool {
	switch r {
	case ConfigHTTPRequestMethodHTTPMethodUnspecified, ConfigHTTPRequestMethodGet, ConfigHTTPRequestMethodPost, ConfigHTTPRequestMethodPut, ConfigHTTPRequestMethodPatch, ConfigHTTPRequestMethodDelete:
		return true
	}
	return false
}

type ConfigHTTPParam struct {
	RequestMethod          param.Field[ConfigHTTPRequestMethod] `json:"requestMethod" api:"required"`
	Headers                param.Field[map[string]string]       `json:"headers"`
	Path                   param.Field[string]                  `json:"path"`
	Query                  param.Field[string]                  `json:"query"`
	RequestBodyContentType param.Field[string]                  `json:"requestBodyContentType"`
	// These are only used when the request method is a POST, PUT, or PATCH
	RequestBodyTemplate param.Field[string] `json:"requestBodyTemplate"`
	// The tool name (commonly an "operation id" in OpenAPI specs) to call on the HTTP
	// adapter. This is used to match the tool spec to the correct endpoint on the HTTP
	// adapter. it will be derived from the name of the tool if not provided.
	ToolName param.Field[string] `json:"toolName"`
}

func (r ConfigHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigMcp struct {
	ToolDescription string        `json:"toolDescription"`
	ToolName        string        `json:"toolName"`
	ToolTitle       string        `json:"toolTitle"`
	JSON            configMcpJSON `json:"-"`
}

// configMcpJSON contains the JSON metadata for the struct [ConfigMcp]
type configMcpJSON struct {
	ToolDescription apijson.Field
	ToolName        apijson.Field
	ToolTitle       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configMcpJSON) RawJSON() string {
	return r.raw
}

type ConfigMcpParam struct {
	ToolDescription param.Field[string] `json:"toolDescription"`
	ToolName        param.Field[string] `json:"toolName"`
	ToolTitle       param.Field[string] `json:"toolTitle"`
}

func (r ConfigMcpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigOpenAPI struct {
	Method      string            `json:"method"`
	OperationID string            `json:"operationId"`
	Path        string            `json:"path"`
	JSON        configOpenAPIJSON `json:"-"`
}

// configOpenAPIJSON contains the JSON metadata for the struct [ConfigOpenAPI]
type configOpenAPIJSON struct {
	Method      apijson.Field
	OperationID apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigOpenAPI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configOpenAPIJSON) RawJSON() string {
	return r.raw
}

type ConfigOpenAPIParam struct {
	Method      param.Field[string] `json:"method"`
	OperationID param.Field[string] `json:"operationId"`
	Path        param.Field[string] `json:"path"`
}

func (r ConfigOpenAPIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Tool struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     ToolSpec                `json:"spec" api:"required"`
	// The current lifecycle state of the tool. Output only. Use the :omit and :restore
	// actions to transition; tool set syncs may also update it.
	State ToolState `json:"state" api:"required"`
	Info  ToolInfo  `json:"info"`
	JSON  toolJSON  `json:"-"`
}

// toolJSON contains the JSON metadata for the struct [Tool]
type toolJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	State       apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolJSON) RawJSON() string {
	return r.raw
}

// The current lifecycle state of the tool. Output only. Use the :omit and :restore
// actions to transition; tool set syncs may also update it.
type ToolState string

const (
	ToolStateStateUnspecified ToolState = "STATE_UNSPECIFIED"
	ToolStateStateAvailable   ToolState = "STATE_AVAILABLE"
	ToolStateStateOmitted     ToolState = "STATE_OMITTED"
	ToolStateStateArchived    ToolState = "STATE_ARCHIVED"
)

func (r ToolState) IsKnown() bool {
	switch r {
	case ToolStateStateUnspecified, ToolStateStateAvailable, ToolStateStateOmitted, ToolStateStateArchived:
		return true
	}
	return false
}

type ToolInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolSet shared.ResourceMetadata `json:"toolSet"`
	JSON    toolInfoJSON            `json:"-"`
}

// toolInfoJSON contains the JSON metadata for the struct [ToolInfo]
type toolInfoJSON struct {
	CreatedBy   apijson.Field
	ToolSet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolInfoJSON) RawJSON() string {
	return r.raw
}

type ToolSpec struct {
	// Config defines the adapter to use for the tool. This is used to determine how
	// the tool is called. For example, if the tool is an HTTP tool, the adapter will
	// be Http. If the tool is an inline tool, the adapter will be Inline.
	Config           ToolSpecConfig         `json:"config" api:"required"`
	Description      string                 `json:"description" api:"required"`
	Parameters       map[string]interface{} `json:"parameters" api:"required"`
	RequiresApproval bool                   `json:"requiresApproval" api:"required"`
	JSON             toolSpecJSON           `json:"-"`
}

// toolSpecJSON contains the JSON metadata for the struct [ToolSpec]
type toolSpecJSON struct {
	Config           apijson.Field
	Description      apijson.Field
	Parameters       apijson.Field
	RequiresApproval apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ToolSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSpecJSON) RawJSON() string {
	return r.raw
}

type ToolSpecParam struct {
	// Config defines the adapter to use for the tool. This is used to determine how
	// the tool is called. For example, if the tool is an HTTP tool, the adapter will
	// be Http. If the tool is an inline tool, the adapter will be Inline.
	Config           param.Field[ToolSpecConfigParam]    `json:"config" api:"required"`
	Description      param.Field[string]                 `json:"description" api:"required"`
	Parameters       param.Field[map[string]interface{}] `json:"parameters" api:"required"`
	RequiresApproval param.Field[bool]                   `json:"requiresApproval" api:"required"`
}

func (r ToolSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Config defines the adapter to use for the tool. This is used to determine how
// the tool is called. For example, if the tool is an HTTP tool, the adapter will
// be Http. If the tool is an inline tool, the adapter will be Inline.
type ToolSpecConfig struct {
	HTTP    ConfigHTTP         `json:"http"`
	Mcp     ConfigMcp          `json:"mcp"`
	OpenAPI ConfigOpenAPI      `json:"openapi"`
	JSON    toolSpecConfigJSON `json:"-"`
}

// toolSpecConfigJSON contains the JSON metadata for the struct [ToolSpecConfig]
type toolSpecConfigJSON struct {
	HTTP        apijson.Field
	Mcp         apijson.Field
	OpenAPI     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSpecConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSpecConfigJSON) RawJSON() string {
	return r.raw
}

// Config defines the adapter to use for the tool. This is used to determine how
// the tool is called. For example, if the tool is an HTTP tool, the adapter will
// be Http. If the tool is an inline tool, the adapter will be Inline.
type ToolSpecConfigParam struct {
	HTTP    param.Field[ConfigHTTPParam]    `json:"http"`
	Mcp     param.Field[ConfigMcpParam]     `json:"mcp"`
	OpenAPI param.Field[ConfigOpenAPIParam] `json:"openapi"`
}

func (r ToolSpecConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetToolNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	Spec     param.Field[ToolSpecParam]                      `json:"spec" api:"required"`
}

func (r ToolSetToolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetToolUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata   param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	Spec       param.Field[ToolSpecParam]                      `json:"spec"`
	UpdateMask param.Field[string]                             `json:"updateMask" format:"field-mask"`
}

func (r ToolSetToolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetToolListParams struct {
	// Filter by bundle_key — return only resources owned by this bundle.
	BundleKey param.Field[string] `query:"bundleKey"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter by tool name (exact match). Multiple values are OR'd together.
	Names param.Field[[]string] `query:"names"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query
	Query param.Field[string] `query:"query"`
	// Filter by approval requirement. Omitted = no filter; true = only tools requiring
	// approval; false = only tools not requiring approval.
	RequiresApproval param.Field[bool] `query:"requiresApproval"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by tool state. Multiple values are OR'd together.
	States param.Field[[]ToolSetToolListParamsState] `query:"states"`
}

// URLQuery serializes [ToolSetToolListParams]'s query parameters as `url.Values`.
func (r ToolSetToolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolSetToolListParamsState string

const (
	ToolSetToolListParamsStateStateUnspecified ToolSetToolListParamsState = "STATE_UNSPECIFIED"
	ToolSetToolListParamsStateStateAvailable   ToolSetToolListParamsState = "STATE_AVAILABLE"
	ToolSetToolListParamsStateStateOmitted     ToolSetToolListParamsState = "STATE_OMITTED"
	ToolSetToolListParamsStateStateArchived    ToolSetToolListParamsState = "STATE_ARCHIVED"
)

func (r ToolSetToolListParamsState) IsKnown() bool {
	switch r {
	case ToolSetToolListParamsStateStateUnspecified, ToolSetToolListParamsStateStateAvailable, ToolSetToolListParamsStateStateOmitted, ToolSetToolListParamsStateStateArchived:
		return true
	}
	return false
}

type ToolSetToolOmitParams struct {
}

func (r ToolSetToolOmitParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetToolRestoreParams struct {
}

func (r ToolSetToolRestoreParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
