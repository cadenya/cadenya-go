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
	"time"
)

// Manage LLM models available to a workspace. Models represent provider and family
// pairs (e.g., "anthropic/claude-sonnet-4.6"). Workspaces are seeded with the
// supported models and you can enable or disable each one.
//
// ModelService contains methods and other services that help with interacting with
// the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.options = opts
	return
}

// Retrieves a model by ID from the workspace
func (r *ModelService) Get(ctx context.Context, id string, query ModelGetParams, opts ...option.RequestOption) (res *Model, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/models/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all models in the workspace
func (r *ModelService) List(ctx context.Context, params ModelListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Model], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/models", url.PathEscape(params.WorkspaceID.Value))
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

// Lists all models in the workspace
func (r *ModelService) ListAutoPaging(ctx context.Context, params ModelListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Model] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Transitions a model to STATE_DISABLED. Fails while agent variations are still
// provisioned on the model; use :swapModelOnVariations to move them first.
func (r *ModelService) Disable(ctx context.Context, id string, body ModelDisableParams, opts ...option.RequestOption) (res *Model, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/models/%s:disable", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions a model to STATE_ENABLED, making it available for agent variations
// in the workspace
func (r *ModelService) Enable(ctx context.Context, id string, body ModelEnableParams, opts ...option.RequestOption) (res *Model, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/models/%s:enable", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reassigns agent variations from one model to another in bulk. Runs
// asynchronously and returns immediately.
func (r *ModelService) Swap(ctx context.Context, params ModelSwapParams, opts ...option.RequestOption) (res *ModelSwapResponse, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/models:swapModelOnVariations", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type Model struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// Model specification
	Spec ModelSpec `json:"spec" api:"required"`
	// Whether the model is usable in this workspace. Output only. Use the :enable and
	// :disable actions to transition.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ENABLED", "STATE_DISABLED".
	State ModelState `json:"state" api:"required"`
	// ModelInfo carries server-derived, read-only details about a model.
	Info ModelInfo `json:"info"`
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
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the model is usable in this workspace. Output only. Use the :enable and
// :disable actions to transition.
type ModelState string

const (
	ModelStateStateUnspecified ModelState = "STATE_UNSPECIFIED"
	ModelStateStateEnabled     ModelState = "STATE_ENABLED"
	ModelStateStateDisabled    ModelState = "STATE_DISABLED"
)

// ModelInfo carries server-derived, read-only details about a model.
type ModelInfo struct {
	// Number of agent variations currently provisioned on this model. Useful for
	// previewing how many variations a swap would affect.
	AgentVariationCount int64 `json:"agentVariationCount"`
	// AIProviderKey is a credential for an AI provider, scoped to a workspace. Most
	// keys are customer-provided (BYOK); Cadenya also provisions promotional keys (see
	// AIProviderKeyInfo.is_promotional), which cannot be modified or deleted by
	// account administrators. The secret value is never returned in responses.
	AIProviderKey AIProviderKey `json:"aiProviderKey"`
	// Represents the last time this model was used in an agent objective
	LastUsedAt time.Time `json:"lastUsedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentVariationCount respjson.Field
		AIProviderKey       respjson.Field
		LastUsedAt          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelInfo) RawJSON() string { return r.JSON.raw }
func (r *ModelInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSpec struct {
	// The model family (e.g., "claude-sonnet-4.6", "gpt-5.4", "gemini-2.5-flash")
	Family string `json:"family" api:"required"`
	// The model provider (e.g., "anthropic", "openai", "google")
	Provider string `json:"provider" api:"required"`
	// Cost per million input tokens in cents (e.g., 300 = $3.00)
	InputPricePerMillionTokens string `json:"inputPricePerMillionTokens"`
	// Maximum number of input tokens the model supports
	MaxInputTokens int64 `json:"maxInputTokens"`
	// Maximum number of output tokens the model can generate
	MaxOutputTokens int64 `json:"maxOutputTokens"`
	// Cost per million output tokens in cents (e.g., 1500 = $15.00)
	OutputPricePerMillionTokens string `json:"outputPricePerMillionTokens"`
	// The model's reasoning capability. Catalog data used to decide whether thinking
	// is requested for objective iterations on this model.
	//
	// Any of "REASONING_UNSPECIFIED", "REASONING_NONE", "REASONING_ADAPTIVE",
	// "REASONING_BUDGET".
	Reasoning ModelSpecReasoning `json:"reasoning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Family                      respjson.Field
		Provider                    respjson.Field
		InputPricePerMillionTokens  respjson.Field
		MaxInputTokens              respjson.Field
		MaxOutputTokens             respjson.Field
		OutputPricePerMillionTokens respjson.Field
		Reasoning                   respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpec) RawJSON() string { return r.JSON.raw }
func (r *ModelSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The model's reasoning capability. Catalog data used to decide whether thinking
// is requested for objective iterations on this model.
type ModelSpecReasoning string

const (
	ModelSpecReasoningReasoningUnspecified ModelSpecReasoning = "REASONING_UNSPECIFIED"
	ModelSpecReasoningReasoningNone        ModelSpecReasoning = "REASONING_NONE"
	ModelSpecReasoningReasoningAdaptive    ModelSpecReasoning = "REASONING_ADAPTIVE"
	ModelSpecReasoningReasoningBudget      ModelSpecReasoning = "REASONING_BUDGET"
)

type ModelSwapResponse = any

type ModelGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ModelListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Filter to models provisioned on a specific AI provider key. Accepts the key's id
	// or an "external_id:"-prefixed slug.
	AIProviderKeyID param.Opt[string] `query:"aiProviderKeyId,omitzero" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, populate each item's info (e.g. the AI provider), at the cost of
	// extra lookups.
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filter models to only ones assigned to an active agent variation/agent. Draft
	// agents count as assigned; archived agents do not. Assignment does not imply
	// recent traffic — see ModelInfo.last_used_at for that.
	IsAssigned param.Opt[bool] `query:"isAssigned,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by a prefix of the model's display name, external id, or id
	// (case-insensitive). A model's external id is the form used in
	// modelConfig.modelId, so a caller holding that can narrow the list by it.
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Free-form search query
	Query param.Opt[string] `query:"query,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// Filter by model state
	//
	// Any of "STATE_UNSPECIFIED", "STATE_ENABLED", "STATE_DISABLED".
	State ModelListParamsState `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ModelListParams]'s query parameters as `url.Values`.
func (r ModelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by model state
type ModelListParamsState string

const (
	ModelListParamsStateStateUnspecified ModelListParamsState = "STATE_UNSPECIFIED"
	ModelListParamsStateStateEnabled     ModelListParamsState = "STATE_ENABLED"
	ModelListParamsStateStateDisabled    ModelListParamsState = "STATE_DISABLED"
)

type ModelDisableParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ModelDisableParams) MarshalJSON() (data []byte, err error) {
	type shadow ModelDisableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelDisableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelEnableParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ModelEnableParams) MarshalJSON() (data []byte, err error) {
	type shadow ModelEnableParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelEnableParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSwapParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// The swaps to perform.
	ModelSwaps []ModelSwapParamsModelSwap `json:"modelSwaps,omitzero"`
	paramObj
}

func (r ModelSwapParams) MarshalJSON() (data []byte, err error) {
	type shadow ModelSwapParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelSwapParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSwapParamsModelSwap struct {
	// The model variations are currently on. Accepts an id or "external_id:" slug.
	CurrentModelID param.Opt[string] `json:"currentModelId,omitzero"`
	// Whether to disable the current model after the swap.
	DisableCurrentAfterSwap param.Opt[bool] `json:"disableCurrentAfterSwap,omitzero"`
	// The model to move variations to. Accepts an id or "external_id:" slug.
	NextModelID param.Opt[string] `json:"nextModelId,omitzero"`
	paramObj
}

func (r ModelSwapParamsModelSwap) MarshalJSON() (data []byte, err error) {
	type shadow ModelSwapParamsModelSwap
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelSwapParamsModelSwap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
