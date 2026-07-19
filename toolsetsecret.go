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
	options []option.RequestOption
}

// NewToolSetSecretService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewToolSetSecretService(opts ...option.RequestOption) (r ToolSetSecretService) {
	r = ToolSetSecretService{}
	r.options = opts
	return
}

// Creates a new secret scoped to the tool set
func (r *ToolSetSecretService) New(ctx context.Context, toolSetID string, params ToolSetSecretNewParams, opts ...option.RequestOption) (res *ToolSetSecret, err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a tool set secret by ID from the tool set
func (r *ToolSetSecretService) Get(ctx context.Context, toolSetID string, id string, query ToolSetSecretGetParams, opts ...option.RequestOption) (res *ToolSetSecret, err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(toolSetID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a secret scoped to the tool set
func (r *ToolSetSecretService) Update(ctx context.Context, toolSetID string, id string, params ToolSetSecretUpdateParams, opts ...option.RequestOption) (res *ToolSetSecret, err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all secrets scoped to the tool set
func (r *ToolSetSecretService) List(ctx context.Context, toolSetID string, params ToolSetSecretListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSetSecret], err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID))
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

// Lists all secrets scoped to the tool set
func (r *ToolSetSecretService) ListAutoPaging(ctx context.Context, toolSetID string, params ToolSetSecretListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSetSecret] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, toolSetID, params, opts...))
}

// Deletes a secret scoped to the tool set
func (r *ToolSetSecretService) Delete(ctx context.Context, toolSetID string, id string, body ToolSetSecretDeleteParams, opts ...option.RequestOption) (err error) {
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
	if toolSetID == "" {
		err = errors.New("missing required toolSetId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/secrets/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(toolSetID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ToolSetSecret struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     ToolSetSecretSpec       `json:"spec" api:"required"`
	// Tool set secret information
	Info ToolSetSecretInfo `json:"info"`
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
func (r ToolSetSecret) RawJSON() string { return r.JSON.raw }
func (r *ToolSetSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetSecretInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy  Profile   `json:"createdBy"`
	LastUsedAt time.Time `json:"lastUsedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		LastUsedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetSecretInfo) RawJSON() string { return r.JSON.raw }
func (r *ToolSetSecretInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetSecretSpec struct {
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetSecretSpec) RawJSON() string { return r.JSON.raw }
func (r *ToolSetSecretSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetSecretSpec to a ToolSetSecretSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetSecretSpecParam.Overrides()
func (r ToolSetSecretSpec) ToParam() ToolSetSecretSpecParam {
	return param.Override[ToolSetSecretSpecParam](json.RawMessage(r.RawJSON()))
}

type ToolSetSecretSpecParam struct {
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r ToolSetSecretSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetSecretSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetSecretSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetSecretNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	Spec     ToolSetSecretSpecParam             `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r ToolSetSecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetSecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetSecretNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetSecretGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetSecretUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Fields to update.
	UpdateMask param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	Spec     ToolSetSecretSpecParam             `json:"spec,omitzero"`
	paramObj
}

func (r ToolSetSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetSecretUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetSecretUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetSecretListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter expression (query param: prefix)
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Free-form search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolSetSecretListParams]'s query parameters as
// `url.Values`.
func (r ToolSetSecretListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolSetSecretDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
