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

// Manage AI agents within a workspace. Agents define AI behavior and tool access.
//
// AgentService contains methods and other services that help with interacting with
// the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	options []option.RequestOption
	// Manage AI agents within a workspace. Agents define AI behavior and tool access.
	Feedback AgentFeedbackService
	// Manage AI agents within a workspace. Agents define AI behavior and tool access.
	WebhookDeliveries AgentWebhookDeliveryService
	// Manage variations of an agent and their tool, sub-agent, and memory layer
	// assignments.
	Variations AgentVariationService
	// Manage recurring schedules attached to agents. Schedules trigger objectives on a
	// cadence defined by AgentScheduleSpec.Schedule.
	Schedules AgentScheduleService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.options = opts
	r.Feedback = NewAgentFeedbackService(opts...)
	r.WebhookDeliveries = NewAgentWebhookDeliveryService(opts...)
	r.Variations = NewAgentVariationService(opts...)
	r.Schedules = NewAgentScheduleService(opts...)
	return
}

// Creates a new agent in the workspace
func (r *AgentService) New(ctx context.Context, params AgentNewParams, opts ...option.RequestOption) (res *Agent, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves an agent by ID from the workspace
func (r *AgentService) Get(ctx context.Context, id string, query AgentGetParams, opts ...option.RequestOption) (res *Agent, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an agent in the workspace
func (r *AgentService) Update(ctx context.Context, id string, params AgentUpdateParams, opts ...option.RequestOption) (res *Agent, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Lists all agents in the workspace
func (r *AgentService) List(ctx context.Context, params AgentListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Agent], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents", url.PathEscape(params.WorkspaceID.Value))
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

// Lists all agents in the workspace
func (r *AgentService) ListAutoPaging(ctx context.Context, params AgentListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Agent] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes an agent from the workspace
func (r *AgentService) Delete(ctx context.Context, id string, body AgentDeleteParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Transitions an agent to STATE_ARCHIVED. Archived agents are hidden from list
// results and cannot be used for objectives; active schedules are paused.
func (r *AgentService) Archive(ctx context.Context, id string, body AgentArchiveParams, opts ...option.RequestOption) (res *Agent, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s:archive", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions an agent to STATE_PUBLISHED, making it available for objectives. The
// agent must have at least one variation.
func (r *AgentService) Publish(ctx context.Context, id string, body AgentPublishParams, opts ...option.RequestOption) (res *Agent, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s:publish", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions an archived agent back to STATE_DRAFT. Publish the agent again to
// make it available for objectives.
func (r *AgentService) Unarchive(ctx context.Context, id string, body AgentUnarchiveParams, opts ...option.RequestOption) (res *Agent, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s:unarchive", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions a published agent back to STATE_DRAFT. Active schedules for the
// agent are paused until it is published again.
func (r *AgentService) Unpublish(ctx context.Context, id string, body AgentUnpublishParams, opts ...option.RequestOption) (res *Agent, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s:unpublish", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Agent resource
type Agent struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// Agent specification (user-provided configuration)
	Spec AgentSpec `json:"spec" api:"required"`
	// The current lifecycle state of the agent. Output only. Agents are created in
	// STATE_DRAFT; use the :publish, :unpublish, :archive, and :unarchive actions to
	// transition between states.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_DRAFT", "STATE_PUBLISHED", "STATE_ARCHIVED".
	State AgentState `json:"state" api:"required"`
	// AgentInfo contains simple information about an agent for display or quick
	// reference
	Info AgentInfo `json:"info"`
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
func (r Agent) RawJSON() string { return r.JSON.raw }
func (r *Agent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current lifecycle state of the agent. Output only. Agents are created in
// STATE_DRAFT; use the :publish, :unpublish, :archive, and :unarchive actions to
// transition between states.
type AgentState string

const (
	AgentStateStateUnspecified AgentState = "STATE_UNSPECIFIED"
	AgentStateStateDraft       AgentState = "STATE_DRAFT"
	AgentStateStatePublished   AgentState = "STATE_PUBLISHED"
	AgentStateStateArchived    AgentState = "STATE_ARCHIVED"
)

// AgentInfo contains simple information about an agent for display or quick
// reference
type AgentInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy      Profile `json:"createdBy"`
	VariationCount int64   `json:"variationCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy      respjson.Field
		VariationCount respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentInfo) RawJSON() string { return r.JSON.raw }
func (r *AgentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Agent specification (user-provided configuration)
type AgentSpec struct {
	// Controls how variations are automatically selected when creating objectives
	// Defaults to RANDOM when unspecified
	//
	// Any of "VARIATION_SELECTION_MODE_UNSPECIFIED",
	// "VARIATION_SELECTION_MODE_RANDOM", "VARIATION_SELECTION_MODE_WEIGHTED".
	VariationSelectionMode AgentSpecVariationSelectionMode `json:"variationSelectionMode" api:"required"`
	// Description of the agent's purpose
	Description string `json:"description"`
	// Enable episodic memory for objectives created for this agent. When true,
	// objective creation requires an episodic_memory key and the system finds or
	// creates a memory layer for that (agent, key) pair, letting the agent store and
	// retrieve memories across objectives that share the key. Memory is agent-level so
	// all variations of the agent share the same layers.
	EnableEpisodicMemory bool `json:"enableEpisodicMemory"`
	// How long episodic memories should be retained. Each new objective slides the
	// layer's expiry forward by this duration, and stored entries expire this long
	// after they are written. If not set, episodic memories are retained indefinitely.
	EpisodicMemoryTtl int64 `json:"episodicMemoryTtl"`
	// Optional output definition for objectives created for this agent. When provided,
	// Cadenya will append a tool to that will be called by the LLM in use by the
	// variant to extract information in the format provided here. Use this option when
	// you want structured data to be created by your objectives.
	OutputDefinition map[string]any `json:"outputDefinition"`
	// SystemPromptDataSchema enforces the shape of system_prompt_data when objectives
	// are created. This is valuable when using liquid formatting in agent variation
	// system prompt templates. The schema is also used when the agent is attached as a
	// sub-agent, as it becomes the tool's input parameter schema. If omitted, the
	// sub-agent schema will be loaded with a simple "prompt" free text string as its
	// schema.
	SystemPromptDataSchema map[string]any `json:"systemPromptDataSchema"`
	// The URL that Cadenya will send events for any objective assigned to the agent.
	WebhookEventsURL string `json:"webhookEventsUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VariationSelectionMode respjson.Field
		Description            respjson.Field
		EnableEpisodicMemory   respjson.Field
		EpisodicMemoryTtl      respjson.Field
		OutputDefinition       respjson.Field
		SystemPromptDataSchema respjson.Field
		WebhookEventsURL       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSpec) RawJSON() string { return r.JSON.raw }
func (r *AgentSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentSpec to a AgentSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentSpecParam.Overrides()
func (r AgentSpec) ToParam() AgentSpecParam {
	return param.Override[AgentSpecParam](json.RawMessage(r.RawJSON()))
}

// Controls how variations are automatically selected when creating objectives
// Defaults to RANDOM when unspecified
type AgentSpecVariationSelectionMode string

const (
	AgentSpecVariationSelectionModeVariationSelectionModeUnspecified AgentSpecVariationSelectionMode = "VARIATION_SELECTION_MODE_UNSPECIFIED"
	AgentSpecVariationSelectionModeVariationSelectionModeRandom      AgentSpecVariationSelectionMode = "VARIATION_SELECTION_MODE_RANDOM"
	AgentSpecVariationSelectionModeVariationSelectionModeWeighted    AgentSpecVariationSelectionMode = "VARIATION_SELECTION_MODE_WEIGHTED"
)

// Agent specification (user-provided configuration)
//
// The property VariationSelectionMode is required.
type AgentSpecParam struct {
	// Controls how variations are automatically selected when creating objectives
	// Defaults to RANDOM when unspecified
	//
	// Any of "VARIATION_SELECTION_MODE_UNSPECIFIED",
	// "VARIATION_SELECTION_MODE_RANDOM", "VARIATION_SELECTION_MODE_WEIGHTED".
	VariationSelectionMode AgentSpecVariationSelectionMode `json:"variationSelectionMode,omitzero" api:"required"`
	// Description of the agent's purpose
	Description param.Opt[string] `json:"description,omitzero"`
	// Enable episodic memory for objectives created for this agent. When true,
	// objective creation requires an episodic_memory key and the system finds or
	// creates a memory layer for that (agent, key) pair, letting the agent store and
	// retrieve memories across objectives that share the key. Memory is agent-level so
	// all variations of the agent share the same layers.
	EnableEpisodicMemory param.Opt[bool] `json:"enableEpisodicMemory,omitzero"`
	// How long episodic memories should be retained. Each new objective slides the
	// layer's expiry forward by this duration, and stored entries expire this long
	// after they are written. If not set, episodic memories are retained indefinitely.
	EpisodicMemoryTtl param.Opt[int64] `json:"episodicMemoryTtl,omitzero"`
	// The URL that Cadenya will send events for any objective assigned to the agent.
	WebhookEventsURL param.Opt[string] `json:"webhookEventsUrl,omitzero"`
	// Optional output definition for objectives created for this agent. When provided,
	// Cadenya will append a tool to that will be called by the LLM in use by the
	// variant to extract information in the format provided here. Use this option when
	// you want structured data to be created by your objectives.
	OutputDefinition map[string]any `json:"outputDefinition,omitzero"`
	// SystemPromptDataSchema enforces the shape of system_prompt_data when objectives
	// are created. This is valuable when using liquid formatting in agent variation
	// system prompt templates. The schema is also used when the agent is attached as a
	// sub-agent, as it becomes the tool's input parameter schema. If omitted, the
	// sub-agent schema will be loaded with a simple "prompt" free text string as its
	// schema.
	SystemPromptDataSchema map[string]any `json:"systemPromptDataSchema,omitzero"`
	paramObj
}

func (r AgentSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Page carries cursor-based pagination state. There is no total: the cursor walks
// the result set without ever counting it, and a count would cost a second query
// on every list.
type Page struct {
	NextCursor string `json:"nextCursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Page) RawJSON() string { return r.JSON.raw }
func (r *Page) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	// Agent specification (user-provided configuration)
	Spec AgentSpecParam `json:"spec,omitzero" api:"required"`
	// Create agent variation request
	DefaultVariation AgentNewParamsDefaultVariation `json:"defaultVariation,omitzero"`
	paramObj
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Create agent variation request
//
// The properties Metadata, Spec are required.
type AgentNewParamsDefaultVariation struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.CreateResourceMetadataParam `json:"metadata,omitzero" api:"required"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec AgentVariationSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r AgentNewParamsDefaultVariation) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParamsDefaultVariation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParamsDefaultVariation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AgentUpdateParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Fields to update
	UpdateMask param.Opt[string] `json:"updateMask,omitzero" format:"field-mask"`
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata shared.UpdateResourceMetadataParam `json:"metadata,omitzero"`
	// Agent specification (user-provided configuration)
	Spec AgentSpecParam `json:"spec,omitzero"`
	paramObj
}

func (r AgentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When true, the `info` field on each returned agent is populated. Requests with
	// this flag count more against your rate limit.
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
	// Filter by agent lifecycle state
	//
	// Any of "STATE_UNSPECIFIED", "STATE_DRAFT", "STATE_PUBLISHED", "STATE_ARCHIVED".
	State AgentListParamsState `query:"state,omitzero" json:"-"`
	// Filter by variation selection mode
	//
	// Any of "VARIATION_SELECTION_MODE_UNSPECIFIED",
	// "VARIATION_SELECTION_MODE_RANDOM", "VARIATION_SELECTION_MODE_WEIGHTED".
	VariationSelectionMode AgentListParamsVariationSelectionMode `query:"variationSelectionMode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by agent lifecycle state
type AgentListParamsState string

const (
	AgentListParamsStateStateUnspecified AgentListParamsState = "STATE_UNSPECIFIED"
	AgentListParamsStateStateDraft       AgentListParamsState = "STATE_DRAFT"
	AgentListParamsStateStatePublished   AgentListParamsState = "STATE_PUBLISHED"
	AgentListParamsStateStateArchived    AgentListParamsState = "STATE_ARCHIVED"
)

// Filter by variation selection mode
type AgentListParamsVariationSelectionMode string

const (
	AgentListParamsVariationSelectionModeVariationSelectionModeUnspecified AgentListParamsVariationSelectionMode = "VARIATION_SELECTION_MODE_UNSPECIFIED"
	AgentListParamsVariationSelectionModeVariationSelectionModeRandom      AgentListParamsVariationSelectionMode = "VARIATION_SELECTION_MODE_RANDOM"
	AgentListParamsVariationSelectionModeVariationSelectionModeWeighted    AgentListParamsVariationSelectionMode = "VARIATION_SELECTION_MODE_WEIGHTED"
)

type AgentDeleteParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type AgentArchiveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r AgentArchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentArchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentArchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentPublishParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r AgentPublishParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentPublishParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentPublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentUnarchiveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r AgentUnarchiveParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentUnarchiveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentUnarchiveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentUnpublishParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r AgentUnpublishParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentUnpublishParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentUnpublishParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
