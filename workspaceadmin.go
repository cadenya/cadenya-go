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
)

// Administer workspaces across the account: create and archive workspaces and
// manage their membership. These operations are account-scoped and require the
// admin role (a token whose profile holds the WorkOS admin role); they live under
// /v1/account/workspaces rather than the workspace-scoped /v1/workspaces tree so
// an admin can manage any workspace in the account, including ones they are not
// themselves a member of.
//
// WorkspaceAdminService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceAdminService] method instead.
type WorkspaceAdminService struct {
	Options []option.RequestOption
	// Administer workspaces across the account: create and archive workspaces and
	// manage their membership. These operations are account-scoped and require the
	// admin role (a token whose profile holds the WorkOS admin role); they live under
	// /v1/account/workspaces rather than the workspace-scoped /v1/workspaces tree so
	// an admin can manage any workspace in the account, including ones they are not
	// themselves a member of.
	Members *WorkspaceAdminMemberService
}

// NewWorkspaceAdminService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceAdminService(opts ...option.RequestOption) (r *WorkspaceAdminService) {
	r = &WorkspaceAdminService{}
	r.Options = opts
	r.Members = NewWorkspaceAdminMemberService(opts...)
	return
}

// Creates a new workspace in the account. Admin only.
func (r *WorkspaceAdminService) New(ctx context.Context, body WorkspaceAdminNewParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account/workspaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a workspace in the account by ID. Admin only.
func (r *WorkspaceAdminService) Get(ctx context.Context, workspaceID string, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account/workspaces/%s", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists every workspace in the account, optionally including archived ones. Admin
// only.
func (r *WorkspaceAdminService) List(ctx context.Context, query WorkspaceAdminListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Workspace], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/account/workspaces"
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

// Lists every workspace in the account, optionally including archived ones. Admin
// only.
func (r *WorkspaceAdminService) ListAutoPaging(ctx context.Context, query WorkspaceAdminListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Workspace] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Archives a workspace (soft delete). The workspace is retained, but any
// subsequent request scoped to it returns a permission error. Admin only.
func (r *WorkspaceAdminService) Archive(ctx context.Context, workspaceID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	path := fmt.Sprintf("v1/account/workspaces/%s", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// A member of a workspace: the profile granted access plus the actor row that
// links it to the workspace. Returned by member list/add operations.
type WorkspaceMember struct {
	// The actor row linking the profile to the workspace (the junction record).
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

type WorkspaceAdminNewParams struct {
	// CreateAccountResourceMetadata contains the user-provided fields for creating an
	// account-scoped resource. Read-only fields (id, account_id, profile_id) are
	// excluded since they are set by the server.
	Metadata param.Field[WorkspaceAdminNewParamsMetadata] `json:"metadata" api:"required"`
	Spec     param.Field[WorkspaceSpecParam]              `json:"spec" api:"required"`
}

func (r WorkspaceAdminNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CreateAccountResourceMetadata contains the user-provided fields for creating an
// account-scoped resource. Read-only fields (id, account_id, profile_id) are
// excluded since they are set by the server.
type WorkspaceAdminNewParamsMetadata struct {
	// Human-readable name for the resource (e.g., "Production API Key", "Staging
	// Workspace")
	Name param.Field[string] `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r WorkspaceAdminNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkspaceAdminListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When true, archived workspaces are included in the results. Defaults to false
	// (active workspaces only).
	IncludeArchived param.Field[bool] `query:"includeArchived"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [WorkspaceAdminListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceAdminListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
