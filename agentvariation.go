// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocadenyacomcadenyasdkgo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cadenya/cadenya-sdk-go/internal/apijson"
	"github.com/cadenya/cadenya-sdk-go/internal/apiquery"
	"github.com/cadenya/cadenya-sdk-go/internal/param"
	"github.com/cadenya/cadenya-sdk-go/internal/requestconfig"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/cadenya/cadenya-sdk-go/packages/pagination"
	"github.com/cadenya/cadenya-sdk-go/shared"
)

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
func (r *AgentVariationService) New(ctx context.Context, agentID string, body AgentVariationNewParams, opts ...option.RequestOption) (res *AgentVariation, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/variations", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a variation by ID from an agent
func (r *AgentVariationService) Get(ctx context.Context, agentID string, id string, opts ...option.RequestOption) (res *AgentVariation, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/variations/%s", agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a variation for an agent
func (r *AgentVariationService) Update(ctx context.Context, agentID string, id string, body AgentVariationUpdateParams, opts ...option.RequestOption) (res *AgentVariation, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/variations/%s", agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all variations for an agent
func (r *AgentVariationService) List(ctx context.Context, agentID string, query AgentVariationListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[AgentVariation], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/%s/variations", agentID)
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
func (r *AgentVariationService) ListAutoPaging(ctx context.Context, agentID string, query AgentVariationListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[AgentVariation] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, agentID, query, opts...))
}

// Deletes a variation from an agent
func (r *AgentVariationService) Delete(ctx context.Context, agentID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agentId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/agents/%s/variations/%s", agentID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
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

// AgentVariation resource
type AgentVariationParam struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata param.Field[shared.ResourceMetadataParam] `json:"metadata" api:"required"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec param.Field[AgentVariationSpecParam] `json:"spec" api:"required"`
}

func (r AgentVariationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AgentVariationInfo provides read-only summary information about a variation
type AgentVariationInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy shared.Profile `json:"createdBy"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Model shared.ResourceMetadata `json:"model"`
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
	CreatedBy     apijson.Field
	Model         apijson.Field
	SubAgentCount apijson.Field
	ToolCount     apijson.Field
	ToolSetCount  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgentVariationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationInfoJSON) RawJSON() string {
	return r.raw
}

// AgentVariationInfo provides read-only summary information about a variation
type AgentVariationInfoParam struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy param.Field[shared.ProfileParam] `json:"createdBy"`
}

func (r AgentVariationInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AgentVariationSpec defines the operational configuration for a variation
type AgentVariationSpec struct {
	// Documents assigned to this variation. Can include individual documents or entire
	// document namespaces (which include all documents in the namespace).
	AgentDocuments []AgentVariationSpecAgentDocument `json:"agentDocuments"`
	// Tools assigned to this variation
	AgentTools []AgentVariationSpecAgentTool `json:"agentTools"`
	// Execution constraints
	Constraints AgentVariationSpecConstraints `json:"constraints"`
	// Human-readable description of what this variation does or when it should be used
	Description string `json:"description"`
	// Enable episodic memory for objectives using this variation. When true, the
	// system automatically creates a document namespace for each objective using the
	// objective's episodic_key as the external_id, allowing the agent to store and
	// retrieve documents specific to that episode.
	EnableEpisodicMemory bool `json:"enableEpisodicMemory"`
	// How long episodic memories should be retained. After this duration, episodic
	// document namespaces can be automatically cleaned up. If not set, episodic
	// memories are retained indefinitely.
	EpisodicMemoryTtl int64 `json:"episodicMemoryTtl"`
	// ModelConfig defines the model configuration for a variation
	ModelConfig AgentVariationSpecModelConfig `json:"modelConfig"`
	// The system prompt for this variation
	Prompt string `json:"prompt"`
	// Tool selection strategy
	ToolSelection AgentVariationSpecToolSelection `json:"toolSelection"`
	// Weight for weighted random selection (>= 0). P(v) = v.weight / sum(all_weights).
	// Only used when the agent's variation_selection_mode is WEIGHTED. A weight of 0
	// means never auto-selected, but can still be chosen explicitly via variation_id
	// on CreateObjectiveRequest.
	Weight int64                  `json:"weight"`
	JSON   agentVariationSpecJSON `json:"-"`
}

// agentVariationSpecJSON contains the JSON metadata for the struct
// [AgentVariationSpec]
type agentVariationSpecJSON struct {
	AgentDocuments       apijson.Field
	AgentTools           apijson.Field
	Constraints          apijson.Field
	Description          apijson.Field
	EnableEpisodicMemory apijson.Field
	EpisodicMemoryTtl    apijson.Field
	ModelConfig          apijson.Field
	Prompt               apijson.Field
	ToolSelection        apijson.Field
	Weight               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AgentVariationSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecJSON) RawJSON() string {
	return r.raw
}

// AgentVariationSpec defines the operational configuration for a variation
type AgentVariationSpecParam struct {
	// Documents assigned to this variation. Can include individual documents or entire
	// document namespaces (which include all documents in the namespace).
	AgentDocuments param.Field[[]AgentVariationSpecAgentDocumentParam] `json:"agentDocuments"`
	// Tools assigned to this variation
	AgentTools param.Field[[]AgentVariationSpecAgentToolParam] `json:"agentTools"`
	// Execution constraints
	Constraints param.Field[AgentVariationSpecConstraintsParam] `json:"constraints"`
	// Human-readable description of what this variation does or when it should be used
	Description param.Field[string] `json:"description"`
	// Enable episodic memory for objectives using this variation. When true, the
	// system automatically creates a document namespace for each objective using the
	// objective's episodic_key as the external_id, allowing the agent to store and
	// retrieve documents specific to that episode.
	EnableEpisodicMemory param.Field[bool] `json:"enableEpisodicMemory"`
	// How long episodic memories should be retained. After this duration, episodic
	// document namespaces can be automatically cleaned up. If not set, episodic
	// memories are retained indefinitely.
	EpisodicMemoryTtl param.Field[int64] `json:"episodicMemoryTtl"`
	// ModelConfig defines the model configuration for a variation
	ModelConfig param.Field[AgentVariationSpecModelConfigParam] `json:"modelConfig"`
	// The system prompt for this variation
	Prompt param.Field[string] `json:"prompt"`
	// Tool selection strategy
	ToolSelection param.Field[AgentVariationSpecToolSelectionParam] `json:"toolSelection"`
	// Weight for weighted random selection (>= 0). P(v) = v.weight / sum(all_weights).
	// Only used when the agent's variation_selection_mode is WEIGHTED. A weight of 0
	// means never auto-selected, but can still be chosen explicitly via variation_id
	// on CreateObjectiveRequest.
	Weight param.Field[int64] `json:"weight"`
}

func (r AgentVariationSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentVariationSpecAgentDocument struct {
	DocumentID string `json:"documentId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	DocumentMetadata    shared.ResourceMetadata `json:"documentMetadata"`
	DocumentNamespaceID string                  `json:"documentNamespaceId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	DocumentNamespaceMetadata shared.ResourceMetadata             `json:"documentNamespaceMetadata"`
	JSON                      agentVariationSpecAgentDocumentJSON `json:"-"`
}

// agentVariationSpecAgentDocumentJSON contains the JSON metadata for the struct
// [AgentVariationSpecAgentDocument]
type agentVariationSpecAgentDocumentJSON struct {
	DocumentID                apijson.Field
	DocumentMetadata          apijson.Field
	DocumentNamespaceID       apijson.Field
	DocumentNamespaceMetadata apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AgentVariationSpecAgentDocument) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecAgentDocumentJSON) RawJSON() string {
	return r.raw
}

type AgentVariationSpecAgentDocumentParam struct {
	DocumentID param.Field[string] `json:"documentId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	DocumentMetadata    param.Field[shared.ResourceMetadataParam] `json:"documentMetadata"`
	DocumentNamespaceID param.Field[string]                       `json:"documentNamespaceId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	DocumentNamespaceMetadata param.Field[shared.ResourceMetadataParam] `json:"documentNamespaceMetadata"`
}

func (r AgentVariationSpecAgentDocumentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentVariationSpecAgentTool struct {
	AgentID string `json:"agentId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	AgentMetadata shared.ResourceMetadata `json:"agentMetadata"`
	ToolID        string                  `json:"toolId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolMetadata shared.ResourceMetadata `json:"toolMetadata"`
	ToolSetID    string                  `json:"toolSetId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolSetMetadata shared.ResourceMetadata         `json:"toolSetMetadata"`
	JSON            agentVariationSpecAgentToolJSON `json:"-"`
}

// agentVariationSpecAgentToolJSON contains the JSON metadata for the struct
// [AgentVariationSpecAgentTool]
type agentVariationSpecAgentToolJSON struct {
	AgentID         apijson.Field
	AgentMetadata   apijson.Field
	ToolID          apijson.Field
	ToolMetadata    apijson.Field
	ToolSetID       apijson.Field
	ToolSetMetadata apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AgentVariationSpecAgentTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecAgentToolJSON) RawJSON() string {
	return r.raw
}

type AgentVariationSpecAgentToolParam struct {
	AgentID param.Field[string] `json:"agentId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	AgentMetadata param.Field[shared.ResourceMetadataParam] `json:"agentMetadata"`
	ToolID        param.Field[string]                       `json:"toolId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolMetadata param.Field[shared.ResourceMetadataParam] `json:"toolMetadata"`
	ToolSetID    param.Field[string]                       `json:"toolSetId"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	ToolSetMetadata param.Field[shared.ResourceMetadataParam] `json:"toolSetMetadata"`
}

func (r AgentVariationSpecAgentToolParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentVariationSpecConstraints struct {
	// The maximum number of sub-objectives that can be created. 0 means no limit.
	MaxSubObjectives int64 `json:"maxSubObjectives"`
	// The maximum number of tool calls that can be made. 0 means no limit.
	MaxToolCalls int64                             `json:"maxToolCalls"`
	JSON         agentVariationSpecConstraintsJSON `json:"-"`
}

// agentVariationSpecConstraintsJSON contains the JSON metadata for the struct
// [AgentVariationSpecConstraints]
type agentVariationSpecConstraintsJSON struct {
	MaxSubObjectives apijson.Field
	MaxToolCalls     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AgentVariationSpecConstraints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecConstraintsJSON) RawJSON() string {
	return r.raw
}

type AgentVariationSpecConstraintsParam struct {
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

type AgentVariationSpecToolSelection struct {
	// AssignedTools is used to indicate that the agent should only use the tools/tool
	// sets that are explicitly assigned to it. Allow discovery is used when the agent
	// thinks it needs to discover more tools.
	AssignedTools ToolSelectionAssignedTools `json:"assignedTools"`
	// AutoDiscovery is used to indicate that the agent should automatically discover
	// tools that are not explicitly assigned to it. Max tools is the maximum number of
	// tools that can be discovered. Hints are optional hints for tool search. These
	// are used in conjunction with the context-aware tool search and can help select
	// the best tools for the task.
	AutoDiscovery ToolSelectionAutoDiscovery          `json:"autoDiscovery"`
	JSON          agentVariationSpecToolSelectionJSON `json:"-"`
}

// agentVariationSpecToolSelectionJSON contains the JSON metadata for the struct
// [AgentVariationSpecToolSelection]
type agentVariationSpecToolSelectionJSON struct {
	AssignedTools apijson.Field
	AutoDiscovery apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AgentVariationSpecToolSelection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationSpecToolSelectionJSON) RawJSON() string {
	return r.raw
}

type AgentVariationSpecToolSelectionParam struct {
	// AssignedTools is used to indicate that the agent should only use the tools/tool
	// sets that are explicitly assigned to it. Allow discovery is used when the agent
	// thinks it needs to discover more tools.
	AssignedTools param.Field[ToolSelectionAssignedToolsParam] `json:"assignedTools"`
	// AutoDiscovery is used to indicate that the agent should automatically discover
	// tools that are not explicitly assigned to it. Max tools is the maximum number of
	// tools that can be discovered. Hints are optional hints for tool search. These
	// are used in conjunction with the context-aware tool search and can help select
	// the best tools for the task.
	AutoDiscovery param.Field[ToolSelectionAutoDiscoveryParam] `json:"autoDiscovery"`
}

func (r AgentVariationSpecToolSelectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AssignedTools is used to indicate that the agent should only use the tools/tool
// sets that are explicitly assigned to it. Allow discovery is used when the agent
// thinks it needs to discover more tools.
type ToolSelectionAssignedTools struct {
	AllowDiscovery bool                           `json:"allowDiscovery"`
	JSON           toolSelectionAssignedToolsJSON `json:"-"`
}

// toolSelectionAssignedToolsJSON contains the JSON metadata for the struct
// [ToolSelectionAssignedTools]
type toolSelectionAssignedToolsJSON struct {
	AllowDiscovery apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ToolSelectionAssignedTools) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSelectionAssignedToolsJSON) RawJSON() string {
	return r.raw
}

// AssignedTools is used to indicate that the agent should only use the tools/tool
// sets that are explicitly assigned to it. Allow discovery is used when the agent
// thinks it needs to discover more tools.
type ToolSelectionAssignedToolsParam struct {
	AllowDiscovery param.Field[bool] `json:"allowDiscovery"`
}

func (r ToolSelectionAssignedToolsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AutoDiscovery is used to indicate that the agent should automatically discover
// tools that are not explicitly assigned to it. Max tools is the maximum number of
// tools that can be discovered. Hints are optional hints for tool search. These
// are used in conjunction with the context-aware tool search and can help select
// the best tools for the task.
type ToolSelectionAutoDiscovery struct {
	Hints    []string                       `json:"hints"`
	MaxTools int64                          `json:"maxTools"`
	JSON     toolSelectionAutoDiscoveryJSON `json:"-"`
}

// toolSelectionAutoDiscoveryJSON contains the JSON metadata for the struct
// [ToolSelectionAutoDiscovery]
type toolSelectionAutoDiscoveryJSON struct {
	Hints       apijson.Field
	MaxTools    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSelectionAutoDiscovery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSelectionAutoDiscoveryJSON) RawJSON() string {
	return r.raw
}

// AutoDiscovery is used to indicate that the agent should automatically discover
// tools that are not explicitly assigned to it. Max tools is the maximum number of
// tools that can be discovered. Hints are optional hints for tool search. These
// are used in conjunction with the context-aware tool search and can help select
// the best tools for the task.
type ToolSelectionAutoDiscoveryParam struct {
	Hints    param.Field[[]string] `json:"hints"`
	MaxTools param.Field[int64]    `json:"maxTools"`
}

func (r ToolSelectionAutoDiscoveryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
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
