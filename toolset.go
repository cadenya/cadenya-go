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
// ToolSetService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolSetService] method instead.
type ToolSetService struct {
	options []option.RequestOption
	// Manage tool sets and the tools they contain. Tool sets group related tools, and
	// tools define specific capabilities available to agents.
	//
	// When a tool set is managed, only API key actors can modify its tools; human
	// (profile) actors cannot.
	Tools ToolSetToolService
	// Manage tool sets and the tools they contain. Tool sets group related tools, and
	// tools define specific capabilities available to agents.
	//
	// When a tool set is managed, only API key actors can modify its tools; human
	// (profile) actors cannot.
	Secrets ToolSetSecretService
}

// NewToolSetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolSetService(opts ...option.RequestOption) (r ToolSetService) {
	r = ToolSetService{}
	r.options = opts
	r.Tools = NewToolSetToolService(opts...)
	r.Secrets = NewToolSetSecretService(opts...)
	return
}

// Creates a new tool set in the workspace
func (r *ToolSetService) New(ctx context.Context, params ToolSetNewParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a tool set by ID from the workspace
func (r *ToolSetService) Get(ctx context.Context, id string, query ToolSetGetParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a tool set in the workspace
func (r *ToolSetService) Update(ctx context.Context, id string, params ToolSetUpdateParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all tool sets in the workspace
func (r *ToolSetService) List(ctx context.Context, params ToolSetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSet], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets", url.PathEscape(params.WorkspaceID.Value))
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

// Lists all tool sets in the workspace
func (r *ToolSetService) ListAutoPaging(ctx context.Context, params ToolSetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSet] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes a tool set in the workspace
func (r *ToolSetService) Delete(ctx context.Context, id string, body ToolSetDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Transitions a tool set to STATE_ARCHIVED. Syncing stops, the tool set is hidden
// from list results, its tools are no longer offered to objectives, and new
// variation assignments are rejected. Existing assignments are retained, and
// history is preserved — unlike delete, archiving works while the tool set is
// still assigned to agent variations.
func (r *ToolSetService) Archive(ctx context.Context, id string, body ToolSetArchiveParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s:archive", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the current OpenAPI specification JSON that has been consumed by the
// tool set. Only applicable to tool sets using the OpenAPI adapter.
func (r *ToolSetService) GetOpenAPISpec(ctx context.Context, toolSetID string, query ToolSetGetOpenAPISpecParams, opts ...option.RequestOption) (res *ToolSetGetOpenAPISpecResponse, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/openapi_spec", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(toolSetID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all events (including sync status) for a tool set
func (r *ToolSetService) ListEvents(ctx context.Context, toolSetID string, params ToolSetListEventsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ToolSetEvent], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s/events", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(toolSetID))
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

// Lists all events (including sync status) for a tool set
func (r *ToolSetService) ListEventsAutoPaging(ctx context.Context, toolSetID string, params ToolSetListEventsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ToolSetEvent] {
	return pagination.NewCursorPaginationAutoPager(r.ListEvents(ctx, toolSetID, params, opts...))
}

// Transitions an archived tool set back to STATE_ACTIVE. Managed tool sets resume
// syncing on their next cycle and their tools become available to objectives
// again.
func (r *ToolSetService) Unarchive(ctx context.Context, id string, body ToolSetUnarchiveParams, opts ...option.RequestOption) (res *ToolSet, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/tool_sets/%s:unarchive", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// ApprovalRequirementFilterUnion contains all possible properties and values from
// [ApprovalRequirementFilterAlways], [ApprovalRequirementFilterOnly].
//
// Use the [ApprovalRequirementFilterUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ApprovalRequirementFilterUnion struct {
	// This field is from variant [ApprovalRequirementFilterAlways].
	Always bool `json:"always"`
	// Any of "always", "only".
	Type string `json:"type"`
	// This field is from variant [ApprovalRequirementFilterOnly].
	Only ToolFilter `json:"only"`
	JSON struct {
		Always respjson.Field
		Type   respjson.Field
		Only   respjson.Field
		raw    string
	} `json:"-"`
}

// anyApprovalRequirementFilter is implemented by each variant of
// [ApprovalRequirementFilterUnion] to add type safety for the return type of
// [ApprovalRequirementFilterUnion.AsAny]
type anyApprovalRequirementFilter interface {
	implApprovalRequirementFilterUnion()
}

func (ApprovalRequirementFilterAlways) implApprovalRequirementFilterUnion() {}
func (ApprovalRequirementFilterOnly) implApprovalRequirementFilterUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ApprovalRequirementFilterUnion.AsAny().(type) {
//	case cadenya.ApprovalRequirementFilterAlways:
//	case cadenya.ApprovalRequirementFilterOnly:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ApprovalRequirementFilterUnion) AsAny() anyApprovalRequirementFilter {
	switch u.Type {
	case "always":
		return u.AsAlways()
	case "only":
		return u.AsOnly()
	}
	return nil
}

func (u ApprovalRequirementFilterUnion) AsAlways() (v ApprovalRequirementFilterAlways) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ApprovalRequirementFilterUnion) AsOnly() (v ApprovalRequirementFilterOnly) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ApprovalRequirementFilterUnion) RawJSON() string { return u.JSON.raw }

func (r *ApprovalRequirementFilterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ApprovalRequirementFilterUnion to a
// ApprovalRequirementFilterUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ApprovalRequirementFilterUnionParam.Overrides()
func (r ApprovalRequirementFilterUnion) ToParam() ApprovalRequirementFilterUnionParam {
	return param.Override[ApprovalRequirementFilterUnionParam](json.RawMessage(r.RawJSON()))
}

func ApprovalRequirementFilterParamOfAlways(always bool) ApprovalRequirementFilterUnionParam {
	var variant ApprovalRequirementFilterAlwaysParam
	variant.Always = always
	return ApprovalRequirementFilterUnionParam{OfAlways: &variant}
}

func ApprovalRequirementFilterParamOfOnly(only ToolFilterParam) ApprovalRequirementFilterUnionParam {
	var variant ApprovalRequirementFilterOnlyParam
	variant.Only = only
	return ApprovalRequirementFilterUnionParam{OfOnly: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ApprovalRequirementFilterUnionParam struct {
	OfAlways *ApprovalRequirementFilterAlwaysParam `json:",omitzero,inline"`
	OfOnly   *ApprovalRequirementFilterOnlyParam   `json:",omitzero,inline"`
	paramUnion
}

func (u ApprovalRequirementFilterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAlways, u.OfOnly)
}
func (u *ApprovalRequirementFilterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ApprovalRequirementFilterUnionParam](
		"type",
		apijson.Discriminator[ApprovalRequirementFilterAlwaysParam]("always"),
		apijson.Discriminator[ApprovalRequirementFilterOnlyParam]("only"),
	)
}

type ApprovalRequirementFilterAlways struct {
	Always bool `json:"always" api:"required"`
	// Any of "always".
	Type ApprovalRequirementFilterAlwaysType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Always      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalRequirementFilterAlways) RawJSON() string { return r.JSON.raw }
func (r *ApprovalRequirementFilterAlways) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ApprovalRequirementFilterAlways to a
// ApprovalRequirementFilterAlwaysParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ApprovalRequirementFilterAlwaysParam.Overrides()
func (r ApprovalRequirementFilterAlways) ToParam() ApprovalRequirementFilterAlwaysParam {
	return param.Override[ApprovalRequirementFilterAlwaysParam](json.RawMessage(r.RawJSON()))
}

type ApprovalRequirementFilterAlwaysType string

const (
	ApprovalRequirementFilterAlwaysTypeAlways ApprovalRequirementFilterAlwaysType = "always"
)

// The properties Always, Type are required.
type ApprovalRequirementFilterAlwaysParam struct {
	Always bool `json:"always" api:"required"`
	// Any of "always".
	Type ApprovalRequirementFilterAlwaysType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ApprovalRequirementFilterAlwaysParam) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalRequirementFilterAlwaysParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalRequirementFilterAlwaysParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalRequirementFilterOnly struct {
	// Top-level filter with simple boolean logic (no nesting)
	Only ToolFilter `json:"only" api:"required"`
	// Any of "only".
	Type ApprovalRequirementFilterOnlyType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Only        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalRequirementFilterOnly) RawJSON() string { return r.JSON.raw }
func (r *ApprovalRequirementFilterOnly) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ApprovalRequirementFilterOnly to a
// ApprovalRequirementFilterOnlyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ApprovalRequirementFilterOnlyParam.Overrides()
func (r ApprovalRequirementFilterOnly) ToParam() ApprovalRequirementFilterOnlyParam {
	return param.Override[ApprovalRequirementFilterOnlyParam](json.RawMessage(r.RawJSON()))
}

type ApprovalRequirementFilterOnlyType string

const (
	ApprovalRequirementFilterOnlyTypeOnly ApprovalRequirementFilterOnlyType = "only"
)

// The properties Only, Type are required.
type ApprovalRequirementFilterOnlyParam struct {
	// Top-level filter with simple boolean logic (no nesting)
	Only ToolFilterParam `json:"only,omitzero" api:"required"`
	// Any of "only".
	Type ApprovalRequirementFilterOnlyType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ApprovalRequirementFilterOnlyParam) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalRequirementFilterOnlyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalRequirementFilterOnlyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Single attribute filter
type AttributeFilter struct {
	// Any of "ATTRIBUTE_UNSPECIFIED", "ATTRIBUTE_NAME", "ATTRIBUTE_TITLE",
	// "ATTRIBUTE_DESCRIPTION".
	Attribute AttributeFilterAttribute `json:"attribute" api:"required"`
	// String matching operations
	Matcher StringMatcherUnion `json:"matcher"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attribute   respjson.Field
		Matcher     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AttributeFilter) RawJSON() string { return r.JSON.raw }
func (r *AttributeFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AttributeFilter to a AttributeFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AttributeFilterParam.Overrides()
func (r AttributeFilter) ToParam() AttributeFilterParam {
	return param.Override[AttributeFilterParam](json.RawMessage(r.RawJSON()))
}

type AttributeFilterAttribute string

const (
	AttributeFilterAttributeAttributeUnspecified AttributeFilterAttribute = "ATTRIBUTE_UNSPECIFIED"
	AttributeFilterAttributeAttributeName        AttributeFilterAttribute = "ATTRIBUTE_NAME"
	AttributeFilterAttributeAttributeTitle       AttributeFilterAttribute = "ATTRIBUTE_TITLE"
	AttributeFilterAttributeAttributeDescription AttributeFilterAttribute = "ATTRIBUTE_DESCRIPTION"
)

// Single attribute filter
//
// The property Attribute is required.
type AttributeFilterParam struct {
	// Any of "ATTRIBUTE_UNSPECIFIED", "ATTRIBUTE_NAME", "ATTRIBUTE_TITLE",
	// "ATTRIBUTE_DESCRIPTION".
	Attribute AttributeFilterAttribute `json:"attribute,omitzero" api:"required"`
	// String matching operations
	Matcher StringMatcherUnionParam `json:"matcher,omitzero"`
	paramObj
}

func (r AttributeFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow AttributeFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AttributeFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// StringMatcherUnion contains all possible properties and values from
// [StringMatcherExact], [StringMatcherStartsWith], [StringMatcherEndsWith],
// [StringMatcherContains], [StringMatcherRegex].
//
// Use the [StringMatcherUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type StringMatcherUnion struct {
	// This field is from variant [StringMatcherExact].
	Exact string `json:"exact"`
	// Any of "exact", "startsWith", "endsWith", "contains", "regex".
	Type          string `json:"type"`
	CaseSensitive bool   `json:"caseSensitive"`
	// This field is from variant [StringMatcherStartsWith].
	StartsWith string `json:"startsWith"`
	// This field is from variant [StringMatcherEndsWith].
	EndsWith string `json:"endsWith"`
	// This field is from variant [StringMatcherContains].
	Contains string `json:"contains"`
	// This field is from variant [StringMatcherRegex].
	Regex string `json:"regex"`
	JSON  struct {
		Exact         respjson.Field
		Type          respjson.Field
		CaseSensitive respjson.Field
		StartsWith    respjson.Field
		EndsWith      respjson.Field
		Contains      respjson.Field
		Regex         respjson.Field
		raw           string
	} `json:"-"`
}

// anyStringMatcher is implemented by each variant of [StringMatcherUnion] to add
// type safety for the return type of [StringMatcherUnion.AsAny]
type anyStringMatcher interface {
	implStringMatcherUnion()
}

func (StringMatcherExact) implStringMatcherUnion()      {}
func (StringMatcherStartsWith) implStringMatcherUnion() {}
func (StringMatcherEndsWith) implStringMatcherUnion()   {}
func (StringMatcherContains) implStringMatcherUnion()   {}
func (StringMatcherRegex) implStringMatcherUnion()      {}

// Use the following switch statement to find the correct variant
//
//	switch variant := StringMatcherUnion.AsAny().(type) {
//	case cadenya.StringMatcherExact:
//	case cadenya.StringMatcherStartsWith:
//	case cadenya.StringMatcherEndsWith:
//	case cadenya.StringMatcherContains:
//	case cadenya.StringMatcherRegex:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u StringMatcherUnion) AsAny() anyStringMatcher {
	switch u.Type {
	case "exact":
		return u.AsExact()
	case "startsWith":
		return u.AsStartsWith()
	case "endsWith":
		return u.AsEndsWith()
	case "contains":
		return u.AsContains()
	case "regex":
		return u.AsRegex()
	}
	return nil
}

func (u StringMatcherUnion) AsExact() (v StringMatcherExact) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StringMatcherUnion) AsStartsWith() (v StringMatcherStartsWith) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StringMatcherUnion) AsEndsWith() (v StringMatcherEndsWith) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StringMatcherUnion) AsContains() (v StringMatcherContains) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StringMatcherUnion) AsRegex() (v StringMatcherRegex) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u StringMatcherUnion) RawJSON() string { return u.JSON.raw }

func (r *StringMatcherUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringMatcherUnion to a StringMatcherUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringMatcherUnionParam.Overrides()
func (r StringMatcherUnion) ToParam() StringMatcherUnionParam {
	return param.Override[StringMatcherUnionParam](json.RawMessage(r.RawJSON()))
}

func StringMatcherParamOfExact(exact string) StringMatcherUnionParam {
	var variant StringMatcherExactParam
	variant.Exact = exact
	return StringMatcherUnionParam{OfExact: &variant}
}

func StringMatcherParamOfStartsWith(startsWith string) StringMatcherUnionParam {
	var variant StringMatcherStartsWithParam
	variant.StartsWith = startsWith
	return StringMatcherUnionParam{OfStartsWith: &variant}
}

func StringMatcherParamOfEndsWith(endsWith string) StringMatcherUnionParam {
	var variant StringMatcherEndsWithParam
	variant.EndsWith = endsWith
	return StringMatcherUnionParam{OfEndsWith: &variant}
}

func StringMatcherParamOfContains(contains string) StringMatcherUnionParam {
	var variant StringMatcherContainsParam
	variant.Contains = contains
	return StringMatcherUnionParam{OfContains: &variant}
}

func StringMatcherParamOfRegex(regex string) StringMatcherUnionParam {
	var variant StringMatcherRegexParam
	variant.Regex = regex
	return StringMatcherUnionParam{OfRegex: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type StringMatcherUnionParam struct {
	OfExact      *StringMatcherExactParam      `json:",omitzero,inline"`
	OfStartsWith *StringMatcherStartsWithParam `json:",omitzero,inline"`
	OfEndsWith   *StringMatcherEndsWithParam   `json:",omitzero,inline"`
	OfContains   *StringMatcherContainsParam   `json:",omitzero,inline"`
	OfRegex      *StringMatcherRegexParam      `json:",omitzero,inline"`
	paramUnion
}

func (u StringMatcherUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExact,
		u.OfStartsWith,
		u.OfEndsWith,
		u.OfContains,
		u.OfRegex)
}
func (u *StringMatcherUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[StringMatcherUnionParam](
		"type",
		apijson.Discriminator[StringMatcherExactParam]("exact"),
		apijson.Discriminator[StringMatcherStartsWithParam]("startsWith"),
		apijson.Discriminator[StringMatcherEndsWithParam]("endsWith"),
		apijson.Discriminator[StringMatcherContainsParam]("contains"),
		apijson.Discriminator[StringMatcherRegexParam]("regex"),
	)
}

type StringMatcherContains struct {
	Contains string `json:"contains" api:"required"`
	// Any of "contains".
	Type          StringMatcherContainsType `json:"type" api:"required"`
	CaseSensitive bool                      `json:"caseSensitive"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contains      respjson.Field
		Type          respjson.Field
		CaseSensitive respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringMatcherContains) RawJSON() string { return r.JSON.raw }
func (r *StringMatcherContains) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringMatcherContains to a StringMatcherContainsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringMatcherContainsParam.Overrides()
func (r StringMatcherContains) ToParam() StringMatcherContainsParam {
	return param.Override[StringMatcherContainsParam](json.RawMessage(r.RawJSON()))
}

type StringMatcherContainsType string

const (
	StringMatcherContainsTypeContains StringMatcherContainsType = "contains"
)

// The properties Contains, Type are required.
type StringMatcherContainsParam struct {
	Contains string `json:"contains" api:"required"`
	// Any of "contains".
	Type          StringMatcherContainsType `json:"type,omitzero" api:"required"`
	CaseSensitive param.Opt[bool]           `json:"caseSensitive,omitzero"`
	paramObj
}

func (r StringMatcherContainsParam) MarshalJSON() (data []byte, err error) {
	type shadow StringMatcherContainsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StringMatcherContainsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringMatcherEndsWith struct {
	EndsWith string `json:"endsWith" api:"required"`
	// Any of "endsWith".
	Type          StringMatcherEndsWithType `json:"type" api:"required"`
	CaseSensitive bool                      `json:"caseSensitive"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndsWith      respjson.Field
		Type          respjson.Field
		CaseSensitive respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringMatcherEndsWith) RawJSON() string { return r.JSON.raw }
func (r *StringMatcherEndsWith) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringMatcherEndsWith to a StringMatcherEndsWithParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringMatcherEndsWithParam.Overrides()
func (r StringMatcherEndsWith) ToParam() StringMatcherEndsWithParam {
	return param.Override[StringMatcherEndsWithParam](json.RawMessage(r.RawJSON()))
}

type StringMatcherEndsWithType string

const (
	StringMatcherEndsWithTypeEndsWith StringMatcherEndsWithType = "endsWith"
)

// The properties EndsWith, Type are required.
type StringMatcherEndsWithParam struct {
	EndsWith string `json:"endsWith" api:"required"`
	// Any of "endsWith".
	Type          StringMatcherEndsWithType `json:"type,omitzero" api:"required"`
	CaseSensitive param.Opt[bool]           `json:"caseSensitive,omitzero"`
	paramObj
}

func (r StringMatcherEndsWithParam) MarshalJSON() (data []byte, err error) {
	type shadow StringMatcherEndsWithParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StringMatcherEndsWithParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringMatcherExact struct {
	Exact string `json:"exact" api:"required"`
	// Any of "exact".
	Type          StringMatcherExactType `json:"type" api:"required"`
	CaseSensitive bool                   `json:"caseSensitive"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exact         respjson.Field
		Type          respjson.Field
		CaseSensitive respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringMatcherExact) RawJSON() string { return r.JSON.raw }
func (r *StringMatcherExact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringMatcherExact to a StringMatcherExactParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringMatcherExactParam.Overrides()
func (r StringMatcherExact) ToParam() StringMatcherExactParam {
	return param.Override[StringMatcherExactParam](json.RawMessage(r.RawJSON()))
}

type StringMatcherExactType string

const (
	StringMatcherExactTypeExact StringMatcherExactType = "exact"
)

// The properties Exact, Type are required.
type StringMatcherExactParam struct {
	Exact string `json:"exact" api:"required"`
	// Any of "exact".
	Type          StringMatcherExactType `json:"type,omitzero" api:"required"`
	CaseSensitive param.Opt[bool]        `json:"caseSensitive,omitzero"`
	paramObj
}

func (r StringMatcherExactParam) MarshalJSON() (data []byte, err error) {
	type shadow StringMatcherExactParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StringMatcherExactParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringMatcherRegex struct {
	Regex string `json:"regex" api:"required"`
	// Any of "regex".
	Type          StringMatcherRegexType `json:"type" api:"required"`
	CaseSensitive bool                   `json:"caseSensitive"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Regex         respjson.Field
		Type          respjson.Field
		CaseSensitive respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringMatcherRegex) RawJSON() string { return r.JSON.raw }
func (r *StringMatcherRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringMatcherRegex to a StringMatcherRegexParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringMatcherRegexParam.Overrides()
func (r StringMatcherRegex) ToParam() StringMatcherRegexParam {
	return param.Override[StringMatcherRegexParam](json.RawMessage(r.RawJSON()))
}

type StringMatcherRegexType string

const (
	StringMatcherRegexTypeRegex StringMatcherRegexType = "regex"
)

// The properties Regex, Type are required.
type StringMatcherRegexParam struct {
	Regex string `json:"regex" api:"required"`
	// Any of "regex".
	Type          StringMatcherRegexType `json:"type,omitzero" api:"required"`
	CaseSensitive param.Opt[bool]        `json:"caseSensitive,omitzero"`
	paramObj
}

func (r StringMatcherRegexParam) MarshalJSON() (data []byte, err error) {
	type shadow StringMatcherRegexParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StringMatcherRegexParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StringMatcherStartsWith struct {
	StartsWith string `json:"startsWith" api:"required"`
	// Any of "startsWith".
	Type          StringMatcherStartsWithType `json:"type" api:"required"`
	CaseSensitive bool                        `json:"caseSensitive"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StartsWith    respjson.Field
		Type          respjson.Field
		CaseSensitive respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StringMatcherStartsWith) RawJSON() string { return r.JSON.raw }
func (r *StringMatcherStartsWith) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this StringMatcherStartsWith to a StringMatcherStartsWithParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// StringMatcherStartsWithParam.Overrides()
func (r StringMatcherStartsWith) ToParam() StringMatcherStartsWithParam {
	return param.Override[StringMatcherStartsWithParam](json.RawMessage(r.RawJSON()))
}

type StringMatcherStartsWithType string

const (
	StringMatcherStartsWithTypeStartsWith StringMatcherStartsWithType = "startsWith"
)

// The properties StartsWith, Type are required.
type StringMatcherStartsWithParam struct {
	StartsWith string `json:"startsWith" api:"required"`
	// Any of "startsWith".
	Type          StringMatcherStartsWithType `json:"type,omitzero" api:"required"`
	CaseSensitive param.Opt[bool]             `json:"caseSensitive,omitzero"`
	paramObj
}

func (r StringMatcherStartsWithParam) MarshalJSON() (data []byte, err error) {
	type shadow StringMatcherStartsWithParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StringMatcherStartsWithParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emitted when a tool set sync operation completes successfully.
type SyncCompleted struct {
	// Optional message with additional details.
	Message string `json:"message"`
	// Number of tools synced.
	ToolsSynced int64 `json:"toolsSynced"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ToolsSynced respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyncCompleted) RawJSON() string { return r.JSON.raw }
func (r *SyncCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emitted when a tool set sync operation fails.
type SyncFailed struct {
	// Indicates this is an error event.
	Error bool `json:"error"`
	// Optional error type/code for programmatic handling.
	ErrorType string `json:"errorType"`
	// Error message describing what went wrong.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ErrorType   respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyncFailed) RawJSON() string { return r.JSON.raw }
func (r *SyncFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Emitted when a tool set sync operation begins.
type SyncStarted struct {
	// Human-readable message describing the start of the sync.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SyncStarted) RawJSON() string { return r.JSON.raw }
func (r *SyncStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-level filter with simple boolean logic (no nesting)
type ToolFilter struct {
	// Any of "OPERATOR_UNSPECIFIED", "OPERATOR_AND", "OPERATOR_OR".
	Operator ToolFilterOperator `json:"operator" api:"required"`
	Filters  []AttributeFilter  `json:"filters"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Operator    respjson.Field
		Filters     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolFilter) RawJSON() string { return r.JSON.raw }
func (r *ToolFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolFilter to a ToolFilterParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolFilterParam.Overrides()
func (r ToolFilter) ToParam() ToolFilterParam {
	return param.Override[ToolFilterParam](json.RawMessage(r.RawJSON()))
}

type ToolFilterOperator string

const (
	ToolFilterOperatorOperatorUnspecified ToolFilterOperator = "OPERATOR_UNSPECIFIED"
	ToolFilterOperatorOperatorAnd         ToolFilterOperator = "OPERATOR_AND"
	ToolFilterOperatorOperatorOr          ToolFilterOperator = "OPERATOR_OR"
)

// Top-level filter with simple boolean logic (no nesting)
//
// The property Operator is required.
type ToolFilterParam struct {
	// Any of "OPERATOR_UNSPECIFIED", "OPERATOR_AND", "OPERATOR_OR".
	Operator ToolFilterOperator     `json:"operator,omitzero" api:"required"`
	Filters  []AttributeFilterParam `json:"filters,omitzero"`
	paramObj
}

func (r ToolFilterParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolFilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolFilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSet struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	Spec     ToolSetSpec             `json:"spec" api:"required"`
	// The current lifecycle state of the tool set. Output only. Tool sets are created
	// STATE_ACTIVE; use the :archive and :unarchive actions to transition between
	// states.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ACTIVE", "STATE_ARCHIVED".
	State ToolSetState `json:"state" api:"required"`
	// Tool set information
	Info ToolSetInfo `json:"info"`
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
func (r ToolSet) RawJSON() string { return r.JSON.raw }
func (r *ToolSet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current lifecycle state of the tool set. Output only. Tool sets are created
// STATE_ACTIVE; use the :archive and :unarchive actions to transition between
// states.
type ToolSetState string

const (
	ToolSetStateStateUnspecified ToolSetState = "STATE_UNSPECIFIED"
	ToolSetStateStateActive      ToolSetState = "STATE_ACTIVE"
	ToolSetStateStateArchived    ToolSetState = "STATE_ARCHIVED"
)

// ToolSetAdapterUnion contains all possible properties and values from
// [ToolSetAdapterMCPVariant], [ToolSetAdapterHTTPVariant],
// [ToolSetAdapterOpenAPIVariant], [ToolSetAdapterBareVariant].
//
// Use the [ToolSetAdapterUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToolSetAdapterUnion struct {
	// This field is from variant [ToolSetAdapterMCPVariant].
	MCP ToolSetAdapterMCP `json:"mcp"`
	// Any of "mcp", "http", "openapi", "bare".
	Type string `json:"type"`
	// This field is from variant [ToolSetAdapterHTTPVariant].
	HTTP ToolSetAdapterHTTP `json:"http"`
	// This field is from variant [ToolSetAdapterOpenAPIVariant].
	OpenAPI ToolSetAdapterOpenAPIUnion `json:"openapi"`
	// This field is from variant [ToolSetAdapterBareVariant].
	Bare ToolSetAdapterBare `json:"bare"`
	JSON struct {
		MCP     respjson.Field
		Type    respjson.Field
		HTTP    respjson.Field
		OpenAPI respjson.Field
		Bare    respjson.Field
		raw     string
	} `json:"-"`
}

// anyToolSetAdapter is implemented by each variant of [ToolSetAdapterUnion] to add
// type safety for the return type of [ToolSetAdapterUnion.AsAny]
type anyToolSetAdapter interface {
	implToolSetAdapterUnion()
}

func (ToolSetAdapterMCPVariant) implToolSetAdapterUnion()     {}
func (ToolSetAdapterHTTPVariant) implToolSetAdapterUnion()    {}
func (ToolSetAdapterOpenAPIVariant) implToolSetAdapterUnion() {}
func (ToolSetAdapterBareVariant) implToolSetAdapterUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ToolSetAdapterUnion.AsAny().(type) {
//	case cadenya.ToolSetAdapterMCPVariant:
//	case cadenya.ToolSetAdapterHTTPVariant:
//	case cadenya.ToolSetAdapterOpenAPIVariant:
//	case cadenya.ToolSetAdapterBareVariant:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ToolSetAdapterUnion) AsAny() anyToolSetAdapter {
	switch u.Type {
	case "mcp":
		return u.AsMCP()
	case "http":
		return u.AsHTTP()
	case "openapi":
		return u.AsOpenAPI()
	case "bare":
		return u.AsBare()
	}
	return nil
}

func (u ToolSetAdapterUnion) AsMCP() (v ToolSetAdapterMCPVariant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSetAdapterUnion) AsHTTP() (v ToolSetAdapterHTTPVariant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSetAdapterUnion) AsOpenAPI() (v ToolSetAdapterOpenAPIVariant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSetAdapterUnion) AsBare() (v ToolSetAdapterBareVariant) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolSetAdapterUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolSetAdapterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterUnion to a ToolSetAdapterUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterUnionParam.Overrides()
func (r ToolSetAdapterUnion) ToParam() ToolSetAdapterUnionParam {
	return param.Override[ToolSetAdapterUnionParam](json.RawMessage(r.RawJSON()))
}

func ToolSetAdapterParamOfMCP(mcp ToolSetAdapterMCPParam) ToolSetAdapterUnionParam {
	var variant ToolSetAdapterMCPVariantParam
	variant.MCP = mcp
	return ToolSetAdapterUnionParam{OfMCP: &variant}
}

func ToolSetAdapterParamOfHTTP(http ToolSetAdapterHTTPParam) ToolSetAdapterUnionParam {
	var variant ToolSetAdapterHTTPVariantParam
	variant.HTTP = http
	return ToolSetAdapterUnionParam{OfHTTP: &variant}
}

func ToolSetAdapterParamOfOpenAPI[
	T ToolSetAdapterOpenAPIURLParam | ToolSetAdapterOpenAPIUploadIDParam,
](openAPI T) ToolSetAdapterUnionParam {
	var variant ToolSetAdapterOpenAPIVariantParam
	switch v := any(openAPI).(type) {
	case ToolSetAdapterOpenAPIURLParam:
		variant.OpenAPI.OfURL = &v
	case ToolSetAdapterOpenAPIUploadIDParam:
		variant.OpenAPI.OfUploadID = &v
	}
	return ToolSetAdapterUnionParam{OfOpenAPI: &variant}
}

func ToolSetAdapterParamOfBare(bare ToolSetAdapterBareParam) ToolSetAdapterUnionParam {
	var variant ToolSetAdapterBareVariantParam
	variant.Bare = bare
	return ToolSetAdapterUnionParam{OfBare: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolSetAdapterUnionParam struct {
	OfMCP     *ToolSetAdapterMCPVariantParam     `json:",omitzero,inline"`
	OfHTTP    *ToolSetAdapterHTTPVariantParam    `json:",omitzero,inline"`
	OfOpenAPI *ToolSetAdapterOpenAPIVariantParam `json:",omitzero,inline"`
	OfBare    *ToolSetAdapterBareVariantParam    `json:",omitzero,inline"`
	paramUnion
}

func (u ToolSetAdapterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMCP, u.OfHTTP, u.OfOpenAPI, u.OfBare)
}
func (u *ToolSetAdapterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ToolSetAdapterUnionParam](
		"type",
		apijson.Discriminator[ToolSetAdapterMCPVariantParam]("mcp"),
		apijson.Discriminator[ToolSetAdapterHTTPVariantParam]("http"),
		apijson.Discriminator[ToolSetAdapterOpenAPIVariantParam]("openapi"),
		apijson.Discriminator[ToolSetAdapterBareVariantParam]("bare"),
	)
}

// Bare tool sets define tools without an execution adapter. A bare tool call
// doesn't fire anything: the objective's workflow pauses and waits for an external
// API consumer to set the tool call's content (e.g. human-in-the-loop tools, or a
// reverse harness that polls for pending tool calls, executes locally, and reports
// results back via SetToolCallContent).
type ToolSetAdapterBare struct {
	// How long to wait for content to be set before the tool call errors. If unset,
	// the call waits indefinitely.
	ContentTimeout int64 `json:"contentTimeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentTimeout respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterBare) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterBare) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterBare to a ToolSetAdapterBareParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterBareParam.Overrides()
func (r ToolSetAdapterBare) ToParam() ToolSetAdapterBareParam {
	return param.Override[ToolSetAdapterBareParam](json.RawMessage(r.RawJSON()))
}

// Bare tool sets define tools without an execution adapter. A bare tool call
// doesn't fire anything: the objective's workflow pauses and waits for an external
// API consumer to set the tool call's content (e.g. human-in-the-loop tools, or a
// reverse harness that polls for pending tool calls, executes locally, and reports
// results back via SetToolCallContent).
type ToolSetAdapterBareParam struct {
	// How long to wait for content to be set before the tool call errors. If unset,
	// the call waits indefinitely.
	ContentTimeout param.Opt[int64] `json:"contentTimeout,omitzero"`
	paramObj
}

func (r ToolSetAdapterBareParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterBareParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterBareParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterBareVariant struct {
	// Bare tool sets define tools without an execution adapter. A bare tool call
	// doesn't fire anything: the objective's workflow pauses and waits for an external
	// API consumer to set the tool call's content (e.g. human-in-the-loop tools, or a
	// reverse harness that polls for pending tool calls, executes locally, and reports
	// results back via SetToolCallContent).
	Bare ToolSetAdapterBare `json:"bare" api:"required"`
	// Any of "bare".
	Type ToolSetAdapterBareVariantType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bare        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterBareVariant) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterBareVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterBareVariant to a
// ToolSetAdapterBareVariantParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterBareVariantParam.Overrides()
func (r ToolSetAdapterBareVariant) ToParam() ToolSetAdapterBareVariantParam {
	return param.Override[ToolSetAdapterBareVariantParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterBareVariantType string

const (
	ToolSetAdapterBareVariantTypeBare ToolSetAdapterBareVariantType = "bare"
)

// The properties Bare, Type are required.
type ToolSetAdapterBareVariantParam struct {
	// Bare tool sets define tools without an execution adapter. A bare tool call
	// doesn't fire anything: the objective's workflow pauses and waits for an external
	// API consumer to set the tool call's content (e.g. human-in-the-loop tools, or a
	// reverse harness that polls for pending tool calls, executes locally, and reports
	// results back via SetToolCallContent).
	Bare ToolSetAdapterBareParam `json:"bare,omitzero" api:"required"`
	// Any of "bare".
	Type ToolSetAdapterBareVariantType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolSetAdapterBareVariantParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterBareVariantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterBareVariantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterHTTP struct {
	BaseURL string            `json:"baseUrl"`
	Headers map[string]string `json:"headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL     respjson.Field
		Headers     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterHTTP) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterHTTP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterHTTP to a ToolSetAdapterHTTPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterHTTPParam.Overrides()
func (r ToolSetAdapterHTTP) ToParam() ToolSetAdapterHTTPParam {
	return param.Override[ToolSetAdapterHTTPParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterHTTPParam struct {
	BaseURL param.Opt[string] `json:"baseUrl,omitzero"`
	Headers map[string]string `json:"headers,omitzero"`
	paramObj
}

func (r ToolSetAdapterHTTPParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterHTTPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterHTTPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterHTTPVariant struct {
	HTTP ToolSetAdapterHTTP `json:"http" api:"required"`
	// Any of "http".
	Type ToolSetAdapterHTTPVariantType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HTTP        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterHTTPVariant) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterHTTPVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterHTTPVariant to a
// ToolSetAdapterHTTPVariantParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterHTTPVariantParam.Overrides()
func (r ToolSetAdapterHTTPVariant) ToParam() ToolSetAdapterHTTPVariantParam {
	return param.Override[ToolSetAdapterHTTPVariantParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterHTTPVariantType string

const (
	ToolSetAdapterHTTPVariantTypeHTTP ToolSetAdapterHTTPVariantType = "http"
)

// The properties HTTP, Type are required.
type ToolSetAdapterHTTPVariantParam struct {
	HTTP ToolSetAdapterHTTPParam `json:"http,omitzero" api:"required"`
	// Any of "http".
	Type ToolSetAdapterHTTPVariantType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolSetAdapterHTTPVariantParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterHTTPVariantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterHTTPVariantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterMCP struct {
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilter        `json:"excludeTools"`
	Headers      map[string]string `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilter `json:"includeTools"`
	// Defines behavior for just-in-time capable tool set adapters (IE: MCP).
	JustInTime ToolSetAdapterMCPJustInTime `json:"justInTime"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilterUnion `json:"toolApprovals"`
	URL           string                         `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExcludeTools  respjson.Field
		Headers       respjson.Field
		IncludeTools  respjson.Field
		JustInTime    respjson.Field
		ToolApprovals respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterMCP) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterMCP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterMCP to a ToolSetAdapterMCPParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterMCPParam.Overrides()
func (r ToolSetAdapterMCP) ToParam() ToolSetAdapterMCPParam {
	return param.Override[ToolSetAdapterMCPParam](json.RawMessage(r.RawJSON()))
}

// Defines behavior for just-in-time capable tool set adapters (IE: MCP).
type ToolSetAdapterMCPJustInTime struct {
	Enabled bool `json:"enabled"`
	// If set, an objective will automatically be failed if tools cannot be loaded in
	// the initial stages of an objective being created. Tools are loaded
	// asynchronously, so this setting is useful for ensuring that an objective
	// continued any further if tools are not available.
	FailObjectiveOnToolListError bool `json:"failObjectiveOnToolListError"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled                      respjson.Field
		FailObjectiveOnToolListError respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterMCPJustInTime) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterMCPJustInTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterMCPParam struct {
	URL param.Opt[string] `json:"url,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilterParam   `json:"excludeTools,omitzero"`
	Headers      map[string]string `json:"headers,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilterParam `json:"includeTools,omitzero"`
	// Defines behavior for just-in-time capable tool set adapters (IE: MCP).
	JustInTime ToolSetAdapterMCPJustInTimeParam `json:"justInTime,omitzero"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilterUnionParam `json:"toolApprovals,omitzero"`
	paramObj
}

func (r ToolSetAdapterMCPParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterMCPParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterMCPParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines behavior for just-in-time capable tool set adapters (IE: MCP).
type ToolSetAdapterMCPJustInTimeParam struct {
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// If set, an objective will automatically be failed if tools cannot be loaded in
	// the initial stages of an objective being created. Tools are loaded
	// asynchronously, so this setting is useful for ensuring that an objective
	// continued any further if tools are not available.
	FailObjectiveOnToolListError param.Opt[bool] `json:"failObjectiveOnToolListError,omitzero"`
	paramObj
}

func (r ToolSetAdapterMCPJustInTimeParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterMCPJustInTimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterMCPJustInTimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterMCPVariant struct {
	MCP ToolSetAdapterMCP `json:"mcp" api:"required"`
	// Any of "mcp".
	Type ToolSetAdapterMCPVariantType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MCP         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterMCPVariant) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterMCPVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterMCPVariant to a
// ToolSetAdapterMCPVariantParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterMCPVariantParam.Overrides()
func (r ToolSetAdapterMCPVariant) ToParam() ToolSetAdapterMCPVariantParam {
	return param.Override[ToolSetAdapterMCPVariantParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterMCPVariantType string

const (
	ToolSetAdapterMCPVariantTypeMCP ToolSetAdapterMCPVariantType = "mcp"
)

// The properties MCP, Type are required.
type ToolSetAdapterMCPVariantParam struct {
	MCP ToolSetAdapterMCPParam `json:"mcp,omitzero" api:"required"`
	// Any of "mcp".
	Type ToolSetAdapterMCPVariantType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolSetAdapterMCPVariantParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterMCPVariantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterMCPVariantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolSetAdapterOpenAPIUnion contains all possible properties and values from
// [ToolSetAdapterOpenAPIURL], [ToolSetAdapterOpenAPIUploadID].
//
// Use the [ToolSetAdapterOpenAPIUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToolSetAdapterOpenAPIUnion struct {
	// Any of "url", "uploadId".
	Type string `json:"type"`
	// This field is from variant [ToolSetAdapterOpenAPIURL].
	URL     string `json:"url"`
	BaseURL string `json:"baseUrl"`
	// This field is from variant [ToolSetAdapterOpenAPIURL].
	ExcludeTools ToolFilter `json:"excludeTools"`
	Headers      string     `json:"headers"`
	// This field is from variant [ToolSetAdapterOpenAPIURL].
	IncludeTools ToolFilter `json:"includeTools"`
	ServerName   string     `json:"serverName"`
	// This field is from variant [ToolSetAdapterOpenAPIURL].
	ToolApprovals ApprovalRequirementFilterUnion `json:"toolApprovals"`
	// This field is from variant [ToolSetAdapterOpenAPIUploadID].
	UploadID string `json:"uploadId"`
	JSON     struct {
		Type          respjson.Field
		URL           respjson.Field
		BaseURL       respjson.Field
		ExcludeTools  respjson.Field
		Headers       respjson.Field
		IncludeTools  respjson.Field
		ServerName    respjson.Field
		ToolApprovals respjson.Field
		UploadID      respjson.Field
		raw           string
	} `json:"-"`
}

// anyToolSetAdapterOpenAPI is implemented by each variant of
// [ToolSetAdapterOpenAPIUnion] to add type safety for the return type of
// [ToolSetAdapterOpenAPIUnion.AsAny]
type anyToolSetAdapterOpenAPI interface {
	implToolSetAdapterOpenAPIUnion()
}

func (ToolSetAdapterOpenAPIURL) implToolSetAdapterOpenAPIUnion()      {}
func (ToolSetAdapterOpenAPIUploadID) implToolSetAdapterOpenAPIUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ToolSetAdapterOpenAPIUnion.AsAny().(type) {
//	case cadenya.ToolSetAdapterOpenAPIURL:
//	case cadenya.ToolSetAdapterOpenAPIUploadID:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ToolSetAdapterOpenAPIUnion) AsAny() anyToolSetAdapterOpenAPI {
	switch u.Type {
	case "url":
		return u.AsURL()
	case "uploadId":
		return u.AsUploadID()
	}
	return nil
}

func (u ToolSetAdapterOpenAPIUnion) AsURL() (v ToolSetAdapterOpenAPIURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSetAdapterOpenAPIUnion) AsUploadID() (v ToolSetAdapterOpenAPIUploadID) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolSetAdapterOpenAPIUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolSetAdapterOpenAPIUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterOpenAPIUnion to a
// ToolSetAdapterOpenAPIUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterOpenAPIUnionParam.Overrides()
func (r ToolSetAdapterOpenAPIUnion) ToParam() ToolSetAdapterOpenAPIUnionParam {
	return param.Override[ToolSetAdapterOpenAPIUnionParam](json.RawMessage(r.RawJSON()))
}

func ToolSetAdapterOpenAPIParamOfURL(url string) ToolSetAdapterOpenAPIUnionParam {
	var variant ToolSetAdapterOpenAPIURLParam
	variant.URL = url
	return ToolSetAdapterOpenAPIUnionParam{OfURL: &variant}
}

func ToolSetAdapterOpenAPIParamOfUploadID(uploadID string) ToolSetAdapterOpenAPIUnionParam {
	var variant ToolSetAdapterOpenAPIUploadIDParam
	variant.UploadID = uploadID
	return ToolSetAdapterOpenAPIUnionParam{OfUploadID: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolSetAdapterOpenAPIUnionParam struct {
	OfURL      *ToolSetAdapterOpenAPIURLParam      `json:",omitzero,inline"`
	OfUploadID *ToolSetAdapterOpenAPIUploadIDParam `json:",omitzero,inline"`
	paramUnion
}

func (u ToolSetAdapterOpenAPIUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfURL, u.OfUploadID)
}
func (u *ToolSetAdapterOpenAPIUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func init() {
	apijson.RegisterUnion[ToolSetAdapterOpenAPIUnionParam](
		"type",
		apijson.Discriminator[ToolSetAdapterOpenAPIURLParam]("url"),
		apijson.Discriminator[ToolSetAdapterOpenAPIUploadIDParam]("uploadId"),
	)
}

type ToolSetAdapterOpenAPIUploadID struct {
	// Any of "uploadId".
	Type ToolSetAdapterOpenAPIUploadIDType `json:"type" api:"required"`
	// ID of a COMPLETE Upload containing the OpenAPI spec document.
	UploadID string `json:"uploadId" api:"required"`
	// Base URL for dispatching tool calls. If set, overrides the server resolved from
	// the spec's servers array.
	BaseURL string `json:"baseUrl"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilter `json:"excludeTools"`
	// Headers sent when fetching the spec from a URL and when dispatching tool calls.
	Headers map[string]string `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilter `json:"includeTools"`
	// Name of the server entry in the spec's servers array (OpenAPI 3.2 server.name
	// field). Used to select which server URL to dispatch to when base_url is not set.
	// If unset, the first server is used. Ignored when base_url is set.
	ServerName string `json:"serverName"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilterUnion `json:"toolApprovals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		UploadID      respjson.Field
		BaseURL       respjson.Field
		ExcludeTools  respjson.Field
		Headers       respjson.Field
		IncludeTools  respjson.Field
		ServerName    respjson.Field
		ToolApprovals respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterOpenAPIUploadID) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterOpenAPIUploadID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterOpenAPIUploadID to a
// ToolSetAdapterOpenAPIUploadIDParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterOpenAPIUploadIDParam.Overrides()
func (r ToolSetAdapterOpenAPIUploadID) ToParam() ToolSetAdapterOpenAPIUploadIDParam {
	return param.Override[ToolSetAdapterOpenAPIUploadIDParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterOpenAPIUploadIDType string

const (
	ToolSetAdapterOpenAPIUploadIDTypeUploadID ToolSetAdapterOpenAPIUploadIDType = "uploadId"
)

// The properties Type, UploadID are required.
type ToolSetAdapterOpenAPIUploadIDParam struct {
	// Any of "uploadId".
	Type ToolSetAdapterOpenAPIUploadIDType `json:"type,omitzero" api:"required"`
	// ID of a COMPLETE Upload containing the OpenAPI spec document.
	UploadID string `json:"uploadId" api:"required"`
	// Base URL for dispatching tool calls. If set, overrides the server resolved from
	// the spec's servers array.
	BaseURL param.Opt[string] `json:"baseUrl,omitzero"`
	// Name of the server entry in the spec's servers array (OpenAPI 3.2 server.name
	// field). Used to select which server URL to dispatch to when base_url is not set.
	// If unset, the first server is used. Ignored when base_url is set.
	ServerName param.Opt[string] `json:"serverName,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilterParam `json:"excludeTools,omitzero"`
	// Headers sent when fetching the spec from a URL and when dispatching tool calls.
	Headers map[string]string `json:"headers,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilterParam `json:"includeTools,omitzero"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilterUnionParam `json:"toolApprovals,omitzero"`
	paramObj
}

func (r ToolSetAdapterOpenAPIUploadIDParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterOpenAPIUploadIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterOpenAPIUploadIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterOpenAPIURL struct {
	// Any of "url".
	Type ToolSetAdapterOpenAPIURLType `json:"type" api:"required"`
	// URL to fetch the OpenAPI spec from. Synced automatically every hour.
	URL string `json:"url" api:"required"`
	// Base URL for dispatching tool calls. If set, overrides the server resolved from
	// the spec's servers array.
	BaseURL string `json:"baseUrl"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilter `json:"excludeTools"`
	// Headers sent when fetching the spec from a URL and when dispatching tool calls.
	Headers map[string]string `json:"headers"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilter `json:"includeTools"`
	// Name of the server entry in the spec's servers array (OpenAPI 3.2 server.name
	// field). Used to select which server URL to dispatch to when base_url is not set.
	// If unset, the first server is used. Ignored when base_url is set.
	ServerName string `json:"serverName"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilterUnion `json:"toolApprovals"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type          respjson.Field
		URL           respjson.Field
		BaseURL       respjson.Field
		ExcludeTools  respjson.Field
		Headers       respjson.Field
		IncludeTools  respjson.Field
		ServerName    respjson.Field
		ToolApprovals respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterOpenAPIURL) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterOpenAPIURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterOpenAPIURL to a
// ToolSetAdapterOpenAPIURLParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterOpenAPIURLParam.Overrides()
func (r ToolSetAdapterOpenAPIURL) ToParam() ToolSetAdapterOpenAPIURLParam {
	return param.Override[ToolSetAdapterOpenAPIURLParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterOpenAPIURLType string

const (
	ToolSetAdapterOpenAPIURLTypeURL ToolSetAdapterOpenAPIURLType = "url"
)

// The properties Type, URL are required.
type ToolSetAdapterOpenAPIURLParam struct {
	// Any of "url".
	Type ToolSetAdapterOpenAPIURLType `json:"type,omitzero" api:"required"`
	// URL to fetch the OpenAPI spec from. Synced automatically every hour.
	URL string `json:"url" api:"required"`
	// Base URL for dispatching tool calls. If set, overrides the server resolved from
	// the spec's servers array.
	BaseURL param.Opt[string] `json:"baseUrl,omitzero"`
	// Name of the server entry in the spec's servers array (OpenAPI 3.2 server.name
	// field). Used to select which server URL to dispatch to when base_url is not set.
	// If unset, the first server is used. Ignored when base_url is set.
	ServerName param.Opt[string] `json:"serverName,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	ExcludeTools ToolFilterParam `json:"excludeTools,omitzero"`
	// Headers sent when fetching the spec from a URL and when dispatching tool calls.
	Headers map[string]string `json:"headers,omitzero"`
	// Top-level filter with simple boolean logic (no nesting)
	IncludeTools ToolFilterParam `json:"includeTools,omitzero"`
	// Approval filters that will automatically set the approval requirement on tools
	// synced from an external source
	ToolApprovals ApprovalRequirementFilterUnionParam `json:"toolApprovals,omitzero"`
	paramObj
}

func (r ToolSetAdapterOpenAPIURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterOpenAPIURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterOpenAPIURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetAdapterOpenAPIVariant struct {
	OpenAPI ToolSetAdapterOpenAPIUnion `json:"openapi" api:"required"`
	// Any of "openapi".
	Type ToolSetAdapterOpenAPIVariantType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OpenAPI     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetAdapterOpenAPIVariant) RawJSON() string { return r.JSON.raw }
func (r *ToolSetAdapterOpenAPIVariant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetAdapterOpenAPIVariant to a
// ToolSetAdapterOpenAPIVariantParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetAdapterOpenAPIVariantParam.Overrides()
func (r ToolSetAdapterOpenAPIVariant) ToParam() ToolSetAdapterOpenAPIVariantParam {
	return param.Override[ToolSetAdapterOpenAPIVariantParam](json.RawMessage(r.RawJSON()))
}

type ToolSetAdapterOpenAPIVariantType string

const (
	ToolSetAdapterOpenAPIVariantTypeOpenAPI ToolSetAdapterOpenAPIVariantType = "openapi"
)

// The properties OpenAPI, Type are required.
type ToolSetAdapterOpenAPIVariantParam struct {
	OpenAPI ToolSetAdapterOpenAPIUnionParam `json:"openapi,omitzero" api:"required"`
	// Any of "openapi".
	Type ToolSetAdapterOpenAPIVariantType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r ToolSetAdapterOpenAPIVariantParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetAdapterOpenAPIVariantParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetAdapterOpenAPIVariantParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single event in the tool set's operation timeline.
type ToolSetEvent struct {
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// Event payload for a tool set operation.
	Event ToolSetEventDataUnion `json:"event"`
	Info  ToolSetEventInfo      `json:"info"`
	// The tool set this event is associated with.
	ToolSetID string `json:"toolSetId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Event       respjson.Field
		Info        respjson.Field
		ToolSetID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetEvent) RawJSON() string { return r.JSON.raw }
func (r *ToolSetEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetEventInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolSet shared.ResourceMetadata `json:"toolSet"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		ToolSet     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetEventInfo) RawJSON() string { return r.JSON.raw }
func (r *ToolSetEventInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolSetEventDataUnion contains all possible properties and values from
// [ToolSetEventDataSyncStarted], [ToolSetEventDataSyncCompleted],
// [ToolSetEventDataSyncFailed].
//
// Use the [ToolSetEventDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToolSetEventDataUnion struct {
	// This field is from variant [ToolSetEventDataSyncStarted].
	SyncStarted SyncStarted `json:"syncStarted"`
	// Any of "syncStarted", "syncCompleted", "syncFailed".
	Type string `json:"type"`
	// This field is from variant [ToolSetEventDataSyncCompleted].
	SyncCompleted SyncCompleted `json:"syncCompleted"`
	// This field is from variant [ToolSetEventDataSyncFailed].
	SyncFailed SyncFailed `json:"syncFailed"`
	JSON       struct {
		SyncStarted   respjson.Field
		Type          respjson.Field
		SyncCompleted respjson.Field
		SyncFailed    respjson.Field
		raw           string
	} `json:"-"`
}

// anyToolSetEventData is implemented by each variant of [ToolSetEventDataUnion] to
// add type safety for the return type of [ToolSetEventDataUnion.AsAny]
type anyToolSetEventData interface {
	implToolSetEventDataUnion()
}

func (ToolSetEventDataSyncStarted) implToolSetEventDataUnion()   {}
func (ToolSetEventDataSyncCompleted) implToolSetEventDataUnion() {}
func (ToolSetEventDataSyncFailed) implToolSetEventDataUnion()    {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ToolSetEventDataUnion.AsAny().(type) {
//	case cadenya.ToolSetEventDataSyncStarted:
//	case cadenya.ToolSetEventDataSyncCompleted:
//	case cadenya.ToolSetEventDataSyncFailed:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ToolSetEventDataUnion) AsAny() anyToolSetEventData {
	switch u.Type {
	case "syncStarted":
		return u.AsSyncStarted()
	case "syncCompleted":
		return u.AsSyncCompleted()
	case "syncFailed":
		return u.AsSyncFailed()
	}
	return nil
}

func (u ToolSetEventDataUnion) AsSyncStarted() (v ToolSetEventDataSyncStarted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSetEventDataUnion) AsSyncCompleted() (v ToolSetEventDataSyncCompleted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolSetEventDataUnion) AsSyncFailed() (v ToolSetEventDataSyncFailed) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolSetEventDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolSetEventDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetEventDataSyncCompleted struct {
	// Emitted when a tool set sync operation completes successfully.
	SyncCompleted SyncCompleted `json:"syncCompleted" api:"required"`
	// Any of "syncCompleted".
	Type ToolSetEventDataSyncCompletedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SyncCompleted respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetEventDataSyncCompleted) RawJSON() string { return r.JSON.raw }
func (r *ToolSetEventDataSyncCompleted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetEventDataSyncCompletedType string

const (
	ToolSetEventDataSyncCompletedTypeSyncCompleted ToolSetEventDataSyncCompletedType = "syncCompleted"
)

type ToolSetEventDataSyncFailed struct {
	// Emitted when a tool set sync operation fails.
	SyncFailed SyncFailed `json:"syncFailed" api:"required"`
	// Any of "syncFailed".
	Type ToolSetEventDataSyncFailedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SyncFailed  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetEventDataSyncFailed) RawJSON() string { return r.JSON.raw }
func (r *ToolSetEventDataSyncFailed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetEventDataSyncFailedType string

const (
	ToolSetEventDataSyncFailedTypeSyncFailed ToolSetEventDataSyncFailedType = "syncFailed"
)

type ToolSetEventDataSyncStarted struct {
	// Emitted when a tool set sync operation begins.
	SyncStarted SyncStarted `json:"syncStarted" api:"required"`
	// Any of "syncStarted".
	Type ToolSetEventDataSyncStartedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SyncStarted respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetEventDataSyncStarted) RawJSON() string { return r.JSON.raw }
func (r *ToolSetEventDataSyncStarted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetEventDataSyncStartedType string

const (
	ToolSetEventDataSyncStartedTypeSyncStarted ToolSetEventDataSyncStartedType = "syncStarted"
)

type ToolSetInfo struct {
	AgentCount     int64 `json:"agentCount"`
	AvailableTools int64 `json:"availableTools"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy    Profile   `json:"createdBy"`
	LastSync     time.Time `json:"lastSync" format:"date-time"`
	OmittedTools int64     `json:"omittedTools"`
	ToolCount    int64     `json:"toolCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentCount     respjson.Field
		AvailableTools respjson.Field
		CreatedBy      respjson.Field
		LastSync       respjson.Field
		OmittedTools   respjson.Field
		ToolCount      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetInfo) RawJSON() string { return r.JSON.raw }
func (r *ToolSetInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetSpec struct {
	Adapter     ToolSetAdapterUnion `json:"adapter"`
	Description string              `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Adapter     respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetSpec) RawJSON() string { return r.JSON.raw }
func (r *ToolSetSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolSetSpec to a ToolSetSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolSetSpecParam.Overrides()
func (r ToolSetSpec) ToParam() ToolSetSpecParam {
	return param.Override[ToolSetSpecParam](json.RawMessage(r.RawJSON()))
}

type ToolSetSpecParam struct {
	Description param.Opt[string]        `json:"description,omitzero"`
	Adapter     ToolSetAdapterUnionParam `json:"adapter,omitzero"`
	paramObj
}

func (r ToolSetSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetGetOpenAPISpecResponse struct {
	// The consumed OpenAPI specification as a JSON string.
	Spec string `json:"spec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolSetGetOpenAPISpecResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolSetGetOpenAPISpecResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	Spec     ToolSetSpecParam                   `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r ToolSetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	UpdateMask  param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	Spec     ToolSetSpecParam                   `json:"spec,omitzero"`
	paramObj
}

func (r ToolSetUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetListParams struct {
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
	// Filter expression (query param: prefix)
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Free-form search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// Filter by tool set lifecycle state. Defaults to STATE_ACTIVE when unspecified;
	// pass STATE_ARCHIVED to list archived tool sets.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ACTIVE", "STATE_ARCHIVED".
	State ToolSetListParamsState `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolSetListParams]'s query parameters as `url.Values`.
func (r ToolSetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by tool set lifecycle state. Defaults to STATE_ACTIVE when unspecified;
// pass STATE_ARCHIVED to list archived tool sets.
type ToolSetListParamsState string

const (
	ToolSetListParamsStateStateUnspecified ToolSetListParamsState = "STATE_UNSPECIFIED"
	ToolSetListParamsStateStateActive      ToolSetListParamsState = "STATE_ACTIVE"
	ToolSetListParamsStateStateArchived    ToolSetListParamsState = "STATE_ARCHIVED"
)

type ToolSetDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetArchiveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ToolSetArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolSetGetOpenAPISpecParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ToolSetListEventsParams struct {
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
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ToolSetListEventsParams]'s query parameters as
// `url.Values`.
func (r ToolSetListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ToolSetUnarchiveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ToolSetUnarchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolSetUnarchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolSetUnarchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
