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
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	return
}

// Retrieves a model by ID from the workspace
func (r *ModelService) Get(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/models/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all models in the workspace
func (r *ModelService) List(ctx context.Context, workspaceID string, query ModelListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Model], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/models", workspaceID)
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

// Lists all models in the workspace
func (r *ModelService) ListAutoPaging(ctx context.Context, workspaceID string, query ModelListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Model] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Transitions a model to STATE_DISABLED. Fails while agent variations are still
// provisioned on the model; use :swapModelOnVariations to move them first.
func (r *ModelService) Disable(ctx context.Context, workspaceID string, id string, body ModelDisableParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/models/%s:disable", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions a model to STATE_ENABLED, making it available for agent variations
// in the workspace
func (r *ModelService) Enable(ctx context.Context, workspaceID string, id string, body ModelEnableParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/models/%s:enable", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reassigns agent variations from one model to another in bulk. Runs
// asynchronously and returns immediately.
func (r *ModelService) Swap(ctx context.Context, workspaceID string, body ModelSwapParams, opts ...option.RequestOption) (res *ModelSwapResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/models:swapModelOnVariations", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type Model struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// Model specification
	Spec ModelSpec `json:"spec" api:"required"`
	// Whether the model is usable in this workspace. Output only. Use the :enable and
	// :disable actions to transition.
	State ModelState `json:"state" api:"required"`
	// ModelInfo carries server-derived, read-only details about a model.
	Info ModelInfo `json:"info"`
	JSON modelJSON `json:"-"`
}

// modelJSON contains the JSON metadata for the struct [Model]
type modelJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	State       apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Model) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelJSON) RawJSON() string {
	return r.raw
}

// Whether the model is usable in this workspace. Output only. Use the :enable and
// :disable actions to transition.
type ModelState string

const (
	ModelStateStateUnspecified ModelState = "STATE_UNSPECIFIED"
	ModelStateStateEnabled     ModelState = "STATE_ENABLED"
	ModelStateStateDisabled    ModelState = "STATE_DISABLED"
)

func (r ModelState) IsKnown() bool {
	switch r {
	case ModelStateStateUnspecified, ModelStateStateEnabled, ModelStateStateDisabled:
		return true
	}
	return false
}

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
	LastUsedAt time.Time     `json:"lastUsedAt" format:"date-time"`
	JSON       modelInfoJSON `json:"-"`
}

// modelInfoJSON contains the JSON metadata for the struct [ModelInfo]
type modelInfoJSON struct {
	AgentVariationCount apijson.Field
	AIProviderKey       apijson.Field
	LastUsedAt          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ModelInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelInfoJSON) RawJSON() string {
	return r.raw
}

type ModelSpec struct {
	// The model family (e.g., "claude-sonnet-4.6", "gpt-5.4", "gemini-2.5-flash")
	Family string `json:"family"`
	// Cost per million input tokens in cents (e.g., 300 = $3.00)
	InputPricePerMillionTokens string `json:"inputPricePerMillionTokens"`
	// Maximum number of input tokens the model supports
	MaxInputTokens int64 `json:"maxInputTokens"`
	// Maximum number of output tokens the model can generate
	MaxOutputTokens int64 `json:"maxOutputTokens"`
	// Cost per million output tokens in cents (e.g., 1500 = $15.00)
	OutputPricePerMillionTokens string `json:"outputPricePerMillionTokens"`
	// The model provider (e.g., "anthropic", "openai", "google")
	Provider string        `json:"provider"`
	JSON     modelSpecJSON `json:"-"`
}

// modelSpecJSON contains the JSON metadata for the struct [ModelSpec]
type modelSpecJSON struct {
	Family                      apijson.Field
	InputPricePerMillionTokens  apijson.Field
	MaxInputTokens              apijson.Field
	MaxOutputTokens             apijson.Field
	OutputPricePerMillionTokens apijson.Field
	Provider                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *ModelSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r modelSpecJSON) RawJSON() string {
	return r.raw
}

type ModelSwapResponse = interface{}

type ModelListParams struct {
	// Filter to models provisioned on a specific AI provider key. Accepts the key's id
	// or an "external_id:"-prefixed slug.
	AIProviderKeyID param.Field[string] `query:"aiProviderKeyId"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When true, populate each item's info (e.g. the AI provider), at the cost of
	// extra lookups.
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Filter models to only ones assigned to an active agent variation/agent. Draft
	// agents count as assigned; archived agents do not. Assignment does not imply
	// recent traffic — see ModelInfo.last_used_at for that.
	IsAssigned param.Field[bool] `query:"isAssigned"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Field[string] `query:"labels"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter by name prefix
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by model state
	State param.Field[ModelListParamsState] `query:"state"`
}

// URLQuery serializes [ModelListParams]'s query parameters as `url.Values`.
func (r ModelListParams) URLQuery() (v url.Values) {
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

func (r ModelListParamsState) IsKnown() bool {
	switch r {
	case ModelListParamsStateStateUnspecified, ModelListParamsStateStateEnabled, ModelListParamsStateStateDisabled:
		return true
	}
	return false
}

type ModelDisableParams struct {
}

func (r ModelDisableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelEnableParams struct {
}

func (r ModelEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelSwapParams struct {
	// The swaps to perform.
	ModelSwaps param.Field[[]ModelSwapParamsModelSwap] `json:"modelSwaps"`
}

func (r ModelSwapParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelSwapParamsModelSwap struct {
	// The model variations are currently on. Accepts an id or "external_id:" slug.
	CurrentModelID param.Field[string] `json:"currentModelId"`
	// Whether to disable the current model after the swap.
	DisableCurrentAfterSwap param.Field[bool] `json:"disableCurrentAfterSwap"`
	// The model to move variations to. Accepts an id or "external_id:" slug.
	NextModelID param.Field[string] `json:"nextModelId"`
}

func (r ModelSwapParamsModelSwap) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
