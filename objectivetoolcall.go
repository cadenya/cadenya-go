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

// ObjectiveToolCallService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveToolCallService] method instead.
type ObjectiveToolCallService struct {
	Options []option.RequestOption
}

// NewObjectiveToolCallService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectiveToolCallService(opts ...option.RequestOption) (r *ObjectiveToolCallService) {
	r = &ObjectiveToolCallService{}
	r.Options = opts
	return
}

// Retrieves a single tool call, including the content the tool returned. Media
// content (images, audio) is served as short-lived signed URLs.
func (r *ObjectiveToolCallService) Get(ctx context.Context, workspaceID string, objectiveID string, toolCallID string, opts ...option.RequestOption) (res *ObjectiveToolCallWithResult, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	if toolCallID == "" {
		err = errors.New("missing required toolCallId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls/%s", workspaceID, objectiveID, toolCallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all tool calls for an objective
func (r *ObjectiveToolCallService) List(ctx context.Context, workspaceID string, objectiveID string, query ObjectiveToolCallListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveToolCall], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls", workspaceID, objectiveID)
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

// Lists all tool calls for an objective
func (r *ObjectiveToolCallService) ListAutoPaging(ctx context.Context, workspaceID string, objectiveID string, query ObjectiveToolCallListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveToolCall] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, objectiveID, query, opts...))
}

// When an agent attempts to use a tool that requires approval, use this endpoint
// to mark it as approved.
func (r *ObjectiveToolCallService) Approve(ctx context.Context, workspaceID string, objectiveID string, toolCallID string, body ObjectiveToolCallApproveParams, opts ...option.RequestOption) (res *ObjectiveToolCall, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	if toolCallID == "" {
		err = errors.New("missing required toolCallId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls/%s:approve", workspaceID, objectiveID, toolCallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// When an agent attempts to use a tool that requires approval, use this endpoint
// to mark it as denied. Use a memo to steer the LLM to a different decision or
// usage of the tool.
func (r *ObjectiveToolCallService) Deny(ctx context.Context, workspaceID string, objectiveID string, toolCallID string, body ObjectiveToolCallDenyParams, opts ...option.RequestOption) (res *ObjectiveToolCall, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	if toolCallID == "" {
		err = errors.New("missing required toolCallId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls/%s:deny", workspaceID, objectiveID, toolCallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// ObjectiveToolCall is a record of a tool call made during an objective's
// execution. Tool calls are mutable — their status changes as they are approved,
// denied, or executed.
type ObjectiveToolCall struct {
	Data            ObjectiveToolCallData            `json:"data" api:"required"`
	ExecutionStatus ObjectiveToolCallExecutionStatus `json:"executionStatus" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// Current status of the tool call
	Status ObjectiveToolCallStatus `json:"status" api:"required"`
	Info   ObjectiveToolCallInfo   `json:"info"`
	JSON   objectiveToolCallJSON   `json:"-"`
}

// objectiveToolCallJSON contains the JSON metadata for the struct
// [ObjectiveToolCall]
type objectiveToolCallJSON struct {
	Data            apijson.Field
	ExecutionStatus apijson.Field
	Metadata        apijson.Field
	Status          apijson.Field
	Info            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObjectiveToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallJSON) RawJSON() string {
	return r.raw
}

type ObjectiveToolCallExecutionStatus string

const (
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusUnspecified ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusPending     ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_PENDING"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusRunning     ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_RUNNING"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusCompleted   ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_COMPLETED"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusErrored     ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_ERRORED"
)

func (r ObjectiveToolCallExecutionStatus) IsKnown() bool {
	switch r {
	case ObjectiveToolCallExecutionStatusToolCallExecutionStatusUnspecified, ObjectiveToolCallExecutionStatusToolCallExecutionStatusPending, ObjectiveToolCallExecutionStatusToolCallExecutionStatusRunning, ObjectiveToolCallExecutionStatusToolCallExecutionStatusCompleted, ObjectiveToolCallExecutionStatusToolCallExecutionStatusErrored:
		return true
	}
	return false
}

// Current status of the tool call
type ObjectiveToolCallStatus string

const (
	ObjectiveToolCallStatusToolCallStatusUnspecified        ObjectiveToolCallStatus = "TOOL_CALL_STATUS_UNSPECIFIED"
	ObjectiveToolCallStatusToolCallStatusAutoApproved       ObjectiveToolCallStatus = "TOOL_CALL_STATUS_AUTO_APPROVED"
	ObjectiveToolCallStatusToolCallStatusWaitingForApproval ObjectiveToolCallStatus = "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL"
	ObjectiveToolCallStatusToolCallStatusApproved           ObjectiveToolCallStatus = "TOOL_CALL_STATUS_APPROVED"
	ObjectiveToolCallStatusToolCallStatusDenied             ObjectiveToolCallStatus = "TOOL_CALL_STATUS_DENIED"
)

func (r ObjectiveToolCallStatus) IsKnown() bool {
	switch r {
	case ObjectiveToolCallStatusToolCallStatusUnspecified, ObjectiveToolCallStatusToolCallStatusAutoApproved, ObjectiveToolCallStatusToolCallStatusWaitingForApproval, ObjectiveToolCallStatusToolCallStatusApproved, ObjectiveToolCallStatusToolCallStatusDenied:
		return true
	}
	return false
}

type ObjectiveToolCallData struct {
	// CallableTool is a union that represents a tool that can be called by an agent.
	// In Cadenya, a tool that is used within an agent objective might be a
	// user-defined tool (IE: MCP, HTTP), another Agent (useful to separate context),
	// or a Cadenya Tool (one Cadenya provides).
	Callable CallableTool `json:"callable" api:"required"`
	// The arguments passed to the tool
	Arguments map[string]interface{} `json:"arguments"`
	// A memo supplied by the reviewer when denying the tool call
	Memo string `json:"memo"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	StatusChangedBy Profile                   `json:"statusChangedBy"`
	JSON            objectiveToolCallDataJSON `json:"-"`
}

// objectiveToolCallDataJSON contains the JSON metadata for the struct
// [ObjectiveToolCallData]
type objectiveToolCallDataJSON struct {
	Callable        apijson.Field
	Arguments       apijson.Field
	Memo            apijson.Field
	StatusChangedBy apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObjectiveToolCallData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallDataJSON) RawJSON() string {
	return r.raw
}

type ObjectiveToolCallInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective shared.OperationMetadata  `json:"objective"`
	JSON      objectiveToolCallInfoJSON `json:"-"`
}

// objectiveToolCallInfoJSON contains the JSON metadata for the struct
// [ObjectiveToolCallInfo]
type objectiveToolCallInfoJSON struct {
	CreatedBy   apijson.Field
	Objective   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveToolCallInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallInfoJSON) RawJSON() string {
	return r.raw
}

// ObjectiveToolCallResult is the content a tool returned after execution. Tools
// can return multiple content blocks, and blocks can be multi-modal (text, image,
// audio). Media blocks are stored by Cadenya and served as short-lived signed URLs
// rather than inline bytes.
type ObjectiveToolCallResult struct {
	Content []ObjectiveToolCallResultContentBlock `json:"content" api:"required"`
	JSON    objectiveToolCallResultJSON           `json:"-"`
}

// objectiveToolCallResultJSON contains the JSON metadata for the struct
// [ObjectiveToolCallResult]
type objectiveToolCallResultJSON struct {
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveToolCallResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallResultJSON) RawJSON() string {
	return r.raw
}

type ObjectiveToolCallResultAudioBlock struct {
	// When the signed URL expires.
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// IANA media type of the stored audio, e.g. audio/wav.
	MimeType string `json:"mimeType" api:"required"`
	// Size of the stored audio in bytes.
	SizeBytes string `json:"sizeBytes" api:"required"`
	// Short-lived signed URL to download the stored audio.
	URL  string                                `json:"url" api:"required"`
	JSON objectiveToolCallResultAudioBlockJSON `json:"-"`
}

// objectiveToolCallResultAudioBlockJSON contains the JSON metadata for the struct
// [ObjectiveToolCallResultAudioBlock]
type objectiveToolCallResultAudioBlockJSON struct {
	ExpiresAt   apijson.Field
	MimeType    apijson.Field
	SizeBytes   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveToolCallResultAudioBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallResultAudioBlockJSON) RawJSON() string {
	return r.raw
}

// ContentBlock is a single block of tool result content. Exactly one of the
// variants is set.
type ObjectiveToolCallResultContentBlock struct {
	Audio ObjectiveToolCallResultAudioBlock       `json:"audio"`
	Image ObjectiveToolCallResultImageBlock       `json:"image"`
	Text  ObjectiveToolCallResultTextBlock        `json:"text"`
	JSON  objectiveToolCallResultContentBlockJSON `json:"-"`
}

// objectiveToolCallResultContentBlockJSON contains the JSON metadata for the
// struct [ObjectiveToolCallResultContentBlock]
type objectiveToolCallResultContentBlockJSON struct {
	Audio       apijson.Field
	Image       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveToolCallResultContentBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallResultContentBlockJSON) RawJSON() string {
	return r.raw
}

type ObjectiveToolCallResultImageBlock struct {
	// When the signed URL expires.
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// IANA media type of the stored image, e.g. image/png.
	MimeType string `json:"mimeType" api:"required"`
	// Size of the stored image in bytes.
	SizeBytes string `json:"sizeBytes" api:"required"`
	// Short-lived signed URL to download the stored image.
	URL  string                                `json:"url" api:"required"`
	JSON objectiveToolCallResultImageBlockJSON `json:"-"`
}

// objectiveToolCallResultImageBlockJSON contains the JSON metadata for the struct
// [ObjectiveToolCallResultImageBlock]
type objectiveToolCallResultImageBlockJSON struct {
	ExpiresAt   apijson.Field
	MimeType    apijson.Field
	SizeBytes   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveToolCallResultImageBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallResultImageBlockJSON) RawJSON() string {
	return r.raw
}

type ObjectiveToolCallResultTextBlock struct {
	Text string                               `json:"text" api:"required"`
	JSON objectiveToolCallResultTextBlockJSON `json:"-"`
}

// objectiveToolCallResultTextBlockJSON contains the JSON metadata for the struct
// [ObjectiveToolCallResultTextBlock]
type objectiveToolCallResultTextBlockJSON struct {
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveToolCallResultTextBlock) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallResultTextBlockJSON) RawJSON() string {
	return r.raw
}

// ObjectiveToolCallWithResult is an ObjectiveToolCall plus the content the tool
// returned. Returned by GetObjectiveToolCall.
type ObjectiveToolCallWithResult struct {
	Data            ObjectiveToolCallData                      `json:"data" api:"required"`
	ExecutionStatus ObjectiveToolCallWithResultExecutionStatus `json:"executionStatus" api:"required"`
	Info            ObjectiveToolCallInfo                      `json:"info" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// Current status of the tool call
	Status ObjectiveToolCallWithResultStatus `json:"status" api:"required"`
	// ObjectiveToolCallResult is the content a tool returned after execution. Tools
	// can return multiple content blocks, and blocks can be multi-modal (text, image,
	// audio). Media blocks are stored by Cadenya and served as short-lived signed URLs
	// rather than inline bytes.
	Result ObjectiveToolCallResult         `json:"result"`
	JSON   objectiveToolCallWithResultJSON `json:"-"`
}

// objectiveToolCallWithResultJSON contains the JSON metadata for the struct
// [ObjectiveToolCallWithResult]
type objectiveToolCallWithResultJSON struct {
	Data            apijson.Field
	ExecutionStatus apijson.Field
	Info            apijson.Field
	Metadata        apijson.Field
	Status          apijson.Field
	Result          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObjectiveToolCallWithResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveToolCallWithResultJSON) RawJSON() string {
	return r.raw
}

type ObjectiveToolCallWithResultExecutionStatus string

const (
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusUnspecified ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusPending     ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_PENDING"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusRunning     ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_RUNNING"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusCompleted   ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_COMPLETED"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusErrored     ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_ERRORED"
)

func (r ObjectiveToolCallWithResultExecutionStatus) IsKnown() bool {
	switch r {
	case ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusUnspecified, ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusPending, ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusRunning, ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusCompleted, ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusErrored:
		return true
	}
	return false
}

// Current status of the tool call
type ObjectiveToolCallWithResultStatus string

const (
	ObjectiveToolCallWithResultStatusToolCallStatusUnspecified        ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_UNSPECIFIED"
	ObjectiveToolCallWithResultStatusToolCallStatusAutoApproved       ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_AUTO_APPROVED"
	ObjectiveToolCallWithResultStatusToolCallStatusWaitingForApproval ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL"
	ObjectiveToolCallWithResultStatusToolCallStatusApproved           ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_APPROVED"
	ObjectiveToolCallWithResultStatusToolCallStatusDenied             ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_DENIED"
)

func (r ObjectiveToolCallWithResultStatus) IsKnown() bool {
	switch r {
	case ObjectiveToolCallWithResultStatusToolCallStatusUnspecified, ObjectiveToolCallWithResultStatusToolCallStatusAutoApproved, ObjectiveToolCallWithResultStatusToolCallStatusWaitingForApproval, ObjectiveToolCallWithResultStatusToolCallStatusApproved, ObjectiveToolCallWithResultStatusToolCallStatusDenied:
		return true
	}
	return false
}

type ObjectiveToolCallListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Filter by tool call status
	Status param.Field[ObjectiveToolCallListParamsStatus] `query:"status"`
}

// URLQuery serializes [ObjectiveToolCallListParams]'s query parameters as
// `url.Values`.
func (r ObjectiveToolCallListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by tool call status
type ObjectiveToolCallListParamsStatus string

const (
	ObjectiveToolCallListParamsStatusToolCallStatusUnspecified        ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_UNSPECIFIED"
	ObjectiveToolCallListParamsStatusToolCallStatusAutoApproved       ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_AUTO_APPROVED"
	ObjectiveToolCallListParamsStatusToolCallStatusWaitingForApproval ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL"
	ObjectiveToolCallListParamsStatusToolCallStatusApproved           ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_APPROVED"
	ObjectiveToolCallListParamsStatusToolCallStatusDenied             ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_DENIED"
)

func (r ObjectiveToolCallListParamsStatus) IsKnown() bool {
	switch r {
	case ObjectiveToolCallListParamsStatusToolCallStatusUnspecified, ObjectiveToolCallListParamsStatusToolCallStatusAutoApproved, ObjectiveToolCallListParamsStatusToolCallStatusWaitingForApproval, ObjectiveToolCallListParamsStatusToolCallStatusApproved, ObjectiveToolCallListParamsStatusToolCallStatusDenied:
		return true
	}
	return false
}

type ObjectiveToolCallApproveParams struct {
}

func (r ObjectiveToolCallApproveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveToolCallDenyParams struct {
	// A memo to associate to the tool call denial. Use a memo to steer the LLM to a
	// different decision or usage of the tool.
	Memo param.Field[string] `json:"memo"`
}

func (r ObjectiveToolCallDenyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
