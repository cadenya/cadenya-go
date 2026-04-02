// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyasdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cadenya/cadenya-sdk-go/internal/apijson"
	"github.com/cadenya/cadenya-sdk-go/internal/apiquery"
	"github.com/cadenya/cadenya-sdk-go/internal/param"
	"github.com/cadenya/cadenya-sdk-go/internal/requestconfig"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/cadenya/cadenya-sdk-go/packages/pagination"
	"github.com/cadenya/cadenya-sdk-go/shared"
)

// WorkspaceSecretService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceSecretService] method instead.
type WorkspaceSecretService struct {
	Options []option.RequestOption
}

// NewWorkspaceSecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceSecretService(opts ...option.RequestOption) (r *WorkspaceSecretService) {
	r = &WorkspaceSecretService{}
	r.Options = opts
	return
}

// Creates a new workspace secret in the workspace
func (r *WorkspaceSecretService) New(ctx context.Context, body WorkspaceSecretNewParams, opts ...option.RequestOption) (res *WorkspaceSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspace_secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a workspace secret by ID from the workspace
func (r *WorkspaceSecretService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkspaceSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspace_secrets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a workspace secret in the workspace
func (r *WorkspaceSecretService) Update(ctx context.Context, id string, body WorkspaceSecretUpdateParams, opts ...option.RequestOption) (res *WorkspaceSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspace_secrets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all workspace secrets in the workspace
func (r *WorkspaceSecretService) List(ctx context.Context, query WorkspaceSecretListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[WorkspaceSecret], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/workspace_secrets"
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

// Lists all workspace secrets in the workspace
func (r *WorkspaceSecretService) ListAutoPaging(ctx context.Context, query WorkspaceSecretListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[WorkspaceSecret] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes a workspace secret from the workspace
func (r *WorkspaceSecretService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspace_secrets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type WorkspaceSecret struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     WorkspaceSecretSpec     `json:"spec" api:"required"`
	// Workspace secret information
	Info WorkspaceSecretInfo `json:"info"`
	JSON workspaceSecretJSON `json:"-"`
}

// workspaceSecretJSON contains the JSON metadata for the struct [WorkspaceSecret]
type workspaceSecretJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceSecretJSON) RawJSON() string {
	return r.raw
}

type WorkspaceSecretInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy  shared.Profile          `json:"createdBy"`
	LastUsedAt time.Time               `json:"lastUsedAt" format:"date-time"`
	JSON       workspaceSecretInfoJSON `json:"-"`
}

// workspaceSecretInfoJSON contains the JSON metadata for the struct
// [WorkspaceSecretInfo]
type workspaceSecretInfoJSON struct {
	CreatedBy   apijson.Field
	LastUsedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceSecretInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceSecretInfoJSON) RawJSON() string {
	return r.raw
}

type WorkspaceSecretSpec struct {
	Value string                  `json:"value"`
	JSON  workspaceSecretSpecJSON `json:"-"`
}

// workspaceSecretSpecJSON contains the JSON metadata for the struct
// [WorkspaceSecretSpec]
type workspaceSecretSpecJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkspaceSecretSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workspaceSecretSpecJSON) RawJSON() string {
	return r.raw
}

type WorkspaceSecretSpecParam struct {
	Value param.Field[string] `json:"value"`
}

func (r WorkspaceSecretSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkspaceSecretNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	Spec     param.Field[WorkspaceSecretSpecParam]           `json:"spec" api:"required"`
}

func (r WorkspaceSecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkspaceSecretUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	Spec     param.Field[WorkspaceSecretSpecParam]           `json:"spec"`
	// Fields to update
	UpdateMask param.Field[string] `json:"updateMask" format:"field-mask"`
}

func (r WorkspaceSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkspaceSecretListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [WorkspaceSecretListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceSecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
