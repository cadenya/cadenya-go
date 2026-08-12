package cadenya

import (
	"context"
	"encoding/json"
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
	options []option.RequestOption
}

// NewAPIKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAPIKeyService(opts ...option.RequestOption) (r APIKeyService) {
	r = APIKeyService{}
	r.options = opts
	return
}

// Creates a new API key in the workspace.
func (r *APIKeyService) New(ctx context.Context, params APIKeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/api_keys", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves an API key by ID.
func (r *APIKeyService) Get(ctx context.Context, id string, query APIKeyGetParams, opts ...option.RequestOption) (res *APIKey, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an API key.
func (r *APIKeyService) Update(ctx context.Context, id string, params APIKeyUpdateParams, opts ...option.RequestOption) (res *APIKey, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists the workspace's API keys.
func (r *APIKeyService) List(ctx context.Context, params APIKeyListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[APIKey], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/api_keys", url.PathEscape(params.WorkspaceID.Value))
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

// Lists the workspace's API keys.
func (r *APIKeyService) ListAutoPaging(ctx context.Context, params APIKeyListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[APIKey] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes an API key.
func (r *APIKeyService) Delete(ctx context.Context, id string, body APIKeyDeleteParams, opts ...option.RequestOption) (err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Disables an API key. While disabled, presenting the key's token fails
// authentication on every endpoint; the key is retained. Idempotent.
func (r *APIKeyService) Disable(ctx context.Context, id string, body APIKeyDisableParams, opts ...option.RequestOption) (res *APIKey, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s:disable", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Re-enables a disabled API key so its token authenticates again. Idempotent.
func (r *APIKeyService) Enable(ctx context.Context, id string, body APIKeyEnableParams, opts ...option.RequestOption) (res *APIKey, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s:enable", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Rotates an API key and returns a new token. All previous tokens for this key are
// invalidated.
func (r *APIKeyService) Rotate(ctx context.Context, id string, body APIKeyRotateParams, opts ...option.RequestOption) (res *APIKey, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/api_keys/%s:rotate", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
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
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ENABLED", "STATE_DISABLED".
	State APIKeyState `json:"state" api:"required"`
	Info  APIKeyInfo  `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		State       respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type APIKeyInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeyInfo) RawJSON() string { return r.JSON.raw }
func (r *APIKeyInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	// Resources: agents, objectives, tools, memory, api_keys, workspaces, widgets,
	// widget_sessions, secrets, account. Verbs: read and manage, where manage implies
	// read — a stored scope set is normalized to drop "x:read" when "x:manage" is
	// present. The secrets and account resources support only manage. "\*" is an
	// explicit full-access grant.
	//
	// Scopes are deny-by-default: a key with an empty list can call only scope-free
	// endpoints. Full access is always an explicit "\*" grant.
	Permissions []string `json:"permissions"`
	// True when this key is managed by the system (i.e. the auto-provisioned global
	// account key). System keys cannot be deleted but can be rotated.
	System bool `json:"system"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Description respjson.Field
		Permissions respjson.Field
		System      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKeySpec) RawJSON() string { return r.JSON.raw }
func (r *APIKeySpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this APIKeySpec to a APIKeySpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// APIKeySpecParam.Overrides()
func (r APIKeySpec) ToParam() APIKeySpecParam {
	return param.Override[APIKeySpecParam](json.RawMessage(r.RawJSON()))
}

// Configuration for an API key.
type APIKeySpecParam struct {
	// Free-form description of what this API key is used for.
	Description param.Opt[string] `json:"description,omitzero"`
	// Scopes granted to this key. Each entry is a colon-separated resource:verb string
	// (e.g. "objectives:manage").
	//
	// Resources: agents, objectives, tools, memory, api_keys, workspaces, widgets,
	// widget_sessions, secrets, account. Verbs: read and manage, where manage implies
	// read — a stored scope set is normalized to drop "x:read" when "x:manage" is
	// present. The secrets and account resources support only manage. "\*" is an
	// explicit full-access grant.
	//
	// Scopes are deny-by-default: a key with an empty list can call only scope-free
	// endpoints. Full access is always an explicit "\*" grant.
	Permissions []string `json:"permissions,omitzero"`
	paramObj
}

func (r APIKeySpecParam) MarshalJSON() (data []byte, err error) {
	type shadow APIKeySpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeySpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateAccountResourceMetadata contains the user-provided fields for creating an
	// account-scoped resource. Read-only fields (id, account_id, profile_id) are
	// excluded since they are set by the server.
	Metadata APIKeyNewParamsMetadata `json:"metadata,omitzero" api:"required"`
	// Configuration for an API key.
	Spec APIKeySpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r APIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CreateAccountResourceMetadata contains the user-provided fields for creating an
// account-scoped resource. Read-only fields (id, account_id, profile_id) are
// excluded since they are set by the server.
//
// The property Name is required.
type APIKeyNewParamsMetadata struct {
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

func (r APIKeyNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyNewParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyNewParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type APIKeyUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Fields to update.
	UpdateMask param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateAccountResourceMetadata contains the user-provided fields for updating an
	// account-scoped resource. Read-only fields (id, account_id, profile_id) are
	// excluded since they are set by the server.
	Metadata APIKeyUpdateParamsMetadata `json:"metadata,omitzero"`
	// Configuration for an API key.
	Spec APIKeySpecParam `json:"spec,omitzero"`
	paramObj
}

func (r APIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UpdateAccountResourceMetadata contains the user-provided fields for updating an
// account-scoped resource. Read-only fields (id, account_id, profile_id) are
// excluded since they are set by the server.
//
// The property Name is required.
type APIKeyUpdateParamsMetadata struct {
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

func (r APIKeyUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyUpdateParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyUpdateParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, included info fields are populated. Requests with this flag count
	// more against your rate limit.
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by ID prefix.
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Free-form search query.
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time).
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [APIKeyListParams]'s query parameters as `url.Values`.
func (r APIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type APIKeyDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type APIKeyDisableParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r APIKeyDisableParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyDisableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyDisableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyEnableParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r APIKeyEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type APIKeyRotateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r APIKeyRotateParams) MarshalJSON() (data []byte, err error) {
	type shadow APIKeyRotateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *APIKeyRotateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
