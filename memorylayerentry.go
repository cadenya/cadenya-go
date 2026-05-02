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
// MemoryLayerEntryService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemoryLayerEntryService] method instead.
type MemoryLayerEntryService struct {
	Options []option.RequestOption
}

// NewMemoryLayerEntryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMemoryLayerEntryService(opts ...option.RequestOption) (r *MemoryLayerEntryService) {
	r = &MemoryLayerEntryService{}
	r.Options = opts
	return
}

// Creates a new entry in a memory layer. Returns the detail view, including the
// resolved content body.
func (r *MemoryLayerEntryService) New(ctx context.Context, memoryLayerID string, body MemoryLayerEntryNewParams, opts ...option.RequestOption) (res *MemoryEntryDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/memory_layers/%s/entries", memoryLayerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a memory entry by ID from a memory layer. Returns the detail view,
// including the content body.
func (r *MemoryLayerEntryService) Get(ctx context.Context, memoryLayerID string, id string, opts ...option.RequestOption) (res *MemoryEntryDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/memory_layers/%s/entries/%s", memoryLayerID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a memory entry in a memory layer. Returns the detail view, including the
// resolved content body.
func (r *MemoryLayerEntryService) Update(ctx context.Context, memoryLayerID string, id string, body MemoryLayerEntryUpdateParams, opts ...option.RequestOption) (res *MemoryEntryDetail, err error) {
	opts = slices.Concat(r.Options, opts)
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/memory_layers/%s/entries/%s", memoryLayerID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all entries in a memory layer
func (r *MemoryLayerEntryService) List(ctx context.Context, memoryLayerID string, query MemoryLayerEntryListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[MemoryEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/memory_layers/%s/entries", memoryLayerID)
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

// Lists all entries in a memory layer
func (r *MemoryLayerEntryService) ListAutoPaging(ctx context.Context, memoryLayerID string, query MemoryLayerEntryListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[MemoryEntry] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, memoryLayerID, query, opts...))
}

// Deletes a memory entry from a memory layer
func (r *MemoryLayerEntryService) Delete(ctx context.Context, memoryLayerID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if memoryLayerID == "" {
		err = errors.New("missing required memoryLayerId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/memory_layers/%s/entries/%s", memoryLayerID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// MemoryEntry is a single keyed value within a MemoryLayer. Entries are addressed
// by their key, which follows the S3 object key safe-character convention (see
// MemoryEntrySpec.key for the full rule). Keys are unique within a single layer;
// the same key may appear in multiple layers, in which case the LIFO stack-walk
// determines which one wins for a given objective.
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
	JSON memoryEntryJSON `json:"-"`
}

// memoryEntryJSON contains the JSON metadata for the struct [MemoryEntry]
type memoryEntryJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryEntryJSON) RawJSON() string {
	return r.raw
}

// MemoryEntryCreateSpec is the input shape for CreateMemoryEntry. It accepts
// either inline content or a reference to a completed Upload; exactly one of the
// two must be set.
type MemoryEntryCreateSpecParam struct {
	// See MemoryEntrySpec.key for the full rule set. Same constraints apply here.
	Key param.Field[string] `json:"key" api:"required"`
	// Inline content, written directly into the entry.
	Content     param.Field[string] `json:"content"`
	Description param.Field[string] `json:"description"`
	// ID of a COMPLETE Upload. The server reads the object from storage, copies its
	// bytes into the entry, and marks the upload consumed.
	UploadID param.Field[string] `json:"uploadId"`
}

func (r MemoryEntryCreateSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Spec MemoryEntrySpec       `json:"spec" api:"required"`
	Info MemoryEntryInfo       `json:"info"`
	JSON memoryEntryDetailJSON `json:"-"`
}

// memoryEntryDetailJSON contains the JSON metadata for the struct
// [MemoryEntryDetail]
type memoryEntryDetailJSON struct {
	Content     apijson.Field
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryEntryDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryEntryDetailJSON) RawJSON() string {
	return r.raw
}

type MemoryEntryInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy Profile `json:"createdBy"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	MemoryLayer shared.ResourceMetadata `json:"memoryLayer"`
	JSON        memoryEntryInfoJSON     `json:"-"`
}

// memoryEntryInfoJSON contains the JSON metadata for the struct [MemoryEntryInfo]
type memoryEntryInfoJSON struct {
	CreatedBy   apijson.Field
	MemoryLayer apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryEntryInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryEntryInfoJSON) RawJSON() string {
	return r.raw
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
	// For skills entries, this key is also the id the model passes to
	// memory_load_skill when it decides to load the entry's content.
	Key string `json:"key" api:"required"`
	// One-line "when to use this" hint shown in the frontmatter manifest for skills
	// entries. The model uses this to decide whether to load the body, so it should be
	// written for the model as the audience. Ignored for layer types that do not
	// advertise frontmatter.
	Description string              `json:"description"`
	JSON        memoryEntrySpecJSON `json:"-"`
}

// memoryEntrySpecJSON contains the JSON metadata for the struct [MemoryEntrySpec]
type memoryEntrySpecJSON struct {
	Key         apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryEntrySpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryEntrySpecJSON) RawJSON() string {
	return r.raw
}

// MemoryEntryUpdateSpec is the input shape for UpdateMemoryEntry. Fields present
// in the request's update_mask are applied; unset fields are left alone. The
// source oneof is optional for updates — omit it to leave the body untouched, or
// set exactly one branch to replace it.
type MemoryEntryUpdateSpecParam struct {
	Content     param.Field[string] `json:"content"`
	Description param.Field[string] `json:"description"`
	Key         param.Field[string] `json:"key"`
	UploadID    param.Field[string] `json:"uploadId"`
}

func (r MemoryEntryUpdateSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemoryLayerEntryNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	// MemoryEntryCreateSpec is the input shape for CreateMemoryEntry. It accepts
	// either inline content or a reference to a completed Upload; exactly one of the
	// two must be set.
	Spec param.Field[MemoryEntryCreateSpecParam] `json:"spec" api:"required"`
}

func (r MemoryLayerEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemoryLayerEntryUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	// MemoryEntryUpdateSpec is the input shape for UpdateMemoryEntry. Fields present
	// in the request's update_mask are applied; unset fields are left alone. The
	// source oneof is optional for updates — omit it to leave the body untouched, or
	// set exactly one branch to replace it.
	Spec       param.Field[MemoryEntryUpdateSpecParam] `json:"spec"`
	UpdateMask param.Field[string]                     `json:"updateMask" format:"field-mask"`
}

func (r MemoryLayerEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemoryLayerEntryListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter by key prefix (e.g., "skills/postmortem/" to list all entries under that
	// hierarchy). Matches against the entry's key, not its name.
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [MemoryLayerEntryListParams]'s query parameters as
// `url.Values`.
func (r MemoryLayerEntryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
