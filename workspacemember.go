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

// Manage workspaces within an account. Workspaces provide organizational grouping
// and isolation for resources such as agents, tools, and API keys. Workspace
// creation, archival, and membership management require an account administrator
// (a token whose profile holds the admin role).
//
// WorkspaceMemberService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceMemberService] method instead.
type WorkspaceMemberService struct {
	Options []option.RequestOption
}

// NewWorkspaceMemberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceMemberService(opts ...option.RequestOption) (r *WorkspaceMemberService) {
	r = &WorkspaceMemberService{}
	r.Options = opts
	return
}

// Lists the members (actors) of a workspace. Requires the admin role.
func (r *WorkspaceMemberService) List(ctx context.Context, workspaceID string, query WorkspaceMemberListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[WorkspaceMember], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/members", workspaceID)
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

// Lists the members (actors) of a workspace. Requires the admin role.
func (r *WorkspaceMemberService) ListAutoPaging(ctx context.Context, workspaceID string, query WorkspaceMemberListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[WorkspaceMember] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Grants a profile access to the workspace by creating an actor that links the
// profile to the workspace. Idempotent — re-adding an active member is a no-op.
// Requires the admin role.
func (r *WorkspaceMemberService) Add(ctx context.Context, workspaceID string, body WorkspaceMemberAddParams, opts ...option.RequestOption) (res *WorkspaceMember, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/members", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Revokes a member's access to the workspace by deactivating their actor. The
// member is immediately cut off; the underlying profile is not deleted. Requires
// the admin role.
func (r *WorkspaceMemberService) Remove(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/members/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WorkspaceMemberListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [WorkspaceMemberListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceMemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceMemberAddParams struct {
	// The existing account profile to add to the workspace.
	ProfileID param.Field[string] `json:"profileId"`
}

func (r WorkspaceMemberAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
