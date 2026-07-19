// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
	"time"
)

// Manage memory layers and their entries. Layers are named containers that can be
// composed into an objective's memory cascade; entries are the keyed values within
// a layer. System-managed layers (e.g., episodic layers created by the runtime)
// cannot be mutated through this API.
//
// MemoryLayerService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemoryLayerService] method instead.
type MemoryLayerService struct {
	options []option.RequestOption
	// Manage memory layers and their entries. Layers are named containers that can be
	// composed into an objective's memory cascade; entries are the keyed values within
	// a layer. System-managed layers (e.g., episodic layers created by the runtime)
	// cannot be mutated through this API.
	Entries MemoryLayerEntryService
}

// NewMemoryLayerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMemoryLayerService(opts ...option.RequestOption) (r MemoryLayerService) {
	r = MemoryLayerService{}
	r.options = opts
	r.Entries = NewMemoryLayerEntryService(opts...)
	return
}

// Creates a new memory layer in the workspace
func (r *MemoryLayerService) New(ctx context.Context, params MemoryLayerNewParams, opts ...option.RequestOption) (res *MemoryLayer, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a memory layer by ID from the workspace
func (r *MemoryLayerService) Get(ctx context.Context, id string, query MemoryLayerGetParams, opts ...option.RequestOption) (res *MemoryLayer, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a memory layer in the workspace
func (r *MemoryLayerService) Update(ctx context.Context, id string, params MemoryLayerUpdateParams, opts ...option.RequestOption) (res *MemoryLayer, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all memory layers in the workspace
func (r *MemoryLayerService) List(ctx context.Context, params MemoryLayerListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[MemoryLayer], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers", url.PathEscape(params.WorkspaceID.Value))
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

// Lists all memory layers in the workspace
func (r *MemoryLayerService) ListAutoPaging(ctx context.Context, params MemoryLayerListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[MemoryLayer] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes a memory layer from the workspace
func (r *MemoryLayerService) Delete(ctx context.Context, id string, body MemoryLayerDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// MemoryLayer is a named container of memory entries that can be composed into an
// objective's memory cascade. Layers are workspace-scoped resources. The layer
// type controls how its entries participate in the agent loop — see
// MemoryLayerType for details.
//
// See "Memory cascade composition" above for how layers compose at lookup time.
type MemoryLayer struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     MemoryLayerSpec         `json:"spec" api:"required"`
	Info     MemoryLayerInfo         `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryLayer) RawJSON() string { return r.JSON.raw }
func (r *MemoryLayer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryLayerInfo struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent shared.ResourceMetadata `json:"agent"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Number of entries currently in this layer.
	EntryCount int64 `json:"entryCount"`
	// Timestamp of the most recent objective that resolved against this layer. Useful
	// for surfacing unused layers in the dashboard.
	LastUsedAt time.Time `json:"lastUsedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent       respjson.Field
		CreatedBy   respjson.Field
		EntryCount  respjson.Field
		LastUsedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryLayerInfo) RawJSON() string { return r.JSON.raw }
func (r *MemoryLayerInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryLayerSpec struct {
	// Any of "MEMORY_LAYER_TYPE_UNSPECIFIED", "MEMORY_LAYER_TYPE_EPISODIC",
	// "MEMORY_LAYER_TYPE_SKILLS".
	Type MemoryLayerSpecType `json:"type" api:"required"`
	// Server-set on episodic layers: the agent this layer belongs to. Unset for
	// non-episodic layers.
	AgentID string `json:"agentId"`
	// Human-readable description of the layer's purpose. Encouraged for user-created
	// layers; system-managed layers may have a generated description.
	Description string `json:"description"`
	// Server-set on episodic layers: the caller-supplied episodic key the layer was
	// created for. Unset for non-episodic layers.
	EpisodicKey string `json:"episodicKey"`
	// For layers with a finite lifetime (e.g., episodic), the time at which the layer
	// becomes eligible for cleanup. Set by the system; unset for persistent layers.
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// Server-set. True for layers managed by the system (e.g., episodic layers created
	// automatically when an objective uses an episodic_key). System-managed layers
	// cannot be assigned to objective cascades via the API and cannot be mutated by
	// clients — their lifecycle is controlled entirely by the runtime.
	SystemManaged bool `json:"systemManaged"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		AgentID       respjson.Field
		Description   respjson.Field
		EpisodicKey   respjson.Field
		ExpiresAt     respjson.Field
		SystemManaged respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryLayerSpec) RawJSON() string { return r.JSON.raw }
func (r *MemoryLayerSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MemoryLayerSpec to a MemoryLayerSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MemoryLayerSpecParam.Overrides()
func (r MemoryLayerSpec) ToParam() MemoryLayerSpecParam {
	return param.Override[MemoryLayerSpecParam](json.RawMessage(r.RawJSON()))
}

type MemoryLayerSpecType string

const (
	MemoryLayerSpecTypeMemoryLayerTypeUnspecified MemoryLayerSpecType = "MEMORY_LAYER_TYPE_UNSPECIFIED"
	MemoryLayerSpecTypeMemoryLayerTypeEpisodic    MemoryLayerSpecType = "MEMORY_LAYER_TYPE_EPISODIC"
	MemoryLayerSpecTypeMemoryLayerTypeSkills      MemoryLayerSpecType = "MEMORY_LAYER_TYPE_SKILLS"
)

// The property Type is required.
type MemoryLayerSpecParam struct {
	// Any of "MEMORY_LAYER_TYPE_UNSPECIFIED", "MEMORY_LAYER_TYPE_EPISODIC",
	// "MEMORY_LAYER_TYPE_SKILLS".
	Type MemoryLayerSpecType `json:"type,omitzero" api:"required"`
	// Human-readable description of the layer's purpose. Encouraged for user-created
	// layers; system-managed layers may have a generated description.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r MemoryLayerSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow MemoryLayerSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryLayerSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryLayerNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	Spec     MemoryLayerSpecParam               `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r MemoryLayerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MemoryLayerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryLayerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryLayerGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type MemoryLayerUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	UpdateMask  param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	Spec     MemoryLayerSpecParam               `json:"spec,omitzero"`
	paramObj
}

func (r MemoryLayerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MemoryLayerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryLayerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryLayerListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Filter to episodic layers belonging to this agent.
	AgentID param.Opt[string] `query:"agentId,omitzero" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter to episodic layers whose episodic key starts with this prefix (e.g.
	// "customer/" matches "customer/42" and "customer/43"). Useful for namespaced
	// keys, similar to a redis key scan.
	EpisodicKeyPrefix param.Opt[string] `query:"episodicKeyPrefix,omitzero" json:"-"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter expression (query param: prefix)
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Free-form search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// Filter by layer type
	//
	// Any of "MEMORY_LAYER_TYPE_UNSPECIFIED", "MEMORY_LAYER_TYPE_EPISODIC",
	// "MEMORY_LAYER_TYPE_SKILLS".
	Type MemoryLayerListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MemoryLayerListParams]'s query parameters as `url.Values`.
func (r MemoryLayerListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by layer type
type MemoryLayerListParamsType string

const (
	MemoryLayerListParamsTypeMemoryLayerTypeUnspecified MemoryLayerListParamsType = "MEMORY_LAYER_TYPE_UNSPECIFIED"
	MemoryLayerListParamsTypeMemoryLayerTypeEpisodic    MemoryLayerListParamsType = "MEMORY_LAYER_TYPE_EPISODIC"
	MemoryLayerListParamsTypeMemoryLayerTypeSkills      MemoryLayerListParamsType = "MEMORY_LAYER_TYPE_SKILLS"
)

type MemoryLayerDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
