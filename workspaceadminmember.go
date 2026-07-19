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
	"net/http"
	"net/url"
	"slices"
)

// Administer workspaces across the account: create and archive workspaces and
// manage their membership. These operations are account-scoped and require the
// admin role (a token whose profile holds the WorkOS admin role); they live under
// /v1/account/workspaces rather than the workspace-scoped /v1/workspaces tree so
// an admin can manage any workspace in the account, including ones they are not
// themselves a member of.
//
// WorkspaceAdminMemberService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceAdminMemberService] method instead.
type WorkspaceAdminMemberService struct {
	options []option.RequestOption
}

// NewWorkspaceAdminMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceAdminMemberService(opts ...option.RequestOption) (r WorkspaceAdminMemberService) {
	r = WorkspaceAdminMemberService{}
	r.options = opts
	return
}

// Lists the members of a workspace. Admin only.
func (r *WorkspaceAdminMemberService) List(ctx context.Context, params WorkspaceAdminMemberListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[WorkspaceMember], err error) {
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
	path := fmt.Sprintf("v1/account/workspaces/%s/members", url.PathEscape(params.WorkspaceID.Value))
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

// Lists the members of a workspace. Admin only.
func (r *WorkspaceAdminMemberService) ListAutoPaging(ctx context.Context, params WorkspaceAdminMemberListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[WorkspaceMember] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Grants a profile access to the workspace by creating (or reactivating) the actor
// that links the profile to the workspace. Accepts either an existing profile_id
// or an email to resolve-or-invite. Idempotent for an already-active member. Admin
// only.
func (r *WorkspaceAdminMemberService) Add(ctx context.Context, params WorkspaceAdminMemberAddParams, opts ...option.RequestOption) (res *WorkspaceMember, err error) {
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
	path := fmt.Sprintf("v1/account/workspaces/%s/members", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Revokes a member's access by deactivating their actor; the member is immediately
// cut off. The underlying profile is not deleted. Admin only.
func (r *WorkspaceAdminMemberService) Remove(ctx context.Context, profileID string, body WorkspaceAdminMemberRemoveParams, opts ...option.RequestOption) (err error) {
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
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return err
	}
	path := fmt.Sprintf("v1/account/workspaces/%s/members/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(profileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WorkspaceAdminMemberListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceAdminMemberListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceAdminMemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceAdminMemberAddParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Email address to add (resolve-or-invite). Mutually exclusive with profile_id.
	Email param.Opt[string] `json:"email,omitzero"`
	// An existing account profile to add. Mutually exclusive with email.
	ProfileID param.Opt[string] `json:"profileId,omitzero"`
	paramObj
}

func (r WorkspaceAdminMemberAddParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceAdminMemberAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceAdminMemberAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceAdminMemberRemoveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
