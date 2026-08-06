// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// ObjectiveToolCallService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveToolCallService] method instead.
type ObjectiveToolCallService struct {
	options []option.RequestOption
}

// NewObjectiveToolCallService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObjectiveToolCallService(opts ...option.RequestOption) (r ObjectiveToolCallService) {
	r = ObjectiveToolCallService{}
	r.options = opts
	return
}

// Retrieves a single tool call, including the content the tool returned. Media
// content (images, audio) is served as short-lived signed URLs.
func (r *ObjectiveToolCallService) Get(ctx context.Context, objectiveID string, toolCallID string, query ObjectiveToolCallGetParams, opts ...option.RequestOption) (res *ObjectiveToolCallWithResult, err error) {
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
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	if toolCallID == "" {
		err = errors.New("missing required toolCallId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(objectiveID), url.PathEscape(toolCallID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all tool calls for an objective
func (r *ObjectiveToolCallService) List(ctx context.Context, objectiveID string, params ObjectiveToolCallListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveToolCall], err error) {
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
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
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

// Lists all tool calls for an objective
func (r *ObjectiveToolCallService) ListAutoPaging(ctx context.Context, objectiveID string, params ObjectiveToolCallListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveToolCall] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, objectiveID, params, opts...))
}

// When an agent attempts to use a tool that requires approval, use this endpoint
// to mark it as approved.
func (r *ObjectiveToolCallService) Approve(ctx context.Context, objectiveID string, toolCallID string, body ObjectiveToolCallApproveParams, opts ...option.RequestOption) (res *ObjectiveToolCall, err error) {
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
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	if toolCallID == "" {
		err = errors.New("missing required toolCallId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls/%s:approve", url.PathEscape(body.WorkspaceID.Value), url.PathEscape(objectiveID), url.PathEscape(toolCallID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// When an agent attempts to use a tool that requires approval, use this endpoint
// to mark it as denied. Use a memo to steer the LLM to a different decision or
// usage of the tool.
func (r *ObjectiveToolCallService) Deny(ctx context.Context, objectiveID string, toolCallID string, params ObjectiveToolCallDenyParams, opts ...option.RequestOption) (res *ObjectiveToolCall, err error) {
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
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	if toolCallID == "" {
		err = errors.New("missing required toolCallId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls/%s:deny", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID), url.PathEscape(toolCallID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// For bare tool calls (tool sets with no execution adapter), sets the content an
// external API consumer supplies for the call — used for human-in-the-loop tools
// and reverse harnesses that execute tools locally and report results back.
func (r *ObjectiveToolCallService) SetContent(ctx context.Context, objectiveID string, toolCallID string, params ObjectiveToolCallSetContentParams, opts ...option.RequestOption) (res *ObjectiveToolCall, err error) {
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
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	if toolCallID == "" {
		err = errors.New("missing required toolCallId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tool_calls/%s:setContent", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID), url.PathEscape(toolCallID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// ObjectiveToolCall is a record of a tool call made during an objective's
// execution. Tool calls are mutable — their status changes as they are approved,
// denied, or executed.
type ObjectiveToolCall struct {
	Data ObjectiveToolCallData `json:"data" api:"required"`
	// Any of "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED",
	// "TOOL_CALL_EXECUTION_STATUS_PENDING", "TOOL_CALL_EXECUTION_STATUS_RUNNING",
	// "TOOL_CALL_EXECUTION_STATUS_COMPLETED", "TOOL_CALL_EXECUTION_STATUS_ERRORED",
	// "TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT".
	ExecutionStatus ObjectiveToolCallExecutionStatus `json:"executionStatus" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// Current status of the tool call
	//
	// Any of "TOOL_CALL_STATUS_UNSPECIFIED", "TOOL_CALL_STATUS_AUTO_APPROVED",
	// "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL", "TOOL_CALL_STATUS_APPROVED",
	// "TOOL_CALL_STATUS_DENIED".
	Status ObjectiveToolCallStatus `json:"status" api:"required"`
	Info   ObjectiveToolCallInfo   `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		ExecutionStatus respjson.Field
		Metadata        respjson.Field
		Status          respjson.Field
		Info            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCall) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallExecutionStatus string

const (
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusUnspecified       ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusPending           ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_PENDING"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusRunning           ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_RUNNING"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusCompleted         ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_COMPLETED"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusErrored           ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_ERRORED"
	ObjectiveToolCallExecutionStatusToolCallExecutionStatusWaitingForContent ObjectiveToolCallExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT"
)

// Current status of the tool call
type ObjectiveToolCallStatus string

const (
	ObjectiveToolCallStatusToolCallStatusUnspecified        ObjectiveToolCallStatus = "TOOL_CALL_STATUS_UNSPECIFIED"
	ObjectiveToolCallStatusToolCallStatusAutoApproved       ObjectiveToolCallStatus = "TOOL_CALL_STATUS_AUTO_APPROVED"
	ObjectiveToolCallStatusToolCallStatusWaitingForApproval ObjectiveToolCallStatus = "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL"
	ObjectiveToolCallStatusToolCallStatusApproved           ObjectiveToolCallStatus = "TOOL_CALL_STATUS_APPROVED"
	ObjectiveToolCallStatusToolCallStatusDenied             ObjectiveToolCallStatus = "TOOL_CALL_STATUS_DENIED"
)

type ObjectiveToolCallData struct {
	// CallableTool is a union that represents a tool that can be called by an agent.
	// In Cadenya, a tool that is used within an agent objective might be a
	// user-defined tool (IE: MCP, HTTP), another Agent (useful to separate context),
	// or a Cadenya Tool (one Cadenya provides).
	Callable CallableTool `json:"callable" api:"required"`
	// The arguments passed to the tool
	Arguments map[string]any `json:"arguments"`
	// A memo supplied by the reviewer when denying the tool call
	Memo string `json:"memo"`
	// List of resolved secrets used by the tool call
	ResolvedSecrets []ResolvedSecret `json:"resolvedSecrets"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	StatusChangedBy Profile `json:"statusChangedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Callable        respjson.Field
		Arguments       respjson.Field
		Memo            respjson.Field
		ResolvedSecrets respjson.Field
		StatusChangedBy respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCallData) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCallData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective shared.OperationMetadata `json:"objective"`
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
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		Objective   respjson.Field
		Tool        respjson.Field
		ToolSet     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCallInfo) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCallInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveToolCallResult is the content a tool returned after execution. Tools
// can return multiple content blocks, and blocks can be multi-modal (text, image,
// audio). Media blocks are stored by Cadenya and served as short-lived signed URLs
// rather than inline bytes.
type ObjectiveToolCallResult struct {
	Content []ObjectiveToolCallResultContentBlock `json:"content" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCallResult) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCallResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallResultAudioBlock struct {
	// When the signed URL expires.
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// IANA media type of the stored audio, e.g. audio/wav.
	MimeType string `json:"mimeType" api:"required"`
	// Size of the stored audio in bytes.
	SizeBytes string `json:"sizeBytes" api:"required"`
	// Short-lived signed URL to download the stored audio.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCallResultAudioBlock) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCallResultAudioBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContentBlock is a single block of tool result content. Exactly one of the
// variants is set.
type ObjectiveToolCallResultContentBlock struct {
	Audio ObjectiveToolCallResultAudioBlock `json:"audio"`
	Image ObjectiveToolCallResultImageBlock `json:"image"`
	Text  ObjectiveToolCallResultTextBlock  `json:"text"`
	// The JSON name of the variant set in `block` (e.g. "text"). Filled by the server;
	// drives the discriminated union in the generated OpenAPI.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Audio       respjson.Field
		Image       respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCallResultContentBlock) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCallResultContentBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallResultImageBlock struct {
	// When the signed URL expires.
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// IANA media type of the stored image, e.g. image/png.
	MimeType string `json:"mimeType" api:"required"`
	// Size of the stored image in bytes.
	SizeBytes string `json:"sizeBytes" api:"required"`
	// Short-lived signed URL to download the stored image.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		MimeType    respjson.Field
		SizeBytes   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCallResultImageBlock) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCallResultImageBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallResultTextBlock struct {
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCallResultTextBlock) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCallResultTextBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveToolCallWithResult is an ObjectiveToolCall plus the content the tool
// returned. Returned by GetObjectiveToolCall.
type ObjectiveToolCallWithResult struct {
	Data ObjectiveToolCallData `json:"data" api:"required"`
	// Any of "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED",
	// "TOOL_CALL_EXECUTION_STATUS_PENDING", "TOOL_CALL_EXECUTION_STATUS_RUNNING",
	// "TOOL_CALL_EXECUTION_STATUS_COMPLETED", "TOOL_CALL_EXECUTION_STATUS_ERRORED",
	// "TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT".
	ExecutionStatus ObjectiveToolCallWithResultExecutionStatus `json:"executionStatus" api:"required"`
	Info            ObjectiveToolCallInfo                      `json:"info" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// Current status of the tool call
	//
	// Any of "TOOL_CALL_STATUS_UNSPECIFIED", "TOOL_CALL_STATUS_AUTO_APPROVED",
	// "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL", "TOOL_CALL_STATUS_APPROVED",
	// "TOOL_CALL_STATUS_DENIED".
	Status ObjectiveToolCallWithResultStatus `json:"status" api:"required"`
	// List of resolved secrets used by the tool call
	ResolvedSecrets []ResolvedSecret `json:"resolvedSecrets"`
	// ObjectiveToolCallResult is the content a tool returned after execution. Tools
	// can return multiple content blocks, and blocks can be multi-modal (text, image,
	// audio). Media blocks are stored by Cadenya and served as short-lived signed URLs
	// rather than inline bytes.
	Result ObjectiveToolCallResult `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		ExecutionStatus respjson.Field
		Info            respjson.Field
		Metadata        respjson.Field
		Status          respjson.Field
		ResolvedSecrets respjson.Field
		Result          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveToolCallWithResult) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveToolCallWithResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallWithResultExecutionStatus string

const (
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusUnspecified       ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusPending           ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_PENDING"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusRunning           ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_RUNNING"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusCompleted         ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_COMPLETED"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusErrored           ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_ERRORED"
	ObjectiveToolCallWithResultExecutionStatusToolCallExecutionStatusWaitingForContent ObjectiveToolCallWithResultExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT"
)

// Current status of the tool call
type ObjectiveToolCallWithResultStatus string

const (
	ObjectiveToolCallWithResultStatusToolCallStatusUnspecified        ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_UNSPECIFIED"
	ObjectiveToolCallWithResultStatusToolCallStatusAutoApproved       ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_AUTO_APPROVED"
	ObjectiveToolCallWithResultStatusToolCallStatusWaitingForApproval ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL"
	ObjectiveToolCallWithResultStatusToolCallStatusApproved           ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_APPROVED"
	ObjectiveToolCallWithResultStatusToolCallStatusDenied             ObjectiveToolCallWithResultStatus = "TOOL_CALL_STATUS_DENIED"
)

// ResolvedSecret is a resolved secret value from the workspace, toolset, or
// objective. When a tool is called, it will rely on secrets in the order of:
//
// - Objective
// - Toolset
// - Workspace
type ResolvedSecret struct {
	Key string `json:"key"`
	// Any of "RESOLVED_SECRET_SOURCE_UNSPECIFIED", "RESOLVED_SECRET_SOURCE_WORKSPACE",
	// "RESOLVED_SECRET_SOURCE_TOOLSET", "RESOLVED_SECRET_SOURCE_OBJECTIVE".
	Source ResolvedSecretSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResolvedSecret) RawJSON() string { return r.JSON.raw }
func (r *ResolvedSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResolvedSecretSource string

const (
	ResolvedSecretSourceResolvedSecretSourceUnspecified ResolvedSecretSource = "RESOLVED_SECRET_SOURCE_UNSPECIFIED"
	ResolvedSecretSourceResolvedSecretSourceWorkspace   ResolvedSecretSource = "RESOLVED_SECRET_SOURCE_WORKSPACE"
	ResolvedSecretSourceResolvedSecretSourceToolset     ResolvedSecretSource = "RESOLVED_SECRET_SOURCE_TOOLSET"
	ResolvedSecretSourceResolvedSecretSourceObjective   ResolvedSecretSource = "RESOLVED_SECRET_SOURCE_OBJECTIVE"
)

// The properties Data, MimeType are required.
type SetToolCallContentRequestAudioBlockParam struct {
	// Base64-encoded audio bytes.
	Data string `json:"data" api:"required" format:"bytes"`
	// IANA media type of the audio, e.g. audio/wav.
	MimeType string `json:"mimeType" api:"required"`
	paramObj
}

func (r SetToolCallContentRequestAudioBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow SetToolCallContentRequestAudioBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetToolCallContentRequestAudioBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContentBlock is a single block of tool call content supplied on input. Exactly
// one of the variants is set.
type SetToolCallContentRequestContentBlockParam struct {
	// The JSON name of the variant set in `block` (e.g. "text"). Required on input;
	// drives the discriminated union in the generated OpenAPI.
	Type  param.Opt[string]                        `json:"type,omitzero"`
	Audio SetToolCallContentRequestAudioBlockParam `json:"audio,omitzero"`
	Image SetToolCallContentRequestImageBlockParam `json:"image,omitzero"`
	Text  SetToolCallContentRequestTextBlockParam  `json:"text,omitzero"`
	paramObj
}

func (r SetToolCallContentRequestContentBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow SetToolCallContentRequestContentBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetToolCallContentRequestContentBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Data, MimeType are required.
type SetToolCallContentRequestImageBlockParam struct {
	// Base64-encoded image bytes.
	Data string `json:"data" api:"required" format:"bytes"`
	// IANA media type of the image, e.g. image/png.
	MimeType string `json:"mimeType" api:"required"`
	paramObj
}

func (r SetToolCallContentRequestImageBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow SetToolCallContentRequestImageBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetToolCallContentRequestImageBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Text is required.
type SetToolCallContentRequestTextBlockParam struct {
	Text string `json:"text" api:"required"`
	paramObj
}

func (r SetToolCallContentRequestTextBlockParam) MarshalJSON() (data []byte, err error) {
	type shadow SetToolCallContentRequestTextBlockParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SetToolCallContentRequestTextBlockParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ObjectiveToolCallListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Opt[bool] `query:"includeInfo,omitzero" json:"-"`
	// Filters by metadata labels. Comma-separated key=value pairs, e.g.
	// "env=prod,team=ai". A resource matches only if every pair matches exactly (AND
	// semantics).
	Labels param.Opt[string] `query:"labels,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by tool call execution status. Useful for reverse-harness polling of bare
	// tool calls waiting for externally supplied content
	// (TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT).
	//
	// Any of "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED",
	// "TOOL_CALL_EXECUTION_STATUS_PENDING", "TOOL_CALL_EXECUTION_STATUS_RUNNING",
	// "TOOL_CALL_EXECUTION_STATUS_COMPLETED", "TOOL_CALL_EXECUTION_STATUS_ERRORED",
	// "TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT".
	ExecutionStatus ObjectiveToolCallListParamsExecutionStatus `query:"executionStatus,omitzero" json:"-"`
	// Filter by tool call status
	//
	// Any of "TOOL_CALL_STATUS_UNSPECIFIED", "TOOL_CALL_STATUS_AUTO_APPROVED",
	// "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL", "TOOL_CALL_STATUS_APPROVED",
	// "TOOL_CALL_STATUS_DENIED".
	Status ObjectiveToolCallListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectiveToolCallListParams]'s query parameters as
// `url.Values`.
func (r ObjectiveToolCallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by tool call execution status. Useful for reverse-harness polling of bare
// tool calls waiting for externally supplied content
// (TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT).
type ObjectiveToolCallListParamsExecutionStatus string

const (
	ObjectiveToolCallListParamsExecutionStatusToolCallExecutionStatusUnspecified       ObjectiveToolCallListParamsExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED"
	ObjectiveToolCallListParamsExecutionStatusToolCallExecutionStatusPending           ObjectiveToolCallListParamsExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_PENDING"
	ObjectiveToolCallListParamsExecutionStatusToolCallExecutionStatusRunning           ObjectiveToolCallListParamsExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_RUNNING"
	ObjectiveToolCallListParamsExecutionStatusToolCallExecutionStatusCompleted         ObjectiveToolCallListParamsExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_COMPLETED"
	ObjectiveToolCallListParamsExecutionStatusToolCallExecutionStatusErrored           ObjectiveToolCallListParamsExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_ERRORED"
	ObjectiveToolCallListParamsExecutionStatusToolCallExecutionStatusWaitingForContent ObjectiveToolCallListParamsExecutionStatus = "TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT"
)

// Filter by tool call status
type ObjectiveToolCallListParamsStatus string

const (
	ObjectiveToolCallListParamsStatusToolCallStatusUnspecified        ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_UNSPECIFIED"
	ObjectiveToolCallListParamsStatusToolCallStatusAutoApproved       ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_AUTO_APPROVED"
	ObjectiveToolCallListParamsStatusToolCallStatusWaitingForApproval ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_WAITING_FOR_APPROVAL"
	ObjectiveToolCallListParamsStatusToolCallStatusApproved           ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_APPROVED"
	ObjectiveToolCallListParamsStatusToolCallStatusDenied             ObjectiveToolCallListParamsStatus = "TOOL_CALL_STATUS_DENIED"
)

type ObjectiveToolCallApproveParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

func (r ObjectiveToolCallApproveParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveToolCallApproveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveToolCallApproveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallDenyParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// A memo to associate to the tool call denial. Use a memo to steer the LLM to a
	// different decision or usage of the tool.
	Memo param.Opt[string] `json:"memo,omitzero"`
	paramObj
}

func (r ObjectiveToolCallDenyParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveToolCallDenyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveToolCallDenyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveToolCallSetContentParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// The content to set on the tool call. Mirrors
	// ObjectiveToolCallResult.ContentBlock but writable: media blocks carry raw data
	// on input where the result-side carries a signed url on output.
	Content []SetToolCallContentRequestContentBlockParam `json:"content,omitzero" api:"required"`
	paramObj
}

func (r ObjectiveToolCallSetContentParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveToolCallSetContentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveToolCallSetContentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
