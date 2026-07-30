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
	"go.cadenya.com/cadenya-go/shared"
	"net/http"
	"net/url"
	"slices"
)

// Read and erase tenants and the subjects under them. Tenants and subjects are
// created by assertion — on objective creation or widget session mint — never
// directly, so this service has no create or update: it exists to enumerate what
// assertions have produced, and to destroy it on request.
//
// TenantService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenantService] method instead.
type TenantService struct {
	options []option.RequestOption
	// Read and erase tenants and the subjects under them. Tenants and subjects are
	// created by assertion — on objective creation or widget session mint — never
	// directly, so this service has no create or update: it exists to enumerate what
	// assertions have produced, and to destroy it on request.
	Subjects TenantSubjectService
}

// NewTenantService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenantService(opts ...option.RequestOption) (r TenantService) {
	r = TenantService{}
	r.options = opts
	r.Subjects = NewTenantSubjectService(opts...)
	return
}

// Retrieves a tenant by its canonical id or by the `external_id:<value>` form the
// customer asserted it under.
func (r *TenantService) Get(ctx context.Context, id string, params TenantGetParams, opts ...option.RequestOption) (res *Tenant, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tenants/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Lists the tenants asserted in a workspace, newest first. `query` matches against
// a tenant's name and its external_id, for type-ahead filters.
func (r *TenantService) List(ctx context.Context, params TenantListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Tenant], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tenants", url.PathEscape(params.WorkspaceID.Value))
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

// Lists the tenants asserted in a workspace, newest first. `query` matches against
// a tenant's name and its external_id, for type-ahead filters.
func (r *TenantService) ListAutoPaging(ctx context.Context, params TenantListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Tenant] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Destroys the tenant, its subjects, every objective associated with it and
// everything reachable from those objectives, and its widget sessions. This is the
// full erasure hammer, wider than `DELETE /widget_sessions`, which removes only
// what widget sessions created. The work runs in the background: this returns the
// tenant in STATE_ERASING rather than a count of what was removed, since a large
// tenant's history cannot be destroyed inside a request. Poll the tenant to follow
// it — STATE_ERASING while it runs, NotFound once it finishes. Erasure is
// terminal; a tenant cannot be recovered once it starts.
func (r *TenantService) Delete(ctx context.Context, id string, body TenantDeleteParams, opts ...option.RequestOption) (res *Tenant, err error) {
	opts = slices.Concat(r.options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	requestconfig.UseDefaultParam(&body.WorkspaceID, precfg.WorkspaceID)
	if body.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tenants/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Subject is a person within a tenant as a readable record. Like Tenant it carries
// no spec — `metadata.external_id` is the customer's key for them, unique within
// the tenant rather than the workspace.
type Subject struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// SubjectInfo provides read-only server-derived data about a subject.
	Info SubjectInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subject) RawJSON() string { return r.JSON.raw }
func (r *Subject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SubjectInfo provides read-only server-derived data about a subject.
type SubjectInfo struct {
	// Number of objectives associated with this subject.
	ObjectiveCount int64 `json:"objectiveCount"`
	// TenantReference is the read-only echo of a resource's tenant association,
	// carrying both Cadenya's canonical id and the customer's own key.
	Tenant TenantReference `json:"tenant"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectiveCount respjson.Field
		Tenant         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubjectInfo) RawJSON() string { return r.JSON.raw }
func (r *SubjectInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tenant is the customer's organization as a readable record rather than an echo.
// It carries no spec: a tenant is never configured, only asserted, so everything
// about it lives in the metadata envelope — `external_id` is the key the customer
// asserted it under, `name` is the most recent name they asserted, and
// `updated_at` is therefore when the tenant was last asserted.
type Tenant struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// The current lifecycle state of the tenant. Output only.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ACTIVE", "STATE_ERASING".
	State TenantState `json:"state" api:"required"`
	// TenantInfo provides read-only server-derived data about a tenant.
	Info TenantInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		State       respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tenant) RawJSON() string { return r.JSON.raw }
func (r *Tenant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current lifecycle state of the tenant. Output only.
type TenantState string

const (
	TenantStateStateUnspecified TenantState = "STATE_UNSPECIFIED"
	TenantStateStateActive      TenantState = "STATE_ACTIVE"
	TenantStateStateErasing     TenantState = "STATE_ERASING"
)

// TenantInfo provides read-only server-derived data about a tenant.
type TenantInfo struct {
	// Number of objectives associated with this tenant, across every surface — widget
	// conversations and objectives created directly against the API alike. This is the
	// footprint a delete would destroy, which is why it is worth the count query that
	// populating `info` costs.
	ObjectiveCount int64 `json:"objectiveCount"`
	// Number of subjects asserted under this tenant.
	SubjectCount int64 `json:"subjectCount"`
	// Number of widget sessions minted for this tenant that still exist.
	WidgetSessionCount int64 `json:"widgetSessionCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ObjectiveCount     respjson.Field
		SubjectCount       respjson.Field
		WidgetSessionCount respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenantInfo) RawJSON() string { return r.JSON.raw }
func (r *TenantInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TenantGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// When true, the `info` field is populated.
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantGetParams]'s query parameters as `url.Values`.
func (r TenantGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, the `info` field on each returned tenant is populated. This costs
	// several count queries per tenant, so it is off by default.
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Substring match against the tenant's name and external_id. Built for type-ahead
	// filter pickers, where the operator knows the customer's own identifier rather
	// than Cadenya's.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time).
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TenantListParams]'s query parameters as `url.Values`.
func (r TenantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TenantDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
