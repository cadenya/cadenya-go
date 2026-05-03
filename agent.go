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

// AgentService manages AI agents at the WORKSPACE level. Agents are
// workspace-scoped resources that define AI behavior and tool access. All
// operations are implicitly scoped to the workspace determined by the JWT token.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// AgentService contains methods and other services that help with interacting with
// the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
	// AgentService manages AI agents at the WORKSPACE level. Agents are
	// workspace-scoped resources that define AI behavior and tool access. All
	// operations are implicitly scoped to the workspace determined by the JWT token.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	Feedback *AgentFeedbackService
	// AgentService manages AI agents at the WORKSPACE level. Agents are
	// workspace-scoped resources that define AI behavior and tool access. All
	// operations are implicitly scoped to the workspace determined by the JWT token.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	WebhookDeliveries *AgentWebhookDeliveryService
	Variations        *AgentVariationService
	// AgentScheduleService manages recurring schedules attached to agents. Schedules
	// trigger objectives on a cadence defined by AgentScheduleSpec.Schedule. All
	// operations are implicitly scoped to the workspace determined by the JWT token.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
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

// Agent resource
type Agent struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata shared.ResourceMetadata `json:"metadata" api:"required"`
	// Agent specification (user-provided configuration)
	Spec AgentSpec `json:"spec" api:"required"`
	// AgentInfo contains simple information about an agent for display or quick
	// reference
	Info AgentInfo `json:"info"`
	JSON agentJSON `json:"-"`
}

// agentJSON contains the JSON metadata for the struct [Agent]
type agentJSON struct {
	Metadata    apijson.Field
	Spec        apijson.Field
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

// Agent resource
type AgentParam struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Metadata param.Field[shared.ResourceMetadataParam] `json:"metadata" api:"required"`
	// Agent specification (user-provided configuration)
	Spec param.Field[AgentSpecParam] `json:"spec" api:"required"`
}

func (r AgentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// AgentInfo contains simple information about an agent for display or quick
// reference
type AgentInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
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

// AgentInfo contains simple information about an agent for display or quick
// reference
type AgentInfoParam struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy param.Field[ProfileParam] `json:"createdBy"`
}

func (r AgentInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Agent specification (user-provided configuration)
type AgentSpec struct {
	// Status of the agent
	Status AgentSpecStatus `json:"status" api:"required"`
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
	InputDataSchema interface{} `json:"inputDataSchema"`
	// The URL that Cadenya will send events for any objective assigned to the agent.
	WebhookEventsURL string        `json:"webhookEventsUrl"`
	JSON             agentSpecJSON `json:"-"`
}

// agentSpecJSON contains the JSON metadata for the struct [AgentSpec]
type agentSpecJSON struct {
	Status                 apijson.Field
	VariationSelectionMode apijson.Field
	Description            apijson.Field
	InputDataSchema        apijson.Field
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

// Status of the agent
type AgentSpecStatus string

const (
	AgentSpecStatusAgentStatusUnspecified AgentSpecStatus = "AGENT_STATUS_UNSPECIFIED"
	AgentSpecStatusAgentStatusDraft       AgentSpecStatus = "AGENT_STATUS_DRAFT"
	AgentSpecStatusAgentStatusPublished   AgentSpecStatus = "AGENT_STATUS_PUBLISHED"
	AgentSpecStatusAgentStatusArchived    AgentSpecStatus = "AGENT_STATUS_ARCHIVED"
)

func (r AgentSpecStatus) IsKnown() bool {
	switch r {
	case AgentSpecStatusAgentStatusUnspecified, AgentSpecStatusAgentStatusDraft, AgentSpecStatusAgentStatusPublished, AgentSpecStatusAgentStatusArchived:
		return true
	}
	return false
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
	// Status of the agent
	Status param.Field[AgentSpecStatus] `json:"status" api:"required"`
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
	InputDataSchema param.Field[interface{}] `json:"inputDataSchema"`
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
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter expression (query param: prefix)
	Prefix param.Field[string] `query:"prefix"`
	// Free-form search query
	Query param.Field[string] `query:"query"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by agent publication status
	Status param.Field[AgentListParamsStatus] `query:"status"`
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

// Filter by agent publication status
type AgentListParamsStatus string

const (
	AgentListParamsStatusAgentStatusUnspecified AgentListParamsStatus = "AGENT_STATUS_UNSPECIFIED"
	AgentListParamsStatusAgentStatusDraft       AgentListParamsStatus = "AGENT_STATUS_DRAFT"
	AgentListParamsStatusAgentStatusPublished   AgentListParamsStatus = "AGENT_STATUS_PUBLISHED"
	AgentListParamsStatusAgentStatusArchived    AgentListParamsStatus = "AGENT_STATUS_ARCHIVED"
)

func (r AgentListParamsStatus) IsKnown() bool {
	switch r {
	case AgentListParamsStatusAgentStatusUnspecified, AgentListParamsStatusAgentStatusDraft, AgentListParamsStatusAgentStatusPublished, AgentListParamsStatusAgentStatusArchived:
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
