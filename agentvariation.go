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
	Options []option.RequestOption
}

// NewAgentVariationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentVariationService(opts ...option.RequestOption) (r *AgentVariationService) {
	r = &AgentVariationService{}
	r.Options = opts
	return
}

// Creates a new variation for an agent
func (r *AgentVariationService) New(ctx context.Context, workspaceID string, agentID string, body AgentVariationNewParams, opts ...option.RequestOption) (res *AgentVariation, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations", workspaceID, agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a variation by ID from an agent
func (r *AgentVariationService) Get(ctx context.Context, workspaceID string, agentID string, id string, opts ...option.RequestOption) (res *AgentVariation, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s", workspaceID, agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a variation for an agent
func (r *AgentVariationService) Update(ctx context.Context, workspaceID string, agentID string, id string, body AgentVariationUpdateParams, opts ...option.RequestOption) (res *AgentVariation, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s", workspaceID, agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all variations for an agent
func (r *AgentVariationService) List(ctx context.Context, workspaceID string, agentID string, query AgentVariationListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[AgentVariation], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations", workspaceID, agentID)
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

// Lists all variations for an agent
func (r *AgentVariationService) ListAutoPaging(ctx context.Context, workspaceID string, agentID string, query AgentVariationListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[AgentVariation] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, agentID, query, opts...))
}

// Deletes a variation from an agent
func (r *AgentVariationService) Delete(ctx context.Context, workspaceID string, agentID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s", workspaceID, agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Assigns a tool, tool set, or sub-agent to a variation. Exactly one target ID
// must be set.
func (r *AgentVariationService) AddAssignment(ctx context.Context, workspaceID string, agentID string, variationID string, body AgentVariationAddAssignmentParams, opts ...option.RequestOption) (res *VariationAssignment, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/assignments", workspaceID, agentID, variationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Attaches a memory layer to a variation at a given position in the variation's
// baseline memory cascade.
func (r *AgentVariationService) AddMemoryLayer(ctx context.Context, workspaceID string, agentID string, variationID string, body AgentVariationAddMemoryLayerParams, opts ...option.RequestOption) (res *VariationMemoryLayerAssignment, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/memory_layer_assignments", workspaceID, agentID, variationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Detaches an assignment from a variation, identified by the assignment ID
// returned when it was added.
func (r *AgentVariationService) RemoveAssignment(ctx context.Context, workspaceID string, agentID string, variationID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/assignments/%s", workspaceID, agentID, variationID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Detaches a memory layer assignment from a variation, identified by the
// assignment id.
func (r *AgentVariationService) RemoveMemoryLayer(ctx context.Context, workspaceID string, agentID string, variationID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/memory_layer_assignments/%s", workspaceID, agentID, variationID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Updates the position of a memory layer assignment on a variation.
func (r *AgentVariationService) UpdateMemoryLayer(ctx context.Context, workspaceID string, agentID string, variationID string, id string, body AgentVariationUpdateMemoryLayerParams, opts ...option.RequestOption) (res *VariationMemoryLayerAssignment, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/variations/%s/memory_layer_assignments/%s", workspaceID, agentID, variationID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
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
	JSON agentVariationJSON `json:"-"`
}

// agentVariationJSON contains the JSON metadata for the struct [AgentVariation]
type agentVariationJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentVariation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationJSON) RawJSON() string {
	return r.raw
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
	ToolSetCount int64                  `json:"toolSetCount"`
	JSON         agentVariationInfoJSON `json:"-"`
}

// agentVariationInfoJSON contains the JSON metadata for the struct
// [AgentVariationInfo]
type agentVariationInfoJSON struct {
	Assignments            apijson.Field
	CreatedBy              apijson.Field
	FeedbackCount          apijson.Field
	MemoryLayerAssignments apijson.Field
	MemoryLayerCount       apijson.Field
	Model                  apijson.Field
	Score                  apijson.Field
	SubAgentCount          apijson.Field
	ToolCount              apijson.Field
	ToolSetCount           apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AgentVariationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationInfoJSON) RawJSON() string {
	return r.raw
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
	SystemPromptTemplate string                 `json:"systemPromptTemplate"`
	JSON                 agentVariationSpecJSON `json:"-"`
}

// agentVariationSpecJSON contains the JSON metadata for the struct
// [AgentVariationSpec]
type agentVariationSpecJSON struct {
	CompactionConfig         apijson.Field
	Constraints              apijson.Field
	Description              apijson.Field
	FirstUserMessageTemplate apijson.Field
	ModelConfig              apijson.Field
	ProgressiveDiscovery     apijson.Field
	SystemPromptTemplate     apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AgentVariationSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecJSON) RawJSON() string {
	return r.raw
}

// AgentVariationSpec defines the operational configuration for a variation
type AgentVariationSpecParam struct {
	// CompactionConfig defines how context window compaction behaves for objectives
	// using this variation.
	CompactionConfig param.Field[AgentVariationSpecCompactionConfigParam] `json:"compactionConfig"`
	// Execution constraints
	Constraints param.Field[AgentVariationSpecConstraintsParam] `json:"constraints"`
	// Human-readable description of what this variation does or when it should be used
	Description param.Field[string] `json:"description"`
	// Liquid template for the first user message of objectives using this variation.
	// Rendered with CreateObjectiveRequest.first_user_message_data into
	// Objective.first_user_message, the first user message in the LLM chat history.
	// CreateObjectiveRequest.first_user_message, when set, overrides the rendered
	// result. If neither this template nor first_user_message is present, objective
	// creation is rejected with InvalidArgument.
	FirstUserMessageTemplate param.Field[string] `json:"firstUserMessageTemplate"`
	// ModelConfig defines the model configuration for a variation
	ModelConfig param.Field[AgentVariationSpecModelConfigParam] `json:"modelConfig"`
	// ProgressiveDiscovery is used to indicate that the agent should automatically
	// discover tools that are not explicitly assigned to it. Max tools is the maximum
	// number of tools that can be discovered per search. Hints are optional hints for
	// tool search. These are used in conjunction with the context-aware tool search
	// and can help select the best tools for the task.
	ProgressiveDiscovery param.Field[AgentVariationSpecProgressiveDiscoveryParam] `json:"progressiveDiscovery"`
	// Liquid template for the system prompt of objectives using this variation.
	// Rendered with CreateObjectiveRequest.system_prompt_data into
	// Objective.system_prompt.
	SystemPromptTemplate param.Field[string] `json:"systemPromptTemplate"`
}

func (r AgentVariationSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	TriggerThreshold float64                                `json:"triggerThreshold"`
	JSON             agentVariationSpecCompactionConfigJSON `json:"-"`
}

// agentVariationSpecCompactionConfigJSON contains the JSON metadata for the struct
// [AgentVariationSpecCompactionConfig]
type agentVariationSpecCompactionConfigJSON struct {
	Summarization      apijson.Field
	ToolResultClearing apijson.Field
	TriggerThreshold   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AgentVariationSpecCompactionConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecCompactionConfigJSON) RawJSON() string {
	return r.raw
}

// CompactionConfig defines how context window compaction behaves for objectives
// using this variation.
type AgentVariationSpecCompactionConfigParam struct {
	// SummarizationStrategy configures LLM-powered summarization of older conversation
	// turns.
	Summarization param.Field[CompactionConfigSummarizationStrategyParam] `json:"summarization"`
	// ToolResultClearingStrategy configures clearing of older tool result content.
	ToolResultClearing param.Field[CompactionConfigToolResultClearingStrategyParam] `json:"toolResultClearing"`
	// Trigger threshold as a percentage of the model's context window (0.0 to 1.0).
	// When input tokens reach this percentage of the model's limit, compaction
	// triggers. Default: 0.75 (75%)
	TriggerThreshold param.Field[float64] `json:"triggerThreshold"`
}

func (r AgentVariationSpecCompactionConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	MaxToolCalls int64                             `json:"maxToolCalls"`
	JSON         agentVariationSpecConstraintsJSON `json:"-"`
}

// agentVariationSpecConstraintsJSON contains the JSON metadata for the struct
// [AgentVariationSpecConstraints]
type agentVariationSpecConstraintsJSON struct {
	InactivityTimeout apijson.Field
	MaxSubObjectives  apijson.Field
	MaxToolCalls      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AgentVariationSpecConstraints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecConstraintsJSON) RawJSON() string {
	return r.raw
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
	InactivityTimeout param.Field[string] `json:"inactivityTimeout"`
	// The maximum number of sub-objectives that can be created. 0 means no limit.
	MaxSubObjectives param.Field[int64] `json:"maxSubObjectives"`
	// The maximum number of tool calls that can be made. 0 means no limit.
	MaxToolCalls param.Field[int64] `json:"maxToolCalls"`
}

func (r AgentVariationSpecConstraintsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ModelConfig defines the model configuration for a variation
type AgentVariationSpecModelConfig struct {
	// The model identifier in family/model format (e.g., "claude/opus-4.6",
	// "claude/sonnet-4.5")
	ModelID string `json:"modelId"`
	// Sampling temperature for model inference (0.0 to 1.0) Lower values produce more
	// deterministic outputs, higher values increase randomness
	Temperature float64                           `json:"temperature"`
	JSON        agentVariationSpecModelConfigJSON `json:"-"`
}

// agentVariationSpecModelConfigJSON contains the JSON metadata for the struct
// [AgentVariationSpecModelConfig]
type agentVariationSpecModelConfigJSON struct {
	ModelID     apijson.Field
	Temperature apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentVariationSpecModelConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecModelConfigJSON) RawJSON() string {
	return r.raw
}

// ModelConfig defines the model configuration for a variation
type AgentVariationSpecModelConfigParam struct {
	// The model identifier in family/model format (e.g., "claude/opus-4.6",
	// "claude/sonnet-4.5")
	ModelID param.Field[string] `json:"modelId"`
	// Sampling temperature for model inference (0.0 to 1.0) Lower values produce more
	// deterministic outputs, higher values increase randomness
	Temperature param.Field[float64] `json:"temperature"`
}

func (r AgentVariationSpecModelConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	MaxTools int64                                      `json:"maxTools"`
	JSON     agentVariationSpecProgressiveDiscoveryJSON `json:"-"`
}

// agentVariationSpecProgressiveDiscoveryJSON contains the JSON metadata for the
// struct [AgentVariationSpecProgressiveDiscovery]
type agentVariationSpecProgressiveDiscoveryJSON struct {
	Hints       apijson.Field
	MaxTools    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentVariationSpecProgressiveDiscovery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecProgressiveDiscoveryJSON) RawJSON() string {
	return r.raw
}

// ProgressiveDiscovery is used to indicate that the agent should automatically
// discover tools that are not explicitly assigned to it. Max tools is the maximum
// number of tools that can be discovered per search. Hints are optional hints for
// tool search. These are used in conjunction with the context-aware tool search
// and can help select the best tools for the task.
type AgentVariationSpecProgressiveDiscoveryParam struct {
	// Free-text guidance appended to the discoverable-tools appendix in the system
	// prompt. Hints steer the model's choice of tool names; they do not filter or rank
	// anything, because tool_search matches names exactly rather than searching.
	Hints param.Field[[]string] `json:"hints"`
	// The most tool names tool_search will load in a single call. Requesting more than
	// this returns an error telling the model to retry in smaller batches -- it is a
	// per-call batch limit, not a ceiling on how many tools an objective may end up
	// with.
	MaxTools param.Field[int64] `json:"maxTools"`
}

func (r AgentVariationSpecProgressiveDiscoveryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SummarizationStrategy configures LLM-powered summarization of older conversation
// turns.
type CompactionConfigSummarizationStrategy struct {
	// Custom instructions that guide what the summarizer preserves. Replaces the
	// default summarization prompt entirely. Example: "Preserve all code snippets,
	// variable names, and technical decisions."
	Instructions string                                    `json:"instructions"`
	JSON         compactionConfigSummarizationStrategyJSON `json:"-"`
}

// compactionConfigSummarizationStrategyJSON contains the JSON metadata for the
// struct [CompactionConfigSummarizationStrategy]
type compactionConfigSummarizationStrategyJSON struct {
	Instructions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CompactionConfigSummarizationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compactionConfigSummarizationStrategyJSON) RawJSON() string {
	return r.raw
}

// SummarizationStrategy configures LLM-powered summarization of older conversation
// turns.
type CompactionConfigSummarizationStrategyParam struct {
	// Custom instructions that guide what the summarizer preserves. Replaces the
	// default summarization prompt entirely. Example: "Preserve all code snippets,
	// variable names, and technical decisions."
	Instructions param.Field[string] `json:"instructions"`
}

func (r CompactionConfigSummarizationStrategyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ToolResultClearingStrategy configures clearing of older tool result content.
type CompactionConfigToolResultClearingStrategy struct {
	// Number of most recent tool call results to keep intact. Older tool results have
	// their content replaced with "[result cleared]" while preserving the assistant
	// tool call message (function name, arguments). Default: 2
	PreserveRecentResults int64                                          `json:"preserveRecentResults"`
	JSON                  compactionConfigToolResultClearingStrategyJSON `json:"-"`
}

// compactionConfigToolResultClearingStrategyJSON contains the JSON metadata for
// the struct [CompactionConfigToolResultClearingStrategy]
type compactionConfigToolResultClearingStrategyJSON struct {
	PreserveRecentResults apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CompactionConfigToolResultClearingStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compactionConfigToolResultClearingStrategyJSON) RawJSON() string {
	return r.raw
}

// ToolResultClearingStrategy configures clearing of older tool result content.
type CompactionConfigToolResultClearingStrategyParam struct {
	// Number of most recent tool call results to keep intact. Older tool results have
	// their content replaced with "[result cleared]" while preserving the assistant
	// tool call message (function name, arguments). Default: 2
	PreserveRecentResults param.Field[int64] `json:"preserveRecentResults"`
}

func (r CompactionConfigToolResultClearingStrategyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	ToolSet shared.BareMetadata     `json:"toolSet"`
	JSON    variationAssignmentJSON `json:"-"`
}

// variationAssignmentJSON contains the JSON metadata for the struct
// [VariationAssignment]
type variationAssignmentJSON struct {
	ID          apijson.Field
	Agent       apijson.Field
	Tool        apijson.Field
	ToolSet     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariationAssignment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variationAssignmentJSON) RawJSON() string {
	return r.raw
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
	Position int64                              `json:"position"`
	JSON     variationMemoryLayerAssignmentJSON `json:"-"`
}

// variationMemoryLayerAssignmentJSON contains the JSON metadata for the struct
// [VariationMemoryLayerAssignment]
type variationMemoryLayerAssignmentJSON struct {
	ID          apijson.Field
	MemoryLayer apijson.Field
	Position    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariationMemoryLayerAssignment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variationMemoryLayerAssignmentJSON) RawJSON() string {
	return r.raw
}

type AgentVariationNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec param.Field[AgentVariationSpecParam] `json:"spec" api:"required"`
}

func (r AgentVariationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentVariationUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec param.Field[AgentVariationSpecParam] `json:"spec"`
	// Fields to update
	UpdateMask param.Field[string] `json:"updateMask" format:"field-mask"`
}

func (r AgentVariationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentVariationListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When true, the `info` field on each returned variation is populated. Requests
	// with this flag count more against your rate limit.
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Field[string] `query:"labels"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [AgentVariationListParams]'s query parameters as
// `url.Values`.
func (r AgentVariationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentVariationAddAssignmentParams struct {
	SubAgentID param.Field[string] `json:"subAgentId"`
	ToolID     param.Field[string] `json:"toolId"`
	ToolSetID  param.Field[string] `json:"toolSetId"`
}

func (r AgentVariationAddAssignmentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentVariationAddMemoryLayerParams struct {
	// Layer to attach. Accepts the canonical `memlyr_…` form or the
	// `external_id:<value>` form.
	MemoryLayerID param.Field[string] `json:"memoryLayerId"`
	// Position in the baseline cascade (lower = more specific). If omitted, the server
	// appends at the most general end (max existing position + 1).
	Position param.Field[int64] `json:"position"`
}

func (r AgentVariationAddMemoryLayerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentVariationUpdateMemoryLayerParams struct {
	// New position. Only field currently updatable on an assignment.
	Position param.Field[int64] `json:"position"`
}

func (r AgentVariationUpdateMemoryLayerParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
