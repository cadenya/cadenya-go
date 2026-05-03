// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cadenya/cadenya-go/internal/apijson"
	"github.com/cadenya/cadenya-go/internal/apiquery"
	"github.com/cadenya/cadenya-go/internal/param"
	"github.com/cadenya/cadenya-go/internal/requestconfig"
	"github.com/cadenya/cadenya-go/option"
	"github.com/cadenya/cadenya-go/packages/pagination"
	"github.com/cadenya/cadenya-go/shared"
)

// MemoryService manages memory layers and their entries at the WORKSPACE level.
// Layers are named containers that can be composed into an objective's memory
// stack; entries are the keyed values within a layer.
//
// All operations are implicitly scoped to the workspace determined by the JWT
// token. System-managed layers (e.g., episodic layers created by the runtime)
// cannot be mutated through this API.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// MemoryLayerService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemoryLayerService] method instead.
type MemoryLayerService struct {
	Options []option.RequestOption
	// MemoryService manages memory layers and their entries at the WORKSPACE level.
	// Layers are named containers that can be composed into an objective's memory
	// stack; entries are the keyed values within a layer.
	//
	// All operations are implicitly scoped to the workspace determined by the JWT
	// token. System-managed layers (e.g., episodic layers created by the runtime)
	// cannot be mutated through this API.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	Entries *MemoryLayerEntryService
}

// NewMemoryLayerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMemoryLayerService(opts ...option.RequestOption) (r *MemoryLayerService) {
	r = &MemoryLayerService{}
	r.Options = opts
	r.Entries = NewMemoryLayerEntryService(opts...)
	return
}

// Creates a new memory layer in the workspace
func (r *MemoryLayerService) New(ctx context.Context, workspaceID string, body MemoryLayerNewParams, opts ...option.RequestOption) (res *MemoryLayer, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a memory layer by ID from the workspace
func (r *MemoryLayerService) Get(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (res *MemoryLayer, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a memory layer in the workspace
func (r *MemoryLayerService) Update(ctx context.Context, workspaceID string, id string, body MemoryLayerUpdateParams, opts ...option.RequestOption) (res *MemoryLayer, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all memory layers in the workspace
func (r *MemoryLayerService) List(ctx context.Context, workspaceID string, query MemoryLayerListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[MemoryLayer], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers", workspaceID)
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

// Lists all memory layers in the workspace
func (r *MemoryLayerService) ListAutoPaging(ctx context.Context, workspaceID string, query MemoryLayerListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[MemoryLayer] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Deletes a memory layer from the workspace
func (r *MemoryLayerService) Delete(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// MemoryLayer is a named container of memory entries that can be composed into an
// objective's memory stack. Layers are workspace-scoped resources. The layer type
// controls how its entries participate in the agent loop — see MemoryLayerType for
// details.
//
// See "Memory stack composition" above for how layers compose at lookup time.
type MemoryLayer struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     MemoryLayerSpec         `json:"spec" api:"required"`
	Info     MemoryLayerInfo         `json:"info"`
	JSON     memoryLayerJSON         `json:"-"`
}

// memoryLayerJSON contains the JSON metadata for the struct [MemoryLayer]
type memoryLayerJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryLayer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryLayerJSON) RawJSON() string {
	return r.raw
}

type MemoryLayerInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy Profile `json:"createdBy"`
	// Number of entries currently in this layer.
	EntryCount int64 `json:"entryCount"`
	// Timestamp of the most recent objective that resolved against this layer. Useful
	// for surfacing unused layers in the dashboard.
	LastUsedAt time.Time           `json:"lastUsedAt" format:"date-time"`
	JSON       memoryLayerInfoJSON `json:"-"`
}

// memoryLayerInfoJSON contains the JSON metadata for the struct [MemoryLayerInfo]
type memoryLayerInfoJSON struct {
	CreatedBy   apijson.Field
	EntryCount  apijson.Field
	LastUsedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryLayerInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryLayerInfoJSON) RawJSON() string {
	return r.raw
}

type MemoryLayerSpec struct {
	Type MemoryLayerSpecType `json:"type" api:"required"`
	// Human-readable description of the layer's purpose. Encouraged for user-created
	// layers; system-managed layers may have a generated description.
	Description string `json:"description"`
	// For layers with a finite lifetime (e.g., episodic), the time at which the layer
	// becomes eligible for cleanup. Set by the system; unset for persistent layers.
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// Server-set. True for layers managed by the system (e.g., episodic layers created
	// automatically when an objective uses an episodic_key). System-managed layers
	// cannot be assigned to objective stacks via the API and cannot be mutated by
	// clients — their lifecycle is controlled entirely by the runtime.
	SystemManaged bool                `json:"systemManaged"`
	JSON          memoryLayerSpecJSON `json:"-"`
}

// memoryLayerSpecJSON contains the JSON metadata for the struct [MemoryLayerSpec]
type memoryLayerSpecJSON struct {
	Type          apijson.Field
	Description   apijson.Field
	ExpiresAt     apijson.Field
	SystemManaged apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MemoryLayerSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryLayerSpecJSON) RawJSON() string {
	return r.raw
}

type MemoryLayerSpecType string

const (
	MemoryLayerSpecTypeMemoryLayerTypeUnspecified MemoryLayerSpecType = "MEMORY_LAYER_TYPE_UNSPECIFIED"
	MemoryLayerSpecTypeMemoryLayerTypeEpisodic    MemoryLayerSpecType = "MEMORY_LAYER_TYPE_EPISODIC"
	MemoryLayerSpecTypeMemoryLayerTypeSkills      MemoryLayerSpecType = "MEMORY_LAYER_TYPE_SKILLS"
)

func (r MemoryLayerSpecType) IsKnown() bool {
	switch r {
	case MemoryLayerSpecTypeMemoryLayerTypeUnspecified, MemoryLayerSpecTypeMemoryLayerTypeEpisodic, MemoryLayerSpecTypeMemoryLayerTypeSkills:
		return true
	}
	return false
}

type MemoryLayerSpecParam struct {
	Type param.Field[MemoryLayerSpecType] `json:"type" api:"required"`
	// Human-readable description of the layer's purpose. Encouraged for user-created
	// layers; system-managed layers may have a generated description.
	Description param.Field[string] `json:"description"`
}

func (r MemoryLayerSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemoryLayerNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	Spec     param.Field[MemoryLayerSpecParam]               `json:"spec" api:"required"`
}

func (r MemoryLayerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemoryLayerUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata   param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	Spec       param.Field[MemoryLayerSpecParam]               `json:"spec"`
	UpdateMask param.Field[string]                             `json:"updateMask" format:"field-mask"`
}

func (r MemoryLayerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemoryLayerListParams struct {
	// Filter by bundle_key — return only resources owned by this bundle.
	BundleKey param.Field[string] `query:"bundleKey"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by layer type
	Type param.Field[MemoryLayerListParamsType] `query:"type"`
}

// URLQuery serializes [MemoryLayerListParams]'s query parameters as `url.Values`.
func (r MemoryLayerListParams) URLQuery() (v url.Values) {
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

func (r MemoryLayerListParamsType) IsKnown() bool {
	switch r {
	case MemoryLayerListParamsTypeMemoryLayerTypeUnspecified, MemoryLayerListParamsTypeMemoryLayerTypeEpisodic, MemoryLayerListParamsTypeMemoryLayerTypeSkills:
		return true
	}
	return false
}
