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

// BulkWorkspaceResources is the workspace-scoped service that applies a
// declarative bundle of workspace resources (tool sets, memory layers, agents,
// variations, assignments, schedules) in one async operation. See
// docs/superpowers/specs/2026-05-02-bulk-workspace-resources-design.md for the
// full design.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// BulkWorkspaceResourceService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkWorkspaceResourceService] method instead.
type BulkWorkspaceResourceService struct {
	Options []option.RequestOption
	// BulkWorkspaceResources is the workspace-scoped service that applies a
	// declarative bundle of workspace resources (tool sets, memory layers, agents,
	// variations, assignments, schedules) in one async operation. See
	// docs/superpowers/specs/2026-05-02-bulk-workspace-resources-design.md for the
	// full design.
	//
	// Authentication: Bearer token (JWT) Scope: Workspace-level operations
	Results *BulkWorkspaceResourceResultService
}

// NewBulkWorkspaceResourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBulkWorkspaceResourceService(opts ...option.RequestOption) (r *BulkWorkspaceResourceService) {
	r = &BulkWorkspaceResourceService{}
	r.Options = opts
	r.Results = NewBulkWorkspaceResourceResultService(opts...)
	return
}

// Retrieves a bulk workspace apply operation by ID.
func (r *BulkWorkspaceResourceService) Get(ctx context.Context, workspaceID string, id string, opts ...option.RequestOption) (res *BulkWorkspaceApply, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/bulk_workspace_applies/%s", workspaceID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists past and in-flight bulk workspace apply operations in the workspace.
func (r *BulkWorkspaceResourceService) List(ctx context.Context, workspaceID string, query BulkWorkspaceResourceListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[BulkWorkspaceApply], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/bulk_workspace_applies", workspaceID)
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

// Lists past and in-flight bulk workspace apply operations in the workspace.
func (r *BulkWorkspaceResourceService) ListAutoPaging(ctx context.Context, workspaceID string, query BulkWorkspaceResourceListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[BulkWorkspaceApply] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, workspaceID, query, opts...))
}

// Asynchronously applies a declarative bundle of workspace resources. Returns the
// operation immediately in PENDING; clients poll Get to track progress.
func (r *BulkWorkspaceResourceService) Apply(ctx context.Context, workspaceID string, body BulkWorkspaceResourceApplyParams, opts ...option.RequestOption) (res *BulkWorkspaceApply, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceID == "" {
		err = errors.New("missing required workspaceId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/bulk_workspace_applies", workspaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AgentEntry struct {
	Name string `json:"name" api:"required"`
	// Agent specification (user-provided configuration)
	Spec   AgentSpec         `json:"spec" api:"required"`
	Labels map[string]string `json:"labels"`
	// Schedules under this agent, keyed by external_id.
	Schedules map[string]AgentScheduleEntry `json:"schedules"`
	// Variations under this agent, keyed by external_id.
	Variations map[string]AgentVariationEntry `json:"variations"`
	JSON       agentEntryJSON                 `json:"-"`
}

// agentEntryJSON contains the JSON metadata for the struct [AgentEntry]
type agentEntryJSON struct {
	Name        apijson.Field
	Spec        apijson.Field
	Labels      apijson.Field
	Schedules   apijson.Field
	Variations  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentEntryJSON) RawJSON() string {
	return r.raw
}

type AgentEntryParam struct {
	Name param.Field[string] `json:"name" api:"required"`
	// Agent specification (user-provided configuration)
	Spec   param.Field[AgentSpecParam]    `json:"spec" api:"required"`
	Labels param.Field[map[string]string] `json:"labels"`
	// Schedules under this agent, keyed by external_id.
	Schedules param.Field[map[string]AgentScheduleEntryParam] `json:"schedules"`
	// Variations under this agent, keyed by external_id.
	Variations param.Field[map[string]AgentVariationEntryParam] `json:"variations"`
}

func (r AgentEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentScheduleEntry struct {
	Name string `json:"name" api:"required"`
	// AgentScheduleSpec is the user-provided configuration for a schedule.
	Spec   AgentScheduleSpec      `json:"spec" api:"required"`
	Labels map[string]string      `json:"labels"`
	JSON   agentScheduleEntryJSON `json:"-"`
}

// agentScheduleEntryJSON contains the JSON metadata for the struct
// [AgentScheduleEntry]
type agentScheduleEntryJSON struct {
	Name        apijson.Field
	Spec        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentScheduleEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentScheduleEntryJSON) RawJSON() string {
	return r.raw
}

type AgentScheduleEntryParam struct {
	Name param.Field[string] `json:"name" api:"required"`
	// AgentScheduleSpec is the user-provided configuration for a schedule.
	Spec   param.Field[AgentScheduleSpecParam] `json:"spec" api:"required"`
	Labels param.Field[map[string]string]      `json:"labels"`
}

func (r AgentScheduleEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AgentVariationEntry struct {
	Name string `json:"name" api:"required"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec AgentVariationSpec `json:"spec" api:"required"`
	// Reconciled list — server adjusts the variation's assignments to exactly this set
	// when the variation is bundle-owned.
	Assignments []VariationAssignmentEntry `json:"assignments"`
	Labels      map[string]string          `json:"labels"`
	// Reconciled list — capped at 10 to match the existing variation
	// memory-layer-assignment cap.
	MemoryLayers []VariationMemoryLayerEntry `json:"memoryLayers"`
	JSON         agentVariationEntryJSON     `json:"-"`
}

// agentVariationEntryJSON contains the JSON metadata for the struct
// [AgentVariationEntry]
type agentVariationEntryJSON struct {
	Name         apijson.Field
	Spec         apijson.Field
	Assignments  apijson.Field
	Labels       apijson.Field
	MemoryLayers apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AgentVariationEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentVariationEntryJSON) RawJSON() string {
	return r.raw
}

type AgentVariationEntryParam struct {
	Name param.Field[string] `json:"name" api:"required"`
	// AgentVariationSpec defines the operational configuration for a variation
	Spec param.Field[AgentVariationSpecParam] `json:"spec" api:"required"`
	// Reconciled list — server adjusts the variation's assignments to exactly this set
	// when the variation is bundle-owned.
	Assignments param.Field[[]VariationAssignmentEntryParam] `json:"assignments"`
	Labels      param.Field[map[string]string]               `json:"labels"`
	// Reconciled list — capped at 10 to match the existing variation
	// memory-layer-assignment cap.
	MemoryLayers param.Field[[]VariationMemoryLayerEntryParam] `json:"memoryLayers"`
}

func (r AgentVariationEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// BulkWorkspaceApply is the operation resource produced by a call to
// BulkWorkspaceResources.Apply. It is operation-typed (uses OperationMetadata,
// like Objective and ObjectiveEvent) and carries the input bundle in `data`, the
// lifecycle state in `status`, and aggregate counts in `info`.
type BulkWorkspaceApply struct {
	Data BulkWorkspaceApplyData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata `json:"metadata" api:"required"`
	Status   BulkWorkspaceApplyStatus `json:"status" api:"required"`
	Info     BulkWorkspaceApplyInfo   `json:"info"`
	JSON     bulkWorkspaceApplyJSON   `json:"-"`
}

// bulkWorkspaceApplyJSON contains the JSON metadata for the struct
// [BulkWorkspaceApply]
type bulkWorkspaceApplyJSON struct {
	Data        apijson.Field
	Metadata    apijson.Field
	Status      apijson.Field
	Info        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApply) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyData struct {
	// Required. Bundle ownership key. Resources created or updated by an Apply have
	// their `metadata.bundle_key` set to this value. On subsequent applies with the
	// same bundle_key, resources currently bearing this bundle_key but absent from the
	// spec are soft-deleted.
	BundleKey string `json:"bundleKey" api:"required"`
	// Agents to upsert, keyed by external_id.
	Agents map[string]AgentEntry `json:"agents"`
	// When true, every agent created or updated by this Apply has its status forced to
	// AGENT_STATUS_PUBLISHED, regardless of the status declared in the agent's
	// AgentSpec. Useful when the bundle represents a production configuration and you
	// want all of its agents live without setting status: AGENT_STATUS_PUBLISHED on
	// each entry.
	//
	// Default false: each agent's AgentSpec.status controls (which is
	// AGENT_STATUS_DRAFT on create when unspecified).
	AutomaticallyPublishAgents bool `json:"automaticallyPublishAgents"`
	// Memory layers to upsert, keyed by external_id.
	MemoryLayers map[string]MemoryLayerEntry `json:"memoryLayers"`
	// Optional URL pointing to the source of this apply (GitHub PR, Jenkins build,
	// GitLab pipeline, etc.). Surfaced in the dashboard so users can jump from an
	// apply back to the change that produced it. Free-form HTTPS URI; not interpreted
	// by the server.
	SourceURL string `json:"sourceUrl"`
	// Tool sets to upsert, keyed by external_id.
	ToolSets map[string]ToolSetEntry    `json:"toolSets"`
	JSON     bulkWorkspaceApplyDataJSON `json:"-"`
}

// bulkWorkspaceApplyDataJSON contains the JSON metadata for the struct
// [BulkWorkspaceApplyData]
type bulkWorkspaceApplyDataJSON struct {
	BundleKey                  apijson.Field
	Agents                     apijson.Field
	AutomaticallyPublishAgents apijson.Field
	MemoryLayers               apijson.Field
	SourceURL                  apijson.Field
	ToolSets                   apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *BulkWorkspaceApplyData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyDataJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyDataParam struct {
	// Required. Bundle ownership key. Resources created or updated by an Apply have
	// their `metadata.bundle_key` set to this value. On subsequent applies with the
	// same bundle_key, resources currently bearing this bundle_key but absent from the
	// spec are soft-deleted.
	BundleKey param.Field[string] `json:"bundleKey" api:"required"`
	// Agents to upsert, keyed by external_id.
	Agents param.Field[map[string]AgentEntryParam] `json:"agents"`
	// When true, every agent created or updated by this Apply has its status forced to
	// AGENT_STATUS_PUBLISHED, regardless of the status declared in the agent's
	// AgentSpec. Useful when the bundle represents a production configuration and you
	// want all of its agents live without setting status: AGENT_STATUS_PUBLISHED on
	// each entry.
	//
	// Default false: each agent's AgentSpec.status controls (which is
	// AGENT_STATUS_DRAFT on create when unspecified).
	AutomaticallyPublishAgents param.Field[bool] `json:"automaticallyPublishAgents"`
	// Memory layers to upsert, keyed by external_id.
	MemoryLayers param.Field[map[string]MemoryLayerEntryParam] `json:"memoryLayers"`
	// Optional URL pointing to the source of this apply (GitHub PR, Jenkins build,
	// GitLab pipeline, etc.). Surfaced in the dashboard so users can jump from an
	// apply back to the change that produced it. Free-form HTTPS URI; not interpreted
	// by the server.
	SourceURL param.Field[string] `json:"sourceUrl"`
	// Tool sets to upsert, keyed by external_id.
	ToolSets param.Field[map[string]ToolSetEntryParam] `json:"toolSets"`
}

func (r BulkWorkspaceApplyDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkWorkspaceApplyInfo struct {
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// Profile represents a human user at the account level. Profiles are
	// account-scoped resources that can be associated with multiple workspaces through
	// the Actor model. Authentication for profiles is handled via SSO/OAuth (WorkOS).
	CreatedBy      Profile                    `json:"createdBy"`
	CreatedCount   int64                      `json:"createdCount"`
	DeletedCount   int64                      `json:"deletedCount"`
	FailedCount    int64                      `json:"failedCount"`
	StartedAt      time.Time                  `json:"startedAt" format:"date-time"`
	TotalCount     int64                      `json:"totalCount"`
	UnchangedCount int64                      `json:"unchangedCount"`
	UpdatedCount   int64                      `json:"updatedCount"`
	JSON           bulkWorkspaceApplyInfoJSON `json:"-"`
}

// bulkWorkspaceApplyInfoJSON contains the JSON metadata for the struct
// [BulkWorkspaceApplyInfo]
type bulkWorkspaceApplyInfoJSON struct {
	CompletedAt    apijson.Field
	CreatedBy      apijson.Field
	CreatedCount   apijson.Field
	DeletedCount   apijson.Field
	FailedCount    apijson.Field
	StartedAt      apijson.Field
	TotalCount     apijson.Field
	UnchangedCount apijson.Field
	UpdatedCount   apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BulkWorkspaceApplyInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyInfoJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyStatus struct {
	State   BulkWorkspaceApplyStatusState `json:"state" api:"required"`
	Message string                        `json:"message"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	PreflightError BulkWorkspaceApplyStatusPreflightError `json:"preflightError"`
	JSON           bulkWorkspaceApplyStatusJSON           `json:"-"`
}

// bulkWorkspaceApplyStatusJSON contains the JSON metadata for the struct
// [BulkWorkspaceApplyStatus]
type bulkWorkspaceApplyStatusJSON struct {
	State          apijson.Field
	Message        apijson.Field
	PreflightError apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BulkWorkspaceApplyStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyStatusJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyStatusState string

const (
	BulkWorkspaceApplyStatusStateStateUnspecified      BulkWorkspaceApplyStatusState = "STATE_UNSPECIFIED"
	BulkWorkspaceApplyStatusStateStatePending          BulkWorkspaceApplyStatusState = "STATE_PENDING"
	BulkWorkspaceApplyStatusStateStateValidating       BulkWorkspaceApplyStatusState = "STATE_VALIDATING"
	BulkWorkspaceApplyStatusStateStateRunning          BulkWorkspaceApplyStatusState = "STATE_RUNNING"
	BulkWorkspaceApplyStatusStateStateSucceeded        BulkWorkspaceApplyStatusState = "STATE_SUCCEEDED"
	BulkWorkspaceApplyStatusStateStatePartiallyApplied BulkWorkspaceApplyStatusState = "STATE_PARTIALLY_APPLIED"
	BulkWorkspaceApplyStatusStateStateFailed           BulkWorkspaceApplyStatusState = "STATE_FAILED"
	BulkWorkspaceApplyStatusStateStateCancelled        BulkWorkspaceApplyStatusState = "STATE_CANCELLED"
)

func (r BulkWorkspaceApplyStatusState) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyStatusStateStateUnspecified, BulkWorkspaceApplyStatusStateStatePending, BulkWorkspaceApplyStatusStateStateValidating, BulkWorkspaceApplyStatusStateStateRunning, BulkWorkspaceApplyStatusStateStateSucceeded, BulkWorkspaceApplyStatusStateStatePartiallyApplied, BulkWorkspaceApplyStatusStateStateFailed, BulkWorkspaceApplyStatusStateStateCancelled:
		return true
	}
	return false
}

// The `Status` type defines a logical error model that is suitable for different
// programming environments, including REST APIs and RPC APIs. It is used by
// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
// data: error code, error message, and error details. You can find out more about
// this error model and how to work with it in the
// [API Design Guide](https://cloud.google.com/apis/design/errors).
type BulkWorkspaceApplyStatusPreflightError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyStatusPreflightErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                     `json:"message"`
	JSON    bulkWorkspaceApplyStatusPreflightErrorJSON `json:"-"`
}

// bulkWorkspaceApplyStatusPreflightErrorJSON contains the JSON metadata for the
// struct [BulkWorkspaceApplyStatusPreflightError]
type bulkWorkspaceApplyStatusPreflightErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyStatusPreflightError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyStatusPreflightErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyStatusPreflightErrorDetail struct {
	// The type of the serialized message.
	Type        string                                           `json:"@type"`
	ExtraFields map[string]interface{}                           `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyStatusPreflightErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyStatusPreflightErrorDetailJSON contains the JSON metadata for
// the struct [BulkWorkspaceApplyStatusPreflightErrorDetail]
type bulkWorkspaceApplyStatusPreflightErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyStatusPreflightErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyStatusPreflightErrorDetailJSON) RawJSON() string {
	return r.raw
}

type MemoryEntryItem struct {
	Key         string              `json:"key" api:"required"`
	Content     string              `json:"content"`
	Description string              `json:"description"`
	UploadID    string              `json:"uploadId"`
	JSON        memoryEntryItemJSON `json:"-"`
}

// memoryEntryItemJSON contains the JSON metadata for the struct [MemoryEntryItem]
type memoryEntryItemJSON struct {
	Key         apijson.Field
	Content     apijson.Field
	Description apijson.Field
	UploadID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryEntryItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryEntryItemJSON) RawJSON() string {
	return r.raw
}

type MemoryEntryItemParam struct {
	Key         param.Field[string] `json:"key" api:"required"`
	Content     param.Field[string] `json:"content"`
	Description param.Field[string] `json:"description"`
	UploadID    param.Field[string] `json:"uploadId"`
}

func (r MemoryEntryItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemoryLayerEntry struct {
	Name string          `json:"name" api:"required"`
	Spec MemoryLayerSpec `json:"spec" api:"required"`
	// Memory entries in this layer, keyed by external_id.
	Entries map[string]MemoryEntryItem `json:"entries"`
	Labels  map[string]string          `json:"labels"`
	JSON    memoryLayerEntryJSON       `json:"-"`
}

// memoryLayerEntryJSON contains the JSON metadata for the struct
// [MemoryLayerEntry]
type memoryLayerEntryJSON struct {
	Name        apijson.Field
	Spec        apijson.Field
	Entries     apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemoryLayerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memoryLayerEntryJSON) RawJSON() string {
	return r.raw
}

type MemoryLayerEntryParam struct {
	Name param.Field[string]               `json:"name" api:"required"`
	Spec param.Field[MemoryLayerSpecParam] `json:"spec" api:"required"`
	// Memory entries in this layer, keyed by external_id.
	Entries param.Field[map[string]MemoryEntryItemParam] `json:"entries"`
	Labels  param.Field[map[string]string]               `json:"labels"`
}

func (r MemoryLayerEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolEntry struct {
	Name   string            `json:"name" api:"required"`
	Spec   ToolSpec          `json:"spec" api:"required"`
	Labels map[string]string `json:"labels"`
	JSON   toolEntryJSON     `json:"-"`
}

// toolEntryJSON contains the JSON metadata for the struct [ToolEntry]
type toolEntryJSON struct {
	Name        apijson.Field
	Spec        apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolEntryJSON) RawJSON() string {
	return r.raw
}

type ToolEntryParam struct {
	Name   param.Field[string]            `json:"name" api:"required"`
	Spec   param.Field[ToolSpecParam]     `json:"spec" api:"required"`
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r ToolEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolSetEntry struct {
	Name   string            `json:"name" api:"required"`
	Spec   ToolSetSpec       `json:"spec" api:"required"`
	Labels map[string]string `json:"labels"`
	// Tools in this tool set, keyed by external_id.
	Tools map[string]ToolEntry `json:"tools"`
	JSON  toolSetEntryJSON     `json:"-"`
}

// toolSetEntryJSON contains the JSON metadata for the struct [ToolSetEntry]
type toolSetEntryJSON struct {
	Name        apijson.Field
	Spec        apijson.Field
	Labels      apijson.Field
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolSetEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolSetEntryJSON) RawJSON() string {
	return r.raw
}

type ToolSetEntryParam struct {
	Name   param.Field[string]            `json:"name" api:"required"`
	Spec   param.Field[ToolSetSpecParam]  `json:"spec" api:"required"`
	Labels param.Field[map[string]string] `json:"labels"`
	// Tools in this tool set, keyed by external_id.
	Tools param.Field[map[string]ToolEntryParam] `json:"tools"`
}

func (r ToolSetEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VariationAssignmentEntry struct {
	SubAgentID string                       `json:"subAgentId"`
	ToolID     string                       `json:"toolId"`
	ToolSetID  string                       `json:"toolSetId"`
	JSON       variationAssignmentEntryJSON `json:"-"`
}

// variationAssignmentEntryJSON contains the JSON metadata for the struct
// [VariationAssignmentEntry]
type variationAssignmentEntryJSON struct {
	SubAgentID  apijson.Field
	ToolID      apijson.Field
	ToolSetID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VariationAssignmentEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variationAssignmentEntryJSON) RawJSON() string {
	return r.raw
}

type VariationAssignmentEntryParam struct {
	SubAgentID param.Field[string] `json:"subAgentId"`
	ToolID     param.Field[string] `json:"toolId"`
	ToolSetID  param.Field[string] `json:"toolSetId"`
}

func (r VariationAssignmentEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VariationMemoryLayerEntry struct {
	// external_id:<value> form. Canonical IDs are rejected.
	MemoryLayerID string                        `json:"memoryLayerId"`
	Position      int64                         `json:"position"`
	JSON          variationMemoryLayerEntryJSON `json:"-"`
}

// variationMemoryLayerEntryJSON contains the JSON metadata for the struct
// [VariationMemoryLayerEntry]
type variationMemoryLayerEntryJSON struct {
	MemoryLayerID apijson.Field
	Position      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *VariationMemoryLayerEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r variationMemoryLayerEntryJSON) RawJSON() string {
	return r.raw
}

type VariationMemoryLayerEntryParam struct {
	// external_id:<value> form. Canonical IDs are rejected.
	MemoryLayerID param.Field[string] `json:"memoryLayerId"`
	Position      param.Field[int64]  `json:"position"`
}

func (r VariationMemoryLayerEntryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BulkWorkspaceResourceListParams struct {
	// Filter by bundle_key — list every apply for a given bundle.
	BundleKey param.Field[string] `query:"bundleKey"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by lifecycle state.
	State param.Field[BulkWorkspaceResourceListParamsState] `query:"state"`
}

// URLQuery serializes [BulkWorkspaceResourceListParams]'s query parameters as
// `url.Values`.
func (r BulkWorkspaceResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by lifecycle state.
type BulkWorkspaceResourceListParamsState string

const (
	BulkWorkspaceResourceListParamsStateStateUnspecified      BulkWorkspaceResourceListParamsState = "STATE_UNSPECIFIED"
	BulkWorkspaceResourceListParamsStateStatePending          BulkWorkspaceResourceListParamsState = "STATE_PENDING"
	BulkWorkspaceResourceListParamsStateStateValidating       BulkWorkspaceResourceListParamsState = "STATE_VALIDATING"
	BulkWorkspaceResourceListParamsStateStateRunning          BulkWorkspaceResourceListParamsState = "STATE_RUNNING"
	BulkWorkspaceResourceListParamsStateStateSucceeded        BulkWorkspaceResourceListParamsState = "STATE_SUCCEEDED"
	BulkWorkspaceResourceListParamsStateStatePartiallyApplied BulkWorkspaceResourceListParamsState = "STATE_PARTIALLY_APPLIED"
	BulkWorkspaceResourceListParamsStateStateFailed           BulkWorkspaceResourceListParamsState = "STATE_FAILED"
	BulkWorkspaceResourceListParamsStateStateCancelled        BulkWorkspaceResourceListParamsState = "STATE_CANCELLED"
)

func (r BulkWorkspaceResourceListParamsState) IsKnown() bool {
	switch r {
	case BulkWorkspaceResourceListParamsStateStateUnspecified, BulkWorkspaceResourceListParamsStateStatePending, BulkWorkspaceResourceListParamsStateStateValidating, BulkWorkspaceResourceListParamsStateStateRunning, BulkWorkspaceResourceListParamsStateStateSucceeded, BulkWorkspaceResourceListParamsStateStatePartiallyApplied, BulkWorkspaceResourceListParamsStateStateFailed, BulkWorkspaceResourceListParamsStateStateCancelled:
		return true
	}
	return false
}

type BulkWorkspaceResourceApplyParams struct {
	Data param.Field[BulkWorkspaceApplyDataParam] `json:"data" api:"required"`
}

func (r BulkWorkspaceResourceApplyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
