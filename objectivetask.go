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

// ObjectiveTaskService contains methods and other services that help with
// interacting with the cadenya API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectiveTaskService] method instead.
type ObjectiveTaskService struct {
	Options []option.RequestOption
}

// NewObjectiveTaskService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectiveTaskService(opts ...option.RequestOption) (r *ObjectiveTaskService) {
	r = &ObjectiveTaskService{}
	r.Options = opts
	return
}

// Retrieves a task by ID from an objective
func (r *ObjectiveTaskService) Get(ctx context.Context, objectiveID string, id string, opts ...option.RequestOption) (res *ObjectiveTask, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/tasks/%s", objectiveID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Lists all tasks for an objective
func (r *ObjectiveTaskService) List(ctx context.Context, objectiveID string, query ObjectiveTaskListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ObjectiveTask], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if objectiveID == "" {
		err = errors.New("missing required objectiveId parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/objectives/%s/tasks", objectiveID)
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

// Lists all tasks for an objective
func (r *ObjectiveTaskService) ListAutoPaging(ctx context.Context, objectiveID string, query ObjectiveTaskListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ObjectiveTask] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, objectiveID, query, opts...))
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
	JSON     objectiveTaskJSON   `json:"-"`
}

// objectiveTaskJSON contains the JSON metadata for the struct [ObjectiveTask]
type objectiveTaskJSON struct {
	Data        apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveTask) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveTaskJSON) RawJSON() string {
	return r.raw
}

type ObjectiveTaskData struct {
	// Whether the task has been completed
	Completed bool `json:"completed" api:"required"`
	// The sequential number of this task within the objective (auto-assigned, 1-based)
	Number int64 `json:"number" api:"required"`
	// Description of the task to be completed
	Task string `json:"task" api:"required"`
	// Timestamp when the task was marked as completed
	CompletedAt time.Time             `json:"completedAt" format:"date-time"`
	JSON        objectiveTaskDataJSON `json:"-"`
}

// objectiveTaskDataJSON contains the JSON metadata for the struct
// [ObjectiveTaskData]
type objectiveTaskDataJSON struct {
	Completed   apijson.Field
	Number      apijson.Field
	Task        apijson.Field
	CompletedAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObjectiveTaskData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r objectiveTaskDataJSON) RawJSON() string {
	return r.raw
}

type ObjectiveTaskListParams struct {
	// Pagination cursor from previous response
	Cursor param.Field[string] `query:"cursor"`
	// Maximum number of results to return
	Limit param.Field[int64] `query:"limit"`
	// Sort order for results
	SortOrder param.Field[string] `query:"sortOrder"`
}

// URLQuery serializes [ObjectiveTaskListParams]'s query parameters as
// `url.Values`.
func (r ObjectiveTaskListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
