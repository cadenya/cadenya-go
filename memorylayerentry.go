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

// Manage memory layers and their entries. Layers are named containers that can be
// composed into an objective's memory cascade; entries are the keyed values within
// a layer. System-managed layers (e.g., episodic layers created by the runtime)
// cannot be mutated through this API.
//
// MemoryLayerEntryService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemoryLayerEntryService] method instead.
type MemoryLayerEntryService struct {
	options []option.RequestOption
}

// NewMemoryLayerEntryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMemoryLayerEntryService(opts ...option.RequestOption) (r MemoryLayerEntryService) {
	r = MemoryLayerEntryService{}
	r.options = opts
	return
}

// Creates a new entry in a memory layer. Returns the detail view, including the
// resolved content body.
func (r *MemoryLayerEntryService) New(ctx context.Context, memoryLayerID string, params MemoryLayerEntryNewParams, opts ...option.RequestOption) (res *MemoryEntryDetail, err error) {
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
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s/entries", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(memoryLayerID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a memory entry by ID from a memory layer. Returns the detail view,
// including the content body.
func (r *MemoryLayerEntryService) Get(ctx context.Context, memoryLayerID string, id string, query MemoryLayerEntryGetParams, opts ...option.RequestOption) (res *MemoryEntryDetail, err error) {
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
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s/entries/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(memoryLayerID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a memory entry in a memory layer. Returns the detail view, including the
// resolved content body.
func (r *MemoryLayerEntryService) Update(ctx context.Context, memoryLayerID string, id string, params MemoryLayerEntryUpdateParams, opts ...option.RequestOption) (res *MemoryEntryDetail, err error) {
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
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s/entries/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(memoryLayerID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all entries in a memory layer
func (r *MemoryLayerEntryService) List(ctx context.Context, memoryLayerID string, params MemoryLayerEntryListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[MemoryEntry], err error) {
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
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s/entries", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(memoryLayerID))
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

// Lists all entries in a memory layer
func (r *MemoryLayerEntryService) ListAutoPaging(ctx context.Context, memoryLayerID string, params MemoryLayerEntryListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[MemoryEntry] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, memoryLayerID, params, opts...))
}

// Deletes a memory entry from a memory layer
func (r *MemoryLayerEntryService) Delete(ctx context.Context, memoryLayerID string, id string, body MemoryLayerEntryDeleteParams, opts ...option.RequestOption) (err error) {
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
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/memory_layers/%s/entries/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(memoryLayerID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// MemoryEntry is a single keyed value within a MemoryLayer. Entries are addressed
// by their key, which follows the S3 object key safe-character convention (see
// MemoryEntrySpec.key for the full rule). Keys are unique within a single layer;
// the same key may appear in multiple layers, in which case the cascade walk
// determines which one wins for a given objective (most specific layer first).
//
// MemoryEntry is the summary shape, returned by ListMemoryEntries. It does not
// carry the entry body — callers that need the body must fetch the entry
// individually via GetMemoryEntry, which returns a MemoryEntryDetail.
type MemoryEntry struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// MemoryEntrySpec is the metadata portion of an entry — the fields that identify
	// and describe it, without the body. It appears on both the summary (MemoryEntry)
	// and detail (MemoryEntryDetail) views.
	Spec MemoryEntrySpec `json:"spec" api:"required"`
	Info MemoryEntryInfo `json:"info"`
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
func (r MemoryEntry) RawJSON() string { return r.JSON.raw }
func (r *MemoryEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryEntryCreateSpec is the input shape for CreateMemoryEntry. It accepts
// either inline content or a reference to a completed Upload; exactly one of the
// two must be set.
//
// The property Key is required.
type MemoryEntryCreateSpecParam struct {
	// See MemoryEntrySpec.key for the full rule set. Same constraints apply here.
	Key string `json:"key" api:"required"`
	// Inline content, written directly into the entry.
	Content     param.Opt[string] `json:"content,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	// The JSON name of the variant set in `source` (e.g. "content"). Required on
	// input; drives the discriminated union in the generated OpenAPI.
	Type param.Opt[string] `json:"type,omitzero"`
	// ID of a COMPLETE Upload. The server reads the object from storage, copies its
	// bytes into the entry, and marks the upload consumed.
	UploadID param.Opt[string] `json:"uploadId,omitzero"`
	paramObj
}

func (r MemoryEntryCreateSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow MemoryEntryCreateSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryEntryCreateSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryEntryDetail is the full representation of an entry, including the resolved
// content body. Returned by GetMemoryEntry, CreateMemoryEntry, and
// UpdateMemoryEntry.
type MemoryEntryDetail struct {
	// The resolved body of the entry. For entries created or updated via an upload_id,
	// this is the ingested content, not the original upload handle. May be empty; an
	// entry with only a key and description is valid (e.g., a stub skill being
	// drafted, or an entry where the frontmatter alone is the payload).
	Content string `json:"content" api:"required"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// MemoryEntrySpec is the metadata portion of an entry — the fields that identify
	// and describe it, without the body. It appears on both the summary (MemoryEntry)
	// and detail (MemoryEntryDetail) views.
	Spec MemoryEntrySpec `json:"spec" api:"required"`
	Info MemoryEntryInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Metadata    respjson.Field
		Spec        respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryEntryDetail) RawJSON() string { return r.JSON.raw }
func (r *MemoryEntryDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryEntryInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	MemoryLayer shared.ResourceMetadata `json:"memoryLayer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		MemoryLayer respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryEntryInfo) RawJSON() string { return r.JSON.raw }
func (r *MemoryEntryInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryEntrySpec is the metadata portion of an entry — the fields that identify
// and describe it, without the body. It appears on both the summary (MemoryEntry)
// and detail (MemoryEntryDetail) views.
type MemoryEntrySpec struct {
	// The lookup key for this entry within its layer. Must conform to the S3 object
	// key safe-characters spec: ASCII alphanumerics and the special characters !, -,
	// \_, ., \*, ', (, ), and /. Forward slashes may be used to suggest hierarchy
	// (e.g., "skills/postmortem/write"), but lookups are flat — the key is a single
	// opaque string, not a path.
	//
	// Additional rules enforced by the service:
	//
	// - May not begin or end with /
	// - May not contain consecutive slashes (//)
	// - May not begin with reserved prefixes (cadenya/, system/)
	// - Case-sensitive
	// - Unique within the parent layer
	//
	// For skills entries, this key is what the model passes to get_memory to load the
	// entry's content.
	Key string `json:"key" api:"required"`
	// One-line "when to use this" hint shown in the frontmatter manifest for skills
	// entries. The model uses this to decide whether to load the body, so it should be
	// written for the model as the audience. Ignored for layer types that do not
	// advertise frontmatter.
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryEntrySpec) RawJSON() string { return r.JSON.raw }
func (r *MemoryEntrySpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryEntryUpdateSpec is the input shape for UpdateMemoryEntry. Fields present
// in the request's update_mask are applied; unset fields are left alone. The
// source oneof is optional for updates — omit it to leave the body untouched, or
// set exactly one branch to replace it.
type MemoryEntryUpdateSpecParam struct {
	Content     param.Opt[string] `json:"content,omitzero"`
	Description param.Opt[string] `json:"description,omitzero"`
	Key         param.Opt[string] `json:"key,omitzero"`
	UploadID    param.Opt[string] `json:"uploadId,omitzero"`
	paramObj
}

func (r MemoryEntryUpdateSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow MemoryEntryUpdateSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryEntryUpdateSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryLayerEntryNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	// MemoryEntryCreateSpec is the input shape for CreateMemoryEntry. It accepts
	// either inline content or a reference to a completed Upload; exactly one of the
	// two must be set.
	Spec MemoryEntryCreateSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r MemoryLayerEntryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MemoryLayerEntryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryLayerEntryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryLayerEntryGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type MemoryLayerEntryUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	UpdateMask  param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	// MemoryEntryUpdateSpec is the input shape for UpdateMemoryEntry. Fields present
	// in the request's update_mask are applied; unset fields are left alone. The
	// source oneof is optional for updates — omit it to leave the body untouched, or
	// set exactly one branch to replace it.
	Spec MemoryEntryUpdateSpecParam `json:"spec,omitzero"`
	paramObj
}

func (r MemoryLayerEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MemoryLayerEntryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryLayerEntryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryLayerEntryListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by key prefix (e.g., "skills/postmortem/" to list all entries under that
	// hierarchy). Matches against the entry's key, not its name.
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Free-form search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MemoryLayerEntryListParams]'s query parameters as
// `url.Values`.
func (r MemoryLayerEntryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MemoryLayerEntryDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
