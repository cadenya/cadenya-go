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

// Manage workspaces within an account. Workspaces provide organizational grouping
// and isolation for resources such as agents, tools, and API keys. Workspace
// creation, archival, and membership management require an account administrator
// (a token whose profile holds the admin role).
//
// WorkspaceService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceService] method instead.
type WorkspaceService struct {
	Options []option.RequestOption
	// Manage workspaces within an account. Workspaces provide organizational grouping
	// and isolation for resources such as agents, tools, and API keys. Workspace
	// creation, archival, and membership management require an account administrator
	// (a token whose profile holds the admin role).
	Members *WorkspaceMemberService
}

// NewWorkspaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceService(opts ...option.RequestOption) (r *WorkspaceService) {
	r = &WorkspaceService{}
	r.Options = opts
	r.Members = NewWorkspaceMemberService(opts...)
	return
}

// Creates a new workspace in the current account. Requires the admin role.
func (r *WorkspaceService) New(ctx context.Context, body WorkspaceNewParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a workspace by ID from the current account.
func (r *WorkspaceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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

// Archives a workspace. This is a soft delete: the workspace is retained but any
// subsequent request scoped to it returns a permission error. Requires the admin
// role.
func (r *WorkspaceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieves the workspace associated with the current API token. Useful for
// workspace-scoped tokens to identify which workspace they belong to.
func (r *WorkspaceService) GetCurrent(ctx context.Context, opts ...option.RequestOption) (res *Workspace, err error) {
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
	// Lifecycle status of the workspace. Archived workspaces reject all requests
	// scoped to them. Server-populated.
	Status WorkspaceStatus `json:"status"`
	JSON   workspaceJSON   `json:"-"`
}

// workspaceJSON contains the JSON metadata for the struct [Workspace]
type workspaceJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
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

// A member of a workspace: the profile granted access plus the actor row that
// links it to the workspace. Returned by member list/add operations.
type WorkspaceMember struct {
	// The actor row linking the profile to the workspace (the junction record). This
	// is the id used to remove the member.
	ActorID string `json:"actorId" api:"required"`
	// The account profile that has access to the workspace.
	ProfileID string `json:"profileId" api:"required"`
	// When the member was added to the workspace.
	AddedAt time.Time `json:"addedAt" format:"date-time"`
	// Email address of the member's profile.
	Email string `json:"email"`
	// Display name of the member's profile.
	Name string              `json:"name"`
	JSON workspaceMemberJSON `json:"-"`
}

// workspaceMemberJSON contains the JSON metadata for the struct [WorkspaceMember]
type workspaceMemberJSON struct {
	ActorID     apijson.Field
	ProfileID   apijson.Field
	AddedAt     apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceMemberJSON) RawJSON() string {
	return r.raw
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

type WorkspaceNewParams struct {
	// CreateAccountResourceMetadata contains the user-provided fields for creating an
	// account-scoped resource. Read-only fields (id, account_id, profile_id) are
	// excluded since they are set by the server.
	Metadata param.Field[WorkspaceNewParamsMetadata] `json:"metadata" api:"required"`
	Spec     param.Field[WorkspaceSpecParam]         `json:"spec" api:"required"`
}

func (r WorkspaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CreateAccountResourceMetadata contains the user-provided fields for creating an
// account-scoped resource. Read-only fields (id, account_id, profile_id) are
// excluded since they are set by the server.
type WorkspaceNewParamsMetadata struct {
	// Human-readable name for the resource (e.g., "Production API Key", "Staging
	// Workspace")
	Name param.Field[string] `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r WorkspaceNewParamsMetadata) MarshalJSON() (data []byte, err error) {
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
