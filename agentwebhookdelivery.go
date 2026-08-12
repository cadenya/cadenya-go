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

// Manage AI agents within a workspace. Agents define AI behavior and tool access.
//
// AgentWebhookDeliveryService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentWebhookDeliveryService] method instead.
type AgentWebhookDeliveryService struct {
	options []option.RequestOption
}

// NewAgentWebhookDeliveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentWebhookDeliveryService(opts ...option.RequestOption) (r AgentWebhookDeliveryService) {
	r = AgentWebhookDeliveryService{}
	r.options = opts
	return
}

// Lists all webhook deliveries for an agent
func (r *AgentWebhookDeliveryService) List(ctx context.Context, agentID string, params AgentWebhookDeliveryListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[WebhookDelivery], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/webhook_deliveries", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(agentID))
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

// Lists all webhook deliveries for an agent
func (r *AgentWebhookDeliveryService) ListAutoPaging(ctx context.Context, agentID string, params AgentWebhookDeliveryListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[WebhookDelivery] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, agentID, params, opts...))
}

type WebhookDelivery struct {
	// Webhook delivery details.
	Data WebhookDeliveryData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDelivery) RawJSON() string { return r.JSON.raw }
func (r *WebhookDelivery) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebhookDeliveryData struct {
	// Related resources
	AgentID      string `json:"agentId" api:"required"`
	AttemptCount int64  `json:"attemptCount" api:"required"`
	// The type of objective event that triggered this webhook delivery
	//
	// Any of "OBJECTIVE_EVENT_TYPE_UNSPECIFIED", "OBJECTIVE_EVENT_TYPE_USER_MESSAGE",
	// "OBJECTIVE_EVENT_TYPE_TOOL_APPROVAL_REQUESTED",
	// "OBJECTIVE_EVENT_TYPE_TOOL_APPROVED", "OBJECTIVE_EVENT_TYPE_TOOL_DENIED",
	// "OBJECTIVE_EVENT_TYPE_TOOL_CALLED", "OBJECTIVE_EVENT_TYPE_ERROR",
	// "OBJECTIVE_EVENT_TYPE_ASSISTANT_MESSAGE", "OBJECTIVE_EVENT_TYPE_TOOL_RESULT",
	// "OBJECTIVE_EVENT_TYPE_TOOL_ERROR",
	// "OBJECTIVE_EVENT_TYPE_CONTEXT_WINDOW_COMPACTED",
	// "OBJECTIVE_EVENT_TYPE_MEMORY_READ", "OBJECTIVE_EVENT_TYPE_CANCELLED",
	// "OBJECTIVE_EVENT_TYPE_SUB_AGENT_SPAWNED",
	// "OBJECTIVE_EVENT_TYPE_SUB_AGENT_UPDATED", "OBJECTIVE_EVENT_TYPE_FINALIZED",
	// "OBJECTIVE_EVENT_TYPE_NOTICE", "OBJECTIVE_EVENT_TYPE_TIMED_OUT",
	// "OBJECTIVE_EVENT_TYPE_REASONING".
	EventType WebhookDeliveryDataEventType `json:"eventType" api:"required"`
	// Response details. The response body is not retained.
	HTTPStatusCode   int64     `json:"httpStatusCode" api:"required"`
	LastAttemptAt    time.Time `json:"lastAttemptAt" api:"required" format:"date-time"`
	LatencyMs        int64     `json:"latencyMs" api:"required"`
	ObjectiveEventID string    `json:"objectiveEventId" api:"required"`
	ObjectiveID      string    `json:"objectiveId" api:"required"`
	// Content length of the response body in bytes
	ResponseContentLength string `json:"responseContentLength" api:"required"`
	// Any of "WEBHOOK_DELIVERY_STATUS_UNSPECIFIED", "WEBHOOK_DELIVERY_STATUS_PENDING",
	// "WEBHOOK_DELIVERY_STATUS_COMPLETED", "WEBHOOK_DELIVERY_STATUS_FAILED",
	// "WEBHOOK_DELIVERY_STATUS_DISABLED".
	Status    WebhookDeliveryDataStatus `json:"status" api:"required"`
	WebhookID string                    `json:"webhookId" api:"required"`
	// Webhook delivery details
	WebhookURL   string `json:"webhookUrl" api:"required"`
	ErrorMessage string `json:"errorMessage"`
	// Response headers received from the webhook endpoint
	ResponseHeaders map[string]string `json:"responseHeaders"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID               respjson.Field
		AttemptCount          respjson.Field
		EventType             respjson.Field
		HTTPStatusCode        respjson.Field
		LastAttemptAt         respjson.Field
		LatencyMs             respjson.Field
		ObjectiveEventID      respjson.Field
		ObjectiveID           respjson.Field
		ResponseContentLength respjson.Field
		Status                respjson.Field
		WebhookID             respjson.Field
		WebhookURL            respjson.Field
		ErrorMessage          respjson.Field
		ResponseHeaders       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookDeliveryData) RawJSON() string { return r.JSON.raw }
func (r *WebhookDeliveryData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of objective event that triggered this webhook delivery
type WebhookDeliveryDataEventType string

const (
	WebhookDeliveryDataEventTypeObjectiveEventTypeUnspecified            WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_UNSPECIFIED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeUserMessage            WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_USER_MESSAGE"
	WebhookDeliveryDataEventTypeObjectiveEventTypeToolApprovalRequested  WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TOOL_APPROVAL_REQUESTED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeToolApproved           WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TOOL_APPROVED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeToolDenied             WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TOOL_DENIED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeToolCalled             WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TOOL_CALLED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeError                  WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_ERROR"
	WebhookDeliveryDataEventTypeObjectiveEventTypeAssistantMessage       WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_ASSISTANT_MESSAGE"
	WebhookDeliveryDataEventTypeObjectiveEventTypeToolResult             WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TOOL_RESULT"
	WebhookDeliveryDataEventTypeObjectiveEventTypeToolError              WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TOOL_ERROR"
	WebhookDeliveryDataEventTypeObjectiveEventTypeContextWindowCompacted WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_CONTEXT_WINDOW_COMPACTED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeMemoryRead             WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_MEMORY_READ"
	WebhookDeliveryDataEventTypeObjectiveEventTypeCancelled              WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_CANCELLED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeSubAgentSpawned        WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_SUB_AGENT_SPAWNED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeSubAgentUpdated        WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_SUB_AGENT_UPDATED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeFinalized              WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_FINALIZED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeNotice                 WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_NOTICE"
	WebhookDeliveryDataEventTypeObjectiveEventTypeTimedOut               WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TIMED_OUT"
	WebhookDeliveryDataEventTypeObjectiveEventTypeReasoning              WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_REASONING"
)

type WebhookDeliveryDataStatus string

const (
	WebhookDeliveryDataStatusWebhookDeliveryStatusUnspecified WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_UNSPECIFIED"
	WebhookDeliveryDataStatusWebhookDeliveryStatusPending     WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_PENDING"
	WebhookDeliveryDataStatusWebhookDeliveryStatusCompleted   WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_COMPLETED"
	WebhookDeliveryDataStatusWebhookDeliveryStatusFailed      WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_FAILED"
	WebhookDeliveryDataStatusWebhookDeliveryStatusDisabled    WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_DISABLED"
)

type AgentWebhookDeliveryListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Optional filter by objective ID
	ObjectiveID param.Opt[string] `query:"objectiveId,omitzero" json:"-"`
	// Optional filter by event type
	//
	// Any of "OBJECTIVE_EVENT_TYPE_UNSPECIFIED", "OBJECTIVE_EVENT_TYPE_USER_MESSAGE",
	// "OBJECTIVE_EVENT_TYPE_TOOL_APPROVAL_REQUESTED",
	// "OBJECTIVE_EVENT_TYPE_TOOL_APPROVED", "OBJECTIVE_EVENT_TYPE_TOOL_DENIED",
	// "OBJECTIVE_EVENT_TYPE_TOOL_CALLED", "OBJECTIVE_EVENT_TYPE_ERROR",
	// "OBJECTIVE_EVENT_TYPE_ASSISTANT_MESSAGE", "OBJECTIVE_EVENT_TYPE_TOOL_RESULT",
	// "OBJECTIVE_EVENT_TYPE_TOOL_ERROR",
	// "OBJECTIVE_EVENT_TYPE_CONTEXT_WINDOW_COMPACTED",
	// "OBJECTIVE_EVENT_TYPE_MEMORY_READ", "OBJECTIVE_EVENT_TYPE_CANCELLED",
	// "OBJECTIVE_EVENT_TYPE_SUB_AGENT_SPAWNED",
	// "OBJECTIVE_EVENT_TYPE_SUB_AGENT_UPDATED", "OBJECTIVE_EVENT_TYPE_FINALIZED",
	// "OBJECTIVE_EVENT_TYPE_NOTICE", "OBJECTIVE_EVENT_TYPE_TIMED_OUT",
	// "OBJECTIVE_EVENT_TYPE_REASONING".
	EventType AgentWebhookDeliveryListParamsEventType `query:"eventType,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentWebhookDeliveryListParams]'s query parameters as
// `url.Values`.
func (r AgentWebhookDeliveryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional filter by event type
type AgentWebhookDeliveryListParamsEventType string

const (
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeUnspecified            AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_UNSPECIFIED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeUserMessage            AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_USER_MESSAGE"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolApprovalRequested  AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TOOL_APPROVAL_REQUESTED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolApproved           AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TOOL_APPROVED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolDenied             AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TOOL_DENIED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolCalled             AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TOOL_CALLED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeError                  AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_ERROR"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeAssistantMessage       AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_ASSISTANT_MESSAGE"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolResult             AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TOOL_RESULT"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolError              AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TOOL_ERROR"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeContextWindowCompacted AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_CONTEXT_WINDOW_COMPACTED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeMemoryRead             AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_MEMORY_READ"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeCancelled              AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_CANCELLED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeSubAgentSpawned        AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_SUB_AGENT_SPAWNED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeSubAgentUpdated        AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_SUB_AGENT_UPDATED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeFinalized              AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_FINALIZED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeNotice                 AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_NOTICE"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeTimedOut               AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TIMED_OUT"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeReasoning              AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_REASONING"
)
