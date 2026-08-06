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
)

// Manage variations of an agent and their tool, sub-agent, and memory layer
// assignments.
//
// AgentVariationService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentVariationService] method instead.
type AgentVariationService struct {
	options []option.RequestOption
}

// NewAgentVariationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentVariationService(opts ...option.RequestOption) (r AgentVariationService) {
	r = AgentVariationService{}
	r.options = opts
	return
}

// Creates a new variation for an agent
func (r *AgentVariationService) New(ctx context.Context, agentID string, params AgentVariationNewParams, opts ...option.RequestOption) (res *AgentVariation, err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves a variation by ID from an agent
func (r *AgentVariationService) Get(ctx context.Context, agentID string, id string, query AgentVariationGetParams, opts ...option.RequestOption) (res *AgentVariation, err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a variation for an agent
func (r *AgentVariationService) Update(ctx context.Context, agentID string, id string, params AgentVariationUpdateParams, opts ...option.RequestOption) (res *AgentVariation, err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all variations for an agent
func (r *AgentVariationService) List(ctx context.Context, agentID string, params AgentVariationListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[AgentVariation], err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID))
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

// Lists all variations for an agent
func (r *AgentVariationService) ListAutoPaging(ctx context.Context, agentID string, params AgentVariationListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[AgentVariation] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, agentID, params, opts...))
}

// Deletes a variation from an agent
func (r *AgentVariationService) Delete(ctx context.Context, agentID string, id string, body AgentVariationDeleteParams, opts ...option.RequestOption) (err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Assigns a tool, tool set, or sub-agent to a variation. Exactly one target ID
// must be set.
func (r *AgentVariationService) AddAssignment(ctx context.Context, agentID string, variationID string, params AgentVariationAddAssignmentParams, opts ...option.RequestOption) (res *VariationAssignment, err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if variationID == "" {
		err = errors.New("missing required variationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/assignments", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(variationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Attaches a memory layer to a variation at a given position in the variation's
// baseline memory cascade.
func (r *AgentVariationService) AddMemoryLayer(ctx context.Context, agentID string, variationID string, params AgentVariationAddMemoryLayerParams, opts ...option.RequestOption) (res *VariationMemoryLayerAssignment, err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if variationID == "" {
		err = errors.New("missing required variationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/memory_layer_assignments", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(variationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Detaches an assignment from a variation, identified by the assignment ID
// returned when it was added.
func (r *AgentVariationService) RemoveAssignment(ctx context.Context, agentID string, variationID string, id string, body AgentVariationRemoveAssignmentParams, opts ...option.RequestOption) (err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return err
	}
	if variationID == "" {
		err = errors.New("missing required variationId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/assignments/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(variationID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Detaches a memory layer assignment from a variation, identified by the
// assignment id.
func (r *AgentVariationService) RemoveMemoryLayer(ctx context.Context, agentID string, variationID string, id string, body AgentVariationRemoveMemoryLayerParams, opts ...option.RequestOption) (err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return err
	}
	if variationID == "" {
		err = errors.New("missing required variationId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/memory_layer_assignments/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(variationID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Updates the position of a memory layer assignment on a variation.
func (r *AgentVariationService) UpdateMemoryLayer(ctx context.Context, agentID string, variationID string, id string, params AgentVariationUpdateMemoryLayerParams, opts ...option.RequestOption) (res *VariationMemoryLayerAssignment, err error) {
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
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if variationID == "" {
		err = errors.New("missing required variationId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/memory_layer_assignments/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID), url.PathEscape(variationID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// AgentVariation resource
type AgentVariation struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec AgentVariationSpec `json:"spec" api:"required"`
	// AgentVariationInfo provides read-only summary information about a variation
	Info AgentVariationInfo `json:"info"`
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
func (r AgentVariation) RawJSON() string { return r.JSON.raw }
func (r *AgentVariation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentVariationInfo provides read-only summary information about a variation
type AgentVariationInfo struct {
	// All tools, tool sets, and sub-agents assigned to this variation. Populated on
	// reads so clients can render a variation's full assignment list without calling
	// the add/remove endpoints just to enumerate.
	Assignments []VariationAssignment `json:"assignments"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Total number of objective feedbacks received for this variation
	FeedbackCount int64 `json:"feedbackCount"`
	// Read-only list of memory layer assignments for this variation, returned in
	// ascending `position` (most specific first — resolution order). Capped at 10
	// entries.
	MemoryLayerAssignments []VariationMemoryLayerAssignment `json:"memoryLayerAssignments"`
	// Count of memory layer assignments.
	MemoryLayerCount int64 `json:"memoryLayerCount"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Model shared.ResourceMetadata `json:"model"`
	// Thompson Sampling score: posterior mean of Beta(ts_alpha, ts_beta). Range [0, 1]
	// where 0.5 = neutral, >0.5 = positive, <0.5 = negative.
	Score float64 `json:"score"`
	// Number of sub-agents assigned to this variation
	SubAgentCount int64 `json:"subAgentCount"`
	// Number of individual tools assigned to this variation
	ToolCount int64 `json:"toolCount"`
	// Number of tool sets assigned to this variation
	ToolSetCount int64 `json:"toolSetCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assignments            respjson.Field
		CreatedBy              respjson.Field
		FeedbackCount          respjson.Field
		MemoryLayerAssignments respjson.Field
		MemoryLayerCount       respjson.Field
		Model                  respjson.Field
		Score                  respjson.Field
		SubAgentCount          respjson.Field
		ToolCount              respjson.Field
		ToolSetCount           respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentVariationInfo) RawJSON() string { return r.JSON.raw }
func (r *AgentVariationInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentVariationSpec defines the operational configuration for a variation
type AgentVariationSpec struct {
	// CompactionConfig defines how context window compaction behaves for objectives
	// using this variation.
	CompactionConfig AgentVariationSpecCompactionConfig `json:"compactionConfig"`
	// Execution constraints
	Constraints AgentVariationSpecConstraints `json:"constraints"`
	// Human-readable description of what this variation does or when it should be used
	Description string `json:"description"`
	// Liquid template for the first user message of objectives using this variation.
	// Rendered with CreateObjectiveRequest.first_user_message_data into
	// Objective.first_user_message, the first user message in the LLM chat history.
	// CreateObjectiveRequest.first_user_message, when set, overrides the rendered
	// result. If neither this template nor first_user_message is present, objective
	// creation is rejected with InvalidArgument.
	FirstUserMessageTemplate string `json:"firstUserMessageTemplate"`
	// ModelConfig defines the model configuration for a variation
	ModelConfig AgentVariationSpecModelConfig `json:"modelConfig"`
	// ProgressiveDiscovery is used to indicate that the agent should automatically
	// discover tools that are not explicitly assigned to it. Max tools is the maximum
	// number of tools that can be discovered per search. Hints are optional hints for
	// tool search. These are used in conjunction with the context-aware tool search
	// and can help select the best tools for the task.
	ProgressiveDiscovery AgentVariationSpecProgressiveDiscovery `json:"progressiveDiscovery"`
	// Liquid template for the system prompt of objectives using this variation.
	// Rendered with CreateObjectiveRequest.system_prompt_data into
	// Objective.system_prompt.
	SystemPromptTemplate string `json:"systemPromptTemplate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompactionConfig         respjson.Field
		Constraints              respjson.Field
		Description              respjson.Field
		FirstUserMessageTemplate respjson.Field
		ModelConfig              respjson.Field
		ProgressiveDiscovery     respjson.Field
		SystemPromptTemplate     respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentVariationSpec) RawJSON() string { return r.JSON.raw }
func (r *AgentVariationSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentVariationSpec to a AgentVariationSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentVariationSpecParam.Overrides()
func (r AgentVariationSpec) ToParam() AgentVariationSpecParam {
	return param.Override[AgentVariationSpecParam](json.RawMessage(r.RawJSON()))
}

// AgentVariationSpec defines the operational configuration for a variation
type AgentVariationSpecParam struct {
	// Human-readable description of what this variation does or when it should be used
	Description param.Opt[string] `json:"description,omitzero"`
	// Liquid template for the first user message of objectives using this variation.
	// Rendered with CreateObjectiveRequest.first_user_message_data into
	// Objective.first_user_message, the first user message in the LLM chat history.
	// CreateObjectiveRequest.first_user_message, when set, overrides the rendered
	// result. If neither this template nor first_user_message is present, objective
	// creation is rejected with InvalidArgument.
	FirstUserMessageTemplate param.Opt[string] `json:"firstUserMessageTemplate,omitzero"`
	// Liquid template for the system prompt of objectives using this variation.
	// Rendered with CreateObjectiveRequest.system_prompt_data into
	// Objective.system_prompt.
	SystemPromptTemplate param.Opt[string] `json:"systemPromptTemplate,omitzero"`
	// CompactionConfig defines how context window compaction behaves for objectives
	// using this variation.
	CompactionConfig AgentVariationSpecCompactionConfigParam `json:"compactionConfig,omitzero"`
	// Execution constraints
	Constraints AgentVariationSpecConstraintsParam `json:"constraints,omitzero"`
	// ModelConfig defines the model configuration for a variation
	ModelConfig AgentVariationSpecModelConfigParam `json:"modelConfig,omitzero"`
	// ProgressiveDiscovery is used to indicate that the agent should automatically
	// discover tools that are not explicitly assigned to it. Max tools is the maximum
	// number of tools that can be discovered per search. Hints are optional hints for
	// tool search. These are used in conjunction with the context-aware tool search
	// and can help select the best tools for the task.
	ProgressiveDiscovery AgentVariationSpecProgressiveDiscoveryParam `json:"progressiveDiscovery,omitzero"`
	paramObj
}

func (r AgentVariationSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CompactionConfig defines how context window compaction behaves for objectives
// using this variation.
type AgentVariationSpecCompactionConfig struct {
	// SummarizationStrategy configures LLM-powered summarization of older conversation
	// turns.
	Summarization CompactionConfigSummarizationStrategy `json:"summarization"`
	// ToolResultClearingStrategy configures clearing of older tool result content.
	ToolResultClearing CompactionConfigToolResultClearingStrategy `json:"toolResultClearing"`
	// Trigger threshold as a percentage of the model's context window (0.0 to 1.0).
	// When input tokens reach this percentage of the model's limit, compaction
	// triggers. Default: 0.75 (75%)
	TriggerThreshold float64 `json:"triggerThreshold"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Summarization      respjson.Field
		ToolResultClearing respjson.Field
		TriggerThreshold   respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentVariationSpecCompactionConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentVariationSpecCompactionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentVariationSpecCompactionConfig to a
// AgentVariationSpecCompactionConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentVariationSpecCompactionConfigParam.Overrides()
func (r AgentVariationSpecCompactionConfig) ToParam() AgentVariationSpecCompactionConfigParam {
	return param.Override[AgentVariationSpecCompactionConfigParam](json.RawMessage(r.RawJSON()))
}

// CompactionConfig defines how context window compaction behaves for objectives
// using this variation.
type AgentVariationSpecCompactionConfigParam struct {
	// Trigger threshold as a percentage of the model's context window (0.0 to 1.0).
	// When input tokens reach this percentage of the model's limit, compaction
	// triggers. Default: 0.75 (75%)
	TriggerThreshold param.Opt[float64] `json:"triggerThreshold,omitzero"`
	// SummarizationStrategy configures LLM-powered summarization of older conversation
	// turns.
	Summarization CompactionConfigSummarizationStrategyParam `json:"summarization,omitzero"`
	// ToolResultClearingStrategy configures clearing of older tool result content.
	ToolResultClearing CompactionConfigToolResultClearingStrategyParam `json:"toolResultClearing,omitzero"`
	paramObj
}

func (r AgentVariationSpecCompactionConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationSpecCompactionConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationSpecCompactionConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentVariationSpecConstraints struct {
	// How long an objective may sit with no activity (no user messages, no LLM calls)
	// before it is finalized as timed out. Between 1 minute and 24 hours, expressed as
	// a duration string in seconds (e.g. "7200s"). When not set, objectives are still
	// swept at the system-wide 24 hour maximum — every objective eventually reaches a
	// terminal state.
	//
	// Note: no gnostic integer hint here on purpose. The Envoy gRPC-JSON transcoder
	// only accepts the canonical protobuf JSON form for Durations — a "<seconds>s"
	// string — so the SDKs must type this as a string (like AgentScheduleSpec.every),
	// not an integer.
	InactivityTimeout string `json:"inactivityTimeout"`
	// The maximum number of sub-objectives that can be created. 0 means no limit.
	MaxSubObjectives int64 `json:"maxSubObjectives"`
	// The maximum number of tool calls that can be made. 0 means no limit.
	MaxToolCalls int64 `json:"maxToolCalls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InactivityTimeout respjson.Field
		MaxSubObjectives  respjson.Field
		MaxToolCalls      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentVariationSpecConstraints) RawJSON() string { return r.JSON.raw }
func (r *AgentVariationSpecConstraints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentVariationSpecConstraints to a
// AgentVariationSpecConstraintsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentVariationSpecConstraintsParam.Overrides()
func (r AgentVariationSpecConstraints) ToParam() AgentVariationSpecConstraintsParam {
	return param.Override[AgentVariationSpecConstraintsParam](json.RawMessage(r.RawJSON()))
}

type AgentVariationSpecConstraintsParam struct {
	// How long an objective may sit with no activity (no user messages, no LLM calls)
	// before it is finalized as timed out. Between 1 minute and 24 hours, expressed as
	// a duration string in seconds (e.g. "7200s"). When not set, objectives are still
	// swept at the system-wide 24 hour maximum — every objective eventually reaches a
	// terminal state.
	//
	// Note: no gnostic integer hint here on purpose. The Envoy gRPC-JSON transcoder
	// only accepts the canonical protobuf JSON form for Durations — a "<seconds>s"
	// string — so the SDKs must type this as a string (like AgentScheduleSpec.every),
	// not an integer.
	InactivityTimeout param.Opt[string] `json:"inactivityTimeout,omitzero"`
	// The maximum number of sub-objectives that can be created. 0 means no limit.
	MaxSubObjectives param.Opt[int64] `json:"maxSubObjectives,omitzero"`
	// The maximum number of tool calls that can be made. 0 means no limit.
	MaxToolCalls param.Opt[int64] `json:"maxToolCalls,omitzero"`
	paramObj
}

func (r AgentVariationSpecConstraintsParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationSpecConstraintsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationSpecConstraintsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ModelConfig defines the model configuration for a variation
type AgentVariationSpecModelConfig struct {
	// The model identifier in family/model format (e.g., "claude/opus-4.6",
	// "claude/sonnet-4.5")
	ModelID string `json:"modelId"`
	// Sampling temperature for model inference (0.0 to 1.0) Lower values produce more
	// deterministic outputs, higher values increase randomness
	Temperature float64 `json:"temperature"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelID     respjson.Field
		Temperature respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentVariationSpecModelConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentVariationSpecModelConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentVariationSpecModelConfig to a
// AgentVariationSpecModelConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentVariationSpecModelConfigParam.Overrides()
func (r AgentVariationSpecModelConfig) ToParam() AgentVariationSpecModelConfigParam {
	return param.Override[AgentVariationSpecModelConfigParam](json.RawMessage(r.RawJSON()))
}

// ModelConfig defines the model configuration for a variation
type AgentVariationSpecModelConfigParam struct {
	// The model identifier in family/model format (e.g., "claude/opus-4.6",
	// "claude/sonnet-4.5")
	ModelID param.Opt[string] `json:"modelId,omitzero"`
	// Sampling temperature for model inference (0.0 to 1.0) Lower values produce more
	// deterministic outputs, higher values increase randomness
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	paramObj
}

func (r AgentVariationSpecModelConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationSpecModelConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationSpecModelConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProgressiveDiscovery is used to indicate that the agent should automatically
// discover tools that are not explicitly assigned to it. Max tools is the maximum
// number of tools that can be discovered per search. Hints are optional hints for
// tool search. These are used in conjunction with the context-aware tool search
// and can help select the best tools for the task.
type AgentVariationSpecProgressiveDiscovery struct {
	// Free-text guidance appended to the discoverable-tools appendix in the system
	// prompt. Hints steer the model's choice of tool names; they do not filter or rank
	// anything, because tool_search matches names exactly rather than searching.
	Hints []string `json:"hints"`
	// The most tool names tool_search will load in a single call. Requesting more than
	// this returns an error telling the model to retry in smaller batches -- it is a
	// per-call batch limit, not a ceiling on how many tools an objective may end up
	// with.
	MaxTools int64 `json:"maxTools"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hints       respjson.Field
		MaxTools    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentVariationSpecProgressiveDiscovery) RawJSON() string { return r.JSON.raw }
func (r *AgentVariationSpecProgressiveDiscovery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentVariationSpecProgressiveDiscovery to a
// AgentVariationSpecProgressiveDiscoveryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentVariationSpecProgressiveDiscoveryParam.Overrides()
func (r AgentVariationSpecProgressiveDiscovery) ToParam() AgentVariationSpecProgressiveDiscoveryParam {
	return param.Override[AgentVariationSpecProgressiveDiscoveryParam](json.RawMessage(r.RawJSON()))
}

// ProgressiveDiscovery is used to indicate that the agent should automatically
// discover tools that are not explicitly assigned to it. Max tools is the maximum
// number of tools that can be discovered per search. Hints are optional hints for
// tool search. These are used in conjunction with the context-aware tool search
// and can help select the best tools for the task.
type AgentVariationSpecProgressiveDiscoveryParam struct {
	// The most tool names tool_search will load in a single call. Requesting more than
	// this returns an error telling the model to retry in smaller batches -- it is a
	// per-call batch limit, not a ceiling on how many tools an objective may end up
	// with.
	MaxTools param.Opt[int64] `json:"maxTools,omitzero"`
	// Free-text guidance appended to the discoverable-tools appendix in the system
	// prompt. Hints steer the model's choice of tool names; they do not filter or rank
	// anything, because tool_search matches names exactly rather than searching.
	Hints []string `json:"hints,omitzero"`
	paramObj
}

func (r AgentVariationSpecProgressiveDiscoveryParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationSpecProgressiveDiscoveryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationSpecProgressiveDiscoveryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SummarizationStrategy configures LLM-powered summarization of older conversation
// turns.
type CompactionConfigSummarizationStrategy struct {
	// Custom instructions that guide what the summarizer preserves. Replaces the
	// default summarization prompt entirely. Example: "Preserve all code snippets,
	// variable names, and technical decisions."
	Instructions string `json:"instructions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instructions respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompactionConfigSummarizationStrategy) RawJSON() string { return r.JSON.raw }
func (r *CompactionConfigSummarizationStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CompactionConfigSummarizationStrategy to a
// CompactionConfigSummarizationStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CompactionConfigSummarizationStrategyParam.Overrides()
func (r CompactionConfigSummarizationStrategy) ToParam() CompactionConfigSummarizationStrategyParam {
	return param.Override[CompactionConfigSummarizationStrategyParam](json.RawMessage(r.RawJSON()))
}

// SummarizationStrategy configures LLM-powered summarization of older conversation
// turns.
type CompactionConfigSummarizationStrategyParam struct {
	// Custom instructions that guide what the summarizer preserves. Replaces the
	// default summarization prompt entirely. Example: "Preserve all code snippets,
	// variable names, and technical decisions."
	Instructions param.Opt[string] `json:"instructions,omitzero"`
	paramObj
}

func (r CompactionConfigSummarizationStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow CompactionConfigSummarizationStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompactionConfigSummarizationStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolResultClearingStrategy configures clearing of older tool result content.
type CompactionConfigToolResultClearingStrategy struct {
	// Number of most recent tool call results to keep intact. Older tool results have
	// their content replaced with "[result cleared]" while preserving the assistant
	// tool call message (function name, arguments). Default: 2
	PreserveRecentResults int64 `json:"preserveRecentResults"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PreserveRecentResults respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompactionConfigToolResultClearingStrategy) RawJSON() string { return r.JSON.raw }
func (r *CompactionConfigToolResultClearingStrategy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CompactionConfigToolResultClearingStrategy to a
// CompactionConfigToolResultClearingStrategyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CompactionConfigToolResultClearingStrategyParam.Overrides()
func (r CompactionConfigToolResultClearingStrategy) ToParam() CompactionConfigToolResultClearingStrategyParam {
	return param.Override[CompactionConfigToolResultClearingStrategyParam](json.RawMessage(r.RawJSON()))
}

// ToolResultClearingStrategy configures clearing of older tool result content.
type CompactionConfigToolResultClearingStrategyParam struct {
	// Number of most recent tool call results to keep intact. Older tool results have
	// their content replaced with "[result cleared]" while preserving the assistant
	// tool call message (function name, arguments). Default: 2
	PreserveRecentResults param.Opt[int64] `json:"preserveRecentResults,omitzero"`
	paramObj
}

func (r CompactionConfigToolResultClearingStrategyParam) MarshalJSON() (data []byte, err error) {
	type shadow CompactionConfigToolResultClearingStrategyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompactionConfigToolResultClearingStrategyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A read-only reference to a single tool, tool set, or sub-agent attached to a
// variation. Read the full set of assignments via
// `AgentVariationInfo.assignments`; mutations go through the dedicated add/remove
// assignment endpoints.
//
// The `id` identifies the assignment itself (not the referenced resource) and is
// the handle used to remove the assignment. It is returned by the add endpoint and
// present on every entry in `AgentVariationInfo.assignments`.
type VariationAssignment struct {
	ID string `json:"id"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Agent shared.BareMetadata `json:"agent"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Tool shared.BareMetadata `json:"tool"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	ToolSet shared.BareMetadata `json:"toolSet"`
	// The JSON name of the variant set in `reference` (e.g. "toolSet"). Filled by the
	// server; drives the discriminated union in the generated OpenAPI.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Agent       respjson.Field
		Tool        respjson.Field
		ToolSet     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VariationAssignment) RawJSON() string { return r.JSON.raw }
func (r *VariationAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// VariationMemoryLayerAssignment attaches a single MemoryLayer to a variation at a
// given position in the variation's baseline memory cascade. A variation has at
// most one assignment per memory_layer_id.
//
// Variations only support whole-layer attachments — entry pinning is an
// objective-level capability.
type VariationMemoryLayerAssignment struct {
	// Assignment row id — handle for removing the assignment. Distinct from the
	// referenced memory layer's id.
	ID string `json:"id"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	MemoryLayer shared.BareMetadata `json:"memoryLayer"`
	// Position in the variation's baseline cascade. Position is specificity,
	// CSS-style: a LOWER position is more specific and is consulted first; the
	// highest-position assignment is the most general fallback. Gaps are fine — only
	// relative position matters. Positions must be unique within a variation; a
	// request that would collide with an existing assignment's position is rejected
	// with InvalidArgument.
	Position int64 `json:"position"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MemoryLayer respjson.Field
		Position    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VariationMemoryLayerAssignment) RawJSON() string { return r.JSON.raw }
func (r *VariationMemoryLayerAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentVariationNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec AgentVariationSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r AgentVariationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentVariationGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AgentVariationUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Fields to update
	UpdateMask param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec AgentVariationSpecParam `json:"spec,omitzero"`
	paramObj
}

func (r AgentVariationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentVariationListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, the `info` field on each returned variation is populated. Requests
	// with this flag count more against your rate limit.
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

// URLQuery serializes [AgentVariationListParams]'s query parameters as
// `url.Values`.
func (r AgentVariationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentVariationDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AgentVariationAddAssignmentParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	SubAgentID  param.Opt[string] `json:"subAgentId,omitzero"`
	ToolID      param.Opt[string] `json:"toolId,omitzero"`
	ToolSetID   param.Opt[string] `json:"toolSetId,omitzero"`
	// The JSON name of the variant set in `target` (e.g. "toolId"). Required on input;
	// drives the discriminated union in the generated OpenAPI.
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r AgentVariationAddAssignmentParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationAddAssignmentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationAddAssignmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentVariationAddMemoryLayerParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Layer to attach. Accepts the canonical `memlyr_…` form or the
	// `external_id:<value>` form.
	MemoryLayerID string `json:"memoryLayerId" api:"required"`
	// Position in the baseline cascade (lower = more specific). If omitted, the server
	// appends at the most general end (max existing position + 1).
	Position param.Opt[int64] `json:"position,omitzero"`
	paramObj
}

func (r AgentVariationAddMemoryLayerParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationAddMemoryLayerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationAddMemoryLayerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentVariationRemoveAssignmentParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AgentVariationRemoveMemoryLayerParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AgentVariationUpdateMemoryLayerParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// New position. Only field currently updatable on an assignment.
	Position param.Opt[int64] `json:"position,omitzero"`
	paramObj
}

func (r AgentVariationUpdateMemoryLayerParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentVariationUpdateMemoryLayerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentVariationUpdateMemoryLayerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
