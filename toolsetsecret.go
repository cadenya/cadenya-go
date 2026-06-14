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

// Manage tool sets and the tools they contain. Tool sets group related tools, and
// tools define specific capabilities available to agents.
//
// When a tool set is managed, only API key actors can modify its tools; human
// (profile) actors cannot.
//
// ToolSetSecretService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolSetSecretService] method instead.
type ToolSetSecretService struct {
	Options []option.RequestOption
}

// NewToolSetSecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolSetSecretService(opts ...option.RequestOption) (r *ToolSetSecretService) {
	r = &ToolSetSecretService{}
	r.Options = opts
	return
}

// Creates a new secret scoped to the tool set
func (r *ToolSetSecretService) New(ctx context.Context, workspaceID string, toolSetID string, body ToolSetSecretNewParams, opts ...option.RequestOption) (res *ToolSetSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets", workspaceID, toolSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a tool set secret by ID from the tool set
func (r *ToolSetSecretService) Get(ctx context.Context, workspaceID string, toolSetID string, id string, opts ...option.RequestOption) (res *ToolSetSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets/%s", workspaceID, toolSetID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a secret scoped to the tool set
func (r *ToolSetSecretService) Update(ctx context.Context, workspaceID string, toolSetID string, id string, body ToolSetSecretUpdateParams, opts ...option.RequestOption) (res *ToolSetSecret, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets/%s", workspaceID, toolSetID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all secrets scoped to the tool set
func (r *ToolSetSecretService) List(ctx context.Context, workspaceID string, toolSetID string, query ToolSetSecretListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSetSecret], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets", workspaceID, toolSetID)
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

// Lists all secrets scoped to the tool set
func (r *ToolSetSecretService) ListAutoPaging(ctx context.Context, workspaceID string, toolSetID string, query ToolSetSecretListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSetSecret] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, toolSetID, query, opts...))
}

// Deletes a secret scoped to the tool set
func (r *ToolSetSecretService) Delete(ctx context.Context, workspaceID string, toolSetID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets/%s", workspaceID, toolSetID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ToolSetSecret struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     ToolSetSecretSpec       `json:"spec" api:"required"`
	// Tool set secret information
	Info ToolSetSecretInfo `json:"info"`
	JSON toolSetSecretJSON `json:"-"`
}

// toolSetSecretJSON contains the JSON metadata for the struct [ToolSetSecret]
type toolSetSecretJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetSecretJSON) RawJSON() string {
	return r.raw
}

type ToolSetSecretInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy  Profile               `json:"createdBy"`
	LastUsedAt time.Time             `json:"lastUsedAt" format:"date-time"`
	JSON       toolSetSecretInfoJSON `json:"-"`
}

// toolSetSecretInfoJSON contains the JSON metadata for the struct
// [ToolSetSecretInfo]
type toolSetSecretInfoJSON struct {
	CreatedBy   apijson.Field
	LastUsedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetSecretInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetSecretInfoJSON) RawJSON() string {
	return r.raw
}

type ToolSetSecretSpec struct {
	Value string                `json:"value"`
	JSON  toolSetSecretSpecJSON `json:"-"`
}

// toolSetSecretSpecJSON contains the JSON metadata for the struct
// [ToolSetSecretSpec]
type toolSetSecretSpecJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetSecretSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetSecretSpecJSON) RawJSON() string {
	return r.raw
}

type ToolSetSecretSpecParam struct {
	Value param.Field[string] `json:"value"`
}

func (r ToolSetSecretSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetSecretNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	Spec     param.Field[ToolSetSecretSpecParam]             `json:"spec" api:"required"`
}

func (r ToolSetSecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetSecretUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	Spec     param.Field[ToolSetSecretSpecParam]             `json:"spec"`
	// Fields to update.
	UpdateMask param.Field[string] `json:"updateMask" format:"field-mask"`
}

func (r ToolSetSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetSecretListParams struct {
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
}

// URLQuery serializes [ToolSetSecretListParams]'s query parameters as
// `url.Values`.
func (r ToolSetSecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
