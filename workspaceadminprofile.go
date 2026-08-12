package cadenya

import (
	"context"
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
// WorkspaceAdminProfileService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceAdminProfileService] method instead.
type WorkspaceAdminProfileService struct {
	options []option.RequestOption
}

// NewWorkspaceAdminProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceAdminProfileService(opts ...option.RequestOption) (r WorkspaceAdminProfileService) {
	r = WorkspaceAdminProfileService{}
	r.options = opts
	return
}

// Searches the account's profiles for a member picker, with free-form name/email
// search and an optional type filter. Account-scoped; admin only.
func (r *WorkspaceAdminProfileService) List(ctx context.Context, query WorkspaceAdminProfileListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Profile], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/account/profiles"
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

// Searches the account's profiles for a member picker, with free-form name/email
// search and an optional type filter. Account-scoped; admin only.
func (r *WorkspaceAdminProfileService) ListAutoPaging(ctx context.Context, query WorkspaceAdminProfileListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Profile] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

type WorkspaceAdminProfileListParams struct {
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Free-form search over profile name and email. Case-insensitive substring match;
	// empty returns all profiles.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceAdminProfileListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceAdminProfileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
