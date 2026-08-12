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

// ObjectiveTaskService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveTaskService] method instead.
type ObjectiveTaskService struct {
	options []option.RequestOption
}

// NewObjectiveTaskService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectiveTaskService(opts ...option.RequestOption) (r ObjectiveTaskService) {
	r = ObjectiveTaskService{}
	r.options = opts
	return
}

// Retrieves a task by ID from an objective
func (r *ObjectiveTaskService) Get(ctx context.Context, objectiveID string, id string, query ObjectiveTaskGetParams, opts ...option.RequestOption) (res *ObjectiveTask, err error) {
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
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tasks/%s", url.PathEscape(query.WorkspaceID.Value), url.PathEscape(objectiveID), url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all tasks for an objective
func (r *ObjectiveTaskService) List(ctx context.Context, objectiveID string, params ObjectiveTaskListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveTask], err error) {
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
	path := fmt.Sprintf("v1/workspaces/%s/objectives/%s/tasks", url.PathEscape(params.WorkspaceID.Value), url.PathEscape(objectiveID))
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

// Lists all tasks for an objective
func (r *ObjectiveTaskService) ListAutoPaging(ctx context.Context, objectiveID string, params ObjectiveTaskListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveTask] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, objectiveID, params, opts...))
}

// ObjectiveTask represents a task within an objective, typically created and
// managed by an AI agent to track progress toward completing the objective.
type ObjectiveTask struct {
	Data ObjectiveTaskData `json:"data" api:"required"`
	// BareMetadata contains the minimal metadata for a resource: the ID and an
	// optional human-readable name. These are used for reference fields where the full
	// metadata (account scoping, timestamps, labels, external IDs) is not needed —
	// e.g., the tool references inside an agent variation spec or the tools assigned
	// to an objective. Both fields are server-populated; clients provide IDs through
	// sibling fields rather than by constructing a BareMetadata themselves.
	Metadata shared.BareMetadata `json:"metadata" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveTask) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveTaskData struct {
	// Whether the task has been completed
	Completed bool `json:"completed" api:"required"`
	// The sequential number of this task within the objective (auto-assigned, 1-based)
	Number int64 `json:"number" api:"required"`
	// Description of the task to be completed
	Task string `json:"task" api:"required"`
	// Timestamp when the task was marked as completed
	CompletedAt time.Time `json:"completedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Completed   respjson.Field
		Number      respjson.Field
		Task        respjson.Field
		CompletedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectiveTaskData) RawJSON() string { return r.JSON.raw }
func (r *ObjectiveTaskData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectiveTaskGetParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	paramObj
}

type ObjectiveTaskListParams struct {
	// Use [option.WithWorkspaceID] on the client to set a global default for this
	// field.
	WorkspaceID param.Opt[string] `path:"workspaceId,omitzero" api:"required" json:"-"`
	// Pagination cursor from previous response
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of results to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Sort order for results
	SortOrder param.Opt[string] `query:"sortOrder,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ObjectiveTaskListParams]'s query parameters as
// `url.Values`.
func (r ObjectiveTaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
