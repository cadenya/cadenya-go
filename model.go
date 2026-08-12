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

type CapabilityMaxOutputTokens = any

// Reasoning / extended thinking (ModelConfig.reasoning_effort). A model that does
// not reason simply omits this capability.
type CapabilityReasoning struct {
	// How reasoning is enabled for this model. Catalog data used to decide whether
	// thinking is requested for objective iterations on this model.
	//
	// Any of "MODE_UNSPECIFIED", "MODE_ADAPTIVE", "MODE_BUDGET".
	Mode CapabilityReasoningMode `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilityReasoning) RawJSON() string { return r.JSON.raw }
func (r *CapabilityReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How reasoning is enabled for this model. Catalog data used to decide whether
// thinking is requested for objective iterations on this model.
type CapabilityReasoningMode string

const (
	CapabilityReasoningModeModeUnspecified CapabilityReasoningMode = "MODE_UNSPECIFIED"
	CapabilityReasoningModeModeAdaptive    CapabilityReasoningMode = "MODE_ADAPTIVE"
	CapabilityReasoningModeModeBudget      CapabilityReasoningMode = "MODE_BUDGET"
)

// Custom stop sequences (ModelConfig.stop_sequences).
type CapabilityStopSequences struct {
	// Maximum number of stop sequences the model accepts per request. 0 means the
	// provider imposes no meaningful limit.
	Limit int64 `json:"limit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CapabilityStopSequences) RawJSON() string { return r.JSON.raw }
func (r *CapabilityStopSequences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CapabilityTemperature = any

type CapabilityTopK = any

type CapabilityTopP = any

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
	// The inference knobs this model supports. Catalog data; drives which ModelConfig
	// fields a variation on this model may set. Reasoning support (and its mode) lives
	// here too, as the "reasoning" capability.
	Capabilities []ModelSpecCapabilityUnion `json:"capabilities"`
	// Cost per million input tokens in cents (e.g., 300 = $3.00)
	InputPricePerMillionTokens string `json:"inputPricePerMillionTokens"`
	// Maximum number of input tokens the model supports
	MaxInputTokens int64 `json:"maxInputTokens"`
	// Maximum number of output tokens the model can generate
	MaxOutputTokens int64 `json:"maxOutputTokens"`
	// Cost per million output tokens in cents (e.g., 1500 = $15.00)
	OutputPricePerMillionTokens string `json:"outputPricePerMillionTokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Family                      respjson.Field
		Provider                    respjson.Field
		Capabilities                respjson.Field
		InputPricePerMillionTokens  respjson.Field
		MaxInputTokens              respjson.Field
		MaxOutputTokens             respjson.Field
		OutputPricePerMillionTokens respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpec) RawJSON() string { return r.JSON.raw }
func (r *ModelSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelSpecCapabilityUnion contains all possible properties and values from
// [ModelSpecCapabilityTemperature], [ModelSpecCapabilityTopP],
// [ModelSpecCapabilityTopK], [ModelSpecCapabilityStopSequences],
// [ModelSpecCapabilityMaxOutputTokens], [ModelSpecCapabilityReasoning].
//
// Use the [ModelSpecCapabilityUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ModelSpecCapabilityUnion struct {
	// This field is from variant [ModelSpecCapabilityTemperature].
	Temperature CapabilityTemperature `json:"temperature"`
	// Any of "temperature", "topP", "topK", "stopSequences", "maxOutputTokens",
	// "reasoning".
	Type string `json:"type"`
	// This field is from variant [ModelSpecCapabilityTopP].
	TopP CapabilityTopP `json:"topP"`
	// This field is from variant [ModelSpecCapabilityTopK].
	TopK CapabilityTopK `json:"topK"`
	// This field is from variant [ModelSpecCapabilityStopSequences].
	StopSequences CapabilityStopSequences `json:"stopSequences"`
	// This field is from variant [ModelSpecCapabilityMaxOutputTokens].
	MaxOutputTokens CapabilityMaxOutputTokens `json:"maxOutputTokens"`
	// This field is from variant [ModelSpecCapabilityReasoning].
	Reasoning CapabilityReasoning `json:"reasoning"`
	JSON      struct {
		Temperature     respjson.Field
		Type            respjson.Field
		TopP            respjson.Field
		TopK            respjson.Field
		StopSequences   respjson.Field
		MaxOutputTokens respjson.Field
		Reasoning       respjson.Field
		raw             string
	} `json:"-"`
}

// anyModelSpecCapability is implemented by each variant of
// [ModelSpecCapabilityUnion] to add type safety for the return type of
// [ModelSpecCapabilityUnion.AsAny]
type anyModelSpecCapability interface {
	implModelSpecCapabilityUnion()
}

func (ModelSpecCapabilityTemperature) implModelSpecCapabilityUnion()     {}
func (ModelSpecCapabilityTopP) implModelSpecCapabilityUnion()            {}
func (ModelSpecCapabilityTopK) implModelSpecCapabilityUnion()            {}
func (ModelSpecCapabilityStopSequences) implModelSpecCapabilityUnion()   {}
func (ModelSpecCapabilityMaxOutputTokens) implModelSpecCapabilityUnion() {}
func (ModelSpecCapabilityReasoning) implModelSpecCapabilityUnion()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ModelSpecCapabilityUnion.AsAny().(type) {
//	case cadenya.ModelSpecCapabilityTemperature:
//	case cadenya.ModelSpecCapabilityTopP:
//	case cadenya.ModelSpecCapabilityTopK:
//	case cadenya.ModelSpecCapabilityStopSequences:
//	case cadenya.ModelSpecCapabilityMaxOutputTokens:
//	case cadenya.ModelSpecCapabilityReasoning:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ModelSpecCapabilityUnion) AsAny() anyModelSpecCapability {
	switch u.Type {
	case "temperature":
		return u.AsTemperature()
	case "topP":
		return u.AsTopP()
	case "topK":
		return u.AsTopK()
	case "stopSequences":
		return u.AsStopSequences()
	case "maxOutputTokens":
		return u.AsMaxOutputTokens()
	case "reasoning":
		return u.AsReasoning()
	}
	return nil
}

func (u ModelSpecCapabilityUnion) AsTemperature() (v ModelSpecCapabilityTemperature) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelSpecCapabilityUnion) AsTopP() (v ModelSpecCapabilityTopP) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelSpecCapabilityUnion) AsTopK() (v ModelSpecCapabilityTopK) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelSpecCapabilityUnion) AsStopSequences() (v ModelSpecCapabilityStopSequences) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelSpecCapabilityUnion) AsMaxOutputTokens() (v ModelSpecCapabilityMaxOutputTokens) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ModelSpecCapabilityUnion) AsReasoning() (v ModelSpecCapabilityReasoning) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ModelSpecCapabilityUnion) RawJSON() string { return u.JSON.raw }

func (r *ModelSpecCapabilityUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSpecCapabilityMaxOutputTokens struct {
	// Per-request output token cap (ModelConfig.max_output_tokens). The effective
	// ceiling is ModelSpec.max_output_tokens.
	MaxOutputTokens CapabilityMaxOutputTokens `json:"maxOutputTokens" api:"required"`
	// Any of "maxOutputTokens".
	Type ModelSpecCapabilityMaxOutputTokensType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxOutputTokens respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpecCapabilityMaxOutputTokens) RawJSON() string { return r.JSON.raw }
func (r *ModelSpecCapabilityMaxOutputTokens) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSpecCapabilityMaxOutputTokensType string

const (
	ModelSpecCapabilityMaxOutputTokensTypeMaxOutputTokens ModelSpecCapabilityMaxOutputTokensType = "maxOutputTokens"
)

type ModelSpecCapabilityReasoning struct {
	// Reasoning / extended thinking (ModelConfig.reasoning_effort). A model that does
	// not reason simply omits this capability.
	Reasoning CapabilityReasoning `json:"reasoning" api:"required"`
	// Any of "reasoning".
	Type ModelSpecCapabilityReasoningType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reasoning   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpecCapabilityReasoning) RawJSON() string { return r.JSON.raw }
func (r *ModelSpecCapabilityReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSpecCapabilityReasoningType string

const (
	ModelSpecCapabilityReasoningTypeReasoning ModelSpecCapabilityReasoningType = "reasoning"
)

type ModelSpecCapabilityStopSequences struct {
	// Custom stop sequences (ModelConfig.stop_sequences).
	StopSequences CapabilityStopSequences `json:"stopSequences" api:"required"`
	// Any of "stopSequences".
	Type ModelSpecCapabilityStopSequencesType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StopSequences respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpecCapabilityStopSequences) RawJSON() string { return r.JSON.raw }
func (r *ModelSpecCapabilityStopSequences) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSpecCapabilityStopSequencesType string

const (
	ModelSpecCapabilityStopSequencesTypeStopSequences ModelSpecCapabilityStopSequencesType = "stopSequences"
)

type ModelSpecCapabilityTemperature struct {
	// Sampling temperature (ModelConfig.temperature).
	Temperature CapabilityTemperature `json:"temperature" api:"required"`
	// Any of "temperature".
	Type ModelSpecCapabilityTemperatureType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Temperature respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpecCapabilityTemperature) RawJSON() string { return r.JSON.raw }
func (r *ModelSpecCapabilityTemperature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSpecCapabilityTemperatureType string

const (
	ModelSpecCapabilityTemperatureTypeTemperature ModelSpecCapabilityTemperatureType = "temperature"
)

type ModelSpecCapabilityTopK struct {
	// Top-k sampling (ModelConfig.top_k).
	TopK CapabilityTopK `json:"topK" api:"required"`
	// Any of "topK".
	Type ModelSpecCapabilityTopKType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TopK        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpecCapabilityTopK) RawJSON() string { return r.JSON.raw }
func (r *ModelSpecCapabilityTopK) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSpecCapabilityTopKType string

const (
	ModelSpecCapabilityTopKTypeTopK ModelSpecCapabilityTopKType = "topK"
)

type ModelSpecCapabilityTopP struct {
	// Nucleus sampling (ModelConfig.top_p).
	TopP CapabilityTopP `json:"topP" api:"required"`
	// Any of "topP".
	Type ModelSpecCapabilityTopPType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TopP        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpecCapabilityTopP) RawJSON() string { return r.JSON.raw }
func (r *ModelSpecCapabilityTopP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSpecCapabilityTopPType string

const (
	ModelSpecCapabilityTopPTypeTopP ModelSpecCapabilityTopPType = "topP"
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
