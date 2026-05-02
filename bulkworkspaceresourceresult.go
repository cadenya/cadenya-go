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

// BulkWorkspaceResources is the workspace-scoped service that applies a
// declarative bundle of workspace resources (tool sets, memory layers, agents,
// variations, assignments, schedules) in one async operation. See
// docs/superpowers/specs/2026-05-02-bulk-workspace-resources-design.md for the
// full design.
//
// Authentication: Bearer token (JWT) Scope: Workspace-level operations
//
// BulkWorkspaceResourceResultService contains methods and other services that help
// with interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkWorkspaceResourceResultService] method instead.
type BulkWorkspaceResourceResultService struct {
	Options []option.RequestOption
}

// NewBulkWorkspaceResourceResultService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewBulkWorkspaceResourceResultService(opts ...option.RequestOption) (r *BulkWorkspaceResourceResultService) {
	r = &BulkWorkspaceResourceResultService{}
	r.Options = opts
	return
}

// Lists each resource action recorded by a bulk workspace apply operation.
func (r *BulkWorkspaceResourceResultService) List(ctx context.Context, bulkWorkspaceApplyID string, query BulkWorkspaceResourceResultListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[BulkWorkspaceApplyResult], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if bulkWorkspaceApplyID == "" {
		err = errors.New("missing required bulkWorkspaceApplyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/bulk_workspace_applies/%s/results", bulkWorkspaceApplyID)
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

// Lists each resource action recorded by a bulk workspace apply operation.
func (r *BulkWorkspaceResourceResultService) ListAutoPaging(ctx context.Context, bulkWorkspaceApplyID string, query BulkWorkspaceResourceResultListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[BulkWorkspaceApplyResult] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, bulkWorkspaceApplyID, query, opts...))
}

// BulkWorkspaceApplyResult is one row of the per-resource result list for a
// BulkWorkspaceApply. Each row is itself an operation (OperationMetadata-typed) so
// it can be paginated, sorted by created_at, and individually addressed. Mirrors
// the Objective → ObjectiveEvent relationship.
type BulkWorkspaceApplyResult struct {
	// BulkWorkspaceApplyResultData carries the outcome for a single resource. The
	// `type` field is the discriminator string that names the populated `outcome`
	// oneof variant (e.g., "toolSet", "memoryEntry"). Every Outcome shell carries an
	// `action` enum and either a resulting resource snapshot (for ACTION_CREATED /
	// ACTION_UPDATED / ACTION_UNCHANGED / ACTION_DELETED) or a google.rpc.Status (for
	// ACTION_FAILED).
	Data BulkWorkspaceApplyResultData `json:"data" api:"required"`
	// Metadata for ephemeral operations and activities (e.g., objectives, executions,
	// runs)
	Metadata shared.OperationMetadata     `json:"metadata" api:"required"`
	JSON     bulkWorkspaceApplyResultJSON `json:"-"`
}

// bulkWorkspaceApplyResultJSON contains the JSON metadata for the struct
// [BulkWorkspaceApplyResult]
type bulkWorkspaceApplyResultJSON struct {
	Data        apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultJSON) RawJSON() string {
	return r.raw
}

// BulkWorkspaceApplyResultData carries the outcome for a single resource. The
// `type` field is the discriminator string that names the populated `outcome`
// oneof variant (e.g., "toolSet", "memoryEntry"). Every Outcome shell carries an
// `action` enum and either a resulting resource snapshot (for ACTION_CREATED /
// ACTION_UPDATED / ACTION_UNCHANGED / ACTION_DELETED) or a google.rpc.Status (for
// ACTION_FAILED).
type BulkWorkspaceApplyResultData struct {
	Agent                BulkWorkspaceApplyResultDataAgentOutcome                `json:"agent"`
	AgentSchedule        BulkWorkspaceApplyResultDataAgentScheduleOutcome        `json:"agentSchedule"`
	AgentVariation       BulkWorkspaceApplyResultDataAgentVariationOutcome       `json:"agentVariation"`
	MemoryEntry          BulkWorkspaceApplyResultDataMemoryEntryOutcome          `json:"memoryEntry"`
	MemoryLayer          BulkWorkspaceApplyResultDataMemoryLayerOutcome          `json:"memoryLayer"`
	Tool                 BulkWorkspaceApplyResultDataToolOutcome                 `json:"tool"`
	ToolSet              BulkWorkspaceApplyResultDataToolSetOutcome              `json:"toolSet"`
	Type                 string                                                  `json:"type"`
	VariationAssignment  BulkWorkspaceApplyResultDataVariationAssignmentOutcome  `json:"variationAssignment"`
	VariationMemoryLayer BulkWorkspaceApplyResultDataVariationMemoryLayerOutcome `json:"variationMemoryLayer"`
	JSON                 bulkWorkspaceApplyResultDataJSON                        `json:"-"`
}

// bulkWorkspaceApplyResultDataJSON contains the JSON metadata for the struct
// [BulkWorkspaceApplyResultData]
type bulkWorkspaceApplyResultDataJSON struct {
	Agent                apijson.Field
	AgentSchedule        apijson.Field
	AgentVariation       apijson.Field
	MemoryEntry          apijson.Field
	MemoryLayer          apijson.Field
	Tool                 apijson.Field
	ToolSet              apijson.Field
	Type                 apijson.Field
	VariationAssignment  apijson.Field
	VariationMemoryLayer apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataAgentOutcome struct {
	Action BulkWorkspaceApplyResultDataAgentOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error      BulkWorkspaceApplyResultDataAgentOutcomeError `json:"error"`
	ExternalID string                                        `json:"externalId"`
	// Agent resource
	Resource Agent                                        `json:"resource"`
	JSON     bulkWorkspaceApplyResultDataAgentOutcomeJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentOutcomeJSON contains the JSON metadata for the
// struct [BulkWorkspaceApplyResultDataAgentOutcome]
type bulkWorkspaceApplyResultDataAgentOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	ExternalID  apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataAgentOutcomeAction string

const (
	BulkWorkspaceApplyResultDataAgentOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataAgentOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataAgentOutcomeActionActionCreated     BulkWorkspaceApplyResultDataAgentOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataAgentOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataAgentOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataAgentOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataAgentOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataAgentOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataAgentOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataAgentOutcomeActionActionFailed      BulkWorkspaceApplyResultDataAgentOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataAgentOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataAgentOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataAgentOutcomeActionActionCreated, BulkWorkspaceApplyResultDataAgentOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataAgentOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataAgentOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataAgentOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataAgentOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataAgentOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                            `json:"message"`
	JSON    bulkWorkspaceApplyResultDataAgentOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentOutcomeErrorJSON contains the JSON metadata for
// the struct [BulkWorkspaceApplyResultDataAgentOutcomeError]
type bulkWorkspaceApplyResultDataAgentOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataAgentOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                  `json:"@type"`
	ExtraFields map[string]interface{}                                  `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataAgentOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentOutcomeErrorDetailJSON contains the JSON
// metadata for the struct [BulkWorkspaceApplyResultDataAgentOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataAgentOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataAgentScheduleOutcome struct {
	Action BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error      BulkWorkspaceApplyResultDataAgentScheduleOutcomeError `json:"error"`
	ExternalID string                                                `json:"externalId"`
	// AgentSchedule resource — a recurring trigger attached to an agent that creates
	// objectives on its cadence.
	Resource AgentSchedule                                        `json:"resource"`
	JSON     bulkWorkspaceApplyResultDataAgentScheduleOutcomeJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentScheduleOutcomeJSON contains the JSON metadata
// for the struct [BulkWorkspaceApplyResultDataAgentScheduleOutcome]
type bulkWorkspaceApplyResultDataAgentScheduleOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	ExternalID  apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentScheduleOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentScheduleOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction string

const (
	BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionCreated     BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionFailed      BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataAgentScheduleOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionCreated, BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataAgentScheduleOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataAgentScheduleOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                                    `json:"message"`
	JSON    bulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorJSON contains the JSON
// metadata for the struct [BulkWorkspaceApplyResultDataAgentScheduleOutcomeError]
type bulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentScheduleOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                          `json:"@type"`
	ExtraFields map[string]interface{}                                          `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorDetailJSON contains the
// JSON metadata for the struct
// [BulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentScheduleOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataAgentVariationOutcome struct {
	Action BulkWorkspaceApplyResultDataAgentVariationOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error      BulkWorkspaceApplyResultDataAgentVariationOutcomeError `json:"error"`
	ExternalID string                                                 `json:"externalId"`
	// AgentVariation resource
	Resource AgentVariation                                        `json:"resource"`
	JSON     bulkWorkspaceApplyResultDataAgentVariationOutcomeJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentVariationOutcomeJSON contains the JSON metadata
// for the struct [BulkWorkspaceApplyResultDataAgentVariationOutcome]
type bulkWorkspaceApplyResultDataAgentVariationOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	ExternalID  apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentVariationOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentVariationOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataAgentVariationOutcomeAction string

const (
	BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataAgentVariationOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionCreated     BulkWorkspaceApplyResultDataAgentVariationOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataAgentVariationOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataAgentVariationOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataAgentVariationOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionFailed      BulkWorkspaceApplyResultDataAgentVariationOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataAgentVariationOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionCreated, BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataAgentVariationOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataAgentVariationOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataAgentVariationOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                                     `json:"message"`
	JSON    bulkWorkspaceApplyResultDataAgentVariationOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentVariationOutcomeErrorJSON contains the JSON
// metadata for the struct [BulkWorkspaceApplyResultDataAgentVariationOutcomeError]
type bulkWorkspaceApplyResultDataAgentVariationOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentVariationOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentVariationOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataAgentVariationOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                           `json:"@type"`
	ExtraFields map[string]interface{}                                           `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataAgentVariationOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataAgentVariationOutcomeErrorDetailJSON contains the
// JSON metadata for the struct
// [BulkWorkspaceApplyResultDataAgentVariationOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataAgentVariationOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataAgentVariationOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataAgentVariationOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataMemoryEntryOutcome struct {
	Action BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error      BulkWorkspaceApplyResultDataMemoryEntryOutcomeError `json:"error"`
	ExternalID string                                              `json:"externalId"`
	// MemoryEntry is a single keyed value within a MemoryLayer. Entries are addressed
	// by their key, which follows the S3 object key safe-character convention (see
	// MemoryEntrySpec.key for the full rule). Keys are unique within a single layer;
	// the same key may appear in multiple layers, in which case the LIFO stack-walk
	// determines which one wins for a given objective.
	//
	// MemoryEntry is the summary shape, returned by ListMemoryEntries. It does not
	// carry the entry body — callers that need the body must fetch the entry
	// individually via GetMemoryEntry, which returns a MemoryEntryDetail.
	Resource MemoryEntry                                        `json:"resource"`
	JSON     bulkWorkspaceApplyResultDataMemoryEntryOutcomeJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataMemoryEntryOutcomeJSON contains the JSON metadata
// for the struct [BulkWorkspaceApplyResultDataMemoryEntryOutcome]
type bulkWorkspaceApplyResultDataMemoryEntryOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	ExternalID  apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataMemoryEntryOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataMemoryEntryOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction string

const (
	BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionCreated     BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionFailed      BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataMemoryEntryOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionCreated, BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataMemoryEntryOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataMemoryEntryOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                                  `json:"message"`
	JSON    bulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorJSON contains the JSON
// metadata for the struct [BulkWorkspaceApplyResultDataMemoryEntryOutcomeError]
type bulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataMemoryEntryOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                        `json:"@type"`
	ExtraFields map[string]interface{}                                        `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorDetailJSON contains the JSON
// metadata for the struct
// [BulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataMemoryEntryOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataMemoryLayerOutcome struct {
	Action BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error      BulkWorkspaceApplyResultDataMemoryLayerOutcomeError `json:"error"`
	ExternalID string                                              `json:"externalId"`
	// MemoryLayer is a named container of memory entries that can be composed into an
	// objective's memory stack. Layers are workspace-scoped resources. The layer type
	// controls how its entries participate in the agent loop — see MemoryLayerType for
	// details.
	//
	// See "Memory stack composition" above for how layers compose at lookup time.
	Resource MemoryLayer                                        `json:"resource"`
	JSON     bulkWorkspaceApplyResultDataMemoryLayerOutcomeJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataMemoryLayerOutcomeJSON contains the JSON metadata
// for the struct [BulkWorkspaceApplyResultDataMemoryLayerOutcome]
type bulkWorkspaceApplyResultDataMemoryLayerOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	ExternalID  apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataMemoryLayerOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataMemoryLayerOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction string

const (
	BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionCreated     BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionFailed      BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataMemoryLayerOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionCreated, BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataMemoryLayerOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataMemoryLayerOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                                  `json:"message"`
	JSON    bulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorJSON contains the JSON
// metadata for the struct [BulkWorkspaceApplyResultDataMemoryLayerOutcomeError]
type bulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataMemoryLayerOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                        `json:"@type"`
	ExtraFields map[string]interface{}                                        `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorDetailJSON contains the JSON
// metadata for the struct
// [BulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataMemoryLayerOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataToolOutcome struct {
	Action BulkWorkspaceApplyResultDataToolOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error      BulkWorkspaceApplyResultDataToolOutcomeError `json:"error"`
	ExternalID string                                       `json:"externalId"`
	Resource   Tool                                         `json:"resource"`
	JSON       bulkWorkspaceApplyResultDataToolOutcomeJSON  `json:"-"`
}

// bulkWorkspaceApplyResultDataToolOutcomeJSON contains the JSON metadata for the
// struct [BulkWorkspaceApplyResultDataToolOutcome]
type bulkWorkspaceApplyResultDataToolOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	ExternalID  apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataToolOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataToolOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataToolOutcomeAction string

const (
	BulkWorkspaceApplyResultDataToolOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataToolOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataToolOutcomeActionActionCreated     BulkWorkspaceApplyResultDataToolOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataToolOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataToolOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataToolOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataToolOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataToolOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataToolOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataToolOutcomeActionActionFailed      BulkWorkspaceApplyResultDataToolOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataToolOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataToolOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataToolOutcomeActionActionCreated, BulkWorkspaceApplyResultDataToolOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataToolOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataToolOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataToolOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataToolOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataToolOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                           `json:"message"`
	JSON    bulkWorkspaceApplyResultDataToolOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataToolOutcomeErrorJSON contains the JSON metadata for
// the struct [BulkWorkspaceApplyResultDataToolOutcomeError]
type bulkWorkspaceApplyResultDataToolOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataToolOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataToolOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataToolOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                 `json:"@type"`
	ExtraFields map[string]interface{}                                 `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataToolOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataToolOutcomeErrorDetailJSON contains the JSON
// metadata for the struct [BulkWorkspaceApplyResultDataToolOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataToolOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataToolOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataToolOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataToolSetOutcome struct {
	Action BulkWorkspaceApplyResultDataToolSetOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error      BulkWorkspaceApplyResultDataToolSetOutcomeError `json:"error"`
	ExternalID string                                          `json:"externalId"`
	Resource   ToolSet                                         `json:"resource"`
	JSON       bulkWorkspaceApplyResultDataToolSetOutcomeJSON  `json:"-"`
}

// bulkWorkspaceApplyResultDataToolSetOutcomeJSON contains the JSON metadata for
// the struct [BulkWorkspaceApplyResultDataToolSetOutcome]
type bulkWorkspaceApplyResultDataToolSetOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	ExternalID  apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataToolSetOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataToolSetOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataToolSetOutcomeAction string

const (
	BulkWorkspaceApplyResultDataToolSetOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataToolSetOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataToolSetOutcomeActionActionCreated     BulkWorkspaceApplyResultDataToolSetOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataToolSetOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataToolSetOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataToolSetOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataToolSetOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataToolSetOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataToolSetOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataToolSetOutcomeActionActionFailed      BulkWorkspaceApplyResultDataToolSetOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataToolSetOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataToolSetOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataToolSetOutcomeActionActionCreated, BulkWorkspaceApplyResultDataToolSetOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataToolSetOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataToolSetOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataToolSetOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataToolSetOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataToolSetOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                              `json:"message"`
	JSON    bulkWorkspaceApplyResultDataToolSetOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataToolSetOutcomeErrorJSON contains the JSON metadata
// for the struct [BulkWorkspaceApplyResultDataToolSetOutcomeError]
type bulkWorkspaceApplyResultDataToolSetOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataToolSetOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataToolSetOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataToolSetOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                    `json:"@type"`
	ExtraFields map[string]interface{}                                    `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataToolSetOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataToolSetOutcomeErrorDetailJSON contains the JSON
// metadata for the struct [BulkWorkspaceApplyResultDataToolSetOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataToolSetOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataToolSetOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataToolSetOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataVariationAssignmentOutcome struct {
	Action BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error BulkWorkspaceApplyResultDataVariationAssignmentOutcomeError `json:"error"`
	// VariationAssignment is a read-only reference to a single tool, tool set, or
	// sub-agent attached to a variation. Clients read the full set of assignments via
	// `AgentVariationInfo.assignments`; mutations go through the dedicated add/remove
	// assignment endpoints under
	// /v1/agents/{agent_id}/variations/{variation_id}/assignments.
	//
	// The `id` identifies the assignment row itself (not the referenced resource) and
	// is the handle used to remove the assignment. It is returned by the add endpoint
	// and present on every entry in AgentVariationInfo.assignments.
	Resource VariationAssignment                                        `json:"resource"`
	JSON     bulkWorkspaceApplyResultDataVariationAssignmentOutcomeJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataVariationAssignmentOutcomeJSON contains the JSON
// metadata for the struct [BulkWorkspaceApplyResultDataVariationAssignmentOutcome]
type bulkWorkspaceApplyResultDataVariationAssignmentOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataVariationAssignmentOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataVariationAssignmentOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction string

const (
	BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionCreated     BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionFailed      BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataVariationAssignmentOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionCreated, BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataVariationAssignmentOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataVariationAssignmentOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                                          `json:"message"`
	JSON    bulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorJSON contains the
// JSON metadata for the struct
// [BulkWorkspaceApplyResultDataVariationAssignmentOutcomeError]
type bulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataVariationAssignmentOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                                `json:"@type"`
	ExtraFields map[string]interface{}                                                `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorDetailJSON contains
// the JSON metadata for the struct
// [BulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataVariationAssignmentOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataVariationMemoryLayerOutcome struct {
	Action BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction `json:"action"`
	// The `Status` type defines a logical error model that is suitable for different
	// programming environments, including REST APIs and RPC APIs. It is used by
	// [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of
	// data: error code, error message, and error details. You can find out more about
	// this error model and how to work with it in the
	// [API Design Guide](https://cloud.google.com/apis/design/errors).
	Error BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeError `json:"error"`
	// VariationMemoryLayerAssignment attaches a single MemoryLayer to a variation at a
	// given position in the variation's baseline memory stack. A variation has at most
	// one assignment per memory_layer_id.
	//
	// Variations only support whole-layer attachments — entry pinning is an
	// objective-level capability.
	Resource VariationMemoryLayerAssignment                              `json:"resource"`
	JSON     bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeJSON contains the JSON
// metadata for the struct
// [BulkWorkspaceApplyResultDataVariationMemoryLayerOutcome]
type bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeJSON struct {
	Action      apijson.Field
	Error       apijson.Field
	Resource    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataVariationMemoryLayerOutcome) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction string

const (
	BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionUnspecified BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionCreated     BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction = "ACTION_CREATED"
	BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionUpdated     BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction = "ACTION_UPDATED"
	BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionUnchanged   BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction = "ACTION_UNCHANGED"
	BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionDeleted     BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction = "ACTION_DELETED"
	BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionFailed      BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction = "ACTION_FAILED"
)

func (r BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionUnspecified, BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionCreated, BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionUpdated, BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionUnchanged, BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionDeleted, BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeActionActionFailed:
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
type BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeError struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code int64 `json:"code"`
	// A list of messages that carry the error details. There is a common set of
	// message types for APIs to use.
	Details []BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorDetail `json:"details"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message string                                                           `json:"message"`
	JSON    bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorJSON contains the
// JSON metadata for the struct
// [BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeError]
type bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorJSON struct {
	Code        apijson.Field
	Details     apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorJSON) RawJSON() string {
	return r.raw
}

// Contains an arbitrary serialized message along with a @type that describes the
// type of the serialized message.
type BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorDetail struct {
	// The type of the serialized message.
	Type        string                                                                 `json:"@type"`
	ExtraFields map[string]interface{}                                                 `json:"-" api:"extrafields"`
	JSON        bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorDetailJSON `json:"-"`
}

// bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorDetailJSON contains
// the JSON metadata for the struct
// [BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorDetail]
type bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorDetailJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bulkWorkspaceApplyResultDataVariationMemoryLayerOutcomeErrorDetailJSON) RawJSON() string {
	return r.raw
}

type BulkWorkspaceResourceResultListParams struct {
	// Filter by action.
	Action param.Field[BulkWorkspaceResourceResultListParamsAction] `query:"action"`
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Sort order for results (asc or desc by creation time)
	SortOrder param.Field[string] `query:"sortOrder"`
	// Filter by data.type discriminator (e.g., "toolSet", "memoryEntry").
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [BulkWorkspaceResourceResultListParams]'s query parameters
// as `url.Values`.
func (r BulkWorkspaceResourceResultListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by action.
type BulkWorkspaceResourceResultListParamsAction string

const (
	BulkWorkspaceResourceResultListParamsActionActionUnspecified BulkWorkspaceResourceResultListParamsAction = "ACTION_UNSPECIFIED"
	BulkWorkspaceResourceResultListParamsActionActionCreated     BulkWorkspaceResourceResultListParamsAction = "ACTION_CREATED"
	BulkWorkspaceResourceResultListParamsActionActionUpdated     BulkWorkspaceResourceResultListParamsAction = "ACTION_UPDATED"
	BulkWorkspaceResourceResultListParamsActionActionUnchanged   BulkWorkspaceResourceResultListParamsAction = "ACTION_UNCHANGED"
	BulkWorkspaceResourceResultListParamsActionActionDeleted     BulkWorkspaceResourceResultListParamsAction = "ACTION_DELETED"
	BulkWorkspaceResourceResultListParamsActionActionFailed      BulkWorkspaceResourceResultListParamsAction = "ACTION_FAILED"
)

func (r BulkWorkspaceResourceResultListParamsAction) IsKnown() bool {
	switch r {
	case BulkWorkspaceResourceResultListParamsActionActionUnspecified, BulkWorkspaceResourceResultListParamsActionActionCreated, BulkWorkspaceResourceResultListParamsActionActionUpdated, BulkWorkspaceResourceResultListParamsActionActionUnchanged, BulkWorkspaceResourceResultListParamsActionActionDeleted, BulkWorkspaceResourceResultListParamsActionActionFailed:
		return true
	}
	return false
}
