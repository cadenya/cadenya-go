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
	"go.cadenya.com/cadenya-go/packages/ssestream"
	"go.cadenya.com/cadenya-go/shared"
	"net/http"
	"net/url"
	"slices"
	"time"
)

// ObjectiveService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveService] method instead.
type ObjectiveService struct {
	options   []option.RequestOption
	Tools     ObjectiveToolService
	ToolCalls ObjectiveToolCallService
	Tasks     ObjectiveTaskService
	Feedback  ObjectiveFeedbackService
}

// NewObjectiveService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectiveService(opts ...option.RequestOption) (r ObjectiveService) {
	r = ObjectiveService{}
	r.options = opts
	r.Tools = NewObjectiveToolService(opts...)
	r.ToolCalls = NewObjectiveToolCallService(opts...)
	r.Tasks = NewObjectiveTaskService(opts...)
	r.Feedback = NewObjectiveFeedbackService(opts...)
	return
}

// Creates a new objective in the workspace
func (r *ObjectiveService) New(ctx context.Context, params ObjectiveNewParams, opts ...option.RequestOption) (res *Objective, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives", url.PathEscape(params.WorkspaceID.Value))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieves an objective by ID from the workspace
func (r *ObjectiveService) Get(ctx context.Context, id string, query ObjectiveGetParams, opts ...option.RequestOption) (res *Objective, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all objectives in the workspace
func (r *ObjectiveService) List(ctx context.Context, params ObjectiveListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Objective], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives", url.PathEscape(params.WorkspaceID.Value))
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

// Lists all objectives in the workspace
func (r *ObjectiveService) ListAutoPaging(ctx context.Context, params ObjectiveListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Objective] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Cancels a running or pending objective. The objective's state will be set to
// STATE_CANCELLED.
func (r *ObjectiveService) Cancel(ctx context.Context, objectiveID string, params ObjectiveCancelParams, opts ...option.RequestOption) (res *Objective, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s:cancel", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Triggers compaction on a running objective. Optionally override the variation's
// compaction config.
func (r *ObjectiveService) Compact(ctx context.Context, objectiveID string, params ObjectiveCompactParams, opts ...option.RequestOption) (res *ObjectiveCompactResponse, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s:compact", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Continues an objective that has completed
func (r *ObjectiveService) Continue(ctx context.Context, objectiveID string, params ObjectiveContinueParams, opts ...option.RequestOption) (res *ObjectiveEvent, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s:continue", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Read-only list of the last five windows of execution for this objective, ordered
// by most recent first
func (r *ObjectiveService) ListContextWindows(ctx context.Context, objectiveID string, params ObjectiveListContextWindowsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveContextWindow], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/context_windows", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
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

// Read-only list of the last five windows of execution for this objective, ordered
// by most recent first
func (r *ObjectiveService) ListContextWindowsAutoPaging(ctx context.Context, objectiveID string, params ObjectiveListContextWindowsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveContextWindow] {
	return pagination.NewCursorPaginationAutoPager(r.ListContextWindows(ctx, objectiveID, params, opts...))
}

// Lists all events for an objective
func (r *ObjectiveService) ListEvents(ctx context.Context, objectiveID string, params ObjectiveListEventsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveEvent], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/events", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
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

// Lists all events for an objective
func (r *ObjectiveService) ListEventsAutoPaging(ctx context.Context, objectiveID string, params ObjectiveListEventsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveEvent] {
	return pagination.NewCursorPaginationAutoPager(r.ListEvents(ctx, objectiveID, params, opts...))
}

// Returns the context-usage breakdown measured for the objective's most recent
// iteration: character lengths per context component (system prompt, memory
// appendices, tool definitions, messages by role) alongside the iteration's input
// token counts.
func (r *ObjectiveService) GetDiagnostics(ctx context.Context, objectiveID string, query ObjectiveGetDiagnosticsParams, opts ...option.RequestOption) (res *ObjectiveGetDiagnosticsResponse, err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/diagnostics", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(objectiveID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Streams events for an objective in real-time using server-sent events (SSE)
func (r *ObjectiveService) StreamEventsStreaming(ctx context.Context, objectiveID string, query ObjectiveStreamEventsParams, opts ...option.RequestOption) (stream *ssestream.Stream[ObjectiveEvent]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/event-stream")}, opts...)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return ssestream.NewStream[ObjectiveEvent](nil, err)
	}
	requestconfig.UseDefaultParam(&query.WorkspaceID, precfg.WorkspaceID)
	if query.WorkspaceID.Value == "" {
		err = errors.New("missing required workspaceId parameter")
		return ssestream.NewStream[ObjectiveEvent](nil, err)
	}
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return ssestream.NewStream[ObjectiveEvent](nil, err)
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/events:stream", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(objectiveID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &raw, opts...)
	return ssestream.NewStream[ObjectiveEvent](ssestream.NewDecoder(raw), err)
}

type AssistantMessage struct {
	Content   string              `json:"content"`
	ToolCalls []AssistantToolCall `json:"toolCalls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantMessage) RawJSON() string { return r.JSON.raw }
func (r *AssistantMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssistantToolCall struct {
	Arguments    string `json:"arguments"`
	FunctionName string `json:"functionName"`
	// CallableTool is a union that represents a tool that can be called by an agent.
	// In Cadenya, a tool that is used within an agent objective might be a
	// user-defined tool (IE: MCP, HTTP), another Agent (useful to separate context),
	// or a Cadenya Tool (one Cadenya provides).
	Tool CallableToolUnion `json:"tool"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments    respjson.Field
		FunctionName respjson.Field
		Tool         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssistantToolCall) RawJSON() string { return r.JSON.raw }
func (r *AssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CallableToolUnion contains all possible properties and values from
// [CallableToolTool], [CallableToolAgent], [CallableToolCadenyaProvidedTool].
//
// Use the [CallableToolUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CallableToolUnion struct {
	// This field is from variant [CallableToolTool].
	Tool shared.ResourceMetadata `json:"tool"`
	// Any of "tool", "agent", "cadenyaProvidedTool".
	Type string `json:"type"`
	// This field is from variant [CallableToolAgent].
	Agent shared.ResourceMetadata `json:"agent"`
	// This field is from variant [CallableToolCadenyaProvidedTool].
	CadenyaProvidedTool shared.ResourceMetadata `json:"cadenyaProvidedTool"`
	JSON                struct {
		Tool                respjson.Field
		Type                respjson.Field
		Agent               respjson.Field
		CadenyaProvidedTool respjson.Field
		raw                 string
	} `json:"-"`
}

// anyCallableTool is implemented by each variant of [CallableToolUnion] to add
// type safety for the return type of [CallableToolUnion.AsAny]
type anyCallableTool interface {
	implCallableToolUnion()
}

func (CallableToolTool) implCallableToolUnion()                {}
func (CallableToolAgent) implCallableToolUnion()               {}
func (CallableToolCadenyaProvidedTool) implCallableToolUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := CallableToolUnion.AsAny().(type) {
//	case cadenya.CallableToolTool:
//	case cadenya.CallableToolAgent:
//	case cadenya.CallableToolCadenyaProvidedTool:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CallableToolUnion) AsAny() anyCallableTool {
	switch u.Type {
	case "tool":
		return u.AsTool()
	case "agent":
		return u.AsAgent()
	case "cadenyaProvidedTool":
		return u.AsCadenyaProvidedTool()
	}
	return nil
}

func (u CallableToolUnion) AsTool() (v CallableToolTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallableToolUnion) AsAgent() (v CallableToolAgent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CallableToolUnion) AsCadenyaProvidedTool() (v CallableToolCadenyaProvidedTool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CallableToolUnion) RawJSON() string { return u.JSON.raw }

func (r *CallableToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallableToolAgent struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent shared.ResourceMetadata `json:"agent" api:"required"`
	// Any of "agent".
	Type CallableToolAgentType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallableToolAgent) RawJSON() string { return r.JSON.raw }
func (r *CallableToolAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallableToolAgentType string

const (
	CallableToolAgentTypeAgent CallableToolAgentType = "agent"
)

type CallableToolCadenyaProvidedTool struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	CadenyaProvidedTool shared.ResourceMetadata `json:"cadenyaProvidedTool" api:"required"`
	// Any of "cadenyaProvidedTool".
	Type CallableToolCadenyaProvidedToolType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CadenyaProvidedTool respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallableToolCadenyaProvidedTool) RawJSON() string { return r.JSON.raw }
func (r *CallableToolCadenyaProvidedTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallableToolCadenyaProvidedToolType string

const (
	CallableToolCadenyaProvidedToolTypeCadenyaProvidedTool CallableToolCadenyaProvidedToolType = "cadenyaProvidedTool"
)

type CallableToolTool struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Tool shared.ResourceMetadata `json:"tool" api:"required"`
	// Any of "tool".
	Type CallableToolToolType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tool        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallableToolTool) RawJSON() string { return r.JSON.raw }
func (r *CallableToolTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallableToolToolType string

const (
	CallableToolToolTypeTool CallableToolToolType = "tool"
)

// ContextLengths is the measured character length of each distinct component of an
// iteration's assembled context window. Values are raw character lengths of the
// component as assembled into the request — token estimates are derived by the
// client against input_tokens (component share = component length / sum of all
// lengths).
//
// New components are added as new fields — wire-compatible; absent components read
// as 0.
type ContextLengths struct {
	// Character length of the chat history messages with the assistant role.
	AssistantMessages int64 `json:"assistantMessages" api:"required"`
	// Character length of the discoverable/available-tools appendix attached to the
	// system prompt.
	AvailableTools int64 `json:"availableTools" api:"required"`
	// Character length of the episodic memory appendix attached to the system prompt.
	EpisodicMemory int64 `json:"episodicMemory" api:"required"`
	// Character length of the skills memory appendix attached to the system prompt.
	SkillsMemory int64 `json:"skillsMemory" api:"required"`
	// Character length of the objective's base system prompt (rendered variation
	// template). Not tokens -- see the message comment.
	SystemPrompt int64 `json:"systemPrompt" api:"required"`
	// Character length of the serialized tool definitions sent with the completion
	// request (names, descriptions, and JSON-schema parameters).
	ToolDefinitions int64 `json:"toolDefinitions" api:"required"`
	// Character length of the tool results present in the chat history.
	ToolResults int64 `json:"toolResults" api:"required"`
	// Character length of the chat history messages with the user role.
	UserMessages int64 `json:"userMessages" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantMessages respjson.Field
		AvailableTools    respjson.Field
		EpisodicMemory    respjson.Field
		SkillsMemory      respjson.Field
		SystemPrompt      respjson.Field
		ToolDefinitions   respjson.Field
		ToolResults       respjson.Field
		UserMessages      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextLengths) RawJSON() string { return r.JSON.raw }
func (r *ContextLengths) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextWindowCompacted struct {
	// Number of messages that were compacted
	MessagesCompacted int64 `json:"messagesCompacted"`
	// The new context window created by this compaction
	NewContextWindow ObjectiveContextWindowData `json:"newContextWindow"`
	// The strategies that were applied during this compaction
	Strategies []string `json:"strategies"`
	// The summary generated by the summarization strategy, if used.
	Summary string `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MessagesCompacted respjson.Field
		NewContextWindow  respjson.Field
		Strategies        respjson.Field
		Summary           respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextWindowCompacted) RawJSON() string { return r.JSON.raw }
func (r *ContextWindowCompacted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryRead is emitted each time the agent resolves a key against the memory
// cascade and loads an entry. Lookups that miss (key not found in any layer) do
// not emit this event.
type MemoryRead struct {
	// The specific entry that was read.
	MemoryEntryID string `json:"memoryEntryId"`
	// The layer the entry resolved to. The top-most layer that contained the key —
	// other layers beneath it that also contained the key are shadowed and not
	// referenced here.
	MemoryLayerID string `json:"memoryLayerId"`
	// Human-readable description of the read, set by the runtime. For example: "Loaded
	// skill", "Resolved context key". Not machine-parsed; intended for UI display
	// alongside the other events in an objective's timeline.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoryEntryID respjson.Field
		MemoryLayerID respjson.Field
		Message       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryRead) RawJSON() string { return r.JSON.raw }
func (r *MemoryRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryReference identifies a memory layer or a specific entry within one, for
// composition into a memory cascade. Used on objectives (where entry pinning is
// permitted).
//
// memory*layer_id accepts both the canonical form (memlyr*…) and the external-id
// form (external_id:my-custom-id). The same applies to memory_entry_id when set.
type MemoryReference struct {
	MemoryLayerID string `json:"memoryLayerId" api:"required"`
	// When set, inserts only this entry from memory_layer_id into the cascade —
	// behaves as a single-entry layer (only this key resolves at this position). The
	// entry must belong to memory_layer_id; mismatches are rejected with
	// InvalidArgument.
	MemoryEntryID string `json:"memoryEntryId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoryLayerID respjson.Field
		MemoryEntryID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryReference) RawJSON() string { return r.JSON.raw }
func (r *MemoryReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MemoryReference to a MemoryReferenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MemoryReferenceParam.Overrides()
func (r MemoryReference) ToParam() MemoryReferenceParam {
	return param.Override[MemoryReferenceParam](json.RawMessage(r.RawJSON()))
}

// MemoryReference identifies a memory layer or a specific entry within one, for
// composition into a memory cascade. Used on objectives (where entry pinning is
// permitted).
//
// memory*layer_id accepts both the canonical form (memlyr*…) and the external-id
// form (external_id:my-custom-id). The same applies to memory_entry_id when set.
//
// The property MemoryLayerID is required.
type MemoryReferenceParam struct {
	MemoryLayerID string `json:"memoryLayerId" api:"required"`
	// When set, inserts only this entry from memory_layer_id into the cascade —
	// behaves as a single-entry layer (only this key resolves at this position). The
	// entry must belong to memory_layer_id; mismatches are rejected with
	// InvalidArgument.
	MemoryEntryID param.Opt[string] `json:"memoryEntryId,omitzero"`
	paramObj
}

func (r MemoryReferenceParam) MarshalJSON() (data []byte, err error) {
	type shadow MemoryReferenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryReferenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Objective is the data for an objective. It contains the snapshotted fields for
// the selected agent and variation. Secrets are returned only with their names,
// and the output definition is copied from the agent's configuration.
type Objective struct {
	// ObjectiveConfigSnapshot is the point-in-time snapshot of the agent, variation,
	// and (when applicable) schedule that an objective was started with.
	ConfigSnapshot ObjectiveConfigSnapshot `json:"configSnapshot" api:"required"`
	// The first user message in the LLM chat history, either provided explicitly at
	// creation or rendered from the variation's first_user_message_template.
	FirstUserMessage string `json:"firstUserMessage" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	// The current lifecycle state of the objective.
	//
	// Any of "STATE_UNSPECIFIED", "STATE_PENDING", "STATE_RUNNING", "STATE_WAITING",
	// "STATE_FAILED", "STATE_CANCELLED", "STATE_FINALIZED", "STATE_TIMED_OUT".
	State ObjectiveState `json:"state" api:"required"`
	// system_prompt is read-only, derived from the selected variation's prompt
	SystemPrompt string `json:"systemPrompt" api:"required"`
	// Episodic is used to configure the episodic memory for the objective
	EpisodicMemory ObjectiveEpisodicMemory `json:"episodicMemory"`
	// Arbitrary data rendered into the variation's first_user_message_template
	FirstUserMessageData map[string]any `json:"firstUserMessageData"`
	// ObjectiveInfo provides read-only aggregated statistics about an objective's
	// execution
	Info ObjectiveInfo `json:"info"`
	// Memory layers/entries layered over the baseline cascade inherited from the
	// selected variation — element-level rules over inherited styles, in CSS terms.
	//
	// Array order is resolution order: EARLIER elements are more specific and are
	// consulted first. Entries pinned via memory_entry_id behave as single-entry
	// layers at their position.
	//
	// System-managed layers (e.g., episodic) cannot be referenced here; they attach
	// themselves automatically based on the episodic key.
	//
	// Size cap: the TOTAL effective cascade (this field + the variation's memory layer
	// assignments) must not exceed 10 entries. A request that would produce a larger
	// cascade is rejected with InvalidArgument.
	MemoryCascade []MemoryReference `json:"memoryCascade"`
	// The output of the objective, populated when the objective completes. Will match
	// the schema of output_json_schema or output_json_inferred. This will only be set
	// if the state of the objective is set to STATE_FINALIZED
	Output map[string]any `json:"output"`
	// A parent objective means the objective was spawned off using a separate agent to
	// complete an objective
	ParentObjectiveID string `json:"parentObjectiveId"`
	// Parameters forced onto this objective's tool calls, as provided at creation. See
	// CreateObjectiveRequest.pinned_parameters for semantics.
	PinnedParameters map[string]string `json:"pinnedParameters"`
	// Secrets that can be used in the headers for tool calls using the secret
	// interpolation format.
	Secrets []ObjectiveSecret `json:"secrets"`
	// Optional human-readable detail about the current state (e.g. a failure reason).
	StateMessage string `json:"stateMessage"`
	// Arbitrary data rendered into the variation's system_prompt_template
	SystemPromptData map[string]any `json:"systemPromptData"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConfigSnapshot       respjson.Field
		FirstUserMessage     respjson.Field
		Metadata             respjson.Field
		State                respjson.Field
		SystemPrompt         respjson.Field
		EpisodicMemory       respjson.Field
		FirstUserMessageData respjson.Field
		Info                 respjson.Field
		MemoryCascade        respjson.Field
		Output               respjson.Field
		ParentObjectiveID    respjson.Field
		PinnedParameters     respjson.Field
		Secrets              respjson.Field
		StateMessage         respjson.Field
		SystemPromptData     respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Objective) RawJSON() string { return r.JSON.raw }
func (r *Objective) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current lifecycle state of the objective.
type ObjectiveState string

const (
	ObjectiveStateStateUnspecified ObjectiveState = "STATE_UNSPECIFIED"
	ObjectiveStateStatePending     ObjectiveState = "STATE_PENDING"
	ObjectiveStateStateRunning     ObjectiveState = "STATE_RUNNING"
	ObjectiveStateStateWaiting     ObjectiveState = "STATE_WAITING"
	ObjectiveStateStateFailed      ObjectiveState = "STATE_FAILED"
	ObjectiveStateStateCancelled   ObjectiveState = "STATE_CANCELLED"
	ObjectiveStateStateFinalized   ObjectiveState = "STATE_FINALIZED"
	ObjectiveStateStateTimedOut    ObjectiveState = "STATE_TIMED_OUT"
)

// Episodic is used to configure the episodic memory for the objective
type ObjectiveEpisodicMemory struct {
	// The caller-supplied episodic key. Objectives created with the same key (for the
	// same agent) share one episodic memory layer.
	Key string `json:"key" api:"required"`
	// The episodic memory layer resolved (created or reused) for this objective's key.
	// Populated by the system at objective creation.
	MemoryLayerID string `json:"memoryLayerId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key           respjson.Field
		MemoryLayerID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEpisodicMemory) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEpisodicMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveConfigSnapshot is the point-in-time snapshot of the agent, variation,
// and (when applicable) schedule that an objective was started with.
type ObjectiveConfigSnapshot struct {
	// Agent resource
	Agent Agent `json:"agent"`
	// AgentSchedule resource — a recurring trigger attached to an agent that creates
	// objectives on its cadence.
	AgentSchedule AgentSchedule `json:"agentSchedule"`
	// AgentVariation resource
	AgentVariation AgentVariation `json:"agentVariation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent          respjson.Field
		AgentSchedule  respjson.Field
		AgentVariation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveConfigSnapshot) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveConfigSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveContextWindow is a window of chat completions that is grouped together
// to prevent context-window overflows. Context windows also allow agents to
// compact their windows and carry on into a new one.
type ObjectiveContextWindow struct {
	Data ObjectiveContextWindowData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata   `json:"metadata" api:"required"`
	Info     ObjectiveContextWindowInfo `json:"info"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Metadata    respjson.Field
		Info        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveContextWindow) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveContextWindow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveContextWindowInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective shared.OperationMetadata `json:"objective"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		Objective   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveContextWindowInfo) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveContextWindowInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveContextWindowData struct {
	// A calculated value for how many completion tokens (output tokens) have been used
	// in this context window
	CompletionTokens int64 `json:"completionTokens"`
	// The objective's ID that this window belongs to
	ObjectiveID string `json:"objectiveId"`
	// The instructions for this window to continue from a previous window's chat
	// history.
	PreviousWindowContinueInstructions string `json:"previousWindowContinueInstructions"`
	// A calculated value for how many prompt tokens (input tokens) have been used in
	// this context window
	PromptTokens int64 `json:"promptTokens"`
	// sequence is a numeric representation of which context window this is. Sequences
	// are useful to perform a max(sequence) on in order to calculate how many context
	// windows an objective has.
	Sequence int64 `json:"sequence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens                   respjson.Field
		ObjectiveID                        respjson.Field
		PreviousWindowContinueInstructions respjson.Field
		PromptTokens                       respjson.Field
		Sequence                           respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveContextWindowData) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveContextWindowData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveDiagnostics is the context-usage breakdown measured for a single
// iteration at request-assembly time. It reports how much of the context window
// each component occupies so tool parameters, memory cascades, and prompts can be
// tuned against real token usage.
type ObjectiveDiagnostics struct {
	// The portion of input_tokens served from the provider's prompt cache. Lets
	// clients distinguish "big but cached" from "big and paid fresh every iteration".
	CachedInputTokens int64 `json:"cachedInputTokens" api:"required"`
	// ContextLengths is the measured character length of each distinct component of an
	// iteration's assembled context window. Values are raw character lengths of the
	// component as assembled into the request — token estimates are derived by the
	// client against input_tokens (component share = component length / sum of all
	// lengths).
	//
	// New components are added as new fields — wire-compatible; absent components read
	// as 0.
	ContextLengths ContextLengths `json:"contextLengths" api:"required"`
	// Input tokens reported by the LLM provider for the iteration's completion.
	InputTokens int64 `json:"inputTokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CachedInputTokens respjson.Field
		ContextLengths    respjson.Field
		InputTokens       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveDiagnostics) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveDiagnostics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveError) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEvent struct {
	Data ObjectiveEventDataUnion `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata        shared.OperationMetadata `json:"metadata" api:"required"`
	ContextWindowID string                   `json:"contextWindowId"`
	// Elapsed time of the work this event records, when it is known at write time
	// (e.g. assistant message generation, tool execution for result/error events).
	// Unset means the event is instantaneous or the duration is not measurable.
	// Serialized as a canonical duration string (e.g. "4.1s"). Always set together
	// with started_at.
	Duration string             `json:"duration"`
	Info     ObjectiveEventInfo `json:"info"`
	// When the work this event records began. Set together with duration, so the work
	// interval is [started_at, started_at + duration]. The event's created_at remains
	// the time the event was persisted.
	StartedAt time.Time `json:"startedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data            respjson.Field
		Metadata        respjson.Field
		ContextWindowID respjson.Field
		Duration        respjson.Field
		Info            respjson.Field
		StartedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEvent) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveEventDataUnion contains all possible properties and values from
// [ObjectiveEventDataUserMessage], [ObjectiveEventDataToolApprovalRequested],
// [ObjectiveEventDataToolApproved], [ObjectiveEventDataToolDenied],
// [ObjectiveEventDataToolCalled], [ObjectiveEventDataError],
// [ObjectiveEventDataAssistantMessage], [ObjectiveEventDataToolResult],
// [ObjectiveEventDataToolError], [ObjectiveEventDataContextWindowCompacted],
// [ObjectiveEventDataMemoryRead], [ObjectiveEventDataCancelled],
// [ObjectiveEventDataSubAgentSpawned], [ObjectiveEventDataSubAgentUpdated],
// [ObjectiveEventDataFinalized], [ObjectiveEventDataNotice],
// [ObjectiveEventDataTimedOut], [ObjectiveEventDataReasoning].
//
// Use the [ObjectiveEventDataUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ObjectiveEventDataUnion struct {
	// Any of "userMessage", "toolApprovalRequested", "toolApproved", "toolDenied",
	// "toolCalled", "error", "assistantMessage", "toolResult", "toolError",
	// "contextWindowCompacted", "memoryRead", "cancelled", "subAgentSpawned",
	// "subAgentUpdated", "finalized", "notice", "timedOut", "reasoning".
	Type string `json:"type"`
	// This field is from variant [ObjectiveEventDataUserMessage].
	UserMessage UserMessage `json:"userMessage"`
	// This field is from variant [ObjectiveEventDataToolApprovalRequested].
	ToolApprovalRequested ToolApprovalRequested `json:"toolApprovalRequested"`
	// This field is from variant [ObjectiveEventDataToolApproved].
	ToolApproved ToolApproved `json:"toolApproved"`
	// This field is from variant [ObjectiveEventDataToolDenied].
	ToolDenied ToolDenied `json:"toolDenied"`
	// This field is from variant [ObjectiveEventDataToolCalled].
	ToolCalled ToolCalled `json:"toolCalled"`
	// This field is from variant [ObjectiveEventDataError].
	Error ObjectiveError `json:"error"`
	// This field is from variant [ObjectiveEventDataAssistantMessage].
	AssistantMessage AssistantMessage `json:"assistantMessage"`
	// This field is from variant [ObjectiveEventDataToolResult].
	ToolResult ToolResult `json:"toolResult"`
	// This field is from variant [ObjectiveEventDataToolError].
	ToolError ToolError `json:"toolError"`
	// This field is from variant [ObjectiveEventDataContextWindowCompacted].
	ContextWindowCompacted ContextWindowCompacted `json:"contextWindowCompacted"`
	// This field is from variant [ObjectiveEventDataMemoryRead].
	MemoryRead MemoryRead `json:"memoryRead"`
	// This field is from variant [ObjectiveEventDataCancelled].
	Cancelled ObjectiveEventDataCancelledCancelled `json:"cancelled"`
	// This field is from variant [ObjectiveEventDataSubAgentSpawned].
	SubAgentSpawned SubAgentSpawned `json:"subAgentSpawned"`
	// This field is from variant [ObjectiveEventDataSubAgentUpdated].
	SubAgentUpdated SubAgentUpdated `json:"subAgentUpdated"`
	// This field is from variant [ObjectiveEventDataFinalized].
	Finalized ObjectiveEventDataFinalizedFinalized `json:"finalized"`
	// This field is from variant [ObjectiveEventDataNotice].
	Notice ObjectiveEventDataNoticeNotice `json:"notice"`
	// This field is from variant [ObjectiveEventDataTimedOut].
	TimedOut ObjectiveEventDataTimedOutTimedOut `json:"timedOut"`
	// This field is from variant [ObjectiveEventDataReasoning].
	Reasoning Reasoning `json:"reasoning"`
	JSON      struct {
		Type                   respjson.Field
		UserMessage            respjson.Field
		ToolApprovalRequested  respjson.Field
		ToolApproved           respjson.Field
		ToolDenied             respjson.Field
		ToolCalled             respjson.Field
		Error                  respjson.Field
		AssistantMessage       respjson.Field
		ToolResult             respjson.Field
		ToolError              respjson.Field
		ContextWindowCompacted respjson.Field
		MemoryRead             respjson.Field
		Cancelled              respjson.Field
		SubAgentSpawned        respjson.Field
		SubAgentUpdated        respjson.Field
		Finalized              respjson.Field
		Notice                 respjson.Field
		TimedOut               respjson.Field
		Reasoning              respjson.Field
		raw                    string
	} `json:"-"`
}

// anyObjectiveEventData is implemented by each variant of
// [ObjectiveEventDataUnion] to add type safety for the return type of
// [ObjectiveEventDataUnion.AsAny]
type anyObjectiveEventData interface {
	implObjectiveEventDataUnion()
}

func (ObjectiveEventDataUserMessage) implObjectiveEventDataUnion()            {}
func (ObjectiveEventDataToolApprovalRequested) implObjectiveEventDataUnion()  {}
func (ObjectiveEventDataToolApproved) implObjectiveEventDataUnion()           {}
func (ObjectiveEventDataToolDenied) implObjectiveEventDataUnion()             {}
func (ObjectiveEventDataToolCalled) implObjectiveEventDataUnion()             {}
func (ObjectiveEventDataError) implObjectiveEventDataUnion()                  {}
func (ObjectiveEventDataAssistantMessage) implObjectiveEventDataUnion()       {}
func (ObjectiveEventDataToolResult) implObjectiveEventDataUnion()             {}
func (ObjectiveEventDataToolError) implObjectiveEventDataUnion()              {}
func (ObjectiveEventDataContextWindowCompacted) implObjectiveEventDataUnion() {}
func (ObjectiveEventDataMemoryRead) implObjectiveEventDataUnion()             {}
func (ObjectiveEventDataCancelled) implObjectiveEventDataUnion()              {}
func (ObjectiveEventDataSubAgentSpawned) implObjectiveEventDataUnion()        {}
func (ObjectiveEventDataSubAgentUpdated) implObjectiveEventDataUnion()        {}
func (ObjectiveEventDataFinalized) implObjectiveEventDataUnion()              {}
func (ObjectiveEventDataNotice) implObjectiveEventDataUnion()                 {}
func (ObjectiveEventDataTimedOut) implObjectiveEventDataUnion()               {}
func (ObjectiveEventDataReasoning) implObjectiveEventDataUnion()              {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ObjectiveEventDataUnion.AsAny().(type) {
//	case cadenya.ObjectiveEventDataUserMessage:
//	case cadenya.ObjectiveEventDataToolApprovalRequested:
//	case cadenya.ObjectiveEventDataToolApproved:
//	case cadenya.ObjectiveEventDataToolDenied:
//	case cadenya.ObjectiveEventDataToolCalled:
//	case cadenya.ObjectiveEventDataError:
//	case cadenya.ObjectiveEventDataAssistantMessage:
//	case cadenya.ObjectiveEventDataToolResult:
//	case cadenya.ObjectiveEventDataToolError:
//	case cadenya.ObjectiveEventDataContextWindowCompacted:
//	case cadenya.ObjectiveEventDataMemoryRead:
//	case cadenya.ObjectiveEventDataCancelled:
//	case cadenya.ObjectiveEventDataSubAgentSpawned:
//	case cadenya.ObjectiveEventDataSubAgentUpdated:
//	case cadenya.ObjectiveEventDataFinalized:
//	case cadenya.ObjectiveEventDataNotice:
//	case cadenya.ObjectiveEventDataTimedOut:
//	case cadenya.ObjectiveEventDataReasoning:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ObjectiveEventDataUnion) AsAny() anyObjectiveEventData {
	switch u.Type {
	case "userMessage":
		return u.AsUserMessage()
	case "toolApprovalRequested":
		return u.AsToolApprovalRequested()
	case "toolApproved":
		return u.AsToolApproved()
	case "toolDenied":
		return u.AsToolDenied()
	case "toolCalled":
		return u.AsToolCalled()
	case "error":
		return u.AsError()
	case "assistantMessage":
		return u.AsAssistantMessage()
	case "toolResult":
		return u.AsToolResult()
	case "toolError":
		return u.AsToolError()
	case "contextWindowCompacted":
		return u.AsContextWindowCompacted()
	case "memoryRead":
		return u.AsMemoryRead()
	case "cancelled":
		return u.AsCancelled()
	case "subAgentSpawned":
		return u.AsSubAgentSpawned()
	case "subAgentUpdated":
		return u.AsSubAgentUpdated()
	case "finalized":
		return u.AsFinalized()
	case "notice":
		return u.AsNotice()
	case "timedOut":
		return u.AsTimedOut()
	case "reasoning":
		return u.AsReasoning()
	}
	return nil
}

func (u ObjectiveEventDataUnion) AsUserMessage() (v ObjectiveEventDataUserMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsToolApprovalRequested() (v ObjectiveEventDataToolApprovalRequested) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsToolApproved() (v ObjectiveEventDataToolApproved) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsToolDenied() (v ObjectiveEventDataToolDenied) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsToolCalled() (v ObjectiveEventDataToolCalled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsError() (v ObjectiveEventDataError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsAssistantMessage() (v ObjectiveEventDataAssistantMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsToolResult() (v ObjectiveEventDataToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsToolError() (v ObjectiveEventDataToolError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsContextWindowCompacted() (v ObjectiveEventDataContextWindowCompacted) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsMemoryRead() (v ObjectiveEventDataMemoryRead) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsCancelled() (v ObjectiveEventDataCancelled) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsSubAgentSpawned() (v ObjectiveEventDataSubAgentSpawned) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsSubAgentUpdated() (v ObjectiveEventDataSubAgentUpdated) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsFinalized() (v ObjectiveEventDataFinalized) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsNotice() (v ObjectiveEventDataNotice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsTimedOut() (v ObjectiveEventDataTimedOut) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectiveEventDataUnion) AsReasoning() (v ObjectiveEventDataReasoning) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectiveEventDataUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectiveEventDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataAssistantMessage struct {
	AssistantMessage AssistantMessage `json:"assistantMessage" api:"required"`
	// Any of "assistantMessage".
	Type ObjectiveEventDataAssistantMessageType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantMessage respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataAssistantMessage) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataAssistantMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataAssistantMessageType string

const (
	ObjectiveEventDataAssistantMessageTypeAssistantMessage ObjectiveEventDataAssistantMessageType = "assistantMessage"
)

type ObjectiveEventDataCancelled struct {
	// ObjectiveCancelled is the terminal event written when an objective is cancelled.
	// After this event, the objective is super-terminal: no further iterations,
	// compaction, or continuation are permitted.
	Cancelled ObjectiveEventDataCancelledCancelled `json:"cancelled" api:"required"`
	// Any of "cancelled".
	Type ObjectiveEventDataCancelledType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataCancelled) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataCancelled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveCancelled is the terminal event written when an objective is cancelled.
// After this event, the objective is super-terminal: no further iterations,
// compaction, or continuation are permitted.
type ObjectiveEventDataCancelledCancelled struct {
	// Optional human-readable note recorded at cancel time. Today the workflow sets
	// "Cancelled" but this field leaves room for richer reasons (e.g. "Cancelled by
	// user", "Cancelled by schedule sweep", "Credit balance exhausted").
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataCancelledCancelled) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataCancelledCancelled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataCancelledType string

const (
	ObjectiveEventDataCancelledTypeCancelled ObjectiveEventDataCancelledType = "cancelled"
)

type ObjectiveEventDataContextWindowCompacted struct {
	ContextWindowCompacted ContextWindowCompacted `json:"contextWindowCompacted" api:"required"`
	// Any of "contextWindowCompacted".
	Type ObjectiveEventDataContextWindowCompactedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContextWindowCompacted respjson.Field
		Type                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataContextWindowCompacted) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataContextWindowCompacted) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataContextWindowCompactedType string

const (
	ObjectiveEventDataContextWindowCompactedTypeContextWindowCompacted ObjectiveEventDataContextWindowCompactedType = "contextWindowCompacted"
)

type ObjectiveEventDataError struct {
	Error ObjectiveError `json:"error" api:"required"`
	// Any of "error".
	Type ObjectiveEventDataErrorType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataError) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataErrorType string

const (
	ObjectiveEventDataErrorTypeError ObjectiveEventDataErrorType = "error"
)

type ObjectiveEventDataFinalized struct {
	// ObjectiveFinalized is the terminal event written when an objective is finalized.
	// After this event, the objective is super-terminal: no further iterations,
	// compaction, or continuation are permitted.
	Finalized ObjectiveEventDataFinalizedFinalized `json:"finalized" api:"required"`
	// Any of "finalized".
	Type ObjectiveEventDataFinalizedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Finalized   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataFinalized) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataFinalized) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveFinalized is the terminal event written when an objective is finalized.
// After this event, the objective is super-terminal: no further iterations,
// compaction, or continuation are permitted.
type ObjectiveEventDataFinalizedFinalized struct {
	// If the objective was created with an output schema, and the agent successfully
	// completed the objective, this field will contain the structured output of the
	// objective.
	Output any `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataFinalizedFinalized) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataFinalizedFinalized) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataFinalizedType string

const (
	ObjectiveEventDataFinalizedTypeFinalized ObjectiveEventDataFinalizedType = "finalized"
)

type ObjectiveEventDataMemoryRead struct {
	// MemoryRead is emitted each time the agent resolves a key against the memory
	// cascade and loads an entry. Lookups that miss (key not found in any layer) do
	// not emit this event.
	MemoryRead MemoryRead `json:"memoryRead" api:"required"`
	// Any of "memoryRead".
	Type ObjectiveEventDataMemoryReadType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoryRead  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataMemoryRead) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataMemoryRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataMemoryReadType string

const (
	ObjectiveEventDataMemoryReadTypeMemoryRead ObjectiveEventDataMemoryReadType = "memoryRead"
)

type ObjectiveEventDataNotice struct {
	// Notice is a non-terminal diagnostic emitted by the runtime when something
	// noteworthy but non-fatal happens during an objective — for example a
	// just-in-time tool set failing to load, or a previously loaded tool being dropped
	// because it was archived. Notices carry no structured payload; they exist to make
	// the objective timeline self-explanatory.
	Notice ObjectiveEventDataNoticeNotice `json:"notice" api:"required"`
	// Any of "notice".
	Type ObjectiveEventDataNoticeType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Notice      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataNotice) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataNotice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Notice is a non-terminal diagnostic emitted by the runtime when something
// noteworthy but non-fatal happens during an objective — for example a
// just-in-time tool set failing to load, or a previously loaded tool being dropped
// because it was archived. Notices carry no structured payload; they exist to make
// the objective timeline self-explanatory.
type ObjectiveEventDataNoticeNotice struct {
	// Stable machine-readable identifier for the notice kind (for example
	// "tool_set_load_failed", "tool_archived"). Clients can switch on it or use it as
	// an i18n key; the message is the English fallback.
	Key string `json:"key"`
	// Any of "LEVEL_UNSPECIFIED", "LEVEL_INFO", "LEVEL_WARN".
	Level string `json:"level"`
	// Human-readable description of what happened.
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Level       respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataNoticeNotice) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataNoticeNotice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataNoticeType string

const (
	ObjectiveEventDataNoticeTypeNotice ObjectiveEventDataNoticeType = "notice"
)

type ObjectiveEventDataReasoning struct {
	// Reasoning carries the human-readable reasoning text a model produced while
	// working on an iteration — extended thinking (Anthropic, Gemini) or reasoning
	// summaries (OpenAI). It is emitted alongside the assistant message from the same
	// model response and is purely informational: the text shown here is never sent
	// back to the model.
	Reasoning Reasoning `json:"reasoning" api:"required"`
	// Any of "reasoning".
	Type ObjectiveEventDataReasoningType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reasoning   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataReasoning) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataReasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataReasoningType string

const (
	ObjectiveEventDataReasoningTypeReasoning ObjectiveEventDataReasoningType = "reasoning"
)

type ObjectiveEventDataSubAgentSpawned struct {
	SubAgentSpawned SubAgentSpawned `json:"subAgentSpawned" api:"required"`
	// Any of "subAgentSpawned".
	Type ObjectiveEventDataSubAgentSpawnedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubAgentSpawned respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataSubAgentSpawned) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataSubAgentSpawned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataSubAgentSpawnedType string

const (
	ObjectiveEventDataSubAgentSpawnedTypeSubAgentSpawned ObjectiveEventDataSubAgentSpawnedType = "subAgentSpawned"
)

type ObjectiveEventDataSubAgentUpdated struct {
	SubAgentUpdated SubAgentUpdated `json:"subAgentUpdated" api:"required"`
	// Any of "subAgentUpdated".
	Type ObjectiveEventDataSubAgentUpdatedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubAgentUpdated respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataSubAgentUpdated) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataSubAgentUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataSubAgentUpdatedType string

const (
	ObjectiveEventDataSubAgentUpdatedTypeSubAgentUpdated ObjectiveEventDataSubAgentUpdatedType = "subAgentUpdated"
)

type ObjectiveEventDataTimedOut struct {
	// ObjectiveTimedOut is the terminal event written when an objective is finalized
	// by the inactivity sweep because it saw no activity (no user messages, no LLM
	// calls) within its variation's inactivity timeout — or the system-wide 24 hour
	// maximum when no timeout is configured. The objective produces no output. After
	// this event, the objective is super-terminal: no further iterations, compaction,
	// or continuation are permitted.
	TimedOut ObjectiveEventDataTimedOutTimedOut `json:"timedOut" api:"required"`
	// Any of "timedOut".
	Type ObjectiveEventDataTimedOutType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimedOut    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataTimedOut) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataTimedOut) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveTimedOut is the terminal event written when an objective is finalized
// by the inactivity sweep because it saw no activity (no user messages, no LLM
// calls) within its variation's inactivity timeout — or the system-wide 24 hour
// maximum when no timeout is configured. The objective produces no output. After
// this event, the objective is super-terminal: no further iterations, compaction,
// or continuation are permitted.
type ObjectiveEventDataTimedOutTimedOut struct {
	// Human-readable note recorded at timeout time (e.g. "Timed out after 2h of
	// inactivity").
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataTimedOutTimedOut) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataTimedOutTimedOut) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataTimedOutType string

const (
	ObjectiveEventDataTimedOutTypeTimedOut ObjectiveEventDataTimedOutType = "timedOut"
)

type ObjectiveEventDataToolApprovalRequested struct {
	ToolApprovalRequested ToolApprovalRequested `json:"toolApprovalRequested" api:"required"`
	// Any of "toolApprovalRequested".
	Type ObjectiveEventDataToolApprovalRequestedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolApprovalRequested respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataToolApprovalRequested) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataToolApprovalRequested) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataToolApprovalRequestedType string

const (
	ObjectiveEventDataToolApprovalRequestedTypeToolApprovalRequested ObjectiveEventDataToolApprovalRequestedType = "toolApprovalRequested"
)

type ObjectiveEventDataToolApproved struct {
	ToolApproved ToolApproved `json:"toolApproved" api:"required"`
	// Any of "toolApproved".
	Type ObjectiveEventDataToolApprovedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolApproved respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataToolApproved) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataToolApproved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataToolApprovedType string

const (
	ObjectiveEventDataToolApprovedTypeToolApproved ObjectiveEventDataToolApprovedType = "toolApproved"
)

type ObjectiveEventDataToolCalled struct {
	ToolCalled ToolCalled `json:"toolCalled" api:"required"`
	// Any of "toolCalled".
	Type ObjectiveEventDataToolCalledType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolCalled  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataToolCalled) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataToolCalled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataToolCalledType string

const (
	ObjectiveEventDataToolCalledTypeToolCalled ObjectiveEventDataToolCalledType = "toolCalled"
)

type ObjectiveEventDataToolDenied struct {
	ToolDenied ToolDenied `json:"toolDenied" api:"required"`
	// Any of "toolDenied".
	Type ObjectiveEventDataToolDeniedType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolDenied  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataToolDenied) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataToolDenied) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataToolDeniedType string

const (
	ObjectiveEventDataToolDeniedTypeToolDenied ObjectiveEventDataToolDeniedType = "toolDenied"
)

type ObjectiveEventDataToolError struct {
	ToolError ToolError `json:"toolError" api:"required"`
	// Any of "toolError".
	Type ObjectiveEventDataToolErrorType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolError   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataToolError) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataToolError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataToolErrorType string

const (
	ObjectiveEventDataToolErrorTypeToolError ObjectiveEventDataToolErrorType = "toolError"
)

type ObjectiveEventDataToolResult struct {
	ToolResult ToolResult `json:"toolResult" api:"required"`
	// Any of "toolResult".
	Type ObjectiveEventDataToolResultType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolResult  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataToolResult) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataToolResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataToolResultType string

const (
	ObjectiveEventDataToolResultTypeToolResult ObjectiveEventDataToolResultType = "toolResult"
)

type ObjectiveEventDataUserMessage struct {
	// Any of "userMessage".
	Type        ObjectiveEventDataUserMessageType `json:"type" api:"required"`
	UserMessage UserMessage                       `json:"userMessage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		UserMessage respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventDataUserMessage) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventDataUserMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveEventDataUserMessageType string

const (
	ObjectiveEventDataUserMessageTypeUserMessage ObjectiveEventDataUserMessageType = "userMessage"
)

type ObjectiveEventInfo struct {
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective shared.OperationMetadata `json:"objective"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		Objective   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveEventInfo) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveEventInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectiveInfo provides read-only aggregated statistics about an objective's
// execution
type ObjectiveInfo struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent shared.ResourceMetadata `json:"agent" api:"required"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	AgentVariation shared.ResourceMetadata `json:"agentVariation" api:"required"`
	// A profile identifies a user or non-human principal (such as an API key) at the
	// account level. Profiles are account-scoped and can be granted access to multiple
	// workspaces.
	CreatedBy Profile `json:"createdBy" api:"required"`
	// ID of the objective's current (most recent) context window. Hydrated on demand;
	// empty when the objective has not yet produced a context window.
	CurrentContextWindowID string `json:"currentContextWindowId" api:"required"`
	// The effective memory cascade at objective creation time: the episodic layer
	// (when present), then Objective.memory_cascade, then the variation's baseline
	// layers by ascending position. Order is resolution order — index 0 is the most
	// specific and is consulted first; the first layer containing a key wins. Returned
	// on reads so clients can see exactly what the objective resolves against without
	// re-joining variation state.
	EffectiveMemoryCascade []MemoryReference `json:"effectiveMemoryCascade" api:"required"`
	// Total number of context windows that this objective has generated
	TotalContextWindows int64 `json:"totalContextWindows" api:"required"`
	// Total number of events generated during this objective's execution
	TotalEvents int64 `json:"totalEvents" api:"required"`
	// Total input tokens consumed across all LLM completions across all context
	// windows
	TotalInputTokens int64 `json:"totalInputTokens" api:"required"`
	TotalIterations  int64 `json:"totalIterations" api:"required"`
	// Total output tokens generated across all LLM completions across all context
	// windows
	TotalOutputTokens int64 `json:"totalOutputTokens" api:"required"`
	// Total number of tool calls made during execution
	TotalToolCalls int64 `json:"totalToolCalls" api:"required"`
	// SubjectReference is the read-only echo of a resource's subject association,
	// carrying both Cadenya's canonical id and the customer's own key.
	Subject SubjectReference `json:"subject"`
	// TenantReference is the read-only echo of a resource's tenant association,
	// carrying both Cadenya's canonical id and the customer's own key.
	Tenant TenantReference `json:"tenant"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Widget shared.BareMetadata `json:"widget"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent                  respjson.Field
		AgentVariation         respjson.Field
		CreatedBy              respjson.Field
		CurrentContextWindowID respjson.Field
		EffectiveMemoryCascade respjson.Field
		TotalContextWindows    respjson.Field
		TotalEvents            respjson.Field
		TotalInputTokens       respjson.Field
		TotalIterations        respjson.Field
		TotalOutputTokens      respjson.Field
		TotalToolCalls         respjson.Field
		Subject                respjson.Field
		Tenant                 respjson.Field
		Widget                 respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveInfo) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveSecret struct {
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveSecret) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reasoning carries the human-readable reasoning text a model produced while
// working on an iteration — extended thinking (Anthropic, Gemini) or reasoning
// summaries (OpenAI). It is emitted alongside the assistant message from the same
// model response and is purely informational: the text shown here is never sent
// back to the model.
type Reasoning struct {
	// The reasoning text. May be a verbatim chain of thought or a provider-generated
	// summary depending on the model.
	Content string `json:"content" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Reasoning) RawJSON() string { return r.JSON.raw }
func (r *Reasoning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAgentSpawned struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent shared.ResourceMetadata `json:"agent"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective shared.OperationMetadata `json:"objective"`
	Task      string                   `json:"task"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent       respjson.Field
		Objective   respjson.Field
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAgentSpawned) RawJSON() string { return r.JSON.raw }
func (r *SubAgentSpawned) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAgentUpdated struct {
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Agent   shared.BareMetadata `json:"agent"`
	Message string              `json:"message"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Objective shared.BareMetadata `json:"objective"`
	// Any of "STATUS_UNSPECIFIED", "STATUS_PENDING", "STATUS_RUNNING",
	// "STATUS_COMPLETED", "STATUS_FAILED", "STATUS_CANCELLED".
	Status SubAgentUpdatedStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent       respjson.Field
		Message     respjson.Field
		Objective   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubAgentUpdated) RawJSON() string { return r.JSON.raw }
func (r *SubAgentUpdated) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SubAgentUpdatedStatus string

const (
	SubAgentUpdatedStatusStatusUnspecified SubAgentUpdatedStatus = "STATUS_UNSPECIFIED"
	SubAgentUpdatedStatusStatusPending     SubAgentUpdatedStatus = "STATUS_PENDING"
	SubAgentUpdatedStatusStatusRunning     SubAgentUpdatedStatus = "STATUS_RUNNING"
	SubAgentUpdatedStatusStatusCompleted   SubAgentUpdatedStatus = "STATUS_COMPLETED"
	SubAgentUpdatedStatusStatusFailed      SubAgentUpdatedStatus = "STATUS_FAILED"
	SubAgentUpdatedStatusStatusCancelled   SubAgentUpdatedStatus = "STATUS_CANCELLED"
)

type ToolApprovalRequested struct {
	// The ID of the objective tool call record. Use this ID with the ApproveToolCall
	// or DenyToolCall RPCs to approve or deny the tool call.
	ToolCallID string `json:"toolCallId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolApprovalRequested) RawJSON() string { return r.JSON.raw }
func (r *ToolApprovalRequested) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolApproved struct {
	// The ID of the objective tool call record that was approved via the
	// ApproveToolCall RPC.
	ToolCallID string `json:"toolCallId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolApproved) RawJSON() string { return r.JSON.raw }
func (r *ToolApproved) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolCalled struct {
	// The arguments passed to the tool.
	Arguments map[string]any `json:"arguments"`
	// Config defines the adapter to use for the tool. This is used to determine how
	// the tool is called. For example, if the tool is an HTTP tool, the adapter will
	// be Http. If the tool is an inline tool, the adapter will be Inline.
	Config ToolSpecConfigUnion `json:"config"`
	// CallableTool is a union that represents a tool that can be called by an agent.
	// In Cadenya, a tool that is used within an agent objective might be a
	// user-defined tool (IE: MCP, HTTP), another Agent (useful to separate context),
	// or a Cadenya Tool (one Cadenya provides).
	Tool CallableToolUnion `json:"tool"`
	// The ID of the objective tool call record that was executed.
	ToolCallID string `json:"toolCallId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Config      respjson.Field
		Tool        respjson.Field
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolCalled) RawJSON() string { return r.JSON.raw }
func (r *ToolCalled) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolDenied struct {
	// The memo provided by the reviewer when denying the tool call. This is passed to
	// the agent to provide further instructions.
	Memo string `json:"memo"`
	// The ID of the objective tool call record that was denied via the DenyToolCall
	// RPC.
	ToolCallID string `json:"toolCallId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Memo        respjson.Field
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolDenied) RawJSON() string { return r.JSON.raw }
func (r *ToolDenied) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolError struct {
	Message string `json:"message"`
	// The ID of the objective tool call record that encountered an error during
	// execution.
	ToolCallID string `json:"toolCallId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolError) RawJSON() string { return r.JSON.raw }
func (r *ToolError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolResult struct {
	// ObjectiveToolCallResult is the content a tool returned after execution. Tools
	// can return multiple content blocks, and blocks can be multi-modal (text, image,
	// audio). Media blocks are stored by Cadenya and served as short-lived signed URLs
	// rather than inline bytes.
	Result     ObjectiveToolCallResult `json:"result" api:"required"`
	ToolCallID string                  `json:"toolCallId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ToolCallID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolResult) RawJSON() string { return r.JSON.raw }
func (r *ToolResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserMessage struct {
	Content string `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMessage) RawJSON() string { return r.JSON.raw }
func (r *UserMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compact objective response
type ObjectiveCompactResponse struct {
	// The new context window created by the compaction
	ContextWindow ObjectiveContextWindowData `json:"contextWindow"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContextWindow respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveCompactResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveCompactResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveGetDiagnosticsResponse struct {
	// ObjectiveDiagnostics is the context-usage breakdown measured for a single
	// iteration at request-assembly time. It reports how much of the context window
	// each component occupies so tool parameters, memory cascades, and prompts can be
	// tuned against real token usage.
	Diagnostics ObjectiveDiagnostics `json:"diagnostics" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Diagnostics respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveGetDiagnosticsResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveGetDiagnosticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveNewParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	AgentID     string            `json:"agentId" api:"required"`
	// Arbitrary data rendered into the selected variation's system_prompt_template
	// (liquid) to produce the objective's system prompt. If the agent has a
	// system_prompt_data_schema, this must satisfy it.
	SystemPromptData map[string]any `json:"systemPromptData,omitzero" api:"required"`
	// Optional explicit first user message for the LLM chat history. When not set, the
	// selected variation's first_user_message_template is rendered with
	// first_user_message_data instead. If neither this field nor a
	// first_user_message_template is present, the request is rejected with
	// InvalidArgument.
	FirstUserMessage param.Opt[string] `json:"firstUserMessage,omitzero"`
	// Optional explicit variation selection. Overrides the agent's
	// variation_selection_mode.
	VariationID param.Opt[string] `json:"variationId,omitzero"`
	// Episodic is used to configure the episodic memory for the objective
	EpisodicMemory ObjectiveNewParamsEpisodicMemory `json:"episodicMemory,omitzero"`
	// Arbitrary data rendered into the selected variation's
	// first_user_message_template (liquid) to produce the first user message. Separate
	// from `system_prompt_data`, which renders the system prompt template.
	FirstUserMessageData map[string]any `json:"firstUserMessageData,omitzero"`
	// Memory layers/entries layered over the baseline cascade inherited from the
	// selected variation — element-level rules over inherited styles, in CSS terms.
	//
	// Array order is resolution order: EARLIER elements are more specific and are
	// consulted first. Entries pinned via memory_entry_id behave as single-entry
	// layers at their position.
	//
	// System-managed layers (e.g., episodic) cannot be referenced here; they attach
	// themselves automatically based on the episodic key.
	//
	// Size cap: the TOTAL effective cascade (this field + the variation's memory layer
	// assignments) must not exceed 10 entries. A request that would produce a larger
	// cascade is rejected with InvalidArgument.
	MemoryCascade []MemoryReferenceParam `json:"memoryCascade,omitzero"`
	// CreateOperationMetadata contains the user-provided fields for creating an
	// operation. Read-only fields (id, account_id, workspace_id, created_at,
	// profile_id) are excluded since they are set by the server.
	Metadata shared.CreateOperationMetadataParam `json:"metadata,omitzero"`
	// Parameters forced onto this objective's tool calls. A pinned parameter is an
	// overlay on a tool's JSON schema: the parameter is removed from what the LLM
	// sees, and its value is always overwritten server-side with the pinned value —
	// the model cannot choose a different value for it.
	PinnedParameters map[string]string `json:"pinnedParameters,omitzero"`
	// Secrets that can be used in the headers for tool calls using the secret
	// interpolation format.
	Secrets []ObjectiveNewParamsSecret `json:"secrets,omitzero"`
	// SubjectAssertion identifies a person within a tenant in the customer's own
	// namespace — typically their user id. Asserting a subject upserts the subject
	// record under the asserted tenant and associates the created resource with it. A
	// subject assertion is only valid alongside a tenant assertion: subject
	// identifiers are scoped to their tenant.
	Subject SubjectAssertionParam `json:"subject,omitzero"`
	// TenantAssertion identifies a tenant in the customer's own namespace — their org,
	// company, or team identifier for an end user. Asserting a tenant upserts the
	// tenant record in the workspace (keyed on `id` as the tenant's external_id) and
	// associates the created resource with it.
	Tenant TenantAssertionParam `json:"tenant,omitzero"`
	paramObj
}

func (r ObjectiveNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Episodic is used to configure the episodic memory for the objective
//
// The property Key is required.
type ObjectiveNewParamsEpisodicMemory struct {
	// The caller-supplied episodic key. Objectives created with the same key (for the
	// same agent) share one episodic memory layer.
	Key string `json:"key" api:"required"`
	paramObj
}

func (r ObjectiveNewParamsEpisodicMemory) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveNewParamsEpisodicMemory
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveNewParamsEpisodicMemory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveNewParamsSecret struct {
	Name  param.Opt[string] `json:"name,omitzero"`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r ObjectiveNewParamsSecret) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveNewParamsSecret
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveNewParamsSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ObjectiveListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Agent ID for filtering
	AgentID param.Opt[string] `query:"agentId,omitzero" json:"-"`
	// Filter to objectives produced by a specific AgentSchedule. Accepts canonical
	// as\_… form or external_id:<value> form.
	AgentScheduleID param.Opt[string] `query:"agentScheduleId,omitzero" json:"-"`
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
	// Optional filters
	ParentObjectiveID param.Opt[string] `query:"parentObjectiveId,omitzero" json:"-"`
	ProfileID         param.Opt[string] `query:"profileId,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// Filter to objectives associated with a subject. Accepts the canonical `subj_…`
	// form or the `external_id:<value>` form; the external_id form is scoped within a
	// tenant and requires `tenant_id` to also be set.
	SubjectID param.Opt[string] `query:"subjectId,omitzero" json:"-"`
	// Filter to objectives associated with a tenant. Accepts the canonical `tenant_…`
	// form or the `external_id:<value>` form.
	TenantID param.Opt[string] `query:"tenantId,omitzero" json:"-"`
	// Filter to objectives whose conversation ran through a widget. Accepts the
	// canonical `wgt_…` form or the `external_id:<value>` form.
	WidgetID param.Opt[string] `query:"widgetId,omitzero" json:"-"`
	// Filter to objectives created by a specific widget session.
	WidgetSessionID param.Opt[string] `query:"widgetSessionId,omitzero" json:"-"`
	// Filter by state
	//
	// Any of "STATE_UNSPECIFIED", "STATE_PENDING", "STATE_RUNNING", "STATE_WAITING",
	// "STATE_FAILED", "STATE_CANCELLED", "STATE_FINALIZED", "STATE_TIMED_OUT".
	State ObjectiveListParamsState `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectiveListParams]'s query parameters as `url.Values`.
func (r ObjectiveListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by state
type ObjectiveListParamsState string

const (
	ObjectiveListParamsStateStateUnspecified ObjectiveListParamsState = "STATE_UNSPECIFIED"
	ObjectiveListParamsStateStatePending     ObjectiveListParamsState = "STATE_PENDING"
	ObjectiveListParamsStateStateRunning     ObjectiveListParamsState = "STATE_RUNNING"
	ObjectiveListParamsStateStateWaiting     ObjectiveListParamsState = "STATE_WAITING"
	ObjectiveListParamsStateStateFailed      ObjectiveListParamsState = "STATE_FAILED"
	ObjectiveListParamsStateStateCancelled   ObjectiveListParamsState = "STATE_CANCELLED"
	ObjectiveListParamsStateStateFinalized   ObjectiveListParamsState = "STATE_FINALIZED"
	ObjectiveListParamsStateStateTimedOut    ObjectiveListParamsState = "STATE_TIMED_OUT"
)

type ObjectiveCancelParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Optional reason for cancellation
	Reason param.Opt[string] `json:"reason,omitzero"`
	paramObj
}

func (r ObjectiveCancelParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveCancelParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveCancelParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveCompactParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// CompactionConfig defines how context window compaction behaves for objectives
	// using this variation.
	CompactionConfig AgentVariationSpecCompactionConfigParam `json:"compactionConfig,omitzero"`
	paramObj
}

func (r ObjectiveCompactParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveCompactParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveCompactParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveContinueParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// The message to continue an objective that has completed (or you are enqueing)
	Message string `json:"message" api:"required"`
	// When set to true, the message will be enqueued for when the agent loop is
	// available to process it.
	Enqueue param.Opt[bool] `json:"enqueue,omitzero"`
	paramObj
}

func (r ObjectiveContinueParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectiveContinueParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectiveContinueParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveListContextWindowsParams struct {
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
	paramObj
}

// URLQuery serializes [ObjectiveListContextWindowsParams]'s query parameters as
// `url.Values`.
func (r ObjectiveListContextWindowsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectiveListEventsParams struct {
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
	// Optional string to fetch events since an ID
	SinceEventID param.Opt[string] `query:"sinceEventId,omitzero" json:"-"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	// Optional context window ID to filter events by
	WindowID param.Opt[string] `query:"windowId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectiveListEventsParams]'s query parameters as
// `url.Values`.
func (r ObjectiveListEventsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectiveGetDiagnosticsParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ObjectiveStreamEventsParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}
