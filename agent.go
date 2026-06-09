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

// Manage AI agents within a workspace. Agents define AI behavior and tool access.
//
// AgentService contains methods and other services that help with interacting with
// the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
	// Manage AI agents within a workspace. Agents define AI behavior and tool access.
	Feedback *AgentFeedbackService
	// Manage AI agents within a workspace. Agents define AI behavior and tool access.
	WebhookDeliveries *AgentWebhookDeliveryService
	// Manage variations of an agent and their tool, sub-agent, and memory layer
	// assignments.
	Variations *AgentVariationService
	// Manage recurring schedules attached to agents. Schedules trigger objectives on a
	// cadence defined by AgentScheduleSpec.Schedule.
	Schedules *AgentScheduleService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r *AgentService) {
	r = &AgentService{}
	r.Options = opts
	r.Feedback = NewAgentFeedbackService(opts...)
	r.WebhookDeliveries = NewAgentWebhookDeliveryService(opts...)
	r.Variations = NewAgentVariationService(opts...)
	r.Schedules = NewAgentScheduleService(opts...)
	return
}

// Creates a new agent in the workspace
func (r *AgentService) New(ctx context.Context, workspaceID string, body AgentNewParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves an agent by ID from the workspace
func (r *AgentService) Get(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an agent in the workspace
func (r *AgentService) Update(ctx context.Context, workspaceID string, id string, body AgentUpdateParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Lists all agents in the workspace
func (r *AgentService) List(ctx context.Context, workspaceID string, query AgentListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Agent], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents", workspaceID)
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

// Lists all agents in the workspace
func (r *AgentService) ListAutoPaging(ctx context.Context, workspaceID string, query AgentListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Agent] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Deletes an agent from the workspace
func (r *AgentService) Delete(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Transitions an agent to STATE_ARCHIVED. Archived agents are hidden from list
// results and cannot be used for objectives; active schedules are paused.
func (r *AgentService) Archive(ctx context.Context, workspaceID string, id string, body AgentArchiveParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s:archive", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions an agent to STATE_PUBLISHED, making it available for objectives. The
// agent must have at least one variation.
func (r *AgentService) Publish(ctx context.Context, workspaceID string, id string, body AgentPublishParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s:publish", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions an archived agent back to STATE_DRAFT. Publish the agent again to
// make it available for objectives.
func (r *AgentService) Unarchive(ctx context.Context, workspaceID string, id string, body AgentUnarchiveParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s:unarchive", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Transitions a published agent back to STATE_DRAFT. Active schedules for the
// agent are paused until it is published again.
func (r *AgentService) Unpublish(ctx context.Context, workspaceID string, id string, body AgentUnpublishParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s:unpublish", workspaceID, id)
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
	State AgentState `json:"state" api:"required"`
	// AgentInfo contains simple information about an agent for display or quick
	// reference
	Info AgentInfo `json:"info"`
	JSON agentJSON `json:"-"`
}

// agentJSON contains the JSON metadata for the struct [Agent]
type agentJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
	State       apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Agent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentJSON) RawJSON() string {
	return r.raw
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

func (r AgentState) IsKnown() bool {
	switch r {
	case AgentStateStateUnspecified, AgentStateStateDraft, AgentStateStatePublished, AgentStateStateArchived:
		return true
	}
	return false
}

// AgentInfo contains simple information about an agent for display or quick
// reference
type AgentInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy      Profile       `json:"createdBy"`
	VariationCount int64         `json:"variationCount"`
	JSON           agentInfoJSON `json:"-"`
}

// agentInfoJSON contains the JSON metadata for the struct [AgentInfo]
type agentInfoJSON struct {
	CreatedBy      apijson.Field
	VariationCount apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AgentInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentInfoJSON) RawJSON() string {
	return r.raw
}

// Agent specification (user-provided configuration)
type AgentSpec struct {
	// Controls how variations are automatically selected when creating objectives
	// Defaults to RANDOM when unspecified
	VariationSelectionMode AgentSpecVariationSelectionMode `json:"variationSelectionMode" api:"required"`
	// Description of the agent's purpose
	Description string `json:"description"`
	// InputDataSchema is used for enforcing a data input when objectives are created.
	// This is valuable when using liquid formatting in agent variation prompts. Input
	// data schema is also valuable when using an agent as a sub-agent, as the schema
	// is used as the tool's input parameter schema. If omitted, the sub-agent schema
	// will be loaded with a simple "prompt" free text string as its schema.
	InputDataSchema map[string]interface{} `json:"inputDataSchema"`
	// Optional output definition for objectives created for this agent. When provided,
	// Cadenya will append a tool to that will be called by the LLM in use by the
	// variant to extract information in the format provided here. Use this option when
	// you want structured data to be created by your objectives.
	OutputDefinition map[string]interface{} `json:"outputDefinition"`
	// The URL that Cadenya will send events for any objective assigned to the agent.
	WebhookEventsURL string        `json:"webhookEventsUrl"`
	JSON             agentSpecJSON `json:"-"`
}

// agentSpecJSON contains the JSON metadata for the struct [AgentSpec]
type agentSpecJSON struct {
	VariationSelectionMode apijson.Field
	Description            apijson.Field
	InputDataSchema        apijson.Field
	OutputDefinition       apijson.Field
	WebhookEventsURL       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AgentSpec) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentSpecJSON) RawJSON() string {
	return r.raw
}

// Controls how variations are automatically selected when creating objectives
// Defaults to RANDOM when unspecified
type AgentSpecVariationSelectionMode string

const (
	AgentSpecVariationSelectionModeVariationSelectionModeUnspecified AgentSpecVariationSelectionMode = "VARIATION_SELECTION_MODE_UNSPECIFIED"
	AgentSpecVariationSelectionModeVariationSelectionModeRandom      AgentSpecVariationSelectionMode = "VARIATION_SELECTION_MODE_RANDOM"
	AgentSpecVariationSelectionModeVariationSelectionModeWeighted    AgentSpecVariationSelectionMode = "VARIATION_SELECTION_MODE_WEIGHTED"
)

func (r AgentSpecVariationSelectionMode) IsKnown() bool {
	switch r {
	case AgentSpecVariationSelectionModeVariationSelectionModeUnspecified, AgentSpecVariationSelectionModeVariationSelectionModeRandom, AgentSpecVariationSelectionModeVariationSelectionModeWeighted:
		return true
	}
	return false
}

// Agent specification (user-provided configuration)
type AgentSpecParam struct {
	// Controls how variations are automatically selected when creating objectives
	// Defaults to RANDOM when unspecified
	VariationSelectionMode param.Field[AgentSpecVariationSelectionMode] `json:"variationSelectionMode" api:"required"`
	// Description of the agent's purpose
	Description param.Field[string] `json:"description"`
	// InputDataSchema is used for enforcing a data input when objectives are created.
	// This is valuable when using liquid formatting in agent variation prompts. Input
	// data schema is also valuable when using an agent as a sub-agent, as the schema
	// is used as the tool's input parameter schema. If omitted, the sub-agent schema
	// will be loaded with a simple "prompt" free text string as its schema.
	InputDataSchema param.Field[map[string]interface{}] `json:"inputDataSchema"`
	// Optional output definition for objectives created for this agent. When provided,
	// Cadenya will append a tool to that will be called by the LLM in use by the
	// variant to extract information in the format provided here. Use this option when
	// you want structured data to be created by your objectives.
	OutputDefinition param.Field[map[string]interface{}] `json:"outputDefinition"`
	// The URL that Cadenya will send events for any objective assigned to the agent.
	WebhookEventsURL param.Field[string] `json:"webhookEventsUrl"`
}

func (r AgentSpecParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Page struct {
	NextCursor string   `json:"nextCursor"`
	Total      int64    `json:"total"`
	JSON       pageJSON `json:"-"`
}

// pageJSON contains the JSON metadata for the struct [Page]
type pageJSON struct {
	NextCursor  apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Page) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageJSON) RawJSON() string {
	return r.raw
}

type AgentNewParams struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	// Agent specification (user-provided configuration)
	Spec param.Field[AgentSpecParam] `json:"spec" api:"required"`
	// Create agent variation request
	DefaultVariation param.Field[AgentNewParamsDefaultVariation] `json:"defaultVariation"`
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Create agent variation request
type AgentNewParamsDefaultVariation struct {
	// CreateResourceMetadata contains the user-provided fields for creating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateResourceMetadataParam] `json:"metadata" api:"required"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec param.Field[AgentVariationSpecParam] `json:"spec" api:"required"`
}

func (r AgentNewParamsDefaultVariation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentUpdateParams struct {
	// UpdateResourceMetadata contains the user-provided fields for updating a
	// workspace-scoped resource. Read-only fields (id, account_id, workspace_id,
	// profile_id, created_at) are excluded since they are set by the server.
	Metadata param.Field[shared.UpdateResourceMetadataParam] `json:"metadata"`
	// Agent specification (user-provided configuration)
	Spec param.Field[AgentSpecParam] `json:"spec"`
	// Fields to update
	UpdateMask param.Field[string] `json:"updateMask" format:"field-mask"`
}

func (r AgentUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentListParams struct {
	// Filter by bundle_key — return only resources owned by this bundle.
	BundleKey param.Field[string] `query:"bundleKey"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When true, the `info` field on each returned agent is populated. Requests with
	// this flag count more against your rate limit.
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by agent lifecycle state
	State param.Field[AgentListParamsState] `query:"state"`
	// Filter by variation selection mode
	VariationSelectionMode param.Field[AgentListParamsVariationSelectionMode] `query:"variationSelectionMode"`
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values) {
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

func (r AgentListParamsState) IsKnown() bool {
	switch r {
	case AgentListParamsStateStateUnspecified, AgentListParamsStateStateDraft, AgentListParamsStateStatePublished, AgentListParamsStateStateArchived:
		return true
	}
	return false
}

// Filter by variation selection mode
type AgentListParamsVariationSelectionMode string

const (
	AgentListParamsVariationSelectionModeVariationSelectionModeUnspecified AgentListParamsVariationSelectionMode = "VARIATION_SELECTION_MODE_UNSPECIFIED"
	AgentListParamsVariationSelectionModeVariationSelectionModeRandom      AgentListParamsVariationSelectionMode = "VARIATION_SELECTION_MODE_RANDOM"
	AgentListParamsVariationSelectionModeVariationSelectionModeWeighted    AgentListParamsVariationSelectionMode = "VARIATION_SELECTION_MODE_WEIGHTED"
)

func (r AgentListParamsVariationSelectionMode) IsKnown() bool {
	switch r {
	case AgentListParamsVariationSelectionModeVariationSelectionModeUnspecified, AgentListParamsVariationSelectionModeVariationSelectionModeRandom, AgentListParamsVariationSelectionModeVariationSelectionModeWeighted:
		return true
	}
	return false
}

type AgentArchiveParams struct {
}

func (r AgentArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentPublishParams struct {
}

func (r AgentPublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentUnarchiveParams struct {
}

func (r AgentUnarchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentUnpublishParams struct {
}

func (r AgentUnpublishParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
