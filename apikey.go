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

// Issue, rotate, disable, and revoke a workspace's API keys. Every key belongs to
// exactly one workspace; the system-managed global account key is managed via
// GlobalAPIKeyService instead.
//
// APIKeyService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAPIKeyService] method instead.
type APIKeyService struct {
	Options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r *APIKeyService) {
	r = &APIKeyService{}
	r.Options = opts
	return
}

// Creates a new API key in the workspace.
func (r *APIKeyService) New(ctx context.Context, workspaceID string, body APIKeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves an API key by ID.
func (r *APIKeyService) Get(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an API key.
func (r *APIKeyService) Update(ctx context.Context, workspaceID string, id string, body APIKeyUpdateParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists the workspace's API keys.
func (r *APIKeyService) List(ctx context.Context, workspaceID string, query APIKeyListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[APIKey], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys", workspaceID)
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

// Lists the workspace's API keys.
func (r *APIKeyService) ListAutoPaging(ctx context.Context, workspaceID string, query APIKeyListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[APIKey] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Deletes an API key.
func (r *APIKeyService) Delete(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Disables an API key. While disabled, presenting the key's token fails
// authentication on every endpoint; the key is retained. Idempotent.
func (r *APIKeyService) Disable(ctx context.Context, workspaceID string, id string, body APIKeyDisableParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s:disable", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Re-enables a disabled API key so its token authenticates again. Idempotent.
func (r *APIKeyService) Enable(ctx context.Context, workspaceID string, id string, body APIKeyEnableParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s:enable", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Rotates an API key and returns a new token. All previous tokens for this key are
// invalidated.
func (r *APIKeyService) Rotate(ctx context.Context, workspaceID string, id string, body APIKeyRotateParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s:rotate", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// An API key. Every key belongs to exactly one workspace and is managed via the
// workspace-scoped API key routes. The only exception is the system-managed global
// account key, which spans all workspaces and is managed via the account
// global_api_key routes.
type APIKey struct {
	// AccountResourceMetadata is used to represent a resource that is associated to an
	// account but not to a workspace.
	Metadata shared.AccountResourceMetadata `json:"metadata" api:"required"`
	// Configuration for an API key.
	Spec APIKeySpec `json:"spec" api:"required"`
	// The current lifecycle state of the API key. Output only. Keys are created
	// STATE_ENABLED; use the :disable and :enable actions to transition between
	// states.
	State APIKeyState `json:"state" api:"required"`
	Info  APIKeyInfo  `json:"info"`
	JSON  apiKeyJSON  `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	State       apijson.Field
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

// The current lifecycle state of the API key. Output only. Keys are created
// STATE_ENABLED; use the :disable and :enable actions to transition between
// states.
type APIKeyState string

const (
	APIKeyStateStateUnspecified APIKeyState = "STATE_UNSPECIFIED"
	APIKeyStateStateEnabled     APIKeyState = "STATE_ENABLED"
	APIKeyStateStateDisabled    APIKeyState = "STATE_DISABLED"
)

func (r APIKeyState) IsKnown() bool {
	switch r {
	case APIKeyStateStateUnspecified, APIKeyStateStateEnabled, APIKeyStateStateDisabled:
		return true
	}
	return false
}

type APIKeyInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile        `json:"createdBy"`
	JSON      apiKeyInfoJSON `json:"-"`
}

// apiKeyInfoJSON contains the JSON metadata for the struct [APIKeyInfo]
type apiKeyInfoJSON struct {
	CreatedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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
	// Scopes granted to this key. Each entry is a colon-separated resource:verb string
	// (e.g. "objectives:manage").
	//
	// Resources: agents, objectives, tools, memory, api_keys, workspaces, secrets,
	// account. Verbs: read and manage, where manage implies read — a stored scope set
	// is normalized to drop "x:read" when "x:manage" is present. The secrets and
	// account resources support only manage. "\*" is an explicit full-access grant.
	//
	// Scopes are deny-by-default: a key with an empty list can call only scope-free
	// endpoints. Full access is always an explicit "\*" grant.
	Permissions []string `json:"permissions"`
	// True when this key is managed by the system (i.e. the auto-provisioned global
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
	// Scopes granted to this key. Each entry is a colon-separated resource:verb string
	// (e.g. "objectives:manage").
	//
	// Resources: agents, objectives, tools, memory, api_keys, workspaces, secrets,
	// account. Verbs: read and manage, where manage implies read — a stored scope set
	// is normalized to drop "x:read" when "x:manage" is present. The secrets and
	// account resources support only manage. "\*" is an explicit full-access grant.
	//
	// Scopes are deny-by-default: a key with an empty list can call only scope-free
	// endpoints. Full access is always an explicit "\*" grant.
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
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"environment": "production", "team":
	// "platform", "version": "v2"}
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
	// Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric
	// characters with "-", "\_", or "." allowed between; keys follow the same shape
	// and additionally accept an optional DNS-subdomain prefix (e.g. "cadenya.com/")
	// of at most 253 characters. Examples: {"environment": "production", "team":
	// "platform", "version": "v2"}
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r APIKeyUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyListParams struct {
	// Pagination cursor from previous response.
	Cursor param.Field[string] `query:"cursor"`
	// When true, included info fields are populated. Requests with this flag count
	// more against your rate limit.
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Field[string] `query:"labels"`
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

type APIKeyDisableParams struct {
}

func (r APIKeyDisableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyEnableParams struct {
}

func (r APIKeyEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type APIKeyRotateParams struct {
}

func (r APIKeyRotateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
