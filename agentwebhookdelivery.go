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

// AgentService manages AI agents at the WORKSPACE level. Agents are
// workspace-scoped resources that define AI behavior and tool access. All
// operations are implicitly scoped to the workspace determined by the JWT token.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// AgentWebhookDeliveryService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentWebhookDeliveryService] method instead.
type AgentWebhookDeliveryService struct {
	Options []option.RequestOption
}

// NewAgentWebhookDeliveryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAgentWebhookDeliveryService(opts ...option.RequestOption) (r *AgentWebhookDeliveryService) {
	r = &AgentWebhookDeliveryService{}
	r.Options = opts
	return
}

// Lists all webhook deliveries for an agent
func (r *AgentWebhookDeliveryService) List(ctx context.Context, workspaceID string, agentID string, query AgentWebhookDeliveryListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[WebhookDelivery], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/agents/%s/webhook_deliveries", workspaceID, agentID)
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

// Lists all webhook deliveries for an agent
func (r *AgentWebhookDeliveryService) ListAutoPaging(ctx context.Context, workspaceID string, agentID string, query AgentWebhookDeliveryListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[WebhookDelivery] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, agentID, query, opts...))
}

type WebhookDelivery struct {
	// Webhook delivery data
	Data WebhookDeliveryData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	JSON     webhookDeliveryJSON      `json:"-"`
}

// webhookDeliveryJSON contains the JSON metadata for the struct [WebhookDelivery]
type webhookDeliveryJSON struct {
	Data        apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDelivery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeliveryJSON) RawJSON() string {
	return r.raw
}

type WebhookDeliveryData struct {
	// Related resources
	AgentID      string `json:"agentId" api:"required"`
	AttemptCount int64  `json:"attemptCount" api:"required"`
	// The type of objective event that triggered this webhook delivery
	EventType WebhookDeliveryDataEventType `json:"eventType" api:"required"`
	// Response details (no response_body to avoid storing large payloads)
	HTTPStatusCode   int64     `json:"httpStatusCode" api:"required"`
	LastAttemptAt    time.Time `json:"lastAttemptAt" api:"required" format:"date-time"`
	LatencyMs        int64     `json:"latencyMs" api:"required"`
	ObjectiveEventID string    `json:"objectiveEventId" api:"required"`
	ObjectiveID      string    `json:"objectiveId" api:"required"`
	// Content length of the response body in bytes
	ResponseContentLength string                    `json:"responseContentLength" api:"required"`
	Status                WebhookDeliveryDataStatus `json:"status" api:"required"`
	WebhookID             string                    `json:"webhookId" api:"required"`
	// Webhook delivery details
	WebhookURL   string `json:"webhookUrl" api:"required"`
	ErrorMessage string `json:"errorMessage"`
	// Response headers received from the webhook endpoint
	ResponseHeaders map[string]string       `json:"responseHeaders"`
	JSON            webhookDeliveryDataJSON `json:"-"`
}

// webhookDeliveryDataJSON contains the JSON metadata for the struct
// [WebhookDeliveryData]
type webhookDeliveryDataJSON struct {
	AgentID               apijson.Field
	AttemptCount          apijson.Field
	EventType             apijson.Field
	HTTPStatusCode        apijson.Field
	LastAttemptAt         apijson.Field
	LatencyMs             apijson.Field
	ObjectiveEventID      apijson.Field
	ObjectiveID           apijson.Field
	ResponseContentLength apijson.Field
	Status                apijson.Field
	WebhookID             apijson.Field
	WebhookURL            apijson.Field
	ErrorMessage          apijson.Field
	ResponseHeaders       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *WebhookDeliveryData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeliveryDataJSON) RawJSON() string {
	return r.raw
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
	WebhookDeliveryDataEventTypeObjectiveEventTypeSubObjectiveCreated    WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_SUB_OBJECTIVE_CREATED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeError                  WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_ERROR"
	WebhookDeliveryDataEventTypeObjectiveEventTypeAssistantMessage       WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_ASSISTANT_MESSAGE"
	WebhookDeliveryDataEventTypeObjectiveEventTypeToolResult             WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TOOL_RESULT"
	WebhookDeliveryDataEventTypeObjectiveEventTypeToolError              WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_TOOL_ERROR"
	WebhookDeliveryDataEventTypeObjectiveEventTypeContextWindowCompacted WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_CONTEXT_WINDOW_COMPACTED"
	WebhookDeliveryDataEventTypeObjectiveEventTypeMemoryRead             WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_MEMORY_READ"
	WebhookDeliveryDataEventTypeObjectiveEventTypeCancelled              WebhookDeliveryDataEventType = "OBJECTIVE_EVENT_TYPE_CANCELLED"
)

func (r WebhookDeliveryDataEventType) IsKnown() bool {
	switch r {
	case WebhookDeliveryDataEventTypeObjectiveEventTypeUnspecified, WebhookDeliveryDataEventTypeObjectiveEventTypeUserMessage, WebhookDeliveryDataEventTypeObjectiveEventTypeToolApprovalRequested, WebhookDeliveryDataEventTypeObjectiveEventTypeToolApproved, WebhookDeliveryDataEventTypeObjectiveEventTypeToolDenied, WebhookDeliveryDataEventTypeObjectiveEventTypeToolCalled, WebhookDeliveryDataEventTypeObjectiveEventTypeSubObjectiveCreated, WebhookDeliveryDataEventTypeObjectiveEventTypeError, WebhookDeliveryDataEventTypeObjectiveEventTypeAssistantMessage, WebhookDeliveryDataEventTypeObjectiveEventTypeToolResult, WebhookDeliveryDataEventTypeObjectiveEventTypeToolError, WebhookDeliveryDataEventTypeObjectiveEventTypeContextWindowCompacted, WebhookDeliveryDataEventTypeObjectiveEventTypeMemoryRead, WebhookDeliveryDataEventTypeObjectiveEventTypeCancelled:
		return true
	}
	return false
}

type WebhookDeliveryDataStatus string

const (
	WebhookDeliveryDataStatusWebhookDeliveryStatusUnspecified WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_UNSPECIFIED"
	WebhookDeliveryDataStatusWebhookDeliveryStatusPending     WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_PENDING"
	WebhookDeliveryDataStatusWebhookDeliveryStatusCompleted   WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_COMPLETED"
	WebhookDeliveryDataStatusWebhookDeliveryStatusFailed      WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_FAILED"
	WebhookDeliveryDataStatusWebhookDeliveryStatusDisabled    WebhookDeliveryDataStatus = "WEBHOOK_DELIVERY_STATUS_DISABLED"
)

func (r WebhookDeliveryDataStatus) IsKnown() bool {
	switch r {
	case WebhookDeliveryDataStatusWebhookDeliveryStatusUnspecified, WebhookDeliveryDataStatusWebhookDeliveryStatusPending, WebhookDeliveryDataStatusWebhookDeliveryStatusCompleted, WebhookDeliveryDataStatusWebhookDeliveryStatusFailed, WebhookDeliveryDataStatusWebhookDeliveryStatusDisabled:
		return true
	}
	return false
}

type AgentWebhookDeliveryListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Optional filter by event type
	EventType param.Field[AgentWebhookDeliveryListParamsEventType] `query:"eventType"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Optional filter by objective ID
	ObjectiveID param.Field[string] `query:"objectiveId"`
}

// URLQuery serializes [AgentWebhookDeliveryListParams]'s query parameters as
// `url.Values`.
func (r AgentWebhookDeliveryListParams) URLQuery() (v url.Values) {
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
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeSubObjectiveCreated    AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_SUB_OBJECTIVE_CREATED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeError                  AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_ERROR"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeAssistantMessage       AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_ASSISTANT_MESSAGE"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolResult             AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TOOL_RESULT"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolError              AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_TOOL_ERROR"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeContextWindowCompacted AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_CONTEXT_WINDOW_COMPACTED"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeMemoryRead             AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_MEMORY_READ"
	AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeCancelled              AgentWebhookDeliveryListParamsEventType = "OBJECTIVE_EVENT_TYPE_CANCELLED"
)

func (r AgentWebhookDeliveryListParamsEventType) IsKnown() bool {
	switch r {
	case AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeUnspecified, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeUserMessage, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolApprovalRequested, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolApproved, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolDenied, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolCalled, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeSubObjectiveCreated, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeError, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeAssistantMessage, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolResult, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeToolError, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeContextWindowCompacted, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeMemoryRead, AgentWebhookDeliveryListParamsEventTypeObjectiveEventTypeCancelled:
		return true
	}
	return false
}
