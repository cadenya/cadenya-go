// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
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
	Options []option.RequestOption
}

// NewWorkspaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceService(opts ...option.RequestOption) (r *WorkspaceService) {
	r = &WorkspaceService{}
	r.Options = opts
	return
}

// Lists all workspaces for the current account
func (r *WorkspaceService) List(ctx context.Context, query WorkspaceListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Workspace], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
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

// Retrieves the workspace associated with the current API token. Useful for
// workspace-scoped tokens to identify which workspace they belong to.
func (r *WorkspaceService) Get(ctx context.Context, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspaces/current"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
	Status WorkspaceStatus `json:"status"`
	JSON   workspaceJSON   `json:"-"`
}

// workspaceJSON contains the JSON metadata for the struct [Workspace]
type workspaceJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Workspace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceJSON) RawJSON() string {
	return r.raw
}

// WorkspaceInfo returns counts
type WorkspaceInfo struct {
	TotalAgents          int64             `json:"totalAgents"`
	TotalAgentVariations int64             `json:"totalAgentVariations"`
	TotalAvailableTools  int64             `json:"totalAvailableTools"`
	TotalMemoryEntries   int64             `json:"totalMemoryEntries"`
	JSON                 workspaceInfoJSON `json:"-"`
}

// workspaceInfoJSON contains the JSON metadata for the struct [WorkspaceInfo]
type workspaceInfoJSON struct {
	TotalAgents          apijson.Field
	TotalAgentVariations apijson.Field
	TotalAvailableTools  apijson.Field
	TotalMemoryEntries   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *WorkspaceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceInfoJSON) RawJSON() string {
	return r.raw
}

// Lifecycle status of the workspace. Archived workspaces reject all requests
// scoped to them. Server-populated.
type WorkspaceStatus string

const (
	WorkspaceStatusStatusEnabled  WorkspaceStatus = "STATUS_ENABLED"
	WorkspaceStatusStatusDisabled WorkspaceStatus = "STATUS_DISABLED"
	WorkspaceStatusStatusArchived WorkspaceStatus = "STATUS_ARCHIVED"
)

func (r WorkspaceStatus) IsKnown() bool {
	switch r {
	case WorkspaceStatusStatusEnabled, WorkspaceStatusStatusDisabled, WorkspaceStatusStatusArchived:
		return true
	}
	return false
}

type WorkspaceSpec struct {
	Description string            `json:"description"`
	JSON        workspaceSpecJSON `json:"-"`
}

// workspaceSpecJSON contains the JSON metadata for the struct [WorkspaceSpec]
type workspaceSpecJSON struct {
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceSpecJSON) RawJSON() string {
	return r.raw
}

type WorkspaceSpecParam struct {
	Description param.Field[string] `json:"description"`
}

func (r WorkspaceSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkspaceListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [WorkspaceListParams]'s query parameters as `url.Values`.
func (r WorkspaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
