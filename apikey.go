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
	"github.com/cadenya/cadenya-go/shared"
)

// Issue, rotate, and revoke API keys for the account, and grant or revoke each
// key's access to individual workspaces.
//
// APIKeyService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
	// Issue, rotate, and revoke API keys for the account, and grant or revoke each
	// key's access to individual workspaces.
	Access *APIKeyAccessService
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r *APIKeyService) {
	r = &APIKeyService{}
	r.Options = opts
	r.Access = NewAPIKeyAccessService(opts...)
	return
}

// Creates a new API key on the account. Optionally grants the key access to one or
// more workspaces via initial_workspace_ids.
func (r *APIKeyService) New(ctx context.Context, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/account/api_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves an API key by ID.
func (r *APIKeyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account/api_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an API key.
func (r *APIKeyService) Update(ctx context.Context, id string, body APIKeyUpdateParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account/api_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all API keys on the account.
func (r *APIKeyService) List(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[APIKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/account/api_keys"
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

// Lists all API keys on the account.
func (r *APIKeyService) ListAutoPaging(ctx context.Context, query APIKeyListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[APIKey] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Deletes an API key.
func (r *APIKeyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/account/api_keys/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Rotates an API key and returns a new token. All previous tokens for this key are
// invalidated.
func (r *APIKeyService) Rotate(ctx context.Context, id string, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/account/api_keys/%s/rotate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

// An API key for the account. Use workspace-association RPCs to grant the key
// access to specific workspaces; a key with zero workspaces is valid but cannot
// access workspace-scoped resources.
type APIKey struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata shared.AccountResourceMetadata `json:"metadata" api:"required"`
	// Configuration for an API key.
	Spec APIKeySpec `json:"spec" api:"required"`
	Info APIKeyInfo `json:"info"`
	JSON apiKeyJSON `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

type APIKeyInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Up to a small number of workspaces this key has access to, intended for display
	// ("Workspace 1, Workspace 2, and 4 more"). Use ListAPIKeyWorkspaces for the full
	// paginated list.
	WorkspacesPreview []shared.BareMetadata `json:"workspacesPreview"`
	// Total number of workspaces this key has access to.
	WorkspacesTotal int64          `json:"workspacesTotal"`
	JSON            apiKeyInfoJSON `json:"-"`
}

// apiKeyInfoJSON contains the JSON metadata for the struct [APIKeyInfo]
type apiKeyInfoJSON struct {
	CreatedBy         apijson.Field
	WorkspacesPreview apijson.Field
	WorkspacesTotal   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *APIKeyInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyInfoJSON) RawJSON() string {
	return r.raw
}

// Configuration for an API key.
type APIKeySpec struct {
	// The bearer token used to authenticate as this API key. Returned only on creation
	// and rotation; subsequent reads omit this field.
	Token string `json:"token"`
	// Free-form description of what this API key is used for.
	Description string `json:"description"`
	// Permissions granted to this key. Each entry is a colon-separated verb:resource
	// string (e.g. "manage:agents"). Currently has no enforced effect; reserved for
	// future fine-grained authorization.
	Permissions []string `json:"permissions"`
	// True when this key is managed by the system (e.g. the auto-provisioned global
	// account key). System keys cannot be deleted but can be rotated.
	System bool           `json:"system"`
	JSON   apiKeySpecJSON `json:"-"`
}

// apiKeySpecJSON contains the JSON metadata for the struct [APIKeySpec]
type apiKeySpecJSON struct {
	Token       apijson.Field
	Description apijson.Field
	Permissions apijson.Field
	System      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIKeySpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeySpecJSON) RawJSON() string {
	return r.raw
}

// Configuration for an API key.
type APIKeySpecParam struct {
	// Free-form description of what this API key is used for.
	Description param.Field[string] `json:"description"`
	// Permissions granted to this key. Each entry is a colon-separated verb:resource
	// string (e.g. "manage:agents"). Currently has no enforced effect; reserved for
	// future fine-grained authorization.
	Permissions param.Field[[]string] `json:"permissions"`
}

func (r APIKeySpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyNewParams struct {
	// CreateAccountResourceMetadata contains the user-provided fields for creating an
	// account-scoped resource. Read-only fields (id, account_id, profile_id) are
	// excluded since they are set by the server.
	Metadata param.Field[APIKeyNewParamsMetadata] `json:"metadata" api:"required"`
	// Configuration for an API key.
	Spec param.Field[APIKeySpecParam] `json:"spec" api:"required"`
	// Workspaces this API key will have access to on creation. Optional — a key can be
	// created with no workspace access and granted later via AddAPIKeyWorkspace.
	InitialWorkspaceIDs param.Field[[]string] `json:"initialWorkspaceIds"`
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// CreateAccountResourceMetadata contains the user-provided fields for creating an
// account-scoped resource. Read-only fields (id, account_id, profile_id) are
// excluded since they are set by the server.
type APIKeyNewParamsMetadata struct {
	// Human-readable name for the resource (e.g., "Production API Key", "Staging
	// Workspace")
	Name param.Field[string] `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r APIKeyNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyUpdateParams struct {
	// UpdateAccountResourceMetadata contains the user-provided fields for updating an
	// account-scoped resource. Read-only fields (id, account_id, profile_id) are
	// excluded since they are set by the server.
	Metadata param.Field[APIKeyUpdateParamsMetadata] `json:"metadata"`
	// Configuration for an API key.
	Spec param.Field[APIKeySpecParam] `json:"spec"`
	// Fields to update.
	UpdateMask param.Field[string] `json:"updateMask" format:"field-mask"`
}

func (r APIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// UpdateAccountResourceMetadata contains the user-provided fields for updating an
// account-scoped resource. Read-only fields (id, account_id, profile_id) are
// excluded since they are set by the server.
type APIKeyUpdateParamsMetadata struct {
	// Human-readable name for the resource (e.g., "Production API Key", "Staging
	// Workspace")
	Name param.Field[string] `json:"name" api:"required"`
	// External ID for the resource (e.g., a workflow ID from an external system)
	ExternalID param.Field[string] `json:"externalId"`
	// Arbitrary key-value pairs for categorization and filtering Examples:
	// {"environment": "production", "team": "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r APIKeyUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyListParams struct {
	// Filter by bundle_key — return only resources owned by this bundle.
	BundleKey param.Field[string] `query:"bundleKey"`
	// Pagination cursor from previous response.
	Cursor param.Field[string] `query:"cursor"`
	// When true, included info fields are populated. Requests with this flag count
	// more against your rate limit.
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return.
	Limit param.Field[int64] `query:"limit"`
	// Filter by ID prefix.
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query.
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time).
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
