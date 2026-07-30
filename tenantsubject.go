// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"go.cadenya.com/cadenya-go/internal/apiquery"
	"go.cadenya.com/cadenya-go/internal/requestconfig"
	"go.cadenya.com/cadenya-go/option"
	"go.cadenya.com/cadenya-go/packages/pagination"
	"go.cadenya.com/cadenya-go/packages/param"
	"net/http"
	"net/url"
	"slices"
)

// Read and erase tenants and the subjects under them. Tenants and subjects are
// created by assertion — on objective creation or widget session mint — never
// directly, so this service has no create or update: it exists to enumerate what
// assertions have produced, and to destroy it on request.
//
// TenantSubjectService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantSubjectService] method instead.
type TenantSubjectService struct {
	options []option.RequestOption
}

// NewTenantSubjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTenantSubjectService(opts ...option.RequestOption) (r TenantSubjectService) {
	r = TenantSubjectService{}
	r.options = opts
	return
}

// Lists the subjects asserted under a tenant. Subjects are only listable through
// their tenant: a subject's external_id is unique within its tenant, not across
// the workspace, so the same key can name different people under different
// tenants.
func (r *TenantSubjectService) List(ctx context.Context, tenantID string, params TenantSubjectListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Subject], err error) {
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
	if tenantID == "" {
		err = errors.New("missing required tenantId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tenants/%s/subjects", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(tenantID))
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

// Lists the subjects asserted under a tenant. Subjects are only listable through
// their tenant: a subject's external_id is unique within its tenant, not across
// the workspace, so the same key can name different people under different
// tenants.
func (r *TenantSubjectService) ListAutoPaging(ctx context.Context, tenantID string, params TenantSubjectListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Subject] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, tenantID, params, opts...))
}

type TenantSubjectListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, the `info` field on each returned subject is populated.
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Substring match against the subject's name and external_id.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time).
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantSubjectListParams]'s query parameters as
// `url.Values`.
func (r TenantSubjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
