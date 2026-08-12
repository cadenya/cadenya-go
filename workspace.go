package cadenya

import (
	"context"
	"encoding/json"
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

// Manage workspaces within an account. Workspaces provide organizational grouping
// and isolation for resources such as agents, tools, and API keys.
//
// This is the workspace-scoped, end-user surface. Administrative operations
// (create / archive workspaces, manage members) live in WorkspaceAdminService
// under /v1/account/workspaces and require the admin role.
//
// WorkspaceService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceService] method instead.
type WorkspaceService struct {
	options []option.RequestOption
}

// NewWorkspaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceService(opts ...option.RequestOption) (r WorkspaceService) {
	r = WorkspaceService{}
	r.options = opts
	return
}

// Lists all workspaces for the current account
func (r *WorkspaceService) List(ctx context.Context, query WorkspaceListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Workspace], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/workspaces"
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

// Lists all workspaces for the current account
func (r *WorkspaceService) ListAutoPaging(ctx context.Context, query WorkspaceListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Workspace] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

type Workspace struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata shared.AccountResourceMetadata `json:"metadata" api:"required"`
	Spec     WorkspaceSpec                  `json:"spec" api:"required"`
	// WorkspaceInfo returns counts
	Info WorkspaceInfo `json:"info"`
	// Lifecycle status of the workspace. Archived workspaces reject all requests
	// scoped to them. Server-populated.
	//
	// Any of "STATUS_ENABLED", "STATUS_DISABLED", "STATUS_ARCHIVED".
	Status WorkspaceStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Info        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workspace) RawJSON() string { return r.JSON.raw }
func (r *Workspace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkspaceInfo returns counts
type WorkspaceInfo struct {
	TotalAgents          int64 `json:"totalAgents"`
	TotalAgentVariations int64 `json:"totalAgentVariations"`
	TotalAvailableTools  int64 `json:"totalAvailableTools"`
	TotalMemoryEntries   int64 `json:"totalMemoryEntries"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TotalAgents          respjson.Field
		TotalAgentVariations respjson.Field
		TotalAvailableTools  respjson.Field
		TotalMemoryEntries   respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceInfo) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle status of the workspace. Archived workspaces reject all requests
// scoped to them. Server-populated.
type WorkspaceStatus string

const (
	WorkspaceStatusStatusEnabled  WorkspaceStatus = "STATUS_ENABLED"
	WorkspaceStatusStatusDisabled WorkspaceStatus = "STATUS_DISABLED"
	WorkspaceStatusStatusArchived WorkspaceStatus = "STATUS_ARCHIVED"
)

type WorkspaceSpec struct {
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceSpec) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkspaceSpec to a WorkspaceSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkspaceSpecParam.Overrides()
func (r WorkspaceSpec) ToParam() WorkspaceSpecParam {
	return param.Override[WorkspaceSpecParam](json.RawMessage(r.RawJSON()))
}

type WorkspaceSpecParam struct {
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r WorkspaceSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceListParams struct {
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

// URLQuery serializes [WorkspaceListParams]'s query parameters as `url.Values`.
func (r WorkspaceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
