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

// ObjectiveService contains methods and other services that help with interacting
// with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveService] method instead.
type ObjectiveService struct {
	Options   []option.RequestOption
	Tools     *ObjectiveToolService
	ToolCalls *ObjectiveToolCallService
	Tasks     *ObjectiveTaskService
	Feedback  *ObjectiveFeedbackService
}

// NewObjectiveService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectiveService(opts ...option.RequestOption) (r *ObjectiveService) {
	r = &ObjectiveService{}
	r.Options = opts
	r.Tools = NewObjectiveToolService(opts...)
	r.ToolCalls = NewObjectiveToolCallService(opts...)
	r.Tasks = NewObjectiveTaskService(opts...)
	r.Feedback = NewObjectiveFeedbackService(opts...)
	return
}

// Creates a new objective in the workspace
func (r *ObjectiveService) New(ctx context.Context, body ObjectiveNewParams, opts ...option.RequestOption) (res *Objective, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/objectives"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves an objective by ID from the workspace
func (r *ObjectiveService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Objective, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all objectives in the workspace
func (r *ObjectiveService) List(ctx context.Context, query ObjectiveListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[Objective], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/objectives"
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

// Lists all objectives in the workspace
func (r *ObjectiveService) ListAutoPaging(ctx context.Context, query ObjectiveListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[Objective] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, query, opts...))
}

// Cancels a running or pending objective. The objective's state will be set to
// STATE_CANCELLED.
func (r *ObjectiveService) Cancel(ctx context.Context, objectiveID string, body ObjectiveCancelParams, opts ...option.RequestOption) (res *Objective, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/cancel", objectiveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Triggers compaction on a running objective. Optionally override the variation's
// compaction config.
func (r *ObjectiveService) Compact(ctx context.Context, objectiveID string, body ObjectiveCompactParams, opts ...option.RequestOption) (res *ObjectiveCompactResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/compact", objectiveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Continues an objective that has completed
func (r *ObjectiveService) Continue(ctx context.Context, objectiveID string, body ObjectiveContinueParams, opts ...option.RequestOption) (res *ObjectiveContinueResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/continue", objectiveID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Read-only list of the last five windows of execution for this objective, ordered
// by most recent first
func (r *ObjectiveService) ListContextWindows(ctx context.Context, objectiveID string, query ObjectiveListContextWindowsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveContextWindow], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/context_windows", objectiveID)
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

// Read-only list of the last five windows of execution for this objective, ordered
// by most recent first
func (r *ObjectiveService) ListContextWindowsAutoPaging(ctx context.Context, objectiveID string, query ObjectiveListContextWindowsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveContextWindow] {
	return pagination.NewCursorPaginationAutoPager(r.ListContextWindows(ctx, objectiveID, query, opts...))
}

// Lists all events for an objective
func (r *ObjectiveService) ListEvents(ctx context.Context, objectiveID string, query ObjectiveListEventsParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveListEventsResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/events", objectiveID)
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

// Lists all events for an objective
func (r *ObjectiveService) ListEventsAutoPaging(ctx context.Context, objectiveID string, query ObjectiveListEventsParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveListEventsResponse] {
	return pagination.NewCursorPaginationAutoPager(r.ListEvents(ctx, objectiveID, query, opts...))
}

type AssistantMessage struct {
	Content   string               `json:"content"`
	ToolCalls []AssistantToolCall  `json:"toolCalls"`
	JSON      assistantMessageJSON `json:"-"`
}

// assistantMessageJSON contains the JSON metadata for the struct
// [AssistantMessage]
type assistantMessageJSON struct {
	Content     apijson.Field
	ToolCalls   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AssistantMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantMessageJSON) RawJSON() string {
	return r.raw
}

type AssistantToolCall struct {
	Arguments    string `json:"arguments"`
	FunctionName string `json:"functionName"`
	// CallableTool is a union that represents a tool that can be called by an agent.
	// In Cadenya, a tool that is used within an agent objective might be a
	// user-defined tool (IE: MCP, HTTP), another Agent (useful to separate context),
	// or a Cadenya Tool (one Cadenya provides).
	Tool CallableTool          `json:"tool"`
	JSON assistantToolCallJSON `json:"-"`
}

// assistantToolCallJSON contains the JSON metadata for the struct
// [AssistantToolCall]
type assistantToolCallJSON struct {
	Arguments    apijson.Field
	FunctionName apijson.Field
	Tool         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AssistantToolCall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r assistantToolCallJSON) RawJSON() string {
	return r.raw
}

// CallableTool is a union that represents a tool that can be called by an agent.
// In Cadenya, a tool that is used within an agent objective might be a
// user-defined tool (IE: MCP, HTTP), another Agent (useful to separate context),
// or a Cadenya Tool (one Cadenya provides).
type CallableTool struct {
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Agent shared.ResourceMetadata `json:"agent"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	CadenyaProvidedTool shared.ResourceMetadata `json:"cadenyaProvidedTool"`
	// Standard metadata for persistent, named resources (e.g., agents, tools, prompts)
	Tool shared.ResourceMetadata `json:"tool"`
	JSON callableToolJSON        `json:"-"`
}

// callableToolJSON contains the JSON metadata for the struct [CallableTool]
type callableToolJSON struct {
	Agent               apijson.Field
	CadenyaProvidedTool apijson.Field
	Tool                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CallableTool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callableToolJSON) RawJSON() string {
	return r.raw
}

type ContextWindowCompacted struct {
	// Number of messages that were compacted
	MessagesCompacted int64 `json:"messagesCompacted"`
	// The new context window created by this compaction
	NewContextWindow ObjectiveContextWindowData `json:"newContextWindow"`
	// The strategies that were applied during this compaction
	Strategies []string `json:"strategies"`
	// The summary generated by the summarization strategy, if used.
	Summary string                     `json:"summary"`
	JSON    contextWindowCompactedJSON `json:"-"`
}

// contextWindowCompactedJSON contains the JSON metadata for the struct
// [ContextWindowCompacted]
type contextWindowCompactedJSON struct {
	MessagesCompacted apijson.Field
	NewContextWindow  apijson.Field
	Strategies        apijson.Field
	Summary           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ContextWindowCompacted) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contextWindowCompactedJSON) RawJSON() string {
	return r.raw
}

// MemoryRead is emitted each time the agent resolves a key against the memory
// stack and loads an entry. Lookups that miss (key not found in any layer) do not
// emit this event.
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
	Message string         `json:"message"`
	JSON    memoryReadJSON `json:"-"`
}

// memoryReadJSON contains the JSON metadata for the struct [MemoryRead]
type memoryReadJSON struct {
	MemoryEntryID apijson.Field
	MemoryLayerID apijson.Field
	Message       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MemoryRead) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryReadJSON) RawJSON() string {
	return r.raw
}

// MemoryReference identifies a memory layer or a specific entry within one, for
// composition into a memory stack. Used on objectives (where entry pinning is
// permitted).
//
// memory*layer_id accepts both the canonical form (memlyr*…) and the external-id
// form (external_id:my-custom-id). The same applies to memory_entry_id.
type MemoryReference struct {
	// When set, pushes only this entry from memory_layer_id onto the stack — behaves
	// as a single-entry layer (only this key resolves at this position). The entry
	// must belong to memory_layer_id; mismatches are rejected with InvalidArgument.
	MemoryEntryID string              `json:"memoryEntryId"`
	MemoryLayerID string              `json:"memoryLayerId"`
	JSON          memoryReferenceJSON `json:"-"`
}

// memoryReferenceJSON contains the JSON metadata for the struct [MemoryReference]
type memoryReferenceJSON struct {
	MemoryEntryID apijson.Field
	MemoryLayerID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MemoryReference) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryReferenceJSON) RawJSON() string {
	return r.raw
}

// MemoryReference identifies a memory layer or a specific entry within one, for
// composition into a memory stack. Used on objectives (where entry pinning is
// permitted).
//
// memory*layer_id accepts both the canonical form (memlyr*…) and the external-id
// form (external_id:my-custom-id). The same applies to memory_entry_id.
type MemoryReferenceParam struct {
	// When set, pushes only this entry from memory_layer_id onto the stack — behaves
	// as a single-entry layer (only this key resolves at this position). The entry
	// must belong to memory_layer_id; mismatches are rejected with InvalidArgument.
	MemoryEntryID param.Field[string] `json:"memoryEntryId"`
	MemoryLayerID param.Field[string] `json:"memoryLayerId"`
}

func (r MemoryReferenceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Objective struct {
	Data ObjectiveData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	Status   ObjectiveStatus          `json:"status" api:"required"`
	// ObjectiveInfo provides read-only aggregated statistics about an objective's
	// execution
	Info ObjectiveInfo `json:"info"`
	// Read-only list of the last five windows of execution for this objective, ordered
	// by most recent first. Is only included in singular RPC calls (GetObjective, for
	// example).
	LastFiveWindows []ObjectiveContextWindow `json:"lastFiveWindows"`
	JSON            objectiveJSON            `json:"-"`
}

// objectiveJSON contains the JSON metadata for the struct [Objective]
type objectiveJSON struct {
	Data            apijson.Field
	Metadata        apijson.Field
	Status          apijson.Field
	Info            apijson.Field
	LastFiveWindows apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Objective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveJSON) RawJSON() string {
	return r.raw
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
	JSON     objectiveContextWindowJSON `json:"-"`
}

// objectiveContextWindowJSON contains the JSON metadata for the struct
// [ObjectiveContextWindow]
type objectiveContextWindowJSON struct {
	Data        apijson.Field
	Metadata    apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveContextWindow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveContextWindowJSON) RawJSON() string {
	return r.raw
}

type ObjectiveContextWindowInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy Profile `json:"createdBy"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective shared.OperationMetadata       `json:"objective"`
	JSON      objectiveContextWindowInfoJSON `json:"-"`
}

// objectiveContextWindowInfoJSON contains the JSON metadata for the struct
// [ObjectiveContextWindowInfo]
type objectiveContextWindowInfoJSON struct {
	CreatedBy   apijson.Field
	Objective   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveContextWindowInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveContextWindowInfoJSON) RawJSON() string {
	return r.raw
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
	Sequence int64                          `json:"sequence"`
	JSON     objectiveContextWindowDataJSON `json:"-"`
}

// objectiveContextWindowDataJSON contains the JSON metadata for the struct
// [ObjectiveContextWindowData]
type objectiveContextWindowDataJSON struct {
	CompletionTokens                   apijson.Field
	ObjectiveID                        apijson.Field
	PreviousWindowContinueInstructions apijson.Field
	PromptTokens                       apijson.Field
	Sequence                           apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ObjectiveContextWindowData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveContextWindowDataJSON) RawJSON() string {
	return r.raw
}

type ObjectiveData struct {
	// Agent resource
	Agent Agent `json:"agent"`
	// Represents a dynamically typed value which can be either null, a number, a
	// string, a boolean, a recursive struct value, or a list of values.
	Data interface{} `json:"data"`
	// The initial message sent to the agent. This becomes the first user message in
	// the LLM chat history.
	InitialMessage string `json:"initialMessage"`
	// Memory layers/entries to push onto this objective's memory stack on top of the
	// baseline stack inherited from the selected variation.
	//
	// See "Memory stack composition" in memory.proto for lookup semantics.
	//
	// Array order is push order: the first element sits lower in the objective's
	// contribution to the stack; the LAST element ends up on top of the effective
	// stack. Entries pinned via memory_entry_id behave as single-entry layers at their
	// position.
	//
	// System-managed layers (e.g., episodic) cannot be referenced here; they attach
	// themselves automatically via the runtime based on episodic_key.
	//
	// Stack size cap: the TOTAL effective stack (variation's memory layers
	//
	//   - this field) must not exceed 10 entries. A request that would produce an
	//     effective stack larger than 10 is rejected with InvalidArgument. Variations
	//     themselves are capped at 10 memory layer assignments, so a variation that is
	//     already "full" leaves no room for objective-level references.
	MemoryStack []MemoryReference `json:"memoryStack"`
	// A parent objective means the objective was spawned off using a separate agent to
	// complete an objective
	ParentObjectiveID string `json:"parentObjectiveId"`
	// Secrets that can be used in the headers for tool calls using the secret
	// interpolation format.
	Secrets []ObjectiveDataSecret `json:"secrets"`
	// ID of the AgentSchedule that produced this objective, when applicable.
	// Read-only; populated by the runtime when the objective is created from a
	// schedule fire. Empty when the objective was created via CreateObjective
	// directly.
	SourceScheduleID string `json:"sourceScheduleId"`
	// system_prompt is read-only, derived from the selected variation's prompt
	SystemPrompt string `json:"systemPrompt"`
	// AgentVariation resource
	Variation AgentVariation    `json:"variation"`
	JSON      objectiveDataJSON `json:"-"`
}

// objectiveDataJSON contains the JSON metadata for the struct [ObjectiveData]
type objectiveDataJSON struct {
	Agent             apijson.Field
	Data              apijson.Field
	InitialMessage    apijson.Field
	MemoryStack       apijson.Field
	ParentObjectiveID apijson.Field
	Secrets           apijson.Field
	SourceScheduleID  apijson.Field
	SystemPrompt      apijson.Field
	Variation         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObjectiveData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveDataJSON) RawJSON() string {
	return r.raw
}

type ObjectiveDataParam struct {
	// Represents a dynamically typed value which can be either null, a number, a
	// string, a boolean, a recursive struct value, or a list of values.
	Data param.Field[interface{}] `json:"data"`
	// The initial message sent to the agent. This becomes the first user message in
	// the LLM chat history.
	InitialMessage param.Field[string] `json:"initialMessage"`
	// Memory layers/entries to push onto this objective's memory stack on top of the
	// baseline stack inherited from the selected variation.
	//
	// See "Memory stack composition" in memory.proto for lookup semantics.
	//
	// Array order is push order: the first element sits lower in the objective's
	// contribution to the stack; the LAST element ends up on top of the effective
	// stack. Entries pinned via memory_entry_id behave as single-entry layers at their
	// position.
	//
	// System-managed layers (e.g., episodic) cannot be referenced here; they attach
	// themselves automatically via the runtime based on episodic_key.
	//
	// Stack size cap: the TOTAL effective stack (variation's memory layers
	//
	//   - this field) must not exceed 10 entries. A request that would produce an
	//     effective stack larger than 10 is rejected with InvalidArgument. Variations
	//     themselves are capped at 10 memory layer assignments, so a variation that is
	//     already "full" leaves no room for objective-level references.
	MemoryStack param.Field[[]MemoryReferenceParam] `json:"memoryStack"`
	// Secrets that can be used in the headers for tool calls using the secret
	// interpolation format.
	Secrets param.Field[[]ObjectiveDataSecretParam] `json:"secrets"`
}

func (r ObjectiveDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveDataSecret struct {
	Name  string                  `json:"name"`
	Value string                  `json:"value"`
	JSON  objectiveDataSecretJSON `json:"-"`
}

// objectiveDataSecretJSON contains the JSON metadata for the struct
// [ObjectiveDataSecret]
type objectiveDataSecretJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveDataSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveDataSecretJSON) RawJSON() string {
	return r.raw
}

type ObjectiveDataSecretParam struct {
	Name  param.Field[string] `json:"name"`
	Value param.Field[string] `json:"value"`
}

func (r ObjectiveDataSecretParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveError struct {
	Message string             `json:"message"`
	Type    string             `json:"type"`
	JSON    objectiveErrorJSON `json:"-"`
}

// objectiveErrorJSON contains the JSON metadata for the struct [ObjectiveError]
type objectiveErrorJSON struct {
	Message     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveErrorJSON) RawJSON() string {
	return r.raw
}

type ObjectiveEventData struct {
	AssistantMessage       AssistantMessage       `json:"assistantMessage"`
	ContextWindowCompacted ContextWindowCompacted `json:"contextWindowCompacted"`
	Error                  ObjectiveError         `json:"error"`
	// MemoryRead is emitted each time the agent resolves a key against the memory
	// stack and loads an entry. Lookups that miss (key not found in any layer) do not
	// emit this event.
	MemoryRead            MemoryRead             `json:"memoryRead"`
	SubObjectiveCreated   SubObjectiveCreated    `json:"subObjectiveCreated"`
	ToolApprovalRequested ToolApprovalRequested  `json:"toolApprovalRequested"`
	ToolApproved          ToolApproved           `json:"toolApproved"`
	ToolCalled            ToolCalled             `json:"toolCalled"`
	ToolDenied            ToolDenied             `json:"toolDenied"`
	ToolError             ToolError              `json:"toolError"`
	ToolResult            ToolResult             `json:"toolResult"`
	Type                  string                 `json:"type"`
	UserMessage           UserMessage            `json:"userMessage"`
	JSON                  objectiveEventDataJSON `json:"-"`
}

// objectiveEventDataJSON contains the JSON metadata for the struct
// [ObjectiveEventData]
type objectiveEventDataJSON struct {
	AssistantMessage       apijson.Field
	ContextWindowCompacted apijson.Field
	Error                  apijson.Field
	MemoryRead             apijson.Field
	SubObjectiveCreated    apijson.Field
	ToolApprovalRequested  apijson.Field
	ToolApproved           apijson.Field
	ToolCalled             apijson.Field
	ToolDenied             apijson.Field
	ToolError              apijson.Field
	ToolResult             apijson.Field
	Type                   apijson.Field
	UserMessage            apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ObjectiveEventData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveEventDataJSON) RawJSON() string {
	return r.raw
}

type ObjectiveEventInfo struct {
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy Profile `json:"createdBy"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Objective shared.OperationMetadata `json:"objective"`
	JSON      objectiveEventInfoJSON   `json:"-"`
}

// objectiveEventInfoJSON contains the JSON metadata for the struct
// [ObjectiveEventInfo]
type objectiveEventInfoJSON struct {
	CreatedBy   apijson.Field
	Objective   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveEventInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveEventInfoJSON) RawJSON() string {
	return r.raw
}

// ObjectiveInfo provides read-only aggregated statistics about an objective's
// execution
type ObjectiveInfo struct {
	// List of callable tools assigned to the agent for this objective Includes tools,
	// agents, and cadenya-provided tools from the agent's configuration
	CallableTools []CallableTool `json:"callableTools"`
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy Profile `json:"createdBy"`
	// The effective memory stack at objective creation time, flattened from the
	// variation's baseline plus ObjectiveData.memory_stack. Order is push order (last
	// = top). Returned on reads so clients can see exactly what stack the objective is
	// using without having to re-join variation state.
	EffectiveMemoryStack []MemoryReference `json:"effectiveMemoryStack"`
	// Total number of context windows that this objective has generated
	TotalContextWindows int64 `json:"totalContextWindows"`
	// Total number of events generated during this objective's execution
	TotalEvents int64 `json:"totalEvents"`
	// Total input tokens consumed across all LLM completions across all context
	// windows
	TotalInputTokens int64 `json:"totalInputTokens"`
	// Total output tokens generated across all LLM completions across all context
	// windows
	TotalOutputTokens int64 `json:"totalOutputTokens"`
	// Total number of tool calls made during execution
	TotalToolCalls int64             `json:"totalToolCalls"`
	JSON           objectiveInfoJSON `json:"-"`
}

// objectiveInfoJSON contains the JSON metadata for the struct [ObjectiveInfo]
type objectiveInfoJSON struct {
	CallableTools        apijson.Field
	CreatedBy            apijson.Field
	EffectiveMemoryStack apijson.Field
	TotalContextWindows  apijson.Field
	TotalEvents          apijson.Field
	TotalInputTokens     apijson.Field
	TotalOutputTokens    apijson.Field
	TotalToolCalls       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ObjectiveInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveInfoJSON) RawJSON() string {
	return r.raw
}

type ObjectiveStatus struct {
	State   ObjectiveStatusState `json:"state" api:"required"`
	Message string               `json:"message"`
	JSON    objectiveStatusJSON  `json:"-"`
}

// objectiveStatusJSON contains the JSON metadata for the struct [ObjectiveStatus]
type objectiveStatusJSON struct {
	State       apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveStatusJSON) RawJSON() string {
	return r.raw
}

type ObjectiveStatusState string

const (
	ObjectiveStatusStateStateUnspecified ObjectiveStatusState = "STATE_UNSPECIFIED"
	ObjectiveStatusStateStatePending     ObjectiveStatusState = "STATE_PENDING"
	ObjectiveStatusStateStateRunning     ObjectiveStatusState = "STATE_RUNNING"
	ObjectiveStatusStateStateCompleted   ObjectiveStatusState = "STATE_COMPLETED"
	ObjectiveStatusStateStateFailed      ObjectiveStatusState = "STATE_FAILED"
	ObjectiveStatusStateStateCancelled   ObjectiveStatusState = "STATE_CANCELLED"
)

func (r ObjectiveStatusState) IsKnown() bool {
	switch r {
	case ObjectiveStatusStateStateUnspecified, ObjectiveStatusStateStatePending, ObjectiveStatusStateStateRunning, ObjectiveStatusStateStateCompleted, ObjectiveStatusStateStateFailed, ObjectiveStatusStateStateCancelled:
		return true
	}
	return false
}

type SubObjectiveCreated struct {
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata"`
	JSON     subObjectiveCreatedJSON  `json:"-"`
}

// subObjectiveCreatedJSON contains the JSON metadata for the struct
// [SubObjectiveCreated]
type subObjectiveCreatedJSON struct {
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubObjectiveCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subObjectiveCreatedJSON) RawJSON() string {
	return r.raw
}

type ToolApprovalRequested struct {
	// The ID of the objective tool call record. Use this ID with the ApproveToolCall
	// or DenyToolCall RPCs to approve or deny the tool call.
	ToolCallID string                    `json:"toolCallId"`
	JSON       toolApprovalRequestedJSON `json:"-"`
}

// toolApprovalRequestedJSON contains the JSON metadata for the struct
// [ToolApprovalRequested]
type toolApprovalRequestedJSON struct {
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolApprovalRequested) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolApprovalRequestedJSON) RawJSON() string {
	return r.raw
}

type ToolApproved struct {
	// The ID of the objective tool call record that was approved via the
	// ApproveToolCall RPC.
	ToolCallID string           `json:"toolCallId"`
	JSON       toolApprovedJSON `json:"-"`
}

// toolApprovedJSON contains the JSON metadata for the struct [ToolApproved]
type toolApprovedJSON struct {
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolApproved) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolApprovedJSON) RawJSON() string {
	return r.raw
}

type ToolCalled struct {
	// The ID of the objective tool call record that was executed.
	ToolCallID string         `json:"toolCallId"`
	JSON       toolCalledJSON `json:"-"`
}

// toolCalledJSON contains the JSON metadata for the struct [ToolCalled]
type toolCalledJSON struct {
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolCalled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolCalledJSON) RawJSON() string {
	return r.raw
}

type ToolDenied struct {
	// The memo provided by the reviewer when denying the tool call. This is passed to
	// the agent to provide further instructions.
	Memo string `json:"memo"`
	// The ID of the objective tool call record that was denied via the DenyToolCall
	// RPC.
	ToolCallID string         `json:"toolCallId"`
	JSON       toolDeniedJSON `json:"-"`
}

// toolDeniedJSON contains the JSON metadata for the struct [ToolDenied]
type toolDeniedJSON struct {
	Memo        apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolDenied) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolDeniedJSON) RawJSON() string {
	return r.raw
}

type ToolError struct {
	Message string `json:"message"`
	// The ID of the objective tool call record that encountered an error during
	// execution.
	ToolCallID string        `json:"toolCallId"`
	JSON       toolErrorJSON `json:"-"`
}

// toolErrorJSON contains the JSON metadata for the struct [ToolError]
type toolErrorJSON struct {
	Message     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolErrorJSON) RawJSON() string {
	return r.raw
}

type ToolResult struct {
	Content    string         `json:"content"`
	ToolCallID string         `json:"toolCallId"`
	JSON       toolResultJSON `json:"-"`
}

// toolResultJSON contains the JSON metadata for the struct [ToolResult]
type toolResultJSON struct {
	Content     apijson.Field
	ToolCallID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolResultJSON) RawJSON() string {
	return r.raw
}

type UserMessage struct {
	Content string          `json:"content"`
	JSON    userMessageJSON `json:"-"`
}

// userMessageJSON contains the JSON metadata for the struct [UserMessage]
type userMessageJSON struct {
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userMessageJSON) RawJSON() string {
	return r.raw
}

// Compact objective response
type ObjectiveCompactResponse struct {
	// The new context window created by the compaction
	ContextWindow ObjectiveContextWindowData   `json:"contextWindow"`
	JSON          objectiveCompactResponseJSON `json:"-"`
}

// objectiveCompactResponseJSON contains the JSON metadata for the struct
// [ObjectiveCompactResponse]
type objectiveCompactResponseJSON struct {
	ContextWindow apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObjectiveCompactResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveCompactResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectiveContinueResponse struct {
	Data ObjectiveEventData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata        shared.OperationMetadata      `json:"metadata" api:"required"`
	ContextWindowID string                        `json:"contextWindowId"`
	Info            ObjectiveEventInfo            `json:"info"`
	JSON            objectiveContinueResponseJSON `json:"-"`
}

// objectiveContinueResponseJSON contains the JSON metadata for the struct
// [ObjectiveContinueResponse]
type objectiveContinueResponseJSON struct {
	Data            apijson.Field
	Metadata        apijson.Field
	ContextWindowID apijson.Field
	Info            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObjectiveContinueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveContinueResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectiveListEventsResponse struct {
	Data ObjectiveEventData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata        shared.OperationMetadata        `json:"metadata" api:"required"`
	ContextWindowID string                          `json:"contextWindowId"`
	Info            ObjectiveEventInfo              `json:"info"`
	JSON            objectiveListEventsResponseJSON `json:"-"`
}

// objectiveListEventsResponseJSON contains the JSON metadata for the struct
// [ObjectiveListEventsResponse]
type objectiveListEventsResponseJSON struct {
	Data            apijson.Field
	Metadata        apijson.Field
	ContextWindowID apijson.Field
	Info            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObjectiveListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveListEventsResponseJSON) RawJSON() string {
	return r.raw
}

type ObjectiveNewParams struct {
	AgentID param.Field[string]             `json:"agentId" api:"required"`
	Data    param.Field[ObjectiveDataParam] `json:"data" api:"required"`
	// CreateOperationMetadata contains the user-provided fields for creating an
	// operation. Read-only fields (id, account_id, workspace_id, created_at,
	// profile_id) are excluded since they are set by the server.
	Metadata param.Field[shared.CreateOperationMetadataParam] `json:"metadata" api:"required"`
	// Optional explicit variation selection. Overrides the agent's
	// variation_selection_mode.
	VariationID param.Field[string] `json:"variationId"`
}

func (r ObjectiveNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveListParams struct {
	// Agent ID for filtering
	AgentID param.Field[string] `query:"agentId"`
	// Filter to objectives produced by a specific AgentSchedule. Matches
	// ObjectiveData.source*schedule_id. Accepts canonical as*… form or
	// external_id:<value> form (see common.proto "Path-parameter ID resolution").
	AgentScheduleID param.Field[string] `query:"agentScheduleId"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Optional filters
	ParentObjectiveID param.Field[string] `query:"parentObjectiveId"`
	ProfileID         param.Field[string] `query:"profileId"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by state
	State param.Field[ObjectiveListParamsState] `query:"state"`
}

// URLQuery serializes [ObjectiveListParams]'s query parameters as `url.Values`.
func (r ObjectiveListParams) URLQuery() (v url.Values) {
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
	ObjectiveListParamsStateStateCompleted   ObjectiveListParamsState = "STATE_COMPLETED"
	ObjectiveListParamsStateStateFailed      ObjectiveListParamsState = "STATE_FAILED"
	ObjectiveListParamsStateStateCancelled   ObjectiveListParamsState = "STATE_CANCELLED"
)

func (r ObjectiveListParamsState) IsKnown() bool {
	switch r {
	case ObjectiveListParamsStateStateUnspecified, ObjectiveListParamsStateStatePending, ObjectiveListParamsStateStateRunning, ObjectiveListParamsStateStateCompleted, ObjectiveListParamsStateStateFailed, ObjectiveListParamsStateStateCancelled:
		return true
	}
	return false
}

type ObjectiveCancelParams struct {
	// Optional reason for cancellation
	Reason param.Field[string] `json:"reason"`
}

func (r ObjectiveCancelParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveCompactParams struct {
	// CompactionConfig defines how context window compaction behaves for objectives
	// using this variation.
	CompactionConfig param.Field[AgentVariationSpecCompactionConfigParam] `json:"compactionConfig"`
}

func (r ObjectiveCompactParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveContinueParams struct {
	// When set to true, the message will be enqueued for when the agent loop is
	// available to process it.
	Enqueue param.Field[bool] `json:"enqueue"`
	// The message to continue an objective that has completed (or you are enqueing)
	Message param.Field[string] `json:"message"`
	// Secrets that should be included with the message. Helpful for when you need to
	// update secrets on the objective (IE: A secret expires and needs to be refreshed)
	Secrets param.Field[[]ObjectiveContinueParamsSecret] `json:"secrets"`
}

func (r ObjectiveContinueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveContinueParamsSecret struct {
	Name  param.Field[string] `json:"name"`
	Value param.Field[string] `json:"value"`
}

func (r ObjectiveContinueParamsSecret) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObjectiveListContextWindowsParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [ObjectiveListContextWindowsParams]'s query parameters as
// `url.Values`.
func (r ObjectiveListContextWindowsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ObjectiveListEventsParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// When set to true you may use more of your alloted API rate-limit
	IncludeInfo param.Field[bool] `query:"includeInfo"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Optional context window ID to filter events by
	WindowID param.Field[string] `query:"windowId"`
}

// URLQuery serializes [ObjectiveListEventsParams]'s query parameters as
// `url.Values`.
func (r ObjectiveListEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
