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
	Options []option.RequestOption
}

// NewWorkspaceAdminMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceAdminMemberService(opts ...option.RequestOption) (r *WorkspaceAdminMemberService) {
	r = &WorkspaceAdminMemberService{}
	r.Options = opts
	return
}

// Lists the members of a workspace. Admin only.
func (r *WorkspaceAdminMemberService) List(ctx context.Context, workspaceID string, query WorkspaceAdminMemberListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[WorkspaceMember], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account/workspaces/%s/members", workspaceID)
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

// Lists the members of a workspace. Admin only.
func (r *WorkspaceAdminMemberService) ListAutoPaging(ctx context.Context, workspaceID string, query WorkspaceAdminMemberListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[WorkspaceMember] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Grants a profile access to the workspace by creating (or reactivating) the actor
// that links the profile to the workspace. Accepts either an existing profile_id
// or an email to resolve-or-invite. Idempotent for an already-active member. Admin
// only.
func (r *WorkspaceAdminMemberService) Add(ctx context.Context, workspaceID string, body WorkspaceAdminMemberAddParams, opts ...option.RequestOption) (res *WorkspaceMember, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account/workspaces/%s/members", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Revokes a member's access by deactivating their actor; the member is immediately
// cut off. The underlying profile is not deleted. Admin only.
func (r *WorkspaceAdminMemberService) Remove(ctx context.Context, workspaceID string, profileID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if profileID == "" {
		err = errors.New("missing required profileId parameter")
		return err
	}
	path := fmt.Sprintf("v1/account/workspaces/%s/members/%s", workspaceID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WorkspaceAdminMemberListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [WorkspaceAdminMemberListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceAdminMemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceAdminMemberAddParams struct {
	// Email address to add (resolve-or-invite). Mutually exclusive with profile_id.
	Email param.Field[string] `json:"email"`
	// An existing account profile to add. Mutually exclusive with email.
	ProfileID param.Field[string] `json:"profileId"`
}

func (r WorkspaceAdminMemberAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
