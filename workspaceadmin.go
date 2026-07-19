// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"go.cadenya.com/cadenya-go/internal/apijson"
	"go.cadenya.com/cadenya-go/internal/apiquery"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/packages/pagination"
	"go.cadenya.com/cadenya-go/packages/param"
	"go.cadenya.com/cadenya-go/packages/respjson"
	"net/http"
	"net/url"
	"slices"
	"time"
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
	options []option.RequestOption
	// Administer workspaces across the account: create and archive workspaces and
	// manage their membership. These operations are account-scoped and require the
	// admin role (a token whose profile holds the WorkOS admin role); they live under
	// /v1/account/workspaces rather than the workspace-scoped /v1/workspaces tree so
	// an admin can manage any workspace in the account, including ones they are not
	// themselves a member of.
	Members WorkspaceAdminMemberService
	// Administer workspaces across the account: create and archive workspaces and
	// manage their membership. These operations are account-scoped and require the
	// admin role (a token whose profile holds the WorkOS admin role); they live under
	// /v1/account/workspaces rather than the workspace-scoped /v1/workspaces tree so
	// an admin can manage any workspace in the account, including ones they are not
	// themselves a member of.
	Profiles WorkspaceAdminProfileService
}

// NewWorkspaceAdminService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceAdminService(opts ...option.RequestOption) (r WorkspaceAdminService) {
	r = WorkspaceAdminService{}
	r.options = opts
	r.Members = NewWorkspaceAdminMemberService(opts...)
	r.Profiles = NewWorkspaceAdminProfileService(opts...)
	return
}

// Creates a new workspace in the account. Admin only.
func (r *WorkspaceAdminService) New(ctx context.Context, body WorkspaceAdminNewParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.options, opts)
	path := "v1/account/workspaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a workspace in the account by ID. Admin only.
func (r *WorkspaceAdminService) Get(ctx context.Context, query WorkspaceAdminGetParams, opts ...option.RequestOption) (res *Workspace, err error) {
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
	path := fmt.Sprintf("v1/account/workspaces/%s", url.PathEscape(query.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a workspace's metadata (e.g. name) and spec. Admin only.
func (r *WorkspaceAdminService) Update(ctx context.Context, params WorkspaceAdminUpdateParams, opts ...option.RequestOption) (res *Workspace, err error) {
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
	path := fmt.Sprintf("v1/account/workspaces/%s", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists every workspace in the account, optionally including archived ones. Admin
// only.
func (r *WorkspaceAdminService) List(ctx context.Context, query WorkspaceAdminListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Workspace], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
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
// subsequent request scoped to it returns a permission error. Archiving the
// account's last active (non-archived) workspace is not allowed and returns
// FailedPrecondition. Admin only.
func (r *WorkspaceAdminService) Archive(ctx context.Context, body WorkspaceAdminArchiveParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/account/workspaces/%s", url.PathEscape(body.WorkspaceID.Value))
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
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActorID     respjson.Field
		ProfileID   respjson.Field
		AddedAt     respjson.Field
		Email       respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceMember) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceMember) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceAdminNewParams struct {
	// CreateAccountResourceMetadata contains the user-provided fields for creating an
	// account-scoped resource. Read-only fields (id, account_id, profile_id) are
	// excluded since they are set by the server.
	Metadata WorkspaceAdminNewParamsMetadata `json:"metadata,omitzero" api:"required"`
	Spec     WorkspaceSpecParam              `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r WorkspaceAdminNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceAdminNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceAdminNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreateAccountResourceMetadata contains the user-provided fields for creating an
// account-scoped resource. Read-only fields (id, account_id, profile_id) are
// excluded since they are set by the server.
//
// The property Name is required.
type WorkspaceAdminNewParamsMetadata struct {
	// Human-readable name for the resource (e.g., "Production API Key", "Staging
	// Workspace")
	Name string `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"environment": "production", "team":
	// "platform", "version": "v2"}
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r WorkspaceAdminNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceAdminNewParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceAdminNewParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceAdminGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type WorkspaceAdminUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Fields to update.
	UpdateMask param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateAccountResourceMetadata contains the user-provided fields for updating an
	// account-scoped resource. Read-only fields (id, account_id, profile_id) are
	// excluded since they are set by the server.
	Metadata WorkspaceAdminUpdateParamsMetadata `json:"metadata,omitzero"`
	Spec     WorkspaceSpecParam                 `json:"spec,omitzero"`
	paramObj
}

func (r WorkspaceAdminUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceAdminUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceAdminUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpdateAccountResourceMetadata contains the user-provided fields for updating an
// account-scoped resource. Read-only fields (id, account_id, profile_id) are
// excluded since they are set by the server.
//
// The property Name is required.
type WorkspaceAdminUpdateParamsMetadata struct {
	// Human-readable name for the resource (e.g., "Production API Key", "Staging
	// Workspace")
	Name string `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"environment": "production", "team":
	// "platform", "version": "v2"}
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r WorkspaceAdminUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceAdminUpdateParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceAdminUpdateParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceAdminListParams struct {
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, archived workspaces are included in the results. Defaults to false
	// (active workspaces only).
	IncludeArchived param.Opt[bool] `query:"includeArchived,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceAdminListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceAdminListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceAdminArchiveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
