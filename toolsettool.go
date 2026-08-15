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
	options []option.RequestOption
}

// NewToolSetToolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolSetToolService(opts ...option.RequestOption) (r ToolSetToolService) {
	r = ToolSetToolService{}
	r.options = opts
	return
}

// Creates a new tool in the tool set
func (r *ToolSetToolService) New(ctx context.Context, toolSetID string, params ToolSetToolNewParams, opts ...option.RequestOption) (res *Tool, err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a tool by ID from the workspace
func (r *ToolSetToolService) Get(ctx context.Context, toolSetID string, id string, query ToolSetToolGetParams, opts ...option.RequestOption) (res *Tool, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(toolSetID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a tool in the tool set
func (r *ToolSetToolService) Update(ctx context.Context, toolSetID string, id string, params ToolSetToolUpdateParams, opts ...option.RequestOption) (res *Tool, err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all tools in the tool set
func (r *ToolSetToolService) List(ctx context.Context, toolSetID string, params ToolSetToolListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Tool], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID))
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

// Lists all tools in the tool set
func (r *ToolSetToolService) ListAutoPaging(ctx context.Context, toolSetID string, params ToolSetToolListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Tool] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, toolSetID, params, opts...))
}

// Deletes a tool in the tool set
func (r *ToolSetToolService) Delete(ctx context.Context, toolSetID string, id string, body ToolSetToolDeleteParams, opts ...option.RequestOption) (err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(toolSetID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Transitions a tool to STATE_OMITTED, excluding it from agent use. Fails if the
// tool is currently assigned to agent variations.
func (r *ToolSetToolService) Omit(ctx context.Context, toolSetID string, id string, body ToolSetToolOmitParams, opts ...option.RequestOption) (res *Tool, err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s:omit", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(toolSetID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions an omitted tool back to STATE_AVAILABLE. For managed tool sets, the
// next sync may omit the tool again if its filters still exclude it.
func (r *ToolSetToolService) Restore(ctx context.Context, toolSetID string, id string, body ToolSetToolRestoreParams, opts ...option.RequestOption) (res *Tool, err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/tools/%s:restore", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(toolSetID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Marks the tool as bare: it has no execution adapter of its own and relies on the
// parent tool set being a Bare tool set. Present so a webhook consumer can tell a
// tool is bare from the tool data alone, without cross-referencing the tool set.
type ConfigBare struct {
	// When set, the tool call's result is recorded immediately as this fixed text
	// instead of parking the call to wait for externally supplied content. The
	// tool_called event is still emitted. Useful for tools whose dispatch is the
	// intent (e.g. a frontend renders a component from the call parameters) but whose
	// LLM turn still needs tool-result content.
	AlwaysSetResult string `json:"alwaysSetResult"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AlwaysSetResult respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigBare) RawJSON() string { return r.JSON.raw }
func (r *ConfigBare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConfigBare to a ConfigBareParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConfigBareParam.Overrides()
func (r ConfigBare) ToParam() ConfigBareParam {
	return param.Override[ConfigBareParam](json.RawMessage(r.RawJSON()))
}

// Marks the tool as bare: it has no execution adapter of its own and relies on the
// parent tool set being a Bare tool set. Present so a webhook consumer can tell a
// tool is bare from the tool data alone, without cross-referencing the tool set.
type ConfigBareParam struct {
	// When set, the tool call's result is recorded immediately as this fixed text
	// instead of parking the call to wait for externally supplied content. The
	// tool_called event is still emitted. Useful for tools whose dispatch is the
	// intent (e.g. a frontend renders a component from the call parameters) but whose
	// LLM turn still needs tool-result content.
	AlwaysSetResult param.Opt[string] `json:"alwaysSetResult,omitzero"`
	paramObj
}

func (r ConfigBareParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigBareParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigBareParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigHTTP struct {
	// Any of "HTTP_METHOD_UNSPECIFIED", "GET", "POST", "PUT", "PATCH", "DELETE".
	RequestMethod          ConfigHTTPRequestMethod `json:"requestMethod" api:"required"`
	Headers                map[string]string       `json:"headers"`
	Path                   string                  `json:"path"`
	Query                  string                  `json:"query"`
	RequestBodyContentType string                  `json:"requestBodyContentType"`
	// These are only used when the request method is a POST, PUT, or PATCH
	RequestBodyTemplate string `json:"requestBodyTemplate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestMethod          respjson.Field
		Headers                respjson.Field
		Path                   respjson.Field
		Query                  respjson.Field
		RequestBodyContentType respjson.Field
		RequestBodyTemplate    respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigHTTP) RawJSON() string { return r.JSON.raw }
func (r *ConfigHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConfigHTTP to a ConfigHTTPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConfigHTTPParam.Overrides()
func (r ConfigHTTP) ToParam() ConfigHTTPParam {
	return param.Override[ConfigHTTPParam](json.RawMessage(r.RawJSON()))
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

// The property RequestMethod is required.
type ConfigHTTPParam struct {
	// Any of "HTTP_METHOD_UNSPECIFIED", "GET", "POST", "PUT", "PATCH", "DELETE".
	RequestMethod          ConfigHTTPRequestMethod `json:"requestMethod,omitzero" api:"required"`
	Path                   param.Opt[string]       `json:"path,omitzero"`
	Query                  param.Opt[string]       `json:"query,omitzero"`
	RequestBodyContentType param.Opt[string]       `json:"requestBodyContentType,omitzero"`
	// These are only used when the request method is a POST, PUT, or PATCH
	RequestBodyTemplate param.Opt[string] `json:"requestBodyTemplate,omitzero"`
	Headers             map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r ConfigHTTPParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigHTTPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigHTTPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigMCP struct {
	// Behavior hints synced from the MCP server's tool definition (ToolAnnotations in
	// the MCP specification). All hints are advisory: servers are not required to send
	// them, and clients should not rely on them for security decisions. Absent hints
	// keep the MCP spec defaults (destructiveHint and openWorldHint default to true;
	// readOnlyHint and idempotentHint default to false).
	Annotations MCPAnnotations `json:"annotations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Annotations respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigMCP) RawJSON() string { return r.JSON.raw }
func (r *ConfigMCP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConfigMCP to a ConfigMCPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConfigMCPParam.Overrides()
func (r ConfigMCP) ToParam() ConfigMCPParam {
	return param.Override[ConfigMCPParam](json.RawMessage(r.RawJSON()))
}

type ConfigMCPParam struct {
	// Behavior hints synced from the MCP server's tool definition (ToolAnnotations in
	// the MCP specification). All hints are advisory: servers are not required to send
	// them, and clients should not rely on them for security decisions. Absent hints
	// keep the MCP spec defaults (destructiveHint and openWorldHint default to true;
	// readOnlyHint and idempotentHint default to false).
	Annotations MCPAnnotationsParam `json:"annotations,omitzero"`
	paramObj
}

func (r ConfigMCPParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigMCPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigMCPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigOpenAPI struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigOpenAPI) RawJSON() string { return r.JSON.raw }
func (r *ConfigOpenAPI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConfigOpenAPI to a ConfigOpenAPIParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConfigOpenAPIParam.Overrides()
func (r ConfigOpenAPI) ToParam() ConfigOpenAPIParam {
	return param.Override[ConfigOpenAPIParam](json.RawMessage(r.RawJSON()))
}

type ConfigOpenAPIParam struct {
	Method param.Opt[string] `json:"method,omitzero"`
	Path   param.Opt[string] `json:"path,omitzero"`
	paramObj
}

func (r ConfigOpenAPIParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigOpenAPIParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigOpenAPIParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Behavior hints synced from the MCP server's tool definition (ToolAnnotations in
// the MCP specification). All hints are advisory: servers are not required to send
// them, and clients should not rely on them for security decisions. Absent hints
// keep the MCP spec defaults (destructiveHint and openWorldHint default to true;
// readOnlyHint and idempotentHint default to false).
type MCPAnnotations struct {
	// If true, the tool may perform destructive updates to its environment. Only
	// meaningful when read_only_hint is false.
	DestructiveHint bool `json:"destructiveHint"`
	// If true, calling the tool repeatedly with the same arguments has no additional
	// effect. Only meaningful when read_only_hint is false.
	IdempotentHint bool `json:"idempotentHint"`
	// If true, the tool may interact with an "open world" of external entities (e.g.
	// web search); if false, its domain is closed.
	OpenWorldHint bool `json:"openWorldHint"`
	// If true, the tool does not modify its environment.
	ReadOnlyHint bool `json:"readOnlyHint"`
	// A human-readable title for the tool.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestructiveHint respjson.Field
		IdempotentHint  respjson.Field
		OpenWorldHint   respjson.Field
		ReadOnlyHint    respjson.Field
		Title           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MCPAnnotations) RawJSON() string { return r.JSON.raw }
func (r *MCPAnnotations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MCPAnnotations to a MCPAnnotationsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MCPAnnotationsParam.Overrides()
func (r MCPAnnotations) ToParam() MCPAnnotationsParam {
	return param.Override[MCPAnnotationsParam](json.RawMessage(r.RawJSON()))
}

// Behavior hints synced from the MCP server's tool definition (ToolAnnotations in
// the MCP specification). All hints are advisory: servers are not required to send
// them, and clients should not rely on them for security decisions. Absent hints
// keep the MCP spec defaults (destructiveHint and openWorldHint default to true;
// readOnlyHint and idempotentHint default to false).
type MCPAnnotationsParam struct {
	// If true, the tool may perform destructive updates to its environment. Only
	// meaningful when read_only_hint is false.
	DestructiveHint param.Opt[bool] `json:"destructiveHint,omitzero"`
	// If true, calling the tool repeatedly with the same arguments has no additional
	// effect. Only meaningful when read_only_hint is false.
	IdempotentHint param.Opt[bool] `json:"idempotentHint,omitzero"`
	// If true, the tool may interact with an "open world" of external entities (e.g.
	// web search); if false, its domain is closed.
	OpenWorldHint param.Opt[bool] `json:"openWorldHint,omitzero"`
	// If true, the tool does not modify its environment.
	ReadOnlyHint param.Opt[bool] `json:"readOnlyHint,omitzero"`
	// A human-readable title for the tool.
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r MCPAnnotationsParam) MarshalJSON() (data []byte, err error) {
	type shadow MCPAnnotationsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MCPAnnotationsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Tool struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     ToolSpec                `json:"spec" api:"required"`
	// The current lifecycle state of the tool. Output only. Use the :omit and :restore
	// actions to transition; tool set syncs may also update it.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_AVAILABLE", "STATE_OMITTED",
	// "STATE_ARCHIVED".
	State ToolState `json:"state" api:"required"`
	Info  ToolInfo  `json:"info"`
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
func (r Tool) RawJSON() string { return r.JSON.raw }
func (r *Tool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type ToolInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Content signature identifying the tool within its tool set: a hash of the
	// sanitized llm_tool_name, description, and canonical parameters. Two tools with
	// the same llm_tool_name but different parameters or description (as MCP servers
	// may return per user) have distinct signatures.
	Signature string `json:"signature"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolSet shared.ResourceMetadata `json:"toolSet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		Signature   respjson.Field
		ToolSet     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolInfo) RawJSON() string { return r.JSON.raw }
func (r *ToolInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSpec struct {
	// Config defines the adapter to use for the tool. This is used to determine how
	// the tool is called. For example, if the tool is an HTTP tool, the adapter will
	// be Http. If the tool is an inline tool, the adapter will be Inline.
	Config      ToolSpecConfigUnion `json:"config" api:"required"`
	Description string              `json:"description" api:"required"`
	// The tool's JSON Schema, as handed to the LLM. Required, but may be the empty
	// object `{}` for a tool that takes no arguments. Requiring it rather than
	// defaulting it means a misspelled field name (`inputSchema`, say) is a 400
	// instead of a silently parameterless tool.
	Parameters       map[string]any `json:"parameters" api:"required"`
	RequiresApproval bool           `json:"requiresApproval" api:"required"`
	// The name provided to the LLM, which may differ from the metadata.name on the
	// tool. LLMs have specific length and format requirements, and tool set sources
	// may not comply with them, so Cadenya does its best to format names into a usable
	// format.
	LlmToolName string `json:"llmToolName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config           respjson.Field
		Description      respjson.Field
		Parameters       respjson.Field
		RequiresApproval respjson.Field
		LlmToolName      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSpec) RawJSON() string { return r.JSON.raw }
func (r *ToolSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSpec to a ToolSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSpecParam.Overrides()
func (r ToolSpec) ToParam() ToolSpecParam {
	return param.Override[ToolSpecParam](json.RawMessage(r.RawJSON()))
}

// The properties Config, Description, Parameters, RequiresApproval are required.
type ToolSpecParam struct {
	// Config defines the adapter to use for the tool. This is used to determine how
	// the tool is called. For example, if the tool is an HTTP tool, the adapter will
	// be Http. If the tool is an inline tool, the adapter will be Inline.
	Config      ToolSpecConfigUnionParam `json:"config,omitzero" api:"required"`
	Description string                   `json:"description" api:"required"`
	// The tool's JSON Schema, as handed to the LLM. Required, but may be the empty
	// object `{}` for a tool that takes no arguments. Requiring it rather than
	// defaulting it means a misspelled field name (`inputSchema`, say) is a 400
	// instead of a silently parameterless tool.
	Parameters       map[string]any `json:"parameters,omitzero" api:"required"`
	RequiresApproval bool           `json:"requiresApproval" api:"required"`
	// The name provided to the LLM, which may differ from the metadata.name on the
	// tool. LLMs have specific length and format requirements, and tool set sources
	// may not comply with them, so Cadenya does its best to format names into a usable
	// format.
	LlmToolName param.Opt[string] `json:"llmToolName,omitzero"`
	paramObj
}

func (r ToolSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolSpecConfigUnion contains all possible properties and values from
// [ToolSpecConfigHTTP], [ToolSpecConfigMCP], [ToolSpecConfigOpenAPI],
// [ToolSpecConfigBare].
//
// Use the [ToolSpecConfigUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToolSpecConfigUnion struct {
	// This field is from variant [ToolSpecConfigHTTP].
	HTTP ConfigHTTP `json:"http"`
	// Any of "http", "mcp", "openapi", "bare".
	Type string `json:"type"`
	// This field is from variant [ToolSpecConfigMCP].
	MCP ConfigMCP `json:"mcp"`
	// This field is from variant [ToolSpecConfigOpenAPI].
	OpenAPI ConfigOpenAPI `json:"openapi"`
	// This field is from variant [ToolSpecConfigBare].
	Bare ConfigBare `json:"bare"`
	JSON struct {
		HTTP    respjson.Field
		Type    respjson.Field
		MCP     respjson.Field
		OpenAPI respjson.Field
		Bare    respjson.Field
		raw     string
	} `json:"-"`
}

// anyToolSpecConfig is implemented by each variant of [ToolSpecConfigUnion] to add
// type safety for the return type of [ToolSpecConfigUnion.AsAny]
type anyToolSpecConfig interface {
	implToolSpecConfigUnion()
}

func (ToolSpecConfigHTTP) implToolSpecConfigUnion()    {}
func (ToolSpecConfigMCP) implToolSpecConfigUnion()     {}
func (ToolSpecConfigOpenAPI) implToolSpecConfigUnion() {}
func (ToolSpecConfigBare) implToolSpecConfigUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ToolSpecConfigUnion.AsAny().(type) {
//	case cadenya.ToolSpecConfigHTTP:
//	case cadenya.ToolSpecConfigMCP:
//	case cadenya.ToolSpecConfigOpenAPI:
//	case cadenya.ToolSpecConfigBare:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ToolSpecConfigUnion) AsAny() anyToolSpecConfig {
	switch u.Type {
	case "http":
		return u.AsHTTP()
	case "mcp":
		return u.AsMCP()
	case "openapi":
		return u.AsOpenAPI()
	case "bare":
		return u.AsBare()
	}
	return nil
}

func (u ToolSpecConfigUnion) AsHTTP() (v ToolSpecConfigHTTP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSpecConfigUnion) AsMCP() (v ToolSpecConfigMCP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSpecConfigUnion) AsOpenAPI() (v ToolSpecConfigOpenAPI) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSpecConfigUnion) AsBare() (v ToolSpecConfigBare) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolSpecConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolSpecConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSpecConfigUnion to a ToolSpecConfigUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSpecConfigUnionParam.Overrides()
func (r ToolSpecConfigUnion) ToParam() ToolSpecConfigUnionParam {
	return param.Override[ToolSpecConfigUnionParam](json.RawMessage(r.RawJSON()))
}

func ToolSpecConfigParamOfHTTP(http ConfigHTTPParam) ToolSpecConfigUnionParam {
	var variant ToolSpecConfigHTTPParam
	variant.HTTP = http
	return ToolSpecConfigUnionParam{OfHTTP: &variant}
}

func ToolSpecConfigParamOfMCP(mcp ConfigMCPParam) ToolSpecConfigUnionParam {
	var variant ToolSpecConfigMCPParam
	variant.MCP = mcp
	return ToolSpecConfigUnionParam{OfMCP: &variant}
}

func ToolSpecConfigParamOfOpenAPI(openAPI ConfigOpenAPIParam) ToolSpecConfigUnionParam {
	var variant ToolSpecConfigOpenAPIParam
	variant.OpenAPI = openAPI
	return ToolSpecConfigUnionParam{OfOpenAPI: &variant}
}

func ToolSpecConfigParamOfBare(bare ConfigBareParam) ToolSpecConfigUnionParam {
	var variant ToolSpecConfigBareParam
	variant.Bare = bare
	return ToolSpecConfigUnionParam{OfBare: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolSpecConfigUnionParam struct {
	OfHTTP    *ToolSpecConfigHTTPParam    `json:",omitzero,inline"`
	OfMCP     *ToolSpecConfigMCPParam     `json:",omitzero,inline"`
	OfOpenAPI *ToolSpecConfigOpenAPIParam `json:",omitzero,inline"`
	OfBare    *ToolSpecConfigBareParam    `json:",omitzero,inline"`
	paramUnion
}

func (u ToolSpecConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfHTTP, u.OfMCP, u.OfOpenAPI, u.OfBare)
}
func (u *ToolSpecConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ToolSpecConfigUnionParam](
		"type",
		apijson.Discriminator[ToolSpecConfigHTTPParam]("http"),
		apijson.Discriminator[ToolSpecConfigMCPParam]("mcp"),
		apijson.Discriminator[ToolSpecConfigOpenAPIParam]("openapi"),
		apijson.Discriminator[ToolSpecConfigBareParam]("bare"),
	)
}

type ToolSpecConfigBare struct {
	// Marks the tool as bare: it has no execution adapter of its own and relies on the
	// parent tool set being a Bare tool set. Present so a webhook consumer can tell a
	// tool is bare from the tool data alone, without cross-referencing the tool set.
	Bare ConfigBare `json:"bare" api:"required"`
	// Any of "bare".
	Type ToolSpecConfigBareType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bare        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSpecConfigBare) RawJSON() string { return r.JSON.raw }
func (r *ToolSpecConfigBare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSpecConfigBare to a ToolSpecConfigBareParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSpecConfigBareParam.Overrides()
func (r ToolSpecConfigBare) ToParam() ToolSpecConfigBareParam {
	return param.Override[ToolSpecConfigBareParam](json.RawMessage(r.RawJSON()))
}

type ToolSpecConfigBareType string

const (
	ToolSpecConfigBareTypeBare ToolSpecConfigBareType = "bare"
)

// The properties Bare, Type are required.
type ToolSpecConfigBareParam struct {
	// Marks the tool as bare: it has no execution adapter of its own and relies on the
	// parent tool set being a Bare tool set. Present so a webhook consumer can tell a
	// tool is bare from the tool data alone, without cross-referencing the tool set.
	Bare ConfigBareParam `json:"bare,omitzero" api:"required"`
	// Any of "bare".
	Type ToolSpecConfigBareType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolSpecConfigBareParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSpecConfigBareParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSpecConfigBareParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSpecConfigHTTP struct {
	HTTP ConfigHTTP `json:"http" api:"required"`
	// Any of "http".
	Type ToolSpecConfigHTTPType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTTP        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSpecConfigHTTP) RawJSON() string { return r.JSON.raw }
func (r *ToolSpecConfigHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSpecConfigHTTP to a ToolSpecConfigHTTPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSpecConfigHTTPParam.Overrides()
func (r ToolSpecConfigHTTP) ToParam() ToolSpecConfigHTTPParam {
	return param.Override[ToolSpecConfigHTTPParam](json.RawMessage(r.RawJSON()))
}

type ToolSpecConfigHTTPType string

const (
	ToolSpecConfigHTTPTypeHTTP ToolSpecConfigHTTPType = "http"
)

// The properties HTTP, Type are required.
type ToolSpecConfigHTTPParam struct {
	HTTP ConfigHTTPParam `json:"http,omitzero" api:"required"`
	// Any of "http".
	Type ToolSpecConfigHTTPType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolSpecConfigHTTPParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSpecConfigHTTPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSpecConfigHTTPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSpecConfigMCP struct {
	MCP ConfigMCP `json:"mcp" api:"required"`
	// Any of "mcp".
	Type ToolSpecConfigMCPType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MCP         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSpecConfigMCP) RawJSON() string { return r.JSON.raw }
func (r *ToolSpecConfigMCP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSpecConfigMCP to a ToolSpecConfigMCPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSpecConfigMCPParam.Overrides()
func (r ToolSpecConfigMCP) ToParam() ToolSpecConfigMCPParam {
	return param.Override[ToolSpecConfigMCPParam](json.RawMessage(r.RawJSON()))
}

type ToolSpecConfigMCPType string

const (
	ToolSpecConfigMCPTypeMCP ToolSpecConfigMCPType = "mcp"
)

// The properties MCP, Type are required.
type ToolSpecConfigMCPParam struct {
	MCP ConfigMCPParam `json:"mcp,omitzero" api:"required"`
	// Any of "mcp".
	Type ToolSpecConfigMCPType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolSpecConfigMCPParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSpecConfigMCPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSpecConfigMCPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSpecConfigOpenAPI struct {
	OpenAPI ConfigOpenAPI `json:"openapi" api:"required"`
	// Any of "openapi".
	Type ToolSpecConfigOpenAPIType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenAPI     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSpecConfigOpenAPI) RawJSON() string { return r.JSON.raw }
func (r *ToolSpecConfigOpenAPI) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSpecConfigOpenAPI to a ToolSpecConfigOpenAPIParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSpecConfigOpenAPIParam.Overrides()
func (r ToolSpecConfigOpenAPI) ToParam() ToolSpecConfigOpenAPIParam {
	return param.Override[ToolSpecConfigOpenAPIParam](json.RawMessage(r.RawJSON()))
}

type ToolSpecConfigOpenAPIType string

const (
	ToolSpecConfigOpenAPITypeOpenAPI ToolSpecConfigOpenAPIType = "openapi"
)

// The properties OpenAPI, Type are required.
type ToolSpecConfigOpenAPIParam struct {
	OpenAPI ConfigOpenAPIParam `json:"openapi,omitzero" api:"required"`
	// Any of "openapi".
	Type ToolSpecConfigOpenAPIType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolSpecConfigOpenAPIParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSpecConfigOpenAPIParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSpecConfigOpenAPIParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetToolNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	Spec     ToolSpecParam                      `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r ToolSetToolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetToolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetToolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetToolGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetToolUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	UpdateMask  param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	Spec     ToolSpecParam                      `json:"spec,omitzero"`
	paramObj
}

func (r ToolSetToolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetToolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetToolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetToolListParams struct {
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
	// Filter by approval requirement. Omitted = no filter; true = only tools requiring
	// approval; false = only tools not requiring approval.
	RequiresApproval param.Opt[bool] `query:"requiresApproval,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// Filter by tool name (exact match). Multiple values are OR'd together.
	Names []string `query:"names,omitzero" json:"-"`
	// Filter by tool state. Multiple values are OR'd together.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_AVAILABLE", "STATE_OMITTED",
	// "STATE_ARCHIVED".
	States []string `query:"states,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolSetToolListParams]'s query parameters as `url.Values`.
func (r ToolSetToolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolSetToolDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetToolOmitParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ToolSetToolOmitParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetToolOmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetToolOmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetToolRestoreParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ToolSetToolRestoreParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetToolRestoreParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetToolRestoreParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
